package resty_tool

import (
	"cld/models"
	"cld/pkg/tool"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const indexUrl = "https://jxw.sylu.edu.cn/xtgl"
const courseUrl = "https://jxw.sylu.edu.cn/kbcx"
const gradeUrl = "https://jxw.sylu.edu.cn/cjcx"

type PublicKey struct {
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

var (
	Error302          = errors.New("Post \"/xtgl/login_slogin.html\": auto redirect is disabled")
	ErrorLapse        = errors.New("Cookie已失效！")
	ErrorCourseNoOpen = errors.New("当前学期课表暂未开放！")
	ErrorGradesNoOpen = errors.New("当前学期暂无成绩！")
)

// 一个请求头
func baseHttpHeaders() map[string]string {
	return map[string]string{
		"User-Agent":    "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0",
		"Content-Type":  "application/x-www-form-urlencoded;charset=uft-8",
		"Cache-Control": "no-cache",
	}
}

// 获取初始cookie与csrftoken
func GetIndexCookieAndCsrfToken(client *resty.Client) (csrftoken string, err error) {
	/*
		学校的教务有时候会抽风,重连多次才会有响应，不然会一直超时，这里使用goto语句优化
		不用担心死循环，一般一次两次就响应成功了，gin框架60秒也会强制推出的
	*/
	retryLimit := 4
	retries := 0

	client.SetTimeout(3 * time.Second)
lable:
	initResp, err := client.R().SetHeaders(baseHttpHeaders()).
		Get(indexUrl + "/login_slogin.html")
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() && retries < retryLimit {
			retries++
			goto lable
		}
		return "", err
	}
	Findcsrftoken := regexp.MustCompile(`id="csrftoken" name="csrftoken" value="([^"]+)"`)
	csrftoken = Findcsrftoken.FindStringSubmatch(string(initResp.Body()))[1]
	client.Cookies = initResp.Cookies()
	return
}

// 获取公匙
func GetPublicKey(client *resty.Client) (publicKey *PublicKey, err error) {
	nowTime := tool.NowTime()
	getPublicKeyResp, err := client.R().SetHeaders(baseHttpHeaders()).
		SetQueryParams(map[string]string{
			"time": nowTime,
			"_":    nowTime,
		}).Get(indexUrl + "/login_getPublicKey.html")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(getPublicKeyResp.String()), &publicKey)
	return
}

// rsa加密
func RsaByPublicKey(password string, publicKey *PublicKey) (string, error) {
	modulusBytes, err := base64.StdEncoding.DecodeString(publicKey.Modulus)
	if err != nil {
		return "", err
	}

	exponentBytes, err := base64.StdEncoding.DecodeString(publicKey.Exponent)
	if err != nil {
		return "", err
	}

	// 解析公钥
	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(modulusBytes),
		E: int(new(big.Int).SetBytes(exponentBytes).Int64()),
	}

	// 加密密码
	bypassword := []byte(password)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, bypassword)
	if err != nil {
		panic(err)
	}

	// Base64 编码加密后的密码
	encryptedPassword := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encryptedPassword, nil
}

func SyluLogin(client *resty.Client, userInfo *models.ParamBind, csrfToken string) (cookies []*http.Cookie, err error) {
	loginresponse, err := client.SetRedirectPolicy(resty.NoRedirectPolicy()).R().SetFormData(map[string]string{
		"csrftoken": csrfToken,
		"language":  "zh_CN",
		"yhm":       userInfo.StudentID,
		"mm":        userInfo.Password,
	}).SetQueryParam("time", tool.NowTime()).SetHeaders(baseHttpHeaders()).
		Post(indexUrl + "/login_slogin.html")

	if err != nil && err.Error() == Error302.Error() {
		return loginresponse.Cookies(), nil
	} else if err != nil {
		return nil, errors.New("服务器连接失败:" + err.Error())
	} else {
		return nil, errors.New("账号或密码错误")
	}
}

func GetCourseByCourseInfo(client *resty.Client, getCourseInfo *models.ParamCourse) (courses *models.ReqCourse, err error) {

	client.SetHostURL(courseUrl)
	defer client.GetClient().CloseIdleConnections()

	formData := map[string]string{
		"xnm":    strconv.Itoa(getCourseInfo.Year),
		"zs":     "1",
		"doType": " app",
		"xqm":    strconv.Itoa(getCourseInfo.Semester),
		"kblx":   "1",
	}
	response, err := client.R().
		SetFormData(formData).
		SetHeader("Cookie", getCourseInfo.Cookie).
		Post("/xskbcxMobile_cxXsKb.html?gnmkdm=N2154")
	if err != nil {
		return nil, err
	}
	if string(response.Body()) == "null" {
		return nil, ErrorLapse
	}
	schedule := new(models.Schedule)
	json.Unmarshal(response.Body(), schedule)
	if len(schedule.KbList) == 0 {
		return nil, ErrorCourseNoOpen
	}

	courses = new(models.ReqCourse)

	time := strings.Split(schedule.RqazcList[0].Rq, "-")
	courses.StartTime.Year, _ = strconv.Atoi(time[0])
	courses.StartTime.Month, _ = strconv.Atoi(time[1])
	courses.StartTime.Day, _ = strconv.Atoi(time[2])

	var course models.JsonCourse
	for _, v := range schedule.KbList {
		course.Name = v.Name
		course.Teacher = v.Teacher
		course.Location = v.Location
		course.Category = v.Category
		course.Method = v.Method
		course.ClassID = v.ID
		course.Section, course.SectionCount = timeToInt(v.Time)
		course.WeekDay, _ = strconv.Atoi(v.WeekDay)
		course.WeekS = parseWeeks(v.WeekS)
		courses.Courses = append(courses.Courses, course)
	}

	return
}

func GetGradesByGradesInfo(client *resty.Client, gradesInfo *models.ParamGrades) (jsongrades []models.JsonGrades, err error) {
	client.SetHostURL(gradeUrl)
	defer client.GetClient().CloseIdleConnections()

	querData := map[string]string{
		"doType": "query",
		"gnmkdm": "N305005",
	}

	formData := map[string]string{
		"xnm":                  strconv.Itoa(gradesInfo.Year),
		"xqm":                  strconv.Itoa(gradesInfo.Semester),
		"queryModel.showCount": "30", //这个参数是成绩的门数，直接拉到30，应该不会有人超过这个数吧
	}

	response, err := client.R().SetQueryParams(querData).
		SetFormData(formData).
		SetHeader("Cookie", gradesInfo.Cookie).
		Post("/cjcx_cxXsgrcj.html")

	if err != nil {
		return nil, err
	}
	if strings.Contains(string(response.Header().Get("Content-Type")), "text/html") {
		return nil, ErrorLapse
	}

	jsongrades = make([]models.JsonGrades, 0)

	Grades := new(models.Grades)
	json.Unmarshal([]byte(response.String()), Grades)

	if len(Grades.Items) < 1 {
		return nil, ErrorGradesNoOpen
	}

	var grade models.JsonGrades
	for _, v := range Grades.Items {
		grade.Name = v.Kcmc
		grade.ClassID = v.JxbID
		grade.Teacher = v.Jsxm
		grade.IsDegree = isDegree(v.Sfxwkc)

		grade.Credits, _ = strconv.ParseFloat(v.Xf, 64)
		grade.GPA, _ = strconv.ParseFloat(v.Jd, 64)
		grade.GradePoints, _ = strconv.ParseFloat(v.Xfjd, 64)
		grade.Fraction, _ = strconv.ParseFloat(v.Bfzcj, 64)
		grade.Grade = v.Cj
		jsongrades = append(jsongrades, grade)
	}

	return
}

func isDegree(degree string) bool {
	if degree == "是" {
		return true
	} else {
		return false
	}
}

func parseWeeks(input string) []int {
	var weeks []int

	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		re := regexp.MustCompile(`(\d+)`)
		bounds := re.FindAllString(r, -1)

		var start int
		var end int
		//有些是1-2周 有些是2周这种 分开看待
		if len(bounds) > 1 {
			start, _ = strconv.Atoi(bounds[0])
			end, _ = strconv.Atoi(bounds[1])
		} else {
			start, _ = strconv.Atoi(bounds[0])
			end = start
		}

		for i := start; i <= end; i++ {
			weeks = append(weeks, i)
		}
	}

	return weeks
}

func timeToInt(time string) (section int, sectionCount int) {
	if len(time) <= 1 {
		var err error
		section, err = strconv.Atoi(time)
		if err != nil {
			section = 0
		}
		sectionCount = 0
		return
	}

	sections := strings.Split(time, "-")
	section, _ = strconv.Atoi(sections[0])
	lastTime, _ := strconv.Atoi(sections[1])
	sectionCount = lastTime - section + 1

	//我觉得应该不会有15-16节吧
	return
}

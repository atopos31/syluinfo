package resty_tool

import (
	"cld/models"
	"cld/pkg/tool"
	"cld/settings"
	"encoding/json"
	"errors"
	"fmt"
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

type Myresty struct {
	*resty.Client
}

// 一个请求头
func baseHttpHeaders() map[string]string {
	return map[string]string{
		"User-Agent":    "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0",
		"Content-Type":  "application/x-www-form-urlencoded;charset=uft-8",
		"Cache-Control": "no-cache",
	}
}

// 创建MyResty对象
func NewMyResty() *Myresty {
	cfg := settings.Conf.Proxy

	restyClient := resty.New()
	restyClient.SetHeaders(baseHttpHeaders())

	if cfg.Host != "" && cfg.Port != "" && cfg.Type != "" {

		pUrl := fmt.Sprintf("%s://%s:%s", cfg.Type, cfg.Host, cfg.Port)
		restyClient.SetProxy(pUrl)

	}

	return &Myresty{
		Client: restyClient,
	}
}

// 通过模拟登录获取Cookie
func (myRes *Myresty) LoginAndGetCookie(studentID string, password string) (cookie string, err error) {
	csrfToken, err := myRes.setIndexCookieAndGetCsrfToken()
	if err != nil {
		return "", err
	}

	publicKey, err := myRes.getPublicKey()
	if err != nil {
		return "", err
	}

	enPass, err := rsaByPublicKey(password, publicKey)
	if err != nil {
		return "", err
	}

	resCookies, err := myRes.syluLogin(studentID, enPass, csrfToken)
	if err != nil {
		return "", err
	}
	cookie = resCookies[1].Name + "=" + resCookies[1].Value
	return
}

// 通过学期信息获取课表
func (myResty *Myresty) GetCourseByCourseInfo(getCourseInfo *models.ParamCourse) (courses *models.ReqCourse, err error) {

	myResty.SetHostURL(courseUrl)
	defer myResty.GetClient().CloseIdleConnections()

	formData := map[string]string{
		"xnm":    strconv.Itoa(getCourseInfo.Year),
		"zs":     "1",
		"doType": " app",
		"xqm":    strconv.Itoa(getCourseInfo.Semester),
		"kblx":   "1",
	}
	response, err := myResty.R().
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

// 通过学期信息获取成绩列表
func (myResty *Myresty) GetGradesByGradesInfo(gradesInfo *models.ParamGrades) (jsongrades []models.JsonGrades, err error) {
	myResty.SetHostURL(gradeUrl)
	defer myResty.GetClient().CloseIdleConnections()

	querData := map[string]string{
		"doType": "query",
		"gnmkdm": "N305005",
	}

	formData := map[string]string{
		"xnm":                  strconv.Itoa(gradesInfo.Year),
		"xqm":                  strconv.Itoa(gradesInfo.Semester),
		"queryModel.showCount": "30", //这个参数是成绩的门数，直接拉到30，应该不会有人超过这个数吧
	}

	response, err := myResty.R().SetQueryParams(querData).
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

// 获取初始Cookie和CsrfToken
func (myRes *Myresty) setIndexCookieAndGetCsrfToken() (csrfToken string, err error) {
	retryLimit := 4
	retries := 0

	myRes.SetTimeout(3 * time.Second)
lable:
	initResp, err := myRes.R().SetHeaders(baseHttpHeaders()).Get(indexUrl + "/login_slogin.html")
	if err != nil {
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() && retries < retryLimit {
			retries++
			goto lable
		}
		return "", err
	}
	myRes.Cookies = initResp.Cookies()

	Findcsrftoken := regexp.MustCompile(`id="csrftoken" name="csrftoken" value="([^"]+)"`)
	csrfToken = Findcsrftoken.FindStringSubmatch(string(initResp.Body()))[1]

	return
}

// 获取公钥
func (myRes *Myresty) getPublicKey() (publicKey *PublicKey, err error) {
	nowTime := tool.NowTime()
	getPublicKeyResp, err := myRes.R().SetHeaders(baseHttpHeaders()).
		SetQueryParams(map[string]string{
			"time": nowTime,
			"_":    nowTime,
		}).Get(indexUrl + "/login_getPublicKey.html")
	if err != nil {
		return nil, err
	}
	//对指针取&避免空指针
	if err := json.Unmarshal(getPublicKeyResp.Body(), &publicKey); err != nil {
		return nil, err
	}
	return
}

func (myResty *Myresty) syluLogin(studentID string, enPass string, csrfToken string) (cookies []*http.Cookie, err error) {
	loginResponse, err := myResty.SetRedirectPolicy(resty.NoRedirectPolicy()).R().SetFormData(map[string]string{
		"csrftoken": csrfToken,
		"language":  "zh_CN",
		"yhm":       studentID,
		"mm":        enPass,
	}).SetQueryParam("time", tool.NowTime()).SetHeaders(baseHttpHeaders()).
		Post(indexUrl + "/login_slogin.html")

	if err != nil && err.Error() == Error302.Error() {
		return loginResponse.Cookies(), nil
	} else if err != nil {
		return nil, errors.New("服务器连接失败:" + err.Error())
	} else {
		return nil, errors.New("账号或密码错误")
	}
}

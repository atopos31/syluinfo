package logic

import (
	collytool "cld/dao/colly_tool"
	"cld/dao/mysql"
	"cld/dao/resty_tool"
	"cld/models"
	"cld/settings"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// 获取代理
func getHttpProxy() string {
	if settings.Conf.Proxy.Host != "" && settings.Conf.Proxy.Port != "" {
		return "http://" + settings.Conf.Proxy.Host + ":" + settings.Conf.Proxy.Port
	}
	return ""
}

// 绑定教务处理
func BintLogin(userInfo *models.ParamBind, userID int64) (userSyluInfo *models.ReqBind, err error) {
	//保存原密码
	oPassword := userInfo.Password

	bindClient := resty.New()

	//设置代理
	proxyURL := getHttpProxy()
	if proxyURL != "" {
		bindClient.SetProxy(proxyURL)
	}
	bindClient.SetCloseConnection(true)

	//获取cookie和csrtoken
	csrfToken, err := resty_tool.GetIndexCookieAndCsrfToken(bindClient)
	if err != nil {
		return nil, err
	}

	//获取公匙
	publicKey, err := resty_tool.GetPublicKey(bindClient)
	if err != nil {
		return nil, err
	}

	//加密
	userInfo.Password, err = resty_tool.RsaByPublicKey(userInfo.Password, publicKey)
	if err != nil {
		return nil, err
	}

	//登录
	cookies, err := resty_tool.SyluLogin(bindClient, userInfo, csrfToken)
	if err != nil {
		return nil, err
	}
	//断开连接
	//bindClient.GetClient().CloseIdleConnections()

	//存储或更新账号密码
	userInfo.Password = oPassword
	if err := mysql.CreateOrUpdateSyluPass(userInfo, userID); err != nil {
		return nil, err
	}

	//更换cookie为string形式
	bindClient.Cookies[0] = cookies[1]
	cookieString := cookiesToString(bindClient.Cookies)

	col := collytool.NewMyCollector()
	syluUser, err := col.GetInforamation(cookieString, userInfo.StudentID)
	if err != nil {
		zap.L().Error("collytool.GetInforamation", zap.String("id", userInfo.StudentID))
		return
	}
	syluUser.Uuid = userID
	if err := mysql.CreateOrUpdateSyluUser(syluUser); err != nil {
		return nil, err
	}
	//构建返回数据
	userSyluInfo = &models.ReqBind{
		Cookie: cookieString,
		SyluInfo: &models.ReqSyluInfo{
			ReUsername: syluUser.ReUsername,
			StudentID:  syluUser.StudentID,
			Grade:      syluUser.Grade,
			Major:      syluUser.Major,
			College:    syluUser.College,
		},
	}

	return
}

// 根据学号获取学期列表
func GetSemeSter(userID int64) (semeList *models.ResSemeSter, err error) {
	syluPass, err := mysql.GetSyluPassByUuid(userID)
	if err != nil {
		return nil, err
	}

	semeList = new(models.ResSemeSter)
	semeList.Index, err = getIndesSeme(syluPass.StudentID)
	if err != nil {
		return nil, err
	}
	semeList.List, err = getSemeList(syluPass.StudentID)
	if err != nil {
		return nil, err
	}

	return

}

// 获取教务cookie
func GetCookie(userID int64) (cookieString string, err error) {
	syluPass, err := mysql.GetSyluPassByUuid(userID)
	if err != nil {
		return "", err
	}

	bindClient := resty.New()
	defer bindClient.GetClient().CloseIdleConnections()
	//设置代理
	bindClient.SetProxy(getHttpProxy())

	//获取cookie和csrtoken
	csrfToken, err := resty_tool.GetIndexCookieAndCsrfToken(bindClient)
	if err != nil {
		return "", err
	}
	//获取公匙
	publicKey, err := resty_tool.GetPublicKey(bindClient)
	if err != nil {
		return "", err
	}

	userInfo := &models.ParamBind{
		StudentID: syluPass.StudentID,
		Password:  syluPass.Password,
	}

	//加密
	userInfo.Password, err = resty_tool.RsaByPublicKey(syluPass.Password, publicKey)
	if err != nil {
		return "", err
	}
	//登录
	cookies, err := resty_tool.SyluLogin(bindClient, userInfo, csrfToken)
	if err != nil {
		return "", err
	}
	//更换cookie为string形式
	bindClient.Cookies[0] = cookies[1]
	cookieString = cookiesToString(bindClient.Cookies)
	return
}

// 获取课程表
func GetCourse(courseInfo *models.ParamCourse) (reqCourse *models.ReqCourse, err error) {
	client := resty.New()
	client.SetProxy(getHttpProxy())
	return resty_tool.GetCourseByCourseInfo(client, courseInfo)
}

// 根据学期获取成绩列表
func GetGrades(gradesInfo *models.ParamGrades) (resGrades *models.ResGrades, err error) {
	resGrades = new(models.ResGrades)
	resGrades.Year = getYear(gradesInfo.Year)
	resGrades.Semester = getSemester(gradesInfo.Semester)

	client := resty.New()
	client.SetProxy(getHttpProxy())
	List, err := resty_tool.GetGradesByGradesInfo(client, gradesInfo)
	if err != nil {
		return nil, err
	}
	resGrades.GradesList = List
	return
}

// 根据课程id获取成绩详情
func GetGradeDetail(gradeDetailInfo *models.ParamGradeDetaile) (resGradeDetail []*models.ResGradeDetail, err error) {
	col := collytool.NewMyCollector()
	return col.GetGradeDetail(gradeDetailInfo)
}

// 获取全部课程平均绩点和学位课平均绩点
func GetGpas(bindGpa *models.ParamGpa) (resGpa *models.ResGpa, err error) {
	col := collytool.NewMyCollector()
	return col.GetGpas(bindGpa.Cookie)
}

// 将cookie转string
func cookiesToString(cookies []*http.Cookie) string {
	var cookieStrings []string
	for _, cookie := range cookies {
		cookieStrings = append(cookieStrings, fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}

	return strings.Join(cookieStrings, "; ")
}

package logic

import (
	collytool "cld/dao/colly_tool"
	"cld/dao/redis"
	restytool "cld/dao/resty_tool"
	"cld/models"

	"go.uber.org/zap"
)

// 绑定教务处理
func BintLogin(syluInfo *models.ParamBind, userID int64) (userSyluInfo *models.ReqBind, err error) {
	myRes := restytool.NewMyResty()
	cookieString, err := myRes.LoginAndGetCookie(syluInfo.StudentID, syluInfo.Password)
	if err != nil {
		zap.L().Error("myRes.LoginAndGetCookie err", zap.String("id", syluInfo.StudentID), zap.Error(err))
		return
	}

	col := collytool.NewMyCollector()
	syluUser, err := col.GetInforamation(cookieString, syluInfo.StudentID)
	if err != nil {
		zap.L().Error("collytool.GetInforamation", zap.String("id", syluInfo.StudentID))
		return
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
func GetSemeSter(userID string) (semeList *models.ResSemeSter, err error) {

	semeList = new(models.ResSemeSter)
	semeList.Index, err = getIndesSeme(userID)
	if err != nil {
		return nil, err
	}
	semeList.List, err = getSemeList(userID)
	if err != nil {
		return nil, err
	}

	return

}

// 获取教务cookie
func GetCookie(syluInfo *models.ParamBind) (cookieString string, err error) {

	myRes := restytool.NewMyResty()
	cookieString, err = myRes.LoginAndGetCookie(syluInfo.StudentID, syluInfo.Password)
	if err != nil {
		zap.L().Error("myRes.LoginAndGetCookie err", zap.String("id", syluInfo.StudentID))
		return "", err
	}

	return
}

// 获取课程表
func GetCourse(courseInfo *models.ParamCourse) (reqCourse *models.ReqCourse, err error) {
	myRes := restytool.NewMyResty()
	reqCourse, err = myRes.GetCourseByCourseInfo(courseInfo)
	if err != nil {
		return nil, err
	}

	return
}

// 根据学期获取成绩列表
func GetGrades(gradesInfo *models.ParamGrades) (resGrades *models.ResGrades, err error) {
	resGrades = new(models.ResGrades)
	resGrades.Year = getYear(gradesInfo.Year)
	resGrades.Semester = getSemester(gradesInfo.Semester)

	myRes := restytool.NewMyResty()
	list, err := myRes.GetGradesByGradesInfo(gradesInfo)
	if err != nil {
		return nil, err
	}
	resGrades.GradesList = list
	return
}

// 根据课程id获取成绩详情
func GetGradeDetail(uuid int64, gradeDetailInfo *models.ParamGradeDetaile) (resGradeDetail []*models.ResGradeDetail, err error) {
	resGradeDetail, err = redis.GetGradeDetail(uuid, gradeDetailInfo.ClassID)
	if err == nil {
		return
	}

	col := collytool.NewMyCollector()
	resGradeDetail, err = col.GetGradeDetail(gradeDetailInfo)

	if err := redis.SetGradeDetail(uuid, gradeDetailInfo.ClassID, resGradeDetail); err != nil {
		zap.L().Error("redis.SetGradeDetail Error:", zap.Error(err))
	}
	return
}

// 获取全部课程平均绩点和学位课平均绩点
func GetGpas(bindGpa *models.ParamGpa) (resGpa *models.ResGpa, err error) {
	col := collytool.NewMyCollector()
	return col.GetGpas(bindGpa.Cookie)
}

// 获取校历
func GetCale(cookie string) (*models.ResSchoolCale, error) {
	col := collytool.NewMyCollector()
	return col.GetSchoolCalendar(cookie)
}

func GetInva(cookie string) (resInva []models.ResInva, err error) {
	myRes := restytool.NewMyResty()
	return myRes.GetInva(cookie)
}

func GetInvaDetail(cookie string, name string) (resInva []models.ResInvaDetail, err error) {
	myRes := restytool.NewMyResty()
	return myRes.GetInvaDetail(cookie, name)
}

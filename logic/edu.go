package logic

import (
	collytool "cld/dao/colly_tool"
	"cld/dao/mysql"
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

	//存储或更新账号密码
	if err := mysql.CreateOrUpdateSyluPass(syluInfo, userID); err != nil {
		return nil, err
	}

	col := collytool.NewMyCollector()
	syluUser, err := col.GetInforamation(cookieString, syluInfo.StudentID)
	if err != nil {
		zap.L().Error("collytool.GetInforamation", zap.String("id", syluInfo.StudentID))
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
	syluInfo, err := mysql.GetSyluPassByUuid(userID)
	if err != nil {
		return "", err
	}

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
func GetGradeDetail(gradeDetailInfo *models.ParamGradeDetaile) (resGradeDetail []*models.ResGradeDetail, err error) {
	col := collytool.NewMyCollector()
	return col.GetGradeDetail(gradeDetailInfo)
}

// 获取全部课程平均绩点和学位课平均绩点
func GetGpas(bindGpa *models.ParamGpa) (resGpa *models.ResGpa, err error) {
	col := collytool.NewMyCollector()
	return col.GetGpas(bindGpa.Cookie)
}

func GetCale(cookie string) (*models.ResSchoolCale, error) {
	col := collytool.NewMyCollector()
	return col.GetSchoolCalendar(cookie)
}

package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	Repassword string `json:"repassword" binding:"required,min=8,eqfield=Password"`
	Captcha    string `json:"captcha" binding:"required"`
}

type ParamLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type ParamReSet struct {
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=8"`
	NewPassword   string `json:"newpassword" binding:"required,min=8,nefield=Password"`
	ReNewPassword string `json:"renewpassword" binding:"required,min=8,eqfield=NewPassword"`
}

type ParamReCover struct {
	Email         string `json:"email" binding:"required,email"`
	Captcha       string `json:"captcha" binding:"required"`
	NewPassword   string `json:"newpassword" binding:"required,min=8"`
	ReNewPassword string `json:"renewpassword" binding:"required,min=8,eqfield=NewPassword"`
}

type ParamBind struct {
	StudentID string `json:"studentID" binding:"required,len=10"`
	Password  string `json:"password" binding:"required"`
}

type ParamCourse struct {
	Cookie   string `json:"cookie"`
	Year     int    `json:"year" binding:"required"`
	Semester int    `json:"semester" binding:"required,oneof=3 12"`
}

type ParamGrades struct {
	Cookie   string `json:"cookie"`
	Year     int    `json:"year" binding:"required"`
	Semester int    `json:"semester" binding:"required,oneof=3 12"`
}

type ParamGradeDetaile struct {
	ClassID  string `json:"classid"`
	Cookie   string `json:"cookie"`
	Year     int    `json:"year" binding:"required"`
	Semester int    `json:"semester" binding:"required,oneof=3 12"`
}

type ParamGpa struct {
	Cookie string `json:"cookie"`
}

package models

type ReqLogin struct {
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Token    string       `json:"token"`
	SyluInfo *ReqSyluInfo `json:"syluinfo"`
}

type ResSemeSter struct {
	Index int             `json:"index"`
	List  []*SemeSterList `json:"list"`
}

type SemeSterList struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
}

type ReqBind struct {
	Cookie   string       `json:"cookie"`
	SyluInfo *ReqSyluInfo `json:"syluinfo"`
}

type ReqSyluInfo struct {
	ReUsername string `json:"reusername" default:"肖嘉兴"`       //真实姓名
	StudentID  string `json:"studentID" default:"2203050212"` //学号
	Grade      string `json:"grade" default:"2022"`           //年级
	College    string `json:"college" default:"信息科学与工程学院"`    //学院
	Major      string `json:"major" default:"计算机科学与技术(0305)"` //专业
}

type ReqCourse struct {
	StartTime Time         `json:"starttime"`
	Courses   []JsonCourse `json:"courses"`
}

type Time struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type JsonCourse struct {
	ClassID         string `json:"classId"`
	Name            string `json:"name"`
	TeachingClasses string `json:"teachingClasses"`
	Teacher         string `json:"teacher"`
	Category        string `json:"category"`
	Method          string `json:"method"`
	Location        string `json:"location"`
	Section         int    `json:"section"`
	SectionCount    int    `json:"sectionCount"`
	WeekDay         int    `json:"weekday"`
	WeekS           []int  `json:"weeks"`
}

type ResGrades struct {
	Year       string       `json:"year"`
	Semester   string       `json:"semester"`
	GradesList []JsonGrades `json:"gradesList"`
}

type JsonGrades struct {
	Name        string  `json:"name"`
	ClassID     string  `json:"classid"`
	Teacher     string  `json:"teacher"`
	IsDegree    bool    `json:"isdegree"`
	Credits     float64 `json:"credits"`     //学分
	GPA         float64 `json:"gpa"`         //绩点
	GradePoints float64 `json:"gradepoints"` //学分绩点
	Fraction    float64 `json:"fraction"`
	Grade       string  `json:"grade"`
}

type ResCosKey struct {
	Bucket        string `json:"bucket"`
	Region        string `json:"region"`
	AllowsPath    string `json:"allowpath"`
	TmpSecretId   string `json:"tmpsecretid"`
	TmpSecretKey  string `json:"tmpsecretkey"`
	SecurityToken string `json:"securitytoken"`
	StartTime     int    `json:"starttime"`
	ExpiredTime   int    `json:"expiredtime"`
}

type ResGradeDetail struct {
	Name   string `json:"name"`
	Weight string `json:"weight"`
	Score  string `json:"score"`
}

type ResGpa struct {
	AllGpa    string `json:"allgpa"`
	DegreeGpa string `json:"degreegpa"`
}

type ResSchoolCale struct {
	Title      string       `json:"title"`
	StartTime  string       `json:"starttime"`
	EndTime    string       `json:"endtime"`
	SchoolCale []SchoolCale `json:"schoolcale"`
}

type SchoolCale struct {
	ID        string `josn:"id"`
	Name      string `json:"name"`
	StartTime string `json:"starttime"`
	EndTime   string `json:"endtime"`
}

type ResInva struct {
	Name  string  `json:"name"`
	Grade float64 `json:"grade"`
}

type ResInvaDetail struct {
	Name  string `json:"name"`
	Grade string `josn:"grade"`
	Time  string `json:"time"`
}

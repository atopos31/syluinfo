package models

type Schedule struct {
	Qsxqj string `json:"qsxqj"`

	Xsxx struct {
		Bjmc  string `json:"BJMC"`
		Xnmc  string `json:"XNMC"`
		XhId  string `json:"XH_ID"`
		Xh    string `json:"XH"`
		Xqmmc string `json:"XQMMC"`
		Jfzt  int    `json:"JFZT"`
		Xm    string `json:"XM"`
		Xqm   string `json:"XQM"`
		Xnm   string `json:"XNM"`
		Kcms  int    `json:"KCMS"`
	} `json:"xsxx"`

	SjkList []Course `json:"sjkList"`

	XqjmcMap map[string]string `json:"xqjmcMap"`

	RqazcList []*DayOff `json:"rqazcList"`

	KbList []Course `json:"kbList"`

	XsbjList []*CourseHour `json:"xsbjList"`
}

type Course struct {
	Category        string `json:"kcxz"`
	Method          string `json:"khfsmc"`
	Name            string `json:"kcmc"`
	TeachingClasses string `json:"jxbzc"`
	Teacher         string `json:"xm"`
	ID              string `json:"kch_id"`
	Location        string `json:"cdmc"`
	Time            string `json:"jcs"`
	WeekDay         string `json:"xqj"`
	WeekS           string `json:"zcd"`
}

type DayOff struct {
	Xqj int    `json:"xqj"`
	Rq  string `json:"rq"`
}

type CourseHour struct {
	Ywxsmc string `json:"ywxsmc"`
	Xslxbj string `json:"xslxbj"`
	Xsmc   string `json:"xsmc"`
	Xsdm   string `json:"xsdm"`
}

type Grades struct {
	CurrentPage   int           `json:"currentPage"`
	CurrentResult int           `json:"currentResult"`
	EntityOrField bool          `json:"entityOrField"`
	Items         []Items       `json:"items"`
	Limit         int           `json:"limit"`
	Offset        int           `json:"offset"`
	PageNo        int           `json:"pageNo"`
	PageSize      int           `json:"pageSize"`
	ShowCount     int           `json:"showCount"`
	SortName      string        `json:"sortName"`
	SortOrder     string        `json:"sortOrder"`
	Sorts         []interface{} `json:"sorts"`
	TotalCount    int           `json:"totalCount"`
	TotalPage     int           `json:"totalPage"`
	TotalResult   int           `json:"totalResult"`
}
type QueryModel struct {
	CurrentPage   int           `json:"currentPage"`
	CurrentResult int           `json:"currentResult"`
	EntityOrField bool          `json:"entityOrField"`
	Limit         int           `json:"limit"`
	Offset        int           `json:"offset"`
	PageNo        int           `json:"pageNo"`
	PageSize      int           `json:"pageSize"`
	ShowCount     int           `json:"showCount"`
	Sorts         []interface{} `json:"sorts"`
	TotalCount    int           `json:"totalCount"`
	TotalPage     int           `json:"totalPage"`
	TotalResult   int           `json:"totalResult"`
}
type UserModel struct {
	Monitor    bool   `json:"monitor"`
	RoleCount  int    `json:"roleCount"`
	RoleKeys   string `json:"roleKeys"`
	RoleValues string `json:"roleValues"`
	Status     int    `json:"status"`
	Usable     bool   `json:"usable"`
}
type Items struct {
	Bfzcj              string     `json:"bfzcj"`
	Bh                 string     `json:"bh"`
	BhID               string     `json:"bh_id"`
	Bj                 string     `json:"bj"`
	Cj                 string     `json:"cj"`
	Cjsfzf             string     `json:"cjsfzf"`
	Date               string     `json:"date"`
	DateDigit          string     `json:"dateDigit"`
	DateDigitSeparator string     `json:"dateDigitSeparator"`
	Day                string     `json:"day"`
	Jd                 string     `json:"jd"`
	JgID               string     `json:"jg_id"`
	Jgmc               string     `json:"jgmc"`
	Jgpxzd             string     `json:"jgpxzd"`
	Jsxm               string     `json:"jsxm"`
	JxbID              string     `json:"jxb_id"`
	Jxbmc              string     `json:"jxbmc"`
	Kcbj               string     `json:"kcbj"`
	Kcgsmc             string     `json:"kcgsmc,omitempty"`
	Kch                string     `json:"kch"`
	KchID              string     `json:"kch_id"`
	Kclbmc             string     `json:"kclbmc"`
	Kcmc               string     `json:"kcmc"`
	Kcxzdm             string     `json:"kcxzdm"`
	Kcxzmc             string     `json:"kcxzmc"`
	Key                string     `json:"key"`
	Khfsmc             string     `json:"khfsmc"`
	Kkbmmc             string     `json:"kkbmmc"`
	Kklxdm             string     `json:"kklxdm"`
	Ksxz               string     `json:"ksxz"`
	Ksxzdm             string     `json:"ksxzdm"`
	Listnav            string     `json:"listnav"`
	LocaleKey          string     `json:"localeKey"`
	Month              string     `json:"month"`
	NjdmID             string     `json:"njdm_id"`
	Njmc               string     `json:"njmc"`
	PageTotal          int        `json:"pageTotal"`
	Pageable           bool       `json:"pageable"`
	QueryModel         QueryModel `json:"queryModel"`
	Rangeable          bool       `json:"rangeable"`
	RowID              string     `json:"row_id"`
	Rwzxs              string     `json:"rwzxs"`
	Sfdkbcx            string     `json:"sfdkbcx"`
	Sfxwkc             string     `json:"sfxwkc"`
	Sfzh               string     `json:"sfzh"`
	Tjrxm              string     `json:"tjrxm"`
	Tjsj               string     `json:"tjsj"`
	TotalResult        string     `json:"totalResult"`
	UserModel          UserModel  `json:"userModel"`
	Xb                 string     `json:"xb"`
	Xbm                string     `json:"xbm"`
	Xf                 string     `json:"xf"`
	Xfjd               string     `json:"xfjd"`
	Xh                 string     `json:"xh"`
	XhID               string     `json:"xh_id"`
	Xm                 string     `json:"xm"`
	Xnm                string     `json:"xnm"`
	Xnmmc              string     `json:"xnmmc"`
	Xqm                string     `json:"xqm"`
	Xqmmc              string     `json:"xqmmc"`
	Year               string     `json:"year"`
	Zsxymc             string     `json:"zsxymc"`
	ZyhID              string     `json:"zyh_id"`
	Zymc               string     `json:"zymc"`
}

type Innovation struct {
	CurrentPage   int               `json:"currentPage"`
	CurrentResult int               `json:"currentResult"`
	EntityOrField bool              `json:"entityOrField"`
	Items         []InnovationItems `json:"items"`
	Limit         int               `json:"limit"`
	Offset        int               `json:"offset"`
	PageNo        int               `json:"pageNo"`
	PageSize      int               `json:"pageSize"`
	ShowCount     int               `json:"showCount"`
	SortName      string            `json:"sortName"`
	SortOrder     string            `json:"sortOrder"`
	Sorts         []any             `json:"sorts"`
	TotalCount    int               `json:"totalCount"`
	TotalPage     int               `json:"totalPage"`
	TotalResult   int               `json:"totalResult"`
}
type InnovationItems struct {
	Hdfz        float64 `json:"hdfz"`
	XhID        string  `json:"xh_id"`
	Sjyxfz      float64 `json:"sjyxfz"`
	Zgfzyq      string  `json:"zgfzyq"`
	Thfzcz      string  `json:"thfzcz"`
	Cj          string  `json:"cj"`
	Xmlbmc      string  `json:"xmlbmc"`
	RowID       int     `json:"row_id"`
	Totalresult int     `json:"totalresult"`
}

type InnovationDetail struct {
	CurrentPage   int                     `json:"currentPage"`
	CurrentResult int                     `json:"currentResult"`
	EntityOrField bool                    `json:"entityOrField"`
	Items         []InnovationDetailItems `json:"items"`
	Limit         int                     `json:"limit"`
	Offset        int                     `json:"offset"`
	PageNo        int                     `json:"pageNo"`
	PageSize      int                     `json:"pageSize"`
	ShowCount     int                     `json:"showCount"`
	SortName      string                  `json:"sortName"`
	SortOrder     string                  `json:"sortOrder"`
	Sorts         []any                   `json:"sorts"`
	TotalCount    int                     `json:"totalCount"`
	TotalPage     int                     `json:"totalPage"`
	TotalResult   int                     `json:"totalResult"`
}
type InnovationDetailItems struct {
	Sbxmmc      string `json:"sbxmmc"`
	XmbmqkID    string `json:"xmbmqk_id"`
	XhID        string `json:"xh_id"`
	Xmnr        string `json:"xmnr"`
	Xmflmc      string `json:"xmflmc"`
	Qrfzsj      string `json:"qrfzsj"`
	XmfzxssqID  string `json:"xmfzxssq_id"`
	Yxfz        string `json:"yxfz"`
	Xmlxmc      string `json:"xmlxmc"`
	Xm          string `json:"xm"`
	Sbfz        string `json:"sbfz"`
	Xmlbmc      string `json:"xmlbmc"`
	RowID       int    `json:"row_id"`
	Totalresult int    `json:"totalresult"`
}

package collytool

import (
	"cld/dao/resty_tool"
	"cld/models"
	"cld/settings"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82"

const informationUrl = "https://jxw.sylu.edu.cn/xsxxxggl"
const gradeUrl = "https://jxw.sylu.edu.cn/cjcx"
const gpaUrl = "https://jxw.sylu.edu.cn/xsxy"
const caleUrl = "https://jxw.sylu.edu.cn/xtgl"

type MyCollector struct {
	*colly.Collector
}

func NewMyCollector() *MyCollector {
	cfg := settings.Conf.Proxy

	collector := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	if cfg.Host != "" && cfg.Port != "" && cfg.Type != "" {

		pUrl := fmt.Sprintf("%s://%s:%s", cfg.Type, cfg.Host, cfg.Port)
		collector.SetProxy(pUrl)

	}

	return &MyCollector{collector}
}

func (c *MyCollector) GetInforamation(cookiestring string, username string) (studentInfo *models.SyluUser, err error) {
	studentInfo = new(models.SyluUser)

	queryParams := url.Values{}
	queryParams.Set("gnmkdm", "N100801")
	queryParams.Set("layout", "default")
	queryParams.Set("su", username)

	c.OnRequest(func(r *colly.Request) {
		r.URL.RawQuery = queryParams.Encode()
		r.Headers.Set("Cookie", cookiestring)
		r.Headers.Set("Connection", "close")
	})

	c.OnHTML("#content_xsxxgl_xsjbxx", func(e *colly.HTMLElement) {
		studentInfo.StudentID = e.ChildText("#col_xh p")
		studentInfo.ReUsername = e.ChildText("#col_xm p")
	})

	c.OnHTML("#content_xsxxgl_xsxjxx", func(e *colly.HTMLElement) {
		studentInfo.Grade = e.ChildText("#col_njdm_id p")
		studentInfo.College = e.ChildText("#col_jg_id p")
		studentInfo.Major = e.ChildText("#col_zyh_id p")
	})

	err = c.Visit(informationUrl + "/xsgrxxwh_cxXsgrxx.html")
	if err != nil {
		fmt.Println("visit err" + err.Error())
	}

	return
}

func (c *MyCollector) GetGradeDetail(bindInfo *models.ParamGradeDetaile) (resGradeDetail []*models.ResGradeDetail, err error) {
	resGradeDetail = make([]*models.ResGradeDetail, 0, 5)
	queryParams := url.Values{}
	queryParams.Add("gnmkdm", "N305005")

	c.OnRequest(func(r *colly.Request) {
		r.URL.RawQuery = queryParams.Encode()
		r.Headers.Add("Cookie", bindInfo.Cookie)
	})

	form := map[string]string{
		"jxb_id": bindInfo.ClassID,
		"xnm":    strconv.Itoa(bindInfo.Year),
		"xqm":    strconv.Itoa(bindInfo.Semester),
	}
	regex := regexp.MustCompile(`【 (.+?) 】`)

	c.OnHTML("table[id=subtab] tbody tr", func(e *colly.HTMLElement) {
		scoreItem := new(models.ResGradeDetail)
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			switch i {
			case 0:
				scoreItem.Name = regex.FindStringSubmatch(el.Text)[1]
			case 1:
				scoreItem.Weight = el.Text[:len(el.Text)-2]
			case 2:
				scoreItem.Score = el.Text[:len(el.Text)-2]
			}
		})
		resGradeDetail = append(resGradeDetail, scoreItem)
	})

	c.OnHTML("title", func(h *colly.HTMLElement) {
		err = resty_tool.ErrorLapse
	})

	errPost := c.Post(gradeUrl+"/cjcx_cxCjxqGjh.html", form)
	if errPost != nil {
		return nil, errPost
	}

	c.Wait()

	return
}

func (c *MyCollector) GetGpas(cookies string) (resGpa *models.ResGpa, err error) {
	queryParams := url.Values{}
	queryParams.Add("gnmkdm", "N105515")

	c.AllowURLRevisit = false
	c.Async = true
	c.MaxDepth = 1

	c.OnRequest(func(r *colly.Request) {
		r.URL.RawQuery = queryParams.Encode()
		r.Headers.Add("Cookie", cookies)
	})

	resGpa = new(models.ResGpa)

	c.OnHTML("#alertBox", func(e *colly.HTMLElement) {
		fonts := e.DOM.Find("font").Find("font")
		resGpa.AllGpa = fonts.Eq(0).Text()
		resGpa.DegreeGpa = fonts.Eq(1).Text()
		return
	})

	c.OnHTML("title", func(h *colly.HTMLElement) {
		err = resty_tool.ErrorLapse
	})

	errVisit := c.Visit(gpaUrl + "/xsxyqk_cxXsxyqkIndex.html")
	if errVisit != nil {
		return nil, errVisit
	}
	c.Wait()

	return
}

func (c *MyCollector) GetSchoolCalendar(cookiestring string) (res *models.ResSchoolCale, err error) {
	queryParams := url.Values{}
	queryParams.Add("gnmkdm", "index")
	queryParams.Add("localeKey", "zh_Cn")

	c.OnRequest(func(r *colly.Request) {
		r.URL.RawQuery = queryParams.Encode()
		r.Headers.Set("Cookie", cookiestring)
		r.Headers.Set("Connection", "close")
	})

	res = new(models.ResSchoolCale)

	c.OnHTML("#rcStr", func(h *colly.HTMLElement) {
		value := h.Attr("value")
		parts := strings.Split(value, "!two!")
		fmt.Println(parts)
		for _, v := range parts {
			if v == "" {
				break
			}
			fields := strings.Split(v, "!one!")
			caleDate := models.SchoolCale{
				ID:        fields[0],
				Name:      fields[1],
				StartTime: fields[2],
				EndTime:   fields[3],
			}

			res.SchoolCale = append(res.SchoolCale, caleDate)
		}
	})

	c.OnHTML("thead tr:first-child th[colspan='24']", func(e *colly.HTMLElement) {
		text := e.Text

		parts := strings.Split(text, "(")
		term := parts[0]

		dateParts := strings.Split(parts[1], "至")
		startDate := dateParts[0]
		endDate := strings.TrimRight(dateParts[1], ")")

		res.Title = term
		res.StartTime = startDate
		res.EndTime = endDate
	})

	c.Visit(caleUrl + "/index_cxAreaSix.html")

	return
}

package collytool

import (
	"cld/dao/resty_tool"
	"cld/models"
	"cld/settings"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

const informationUrl = "https://jxw.sylu.edu.cn/xsxxxggl"
const gradeUrl = "https://jxw.sylu.edu.cn/cjcx"

type MyCollector struct {
	*colly.Collector
}

func getProxyURL() (*url.URL, error) {
	if settings.Conf.Proxy.Host != "" && settings.Conf.Proxy.Port != "" {
		return url.Parse("http://" + settings.Conf.Proxy.Host + ":" + settings.Conf.Proxy.Port)
	}
	return nil, nil
}

func NewMyCollector() *MyCollector {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82"),
	)

	proxyURL, _ := getProxyURL()
	if proxyURL != nil {
		c.SetProxyFunc(func(r *http.Request) (*url.URL, error) {
			return proxyURL, nil
		})
	}

	return &MyCollector{c}
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
		"xnm":    bindInfo.Year,
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

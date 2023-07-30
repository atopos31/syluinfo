package collytool

import (
	"cld/models"
	"cld/settings"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gocolly/colly"
)

const informationUrl = "https://jxw.sylu.edu.cn/xsxxxggl"

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

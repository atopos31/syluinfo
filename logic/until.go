package logic

import (
	"cld/models"
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
)

// 获取当前学期
func GetNowSem() (semyear, sem int) {
	now := time.Now()
	month := now.Month()
	if month >= 2 && month <= 7 {
		sem = 12
	} else {
		sem = 3
	}

	semyear = now.Year()

	if sem == 12 {
		semyear--
	}
	return
}

// 获取学期列表
func getSemeList(id string) ([]*models.SemeSterList, error) {
	firstTwo := id[:2]
	indexYear := "20" + firstTwo

	num, err := strconv.Atoi(indexYear)
	if err != nil {
		zap.L().Error("getSemeList Error", zap.String("ID", id))
		return nil, err
	}

	semList := make([]*models.SemeSterList, 0, 8)
	var month = 3
	for i := 0; i < 8; i++ {
		var semeSter models.SemeSterList
		semeSter.Name = getYear(num) + " " + getSemester(month)
		semeSter.Year = num
		semeSter.Month = month
		semList = append(semList, &semeSter)
		if month == 3 {
			month = 12
		} else {
			month = 3
			num++
		}
	}

	return semList, nil

}

// 获取当前时间为第几个学期
func getIndesSeme(id string) (int, error) {
	now := time.Now()
	firstTwo := id[:2]

	num, err := strconv.Atoi(firstTwo)
	if err != nil {
		return 0, nil
	}

	year := now.Year()
	month := int(now.Month())

	yearFormated := year % 100
	monthFormated := month

	var semenum int
	if monthFormated >= 3 && monthFormated <= 8 {
		semenum = 0
	} else {
		semenum = 1
	}

	return ((yearFormated-num)*2 + semenum) - 1, nil
}

// 获取学年字符串
func getYear(year int) string {
	start := strconv.Itoa(year)
	end := strconv.Itoa(year + 1)

	return fmt.Sprintf("%s-%s学年", start, end)
}

// 获取学期字符串
func getSemester(mounth int) string {
	if mounth == 3 {
		return "第一学期"
	} else {
		return "第二学期"
	}
}

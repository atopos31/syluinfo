package logic

import (
	"cld/models"
	"fmt"
	"strconv"
	"time"
)

func getSemeList(id string) ([]*models.SemeSterList, error) {
	firstTwo := id[:2]
	indexYear := "20" + firstTwo

	num, err := strconv.Atoi(indexYear)
	if err != nil {
		return nil, err
	}
	var month = 3
	semList := make([]*models.SemeSterList, 0, 8)

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

func getYear(year int) string {
	start := strconv.Itoa(year)
	end := strconv.Itoa(year + 1)

	return fmt.Sprintf("%s-%s学年", start, end)
}

func getSemester(mounth int) string {
	if mounth == 3 {
		return "第一学期"
	} else {
		return "第二学期"
	}
}

package logic

import (
	"fmt"
	"strconv"
)

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

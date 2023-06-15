package window

import (
	"strconv"
	"time"
)

func getInvoiceDate() (string, string, string, string) {
	year, month, _ := time.Now().Date()

	lastDayOfMonth := daysIn(month, year)
	stringMonth := strconv.Itoa(int(month))

	mon := func() string {
		if month > 9 {
			return stringMonth
		} else {
			return "0" + stringMonth
		}
	}()

	return strconv.Itoa(lastDayOfMonth) + "-" + mon + "-" + strconv.Itoa(year), strconv.Itoa(int(year)), stringMonth, month.String()
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

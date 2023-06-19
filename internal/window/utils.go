package window

import (
	"reflect"
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

func iterateStructFields(
	value reflect.Value,
	callback func(fieldName string, fieldValue interface{}),
	sectionCallback ...func(fieldName string, fieldValue interface{}),
) {
	// Iterate through the fields of the struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)

		// Check if the field is a nested struct
		if fieldValue.Kind() == reflect.Struct {
			if len(sectionCallback) > 0 {
				sectionCallback[0](field.Name, fieldValue.Interface())
			}
			// Recursively iterate through the nested struct fields
			iterateStructFields(fieldValue, callback, sectionCallback...)
		} else {
			// Call the callback function with the field name and value
			callback(field.Name, fieldValue.Interface())
		}
	}
}

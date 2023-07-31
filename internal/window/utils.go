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

func createStructFromSlice(slice []string, structType reflect.Type) (reflect.Value, []string) {
	structValue := reflect.New(structType).Elem()
	localSlice := slice

	// For each field in the struct, check if it's a nested struct and recursively create it
	for i := 0; i < structValue.NumField() && i < len(slice); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i).Type
		if fieldType.Kind() == reflect.Struct {
			// Nested struct, recursively create it
			nestedStruct, curSlice := createStructFromSlice(localSlice, fieldType)
			localSlice = curSlice
			field.Set(nestedStruct)
		} else {
			// Convert the slice element to the appropriate type and set the value
			sliceValue := reflect.ValueOf(localSlice[0])
			convertedValue := convertStringToType(sliceValue, fieldType)
			field.Set(convertedValue)
			localSlice = localSlice[1:]
		}
	}

	return structValue, localSlice
}

func convertStringToType(value reflect.Value, targetType reflect.Type) reflect.Value {
	switch targetType.Kind() {
	case reflect.String:
		return value.Convert(targetType)
	case reflect.Int:
		// Handle int parsing from string
		intValue, _ := strconv.Atoi(value.String())
		return reflect.ValueOf(intValue).Convert(targetType)
	// Add more cases for other data types as needed
	default:
		return reflect.Zero(targetType)
	}
}

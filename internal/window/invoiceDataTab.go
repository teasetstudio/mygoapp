package window

import (
	"fmt"
	"mygoapp/internal/config"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getInvoiceDataContainer() *fyne.Container {
	invoiceDataContainer := invoiceDataContainer()

	grid := container.NewVBox(invoiceDataContainer)
	return grid
}

func invoiceDataContainer() *fyne.Container {
	form := container.New(layout.NewFormLayout())

	// Get the value of the struct
	invoiceData := reflect.ValueOf(config.InvoiceData)
	var fieldNames []string

	// Define a callback function to handle each field
	itemCallback := func(fieldName string, fieldValue interface{}) {
		fieldNames = append(fieldNames, fieldName)
		label := widget.NewLabel(fieldName)
		input := widget.NewEntry()
		fieldStringValue := fmt.Sprintf("%v", fieldValue)
		input.SetText(fieldStringValue)
		form.Add(label)
		form.Add(input)
	}
	sectionCallback := func(fieldName string, fieldValue interface{}) {

		headingLabel := widget.NewLabelWithStyle(fieldName, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
		headingLabel.TextStyle = fyne.TextStyle{Bold: true}

		dummyLabel := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})

		form.Add(headingLabel)
		form.Add(dummyLabel)
	}

	// Iterate through the fields of the struct, invoking the callback function
	iterateStructFields(invoiceData, itemCallback, sectionCallback)

	// Save button
	saveButton := widget.NewButton("Save", func() {
		data := []string{}

		children := form.Objects
		for _, child := range children {
			// Check if the child is a widget.Entry (text input) or widget.Check (checkbox)
			switch c := child.(type) {
			// case *widget.Label:
			// 	fmt.Println("Label:", c.Text)
			case *widget.Entry:
				data = append(data, c.Text)
			}
		}

		// Create a struct from the slice of strings
		invoiceData, _ := createStructFromSlice(data, reflect.TypeOf(config.InvoiceDataType{}))
		// Convert the reflect.Value back to InvoiceDataType
		if invoice, ok := invoiceData.Interface().(config.InvoiceDataType); ok {
			config.SetYAMLInvoiceDataFile(invoice)
			config.InvoiceData = invoice
		} else {
			fmt.Println("Failed to convert the reflect.Value to InvoiceDataType.")
		}
	})

	// Title label for the section
	titleLabel := widget.NewLabel("Input Section")
	titleLabel.Alignment = fyne.TextAlignCenter

	// Layout the elements
	content := container.NewVBox(
		// titleLabel,
		form,
		saveButton,
	)

	container := container.NewVBox(content)

	return container
}

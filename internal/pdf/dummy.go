package pdf

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func GetDummyPDF() {
	pdf := gofpdf.New("P", "mm", "A4", "") // Create a new PDF instance
	pdf.AddPage()                          // Add a page

	// Load the DejaVuSans font file
	pdf.AddFont("DejaVuSans", "", "DejaVuSans.ttf")

	// Set the DejaVuSans font as the default font
	pdf.SetFont("DejaVuSans", "", 12)

	// Set font and styling for invoice
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "INVOICE")
	pdf.Ln(20) // Add a new line

	// Add seller and buyer information on one row
	pdf.SetXY(10, 30)
	pdf.MultiCell(80, 10, "Seller:\n\nCompany Name\nAddress Line 1\nAddress Line 2\nCity, State, Zip Code", "", "L", false)
	pdf.SetXY(120, 30)
	pdf.MultiCell(80, 10, "Buyer:\n\nCompany Name\nAddress Line 1\nAddress Line 2\nCity, State, Zip Code", "", "R", false)
	pdf.Ln(20) // Add a new line

	// Add issue date and sale date
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(80, 10, "Data wystawienia: ", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(80, 10, time.Now().Format("02-01-2006"), "", 0, "L", false, 0, "")
	pdf.Ln(5) // Add a new line

	// Add issue date and sale date
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(80, 10, "Data sprzedaży: ", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(80, 10, time.Now().Format("02-01-2006"), "", 0, "L", false, 0, "")
	pdf.Ln(20) // Add a new line

	// Add invoice table headers
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Item")
	pdf.Cell(40, 10, "Quantity")
	pdf.Cell(40, 10, "Price")
	pdf.Cell(40, 10, "Total")
	pdf.Ln(10) // Add a new line

	// Add invoice table rows
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Item 1")
	pdf.Cell(40, 10, "2")
	pdf.Cell(40, 10, "$10")
	pdf.Cell(40, 10, "$20")
	pdf.Ln(10) // Add a new line

	pdf.Cell(40, 10, "Item 2")
	pdf.Cell(40, 10, "3")
	pdf.Cell(40, 10, "$15")
	pdf.Cell(40, 10, "$45")
	pdf.Ln(10) // Add a new line

	// Calculate VAT and total
	subtotal := 65.0
	vatRate := 0.2
	vat := subtotal * vatRate
	total := subtotal + vat

	// Add VAT and total to the invoice
	pdf.Cell(40, 10, "")
	pdf.Cell(40, 10, "")
	pdf.Cell(40, 10, "VAT (20%):")
	pdf.Cell(40, 10, fmt.Sprintf("$%.2f", vat))
	pdf.Ln(10) // Add a new line

	pdf.Cell(40, 10, "")
	pdf.Cell(40, 10, "")
	pdf.Cell(40, 10, "Total (incl. VAT):")
	pdf.Cell(40, 10, fmt.Sprintf("$%.2f", total))
	pdf.Ln(10) // Add a new line

	err := pdf.OutputFileAndClose("invoice.pdf") // Save the PDF file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Invoice created successfully.")
}

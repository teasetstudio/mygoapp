package pdf_invoice

import (
	"fmt"
	"mygoapp/internal/config"

	"github.com/jung-kurt/gofpdf"
)

func GetPdfInvoice(invoiceDate string, invoiceConfig *config.InvoiceDataType) *gofpdf.Fpdf {
	fontsPath := getFontFolder()

	pdf := gofpdf.New("P", "mm", "A4", fontsPath)
	pdf.AddPage()
	pdf.AddUTF8Font("DejaVu", "", "DejaVuSansCondensed.ttf")
	pdf.AddUTF8Font("DejaVu", "B", "DejaVuSansCondensed-Bold.ttf")

	var products []config.TowarType
	products = append(products, invoiceConfig.Towar)
	if invoiceDate == "" {
		invoiceDate = getCurrentDate()
	}

	invoiceNr, err := getLastDayOfMonth(invoiceDate)
	if err != nil {
		fmt.Printf("'%s' does not match the format 'DD-MM-YYYY'. Error: %v\n", invoiceDate, err)
	}

	dataWystawienia := invoiceNr
	dataSprzedazy := invoiceNr
	tableContent, totalPrice, tableSummaryContent := getTableContent(products)
	totalPriceStr, totalPriceInPolish := getFooterContent(totalPrice)
	konta := config.Sprzedawca.Konta

	setHeader(pdf, invoiceNr)
	setFirms(pdf, config.Sprzedawca, config.Nabywca)
	setDates(pdf, dataWystawienia, dataSprzedazy)
	setInvoiceTable(pdf, tableContent, tableSummaryContent)
	setFooter(pdf, totalPriceStr, totalPriceInPolish, konta)

	return pdf
}

package pdf_invoice

import (
	"mygoapp/internal/config"

	"github.com/jung-kurt/gofpdf"
)

func GetPdfInvoice(invoiceDate string, invoiceConfig *config.InvoiceDataType) *gofpdf.Fpdf {
	isFontsInstalled()

	pdf := gofpdf.New("P", "mm", "A4", config.FontsDir)
	pdf.AddPage()
	pdf.AddUTF8Font(primaryFontFamily, "", fontfileNames[0])
	pdf.AddUTF8Font(primaryFontFamily, "B", fontfileNames[1])

	var products []config.TowarType
	products = append(products, invoiceConfig.Towar)
	if invoiceDate == "" {
		invoiceDate = getCurrentDate()
	}

	invoiceNr, err := getLastDayOfMonth(invoiceDate)
	if err != nil {
		panic("does not match the format")
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

package pdf

import (
	"log"

	"mygoapp/internal/config"
	pdf_invoice "mygoapp/internal/pdf/invoice"
)

func GetPdfInvoice(invoiceDate string, downloadFilePath string, invoiceConfig *config.InvoiceDataType) {
	invoice := pdf_invoice.GetPdfInvoice(invoiceDate, invoiceConfig)

	if downloadFilePath == "" {
		downloadFilePath = "invoice.pdf"
	}
	err := invoice.OutputFileAndClose(downloadFilePath)
	if err != nil {
		log.Fatal(err)
	}
}

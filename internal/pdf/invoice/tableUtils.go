package pdf_invoice

import (
	"mygoapp/internal/config"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

var headerRowHeight = 3.2
var headerCellYOffset = 4.0
var rowHeight = 3.6
var cellYOffset = 4.0

var header = []string{"Lp.", "Nazwa towaru lub usługi", "Jm.", "Ilość", "Cena netto", "Wartość netto", "Stawka VAT", "Kwota VAT", "Wartość brutto"}
var colWidths = []float64{8, 68, 13, 13, 18, 18, 16, 18, 18}

func calcLineHeight(linesNum int, maxLinesNum int, rowHeight float64) float64 {
	return float64(maxLinesNum/linesNum) * rowHeight
}

func getMaxCellHeight(row []string, colWidths []float64, pdf *gofpdf.Fpdf) (float64, int) {
	maxHeight := 0.0
	maxLines := 1
	for i, col := range row {
		lines := getLinesNum(col, colWidths[i], pdf)
		cellHeight := float64(lines) * rowHeight
		if cellHeight > maxHeight {
			maxHeight = cellHeight
			maxLines = lines
		}
	}
	return maxHeight, maxLines
}

func getLinesNum(col string, width float64, pdf *gofpdf.Fpdf) int {
	return len(pdf.SplitLines([]byte(col), width))
}

func getTableContent(products []config.TowarType) (tableContent [][]string, totalPrice float64, tableSummaryContent [][]string) {
	totalPriceNoVat := 0.0
	totalVatPrice := 0.0
	for i, product := range products {
		productPrice := product.Cena * float64(product.Amount)
		vatAmount := productPrice * (float64(product.Vat) / 100)
		productTotalPrice := productPrice + vatAmount
		index := strconv.Itoa(i + 1)
		amount := strconv.Itoa(product.Amount)
		tableContent = append(tableContent, []string{
			index, product.Nazwa, "szt.", amount,
			formatSumToString(product.Cena),
			formatSumToString(productPrice),
			strconv.Itoa(product.Vat) + "%",
			formatSumToString(vatAmount),
			formatSumToString(productTotalPrice),
		})

		totalPriceNoVat += productPrice
		totalVatPrice += vatAmount
		totalPrice += productTotalPrice
	}

	tableSummaryContent = getTableSummaryContent(totalPriceNoVat, totalVatPrice, totalPrice)

	return
}

func getTableSummaryContent(totalPriceNoVat float64, totalVatPrice float64, totalPrice float64) [][]string {
	return [][]string{
		{"W tym", formatSumToString(totalPriceNoVat), "23%", formatSumToString(totalVatPrice), formatSumToString(totalPrice)},
		{"Razem", formatSumToString(totalPriceNoVat), " ", formatSumToString(totalVatPrice), formatSumToString(totalPrice)},
		// ... Add more rows as needed
	}
}

func getCellOffset(cellNum int, colWidths []float64) (float64, []float64) {
	if cellNum > len(colWidths) {
		return 0, colWidths
	}
	leftColWidths := colWidths
	cellOffset := 0.0

	for i := 0; i < cellNum+1; i++ {
		cellOffset += leftColWidths[0]
		leftColWidths = leftColWidths[1:]
	}
	return cellOffset, leftColWidths
}

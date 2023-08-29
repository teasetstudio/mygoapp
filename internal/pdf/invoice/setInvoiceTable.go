package pdf_invoice

import (
	"github.com/jung-kurt/gofpdf"
)

func setInvoiceTable(pdf *gofpdf.Fpdf, content [][]string, summaryContent [][]string) {
	// HEADER ROW
	pdf.SetFont("DejaVu", "", 9)
	pdf.SetTextColor(white())
	x, y := pdf.GetXY()
	y += 15
	pdf.SetFillColor(primary())
	pdf.SetDrawColor(primary())
	maxRowHeight, maxLinesNum := getMaxCellHeight(header, colWidths, pdf)
	pdf.LinearGradient(x, y, 190, maxRowHeight+headerCellYOffset, 12, 36, 61, 70, 92, 121, 0.5, 0.5, 1, 1)
	for i, h := range header {
		linesNum := getLinesNum(h, colWidths[i], pdf)
		lineHeight := calcLineHeight(linesNum, maxLinesNum, headerRowHeight)
		pdf.SetXY(x, y+(headerCellYOffset/2)+0.5)
		pdf.MultiCell(colWidths[i], lineHeight, h, "", "C", false)
		x += colWidths[i]
	}
	pdf.SetY(y + maxRowHeight + headerCellYOffset)

	// TABLE CONTENT
	pdf.SetTextColor(darkgray())
	for _, row := range content {
		maxRowHeight, maxLinesNum := getMaxCellHeight(row, colWidths, pdf)
		x, y = pdf.GetXY()
		for i, col := range row {
			linesNum := getLinesNum(col, colWidths[i], pdf)
			lineHeight := calcLineHeight(linesNum, maxLinesNum, rowHeight)
			// pdf.SetXY(x, y)
			pdf.Rect(x, y, colWidths[i], maxRowHeight+cellYOffset, "D")
			pdf.SetXY(x, y+(cellYOffset/2))
			var align string
			switch i {
			case 1:
				align = "L"
			default:
				align = "C"
			}
			pdf.MultiCell(colWidths[i], lineHeight, col, "", align, false)
			x += colWidths[i]
		}
		pdf.SetY(y + maxRowHeight + cellYOffset)
	}

	setInvoiceTableSummary(pdf, summaryContent)
}

func setInvoiceTableSummary(pdf *gofpdf.Fpdf, content [][]string) {
	pdf.SetFont("DejaVu", "", 9)
	pdf.SetTextColor(darkgray())
	cellOffset, leftColWidths := getCellOffset(3, colWidths)
	for _, row := range content {
		maxRowHeight, maxLinesNum := getMaxCellHeight(row, leftColWidths, pdf)
		x, y := pdf.GetXY()
		x += cellOffset
		pdf.SetXY(x, y)
		for i, col := range row {
			linesNum := getLinesNum(col, leftColWidths[i], pdf)
			lineHeight := calcLineHeight(linesNum, maxLinesNum, rowHeight)
			textAlign := "C"
			if i != 0 {
				pdf.Rect(x, y, leftColWidths[i], maxRowHeight+cellYOffset, "D")
				pdf.SetFontStyle("")
			} else {
				pdf.SetFontStyle("B")
				textAlign = "R"
			}
			pdf.SetXY(x, y+(cellYOffset/2))
			pdf.MultiCell(leftColWidths[i], lineHeight, col, "", textAlign, false)
			x += leftColWidths[i]
		}
		pdf.SetY(y + maxRowHeight + cellYOffset)
	}
}

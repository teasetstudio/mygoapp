package pdf_invoice

import "github.com/jung-kurt/gofpdf"

func setFooter(pdf *gofpdf.Fpdf, sum string, sumText string, konta string) {
	pdf.SetFont(primaryFontFamily, "", 16)
	pdf.SetTextColor(darkgray())
	text := "Do zapłaty " + sum + " PLN"
	initX, initY := pdf.GetXY()
	initY += 10

	x := initX
	y := initY
	colIndexOffset := 4
	xOffset := x
	for i := 0; i < len(colWidths) && i < colIndexOffset; i++ {
		xOffset += colWidths[i]
	}
	width := 0.0
	startIndex := len(colWidths) - colIndexOffset - 1
	for i := startIndex; i < len(colWidths); i++ {
		width += colWidths[i]
	}
	y += 6
	x = xOffset
	pdf.SetFillColor(primary())
	pdf.Rect(x, y, width, 10, "F")
	y += 7
	x += 11
	pdf.SetTextColor(white())
	pdf.Text(x, y, text)
	y += 8
	x -= 6
	pdf.SetFont(primaryFontFamily, "", 8)
	pdf.SetTextColor(primary())
	pdf.Text(x, y, "Słownie")

	y -= 3.5
	x += 13
	pdf.SetXY(x, y)
	pdf.MultiCell(width-18, 5, sumText, "", "L", false)

	// LEFT SIDE
	x = initX
	y = initY + 13
	pdf.Text(x, y, "Sposób płatności")
	x += 25
	pdf.SetFontStyle("B")
	pdf.Text(x, y, "przelew")
	x -= 25
	y += 2.7
	pdf.SetFontStyle("")
	pdf.Rect(x, y, 80, 0.3, "F")
	y += 5
	pdf.Text(x, y, "Numer konta")
	x += 25
	pdf.SetFontStyle("B")
	pdf.Text(x, y, konta)

	// Pages
	pdf.SetFont(primaryFontFamily, "", 10)
	pdf.SetTextColor(gray())
	pdf.Text(195, 292, "1 / 1")
}

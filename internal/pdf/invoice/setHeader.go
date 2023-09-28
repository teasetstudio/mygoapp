package pdf_invoice

import "github.com/jung-kurt/gofpdf"

func setHeader(pdf *gofpdf.Fpdf, invoiceNr string) {
	drawBackground(pdf)
	pdf.SetFont(primaryFontFamily, "B", 30)
	pdf.SetTextColor(white())

	initX, initY := pdf.GetXY()
	y := initY + 5
	x := initX

	pdf.Text(x, y, "FAKTURA")

	pdf.SetFont(primaryFontFamily, "B", 14)
	pdf.SetTextColor(primary())
	numberText := "Nr: " + invoiceNr
	x = getDefaultX_RightAlingedText(pdf, numberText)
	y += 20
	pdf.Text(x, y, numberText)
	y += 6
	pdf.SetY(y)
}

func drawBackground(pdf *gofpdf.Fpdf) {
	x, y := pdf.GetXY()

	// PRIMARY
	pdf.SetFillColor(primary())
	width := 210.0
	height := 297.0
	pdf.Rect(0, 0, width, height, "F")

	// SECONDARY
	r1 := 580.0
	pdf.SetFillColor(secondary())
	pdf.Circle(x+365, y+490, r1, "F")

	// GRAY
	r2 := 530.0
	pdf.SetFillColor(lightgray())
	pdf.Circle(x+260, y+515, r2, "F")

	// WHITE
	r3 := 450.0
	pdf.SetFillColor(white())
	pdf.Circle(x+215, y+450, r3, "F")

}

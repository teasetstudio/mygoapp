package pdf_invoice

import "github.com/jung-kurt/gofpdf"

func setDates(pdf *gofpdf.Fpdf, issueDate string, saleDate string) {
	pdf.SetFont("DejaVu", "", 11)
	pdf.SetTextColor(darkgray())

	x, y := pdf.GetXY()
	y += 5
	pdf.Text(x, y, "Data wystawienia:")
	pdf.Text(x+34, y, issueDate)

	y += 5
	pdf.Text(x, y, "Data sprzedaży:")
	pdf.Text(x+34, y, issueDate)
	y += 5
	pdf.SetY(y)
}

package pdf_invoice

import (
	"mygoapp/internal/config"

	"github.com/jung-kurt/gofpdf"
)

func setFirms(pdf *gofpdf.Fpdf, sprzedawca config.Firma, nabywca config.Firma) {
	y := pdf.GetY()
	pdf.SetY(y + 30)

	yS := setSprzedawca(pdf, sprzedawca)
	yN := setNabywca(pdf, nabywca)
	y = yN
	if yS > yN {
		y = yS
	}
	y += 5
	pdf.SetY(y)
}

func setSprzedawca(pdf *gofpdf.Fpdf, firma config.Firma) float64 {
	title := "Sprzedawca:"
	x, y := pdf.GetXY()

	pdf.SetFont("DejaVu", "B", 12)
	pdf.SetTextColor(primary())
	pdf.Text(x, y, title)

	y += 2
	pdf.SetFont("DejaVu", "", 11)
	pdf.SetTextColor(darkgray())
	y = drawMultiLineText(pdf, firma.Nazwa, 35, 5, x, y, false)

	y += 5
	pdf.Text(x, y, "NIP: "+firma.Nip)

	y += 5
	address := firma.Ulica + " " + firma.Nr_budynku
	if firma.Lokalu != "" {
		address += " / " + firma.Lokalu
	}
	pdf.Text(x, y, address)

	y += 5
	kodAndCity := firma.Kod + " " + firma.Miasto
	pdf.Text(x, y, kodAndCity)

	return y
}

func setNabywca(pdf *gofpdf.Fpdf, firma config.Firma) float64 {
	pageWidth, _ := pdf.GetPageSize()
	marginRight := 10.0
	offsetX := pageWidth - marginRight
	title := "Nabywca:"
	_, y := pdf.GetXY()
	x := getX_RightAlingedText(pdf, title, offsetX-3)
	pdf.SetFont("DejaVu", "B", 12)
	pdf.SetTextColor(primary())
	pdf.Text(x, y, title)

	y += 2
	pdf.SetFont("DejaVu", "", 11)
	pdf.SetTextColor(darkgray())
	y = drawMultiLineText(pdf, firma.Nazwa, 35, 5, offsetX, y, true)

	y += 5
	nipText := "NIP: " + firma.Nip
	x = getX_RightAlingedText(pdf, nipText, offsetX)
	pdf.Text(x, y, nipText)

	y += 5
	address := firma.Ulica + " " + firma.Nr_budynku
	if firma.Lokalu != "" {
		address += " / " + firma.Lokalu
	}
	x = getX_RightAlingedText(pdf, address, offsetX)
	pdf.Text(x, y, address)

	y += 5
	kodAndCity := firma.Kod + " " + firma.Miasto
	x = getX_RightAlingedText(pdf, kodAndCity, offsetX)
	pdf.Text(x, y, kodAndCity)

	return y
}

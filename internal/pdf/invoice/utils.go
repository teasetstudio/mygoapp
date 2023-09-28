package pdf_invoice

import (
	"fmt"
	"mygoapp/internal/config"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

var primaryFontFamily = "DejaVu"
var fontfileNames = []string{"DejaVuSansCondensed.ttf", "DejaVuSansCondensed-Bold.ttf"}

func darkgray() (int, int, int) {
	return 70, 70, 70
}

func gray() (int, int, int) {
	return 170, 170, 170
}

func lightgray() (int, int, int) {
	return 225, 225, 225
}

func primary() (int, int, int) {
	return 12, 36, 61
}

func secondary() (int, int, int) {
	return 22, 46, 71
}

func white() (int, int, int) {
	return 250, 250, 250
}

func drawMultiLineText(pdf *gofpdf.Fpdf, text string, maxLenght int, offset float64, x float64, y float64, isAlignRight bool) float64 {
	lines := splitString(text, maxLenght)
	for _, line := range lines {
		y += offset
		xx := x
		if isAlignRight {
			xx = getX_RightAlingedText(pdf, line, x)
		}
		pdf.Text(xx, y, line)
	}
	return y
}

func isFontsInstalled() {
	missingFonts := []string{}
	for _, fontName := range fontfileNames {
		fontPath := filepath.Join(config.FontsDir, fontName)
		if !isFileExists(fontPath) {
			missingFonts = append(missingFonts, fontName)
		}
	}
	if len(missingFonts) > 0 {
		message := fmt.Sprintf("The following fonts are missing: %s", strings.Join(missingFonts, ", "))
		panic(message)
	}
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getX_RightAlingedText(pdf *gofpdf.Fpdf, text string, initX float64) float64 {
	stringWidth := pdf.GetStringWidth(text)
	return initX - stringWidth
}

func getDefaultX_RightAlingedText(pdf *gofpdf.Fpdf, text string) float64 {
	pageWidth, _ := pdf.GetPageSize()
	marginRight := 10.0
	offsetX := pageWidth - marginRight
	return getX_RightAlingedText(pdf, text, offsetX)
}

func splitString(input string, maxLen int) []string {
	var result []string

	words := strings.Fields(input)
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word) <= maxLen {
			currentLine += " " + word
		} else {
			result = append(result, strings.TrimSpace(currentLine))
			currentLine = word
		}
	}

	if currentLine != "" {
		result = append(result, strings.TrimSpace(currentLine))
	}

	// Handle words that are longer than maxLen
	for i, line := range result {
		if len(line) > maxLen {
			firstPart := line[:maxLen]
			remainingPart := line[maxLen:]
			result[i] = firstPart
			result = append(result[:i+1], append([]string{remainingPart}, result[i+1:]...)...)
		}
	}

	return result
}

func getFooterContent(sum float64) (string, string) {
	beforeDot, afterDot := splitFloat(sum)
	sumText := numberToPolish(beforeDot) + " " + afterDot + "/100 PLN"
	return formatSumToString(sum), sumText
}

func formatSumToString(n float64) string {
	// Convert the float64 to string with 2 decimal places
	str := fmt.Sprintf("%.2f", n)

	// Split the string into whole and fractional parts
	parts := strings.Split(str, ".")

	whole := parts[0]
	formattedWhole := ""

	// Add spaces every third digit from the right
	counter := 0
	for i := len(whole) - 1; i >= 0; i-- {
		counter++
		formattedWhole = string(whole[i]) + formattedWhole
		if counter%3 == 0 && i != 0 {
			formattedWhole = " " + formattedWhole
		}
	}

	// Combine the whole and fractional parts, replacing "." with ","
	return formattedWhole + "," + parts[1]
}

func numberToPolish(num int) string {
	if num < 0 || num > 99999 {
		return "Number out of range"
	}

	if num == 0 {
		return "zero"
	}

	units := []string{"", "jeden", "dwa", "trzy", "cztery", "pięć", "sześć", "siedem", "osiem", "dziewięć"}
	teens := []string{"dziesięć", "jedenaście", "dwanaście", "trzynaście", "czternaście", "piętnaście", "szesnaście", "siedemnaście", "osiemnaście", "dziewiętnaście"}
	tens := []string{"", "", "dwadzieścia", "trzydzieści", "czterdzieści", "pięćdziesiąt", "sześćdziesiąt", "siedemdziesiąt", "osiemdziesiąt", "dziewięćdziesiąt"}
	thousands := []string{"", "tysiąc", "tysiące", "tysiące", "tysiące", "tysięcy", "tysięcy", "tysięcy", "tysięcy", "tysięcy"}

	result := ""

	// Handle thousands
	thous := num / 1000
	if thous > 0 {
		if thous >= 10 && thous <= 19 {
			result += teens[thous-10] + " tysięcy "
		} else {
			if thous/10 > 0 {
				result += tens[thous/10] + " "
				if thous%10 == 0 {
					result += "tysięcy "
				}
			}
			if thous%10 > 0 {
				result += units[thous%10] + " "
				result += thousands[thous%10] + " "
			}
		}
	}

	// Remaining number after processing thousands
	num %= 1000

	// Handle hundreds
	// Handle hundreds
	hunds := num / 100
	if hunds > 0 {
		switch hunds {
		case 1:
			result += "sto "
		case 2:
			result += "dwieście "
		case 3:
			result += "trzysta "
		case 4:
			result += "czterysta "
		case 5:
			result += "pięćset "
		case 6:
			result += "sześćset "
		case 7:
			result += "siedemset "
		case 8:
			result += "osiemset "
		case 9:
			result += "dziewięćset "
		}
	}

	// Remaining number after processing hundreds
	num %= 100

	// Handle tens and units
	if num >= 10 && num <= 19 {
		result += teens[num-10]
	} else {
		tensPlace := num / 10
		if tensPlace > 0 {
			result += tens[tensPlace] + " "
		}
		if num%10 > 0 {
			result += units[num%10]
		}
	}

	return strings.TrimSpace(result)
}

func splitFloat(num float64) (beforeDot int, afterDot string) {
	strNum := strconv.FormatFloat(num, 'f', 2, 64)
	parts := strings.Split(strNum, ".")
	beforeDot, _ = strconv.Atoi(parts[0])
	if len(parts) == 2 {
		afterDot = parts[1]
	}
	return
}
func getLastDayOfMonth(dateStr string) (string, error) {
	// Parse the given date string
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return "", err
	}

	// Add one month to the current month and get the zeroth day
	// (this will give us the last day of the current month)
	lastDay := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location())
	return lastDay.Format("02-01-2006"), nil
}

func getCurrentDate() string {
	currentDate := time.Now()
	return currentDate.Format("02-01-2006")
}

package scrap

import (
	"log"
	"mygoapp/internal/config"

	"github.com/mxschmitt/playwright-go"
)

func fakturowoLogin(page playwright.Page, invoiceConfig *config.InvoiceConfigType) {
	_, err := page.Goto("https://www.fakturowo.pl/logowanie", playwright.PageGotoOptions{Timeout: playwright.Float(100000)})
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	email, err := page.QuerySelector("#email")
	if err != nil {
		log.Fatalf("could not get email: %v", err)
	}
	email.Fill(invoiceConfig.User.Email)

	pass, err := page.QuerySelector("#password")
	if err != nil {
		log.Fatalf("could not get password: %v", err)
	}
	pass.Fill(invoiceConfig.User.Pass)

	submitBtn, err := page.QuerySelector("#form button[type=submit]")
	if err != nil {
		log.Fatalf("could not get submit btn: %v", err)
	}
	submitBtn.Click()
}

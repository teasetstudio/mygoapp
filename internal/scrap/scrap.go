package scrap

import (
	"log"
	"time"

	"mygoapp/internal/config"

	"github.com/mxschmitt/playwright-go"
)

func Start(invoiceDate string) string {
	pw, err := playwright.Run(&playwright.RunOptions{Browsers: []string{"chromium"}})
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage(playwright.BrowserNewContextOptions{AcceptDownloads: playwright.Bool(true)})
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.fakturowo.pl/logowanie", playwright.PageGotoOptions{Timeout: playwright.Float(100000)}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	email, err := page.QuerySelector("#email")
	if err != nil {
		log.Fatalf("could not get email: %v", err)
	}
	email.Fill(config.User.Email)

	pass, err := page.QuerySelector("#password")
	if err != nil {
		log.Fatalf("could not get password: %v", err)
	}
	pass.Fill(config.User.Pass)

	submitBtn, err := page.QuerySelector("#form button[type=submit]")
	if err != nil {
		log.Fatalf("could not get submit btn: %v", err)
	}
	submitBtn.Click()

	if _, err = page.Goto("https://www.fakturowo.pl/wystaw", playwright.PageGotoOptions{Timeout: playwright.Float(100000)}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// p:contains("example")
	// nabywcaInput, err := page.QuerySelector("#nazwa_nabywca")
	// if err != nil {
	// 	log.Fatalf("could not get Nazwa: %v", err)
	// }
	// nabywcaInput.Fill(config.Nabywca.Nazwa)
	page.Fill("#nazwa_nabywca", config.Nabywca.Nazwa)
	page.Fill("#nip_nabywca", config.Nabywca.Nip)
	page.Fill("#ulica_nabywca", config.Nabywca.Ulica)
	page.Fill("#budynek_nabywca", config.Nabywca.Nr_budynku)
	page.Fill("#miasto_nabywca", config.Nabywca.Miasto)
	page.Fill("#kod_nabywca", config.Nabywca.Kod)
	page.Fill("#nazwa_0", config.Towar.Nazwa)
	page.Fill("#cena_netto_0", config.Towar.Cena)

	podpisCheckbox1, err := page.QuerySelector("input[name=\"sprzedawca[pokaz_podpis]\"]")
	if err != nil {
		log.Fatalf("could not get podpisCheckbox1: %v", err)
	}
	if is, _ := podpisCheckbox1.IsChecked(); is {
		podpisCheckbox1.Click()
	}

	podpisCheckbox2, _ := page.QuerySelector("input[name=\"nabywca[pokaz_podpis]\"]")
	if is, _ := podpisCheckbox2.IsChecked(); is {
		podpisCheckbox2.Click()
	}

	terminCheckbox, _ := page.QuerySelector("input[name=\"pokaz_termin\"]")
	if is, _ := terminCheckbox.IsChecked(); is {
		terminCheckbox.Click()
	}

	uwagiCheckbox, _ := page.QuerySelector("input[name=\"pokaz_uwagi\"]")
	if is, _ := uwagiCheckbox.IsChecked(); is {
		uwagiCheckbox.Click()
	}

	pokazCheckbox, _ := page.QuerySelector("input[name=pokaz_miejsce]")
	if is, _ := pokazCheckbox.IsChecked(); is {
		pokazCheckbox.Click()
	}

	page.Fill("input[name=numer]", invoiceDate)
	page.Click("#date1")
	for i := 0; i < 10; i++ {
		page.Keyboard().Press("Backspace")
	}
	for _, letter := range invoiceDate {
		page.Keyboard().Press(string(letter))
	}
	page.Fill("#nodate1", invoiceDate)
	page.Click("#pobierz_i_zapisz")

	downloadPath := config.DownloadDir + "\\" + invoiceDate + "-faktura.pdf"

	page.On("download", func(download playwright.Download) {
		download.SaveAs(downloadPath)
		download.Delete()
	})
	time.Sleep(time.Second * 2)
	pw.Stop()
	return downloadPath
}

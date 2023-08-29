package config

type Firma struct {
	Nazwa      string `yaml:"nazwa"`
	Nip        string `yaml:"nip"`
	Ulica      string `yaml:"ulica"`
	Nr_budynku string `yaml:"nr_budynku"`
	Lokalu     string `yaml:"lokalu"`
	Miasto     string `yaml:"miasto"`
	Kod        string `yaml:"kod"`
	Konta      string `yaml:"konta"`
}

type TowarType struct {
	Nazwa  string  `yaml:"nazwa"`
	Cena   float64 `yaml:"cena"`
	Amount int     `yaml:"amount"`
	Vat    int     `yaml:"vat"`
}

type UserType struct {
	Email          string `yaml:"email"`
	Pass           string `yaml:"pass"`
	SenderEmail    string `yaml:"sender_email"`
	SenderPassword string `yaml:"sender_password"`
	RecipientEmail string `yaml:"recipient_email"`
}

type InvoiceDataType struct {
	Sprzedawca Firma     `yaml:"sprzedawca"`
	Nabywca    Firma     `yaml:"nabywca"`
	Towar      TowarType `yaml:"towar"`
	User       UserType  `yaml:"user"`
}

var DefaultNabywca = Firma{
	Nazwa:      "GODEL TECHNOLOGIES EUROPE SPÓŁKA ZOGRANICZONĄ ODPOWIEDZIALNOŚCIĄ",
	Nip:        "7252307374",
	Ulica:      "Ogrodowa",
	Nr_budynku: "8",
	Miasto:     "Łódź",
	Kod:        "91-062",
}

var Nabywca = Firma{
	Nazwa:      DefaultNabywca.Nazwa,
	Nip:        DefaultNabywca.Nip,
	Ulica:      DefaultNabywca.Ulica,
	Nr_budynku: DefaultNabywca.Nr_budynku,
	Miasto:     DefaultNabywca.Miasto,
	Kod:        DefaultNabywca.Kod,
}

var DefaultSprzedawca = Firma{
	Nazwa:      "Ivan Tichkevitch",
	Nip:        "5833464281",
	Ulica:      "Jana Heweliusza",
	Nr_budynku: "11",
	Lokalu:     "819",
	Miasto:     "Gdańsk",
	Kod:        "80-890",
	Konta:      "40 1020 1811 0000 0202 0416 4158",
}

var Sprzedawca = Firma{
	Nazwa:      DefaultSprzedawca.Nazwa,
	Nip:        DefaultSprzedawca.Nip,
	Ulica:      DefaultSprzedawca.Ulica,
	Nr_budynku: DefaultSprzedawca.Nr_budynku,
	Lokalu:     DefaultSprzedawca.Lokalu,
	Miasto:     DefaultSprzedawca.Miasto,
	Kod:        DefaultSprzedawca.Kod,
	Konta:      DefaultSprzedawca.Konta,
}

var DefaultTowar = TowarType{
	Nazwa:  "Software development services",
	Cena:   19000.00,
	Amount: 1,
	Vat:    23,
}

var Towar = TowarType{
	Nazwa:  DefaultTowar.Nazwa,
	Cena:   DefaultTowar.Cena,
	Amount: DefaultTowar.Amount,
	Vat:    DefaultTowar.Vat,
}

var DefaultUser = UserType{
	Email:          "spam9000me@gmail.com",
	Pass:           "",
	SenderEmail:    "ivocabulary9000@gmail.com",
	SenderPassword: "ohfnffkbyiesjcda",
	// RecipientEmail: "spam9000me@gmail.com",
	RecipientEmail: "i.tichkevitch@godeltech.com",
}

var User = UserType{
	Email:          DefaultUser.Email,
	Pass:           DefaultUser.Pass,
	SenderEmail:    DefaultUser.SenderEmail,
	SenderPassword: DefaultUser.SenderPassword,
	RecipientEmail: DefaultUser.RecipientEmail,
}

var DefaultInvoiceData = InvoiceDataType{
	Sprzedawca: DefaultSprzedawca,
	Nabywca:    DefaultNabywca,
	Towar:      DefaultTowar,
	User:       DefaultUser,
}

var InvoiceData = InvoiceDataType{
	Sprzedawca: Sprzedawca,
	Nabywca:    Nabywca,
	Towar:      Towar,
	User:       User,
}

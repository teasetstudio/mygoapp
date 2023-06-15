package config

type firma struct {
	Nazwa      string `yaml:"nazwa"`
	Nip        string `yaml:"nip"`
	Ulica      string `yaml:"ulica"`
	Nr_budynku string `yaml:"nr_budynku"`
	Lokalu     string `yaml:"lokalu"`
	Miasto     string `yaml:"miasto"`
	Kod        string `yaml:"kod"`
	Konta      string `yaml:"konta"`
}

type towar struct {
	Nazwa string `yaml:"nazwa"`
	Cena  string `yaml:"cena"`
}

type user struct {
	Email          string `yaml:"email"`
	Pass           string `yaml:"pass"`
	SenderEmail    string `yaml:"sender_email"`
	SenderPassword string `yaml:"sender_password"`
	RecipientEmail string `yaml:"recipient_email"`
}

type InvoiceConfigType struct {
	Sprzedawca firma `yaml:"sprzedawca"`
	Nabywca    firma `yaml:"nabywca"`
	Towar      towar `yaml:"towar"`
	User       user  `yaml:"user"`
}

var DefaultNabywca = firma{
	Nazwa:      "GODEL TECHNOLOGIES EUROPE SPÓŁKA ZOGRANICZONĄ ODPOWIEDZIALNOŚCIĄ",
	Nip:        "7252307374",
	Ulica:      "Ogrodowa",
	Nr_budynku: "8",
	Miasto:     "Łódź",
	Kod:        "91-062",
}

var Nabywca = firma{
	Nazwa:      DefaultNabywca.Nazwa,
	Nip:        DefaultNabywca.Nip,
	Ulica:      DefaultNabywca.Ulica,
	Nr_budynku: DefaultNabywca.Nr_budynku,
	Miasto:     DefaultNabywca.Miasto,
	Kod:        DefaultNabywca.Kod,
}

var DefaultSprzedawca = firma{
	Nazwa:      "Ivan Tichkevitch",
	Nip:        "5833464281",
	Ulica:      "Jana Heweliusza",
	Nr_budynku: "11",
	Lokalu:     "819",
	Miasto:     "Gdańsk",
	Kod:        "80-890",
	Konta:      "40 1020 1811 0000 0202 0416 4158",
}

var Sprzedawca = firma{
	Nazwa:      DefaultSprzedawca.Nazwa,
	Nip:        DefaultSprzedawca.Nip,
	Ulica:      DefaultSprzedawca.Ulica,
	Nr_budynku: DefaultSprzedawca.Nr_budynku,
	Lokalu:     DefaultSprzedawca.Lokalu,
	Miasto:     DefaultSprzedawca.Miasto,
	Kod:        DefaultSprzedawca.Kod,
	Konta:      DefaultSprzedawca.Konta,
}

var DefaultTowar = towar{
	Nazwa: "Software development services",
	Cena:  "19000",
}

var Towar = towar{
	Nazwa: DefaultTowar.Nazwa,
	Cena:  DefaultTowar.Cena,
}

var DefaultUser = user{
	Email:          "spam9000me@gmail.com",
	Pass:           "1q2w3e1q2w3eQ1",
	SenderEmail:    "ivocabulary9000@gmail.com",
	SenderPassword: "ohfnffkbyiesjcda",
	// RecipientEmail: "spam9000me@gmail.com",
	RecipientEmail: "i.tichkevitch@godeltech.com",
}

var User = user{
	Email:          DefaultUser.Email,
	Pass:           DefaultUser.Pass,
	SenderEmail:    DefaultUser.SenderEmail,
	SenderPassword: DefaultUser.SenderPassword,
	RecipientEmail: DefaultUser.RecipientEmail,
}

var DefaultInvoiceConfig = InvoiceConfigType{
	Sprzedawca: DefaultSprzedawca,
	Nabywca:    DefaultNabywca,
	Towar:      DefaultTowar,
	User:       DefaultUser,
}

var InvoiceConfig = InvoiceConfigType{
	Sprzedawca: Sprzedawca,
	Nabywca:    Nabywca,
	Towar:      Towar,
	User:       User,
}

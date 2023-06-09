package config

type nabywca struct {
	Nazwa      string
	Nip        string
	Ulica      string
	Nr_budynku string
	Miasto     string
	Kod        string
}

var Nabywca = nabywca{
	Nazwa:      "GODEL TECHNOLOGIES EUROPE SPÓŁKA ZOGRANICZONĄ ODPOWIEDZIALNOŚCIĄ",
	Nip:        "7252307374",
	Ulica:      "Ogrodowa",
	Nr_budynku: "8",
	Miasto:     "Łódź",
	Kod:        "91-062",
}

type towar struct {
	Nazwa string
	Cena  string
}

var Towar = towar{
	Nazwa: "Software development services",
	Cena:  "19000",
}

type user struct {
	Email string
	Pass  string
}

var User = user{
	Email: "spam9000me@gmail.com",
	Pass:  "1q2w3e1q2w3eQ1",
}

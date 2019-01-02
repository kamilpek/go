package main

const imgwURL = "https://danepubliczne.imgw.pl/api/data/synop/format/xml"

type Dane struct {
	Pomiary []*Pomiar `xml:"item"`
}

type Pomiar struct {
	Stacja      string  `xml:"stacja"`
	Data        string  `xml:"data_pomiaru"`
	Godzina     int16   `xml:"godzina_pomiaru"`
	Temperatura float32 `xml:"temperatura"`
	Wiatr       float32 `xml:"predkosc_wiatru"`
	Wilgotnosc  float32 `xml:"wilgotnosc_wzgledna"`
	Opad        float32 `xml:"suma_opadu"`
	Cisnienie   float32 `xml:"cisnienie"`
}

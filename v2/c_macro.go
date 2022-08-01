package v2

type IDPMacro struct {
	FactSheet   string `json:"fsht"`
	Description string `json:"desc"`
	LastUpdate  string `json:"last"`
	Data        macros `json:"data"`
}

type macros struct {
	KR1Y      macroRows `json:"kr1y"`
	KR3Y      macroRows `json:"kr3y"`
	KR5Y      macroRows `json:"kr5y"`
	IFD1Y     macroRows `json:"ifd1y"`
	CD91D     macroRows `json:"cd91d"`
	CP91D     macroRows `json:"cp91d"`
	FB6M      macroRows `json:"fb6m"`
	FB1Y      macroRows `json:"fb1y"`
	FB3Y      macroRows `json:"fb3y"`
	KORIBOR3M macroRows `json:"koribor3m"`
	Feds      macroRows `json:"ffr"`
}

type newMacroPost struct {
	// Domestic
	KR1Y          macroRows `json:"kr1y"`
	KR2Y          macroRows `json:"kr2y"`
	KR3Y          macroRows `json:"kr3y"`
	KR5Y          macroRows `json:"kr5y"`
	KR10Y         macroRows `json:"kr10y"`
	KR20Y         macroRows `json:"kr20y"`
	KR30Y         macroRows `json:"kr30y"`
	KR50Y         macroRows `json:"kr50y"`
	K1MBS5Y       macroRows `json:"k1mbs5y"`
	MSB91D        macroRows `json:"msb91d"`
	MSB1Y         macroRows `json:"msb1y"`
	MSB2Y         macroRows `json:"msb2y"`
	KELEC3Y       macroRows `json:"kelec3y"`
	IFB1Y         macroRows `json:"ifb1y"`
	Corpbond3Yaa  macroRows `json:"corpbond3yaa"`
	Corpbond3Ybbb macroRows `json:"corpbond3ybbb"`
	CD91D         macroRows `json:"cd91d"`
	CP91D         macroRows `json:"cp91d"`
	KORIBOR1W     macroRows `json:"koribor1w"`
	KORIBOR1M     macroRows `json:"koribor1m"`
	KORIBOR2M     macroRows `json:"koribor2m"`
	KORIBOR3M     macroRows `json:"koribor3m"`
	KORIBOR6M     macroRows `json:"koribor6m"`
	KORIBOR12M    macroRows `json:"koribor12m"`
	FB6M          macroRows `json:"fb6m"`
	FB1Y          macroRows `json:"fb1y"`
	FB3Y          macroRows `json:"fb3y"`

	// Foreign
	Feds macroRows `json:"ffr"`
}

type macroRows []macroRow

type macroRow struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

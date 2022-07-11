package apis

type IDPAsset struct {
	FromSheet  string   `json:"fsht"`
	Desc       string   `json:"desc"`
	LastUpdate string   `json:"last"`
	Data       []assets `json:"data"`
}

type IDPDebt struct {
	FromSheet      string            `json:"fsht"`
	Desc           string            `json:"desc"`
	LastUpdate     string            `json:"last"`
	RowCount       int               `json:"rC"`
	Data           []debts           `json:"data"`
	DataGraphLeft  []debtsGraphLeft  `json:"datag1"`
	DataGraphRight []debtsGraphRight `json:"datag2"`
}

type IDPSingle struct {
	FromSheet  string `json:"fsht"`
	Desc       string `json:"desc"`
	LastUpdate string `json:"last"`
	Data       debts  `json:"data"`
}

type IDPMacro struct {
	FromSheet  string `json:"fsht"`
	Desc       string `json:"desc"`
	LastUpdate string `json:"last"`
	Data       macros `json:"data"`
}

type IDPModelInfo struct {
	FromSheet  string      `json:"fsht"`
	Desc       string      `json:"desc"`
	LastUpdate string      `json:"last"`
	Data       []modelmeta `json:"data"`
}

type modelmeta struct {
	ModelIndex int     `json:"num"`
	ModelStDev float64 `json:"sd"`
	InputPair  string  `json:"inp_pair"`
}

type IDPModelCoef struct {
	FromSheet  string        `json:"fsht"`
	Desc       string        `json:"desc"`
	LastUpdate string        `json:"last"`
	Data       []coefficient `json:"data"`
}

type coefficient struct {
	VarName   string  `json:"c_name_en"`
	Coef      float64 `json:"val"`
	CoefIndex int     `json:"list_num"`
}

type ModelPrediction struct {
	BankFix []float64 `json:"bankfix"`
	InsFix  []float64 `json:"insfix"`
	EtcFix  []float64 `json:"etcfix"`

	BankFloat []float64 `json:"bankfloat"`
	InsFloat  []float64 `json:"insfloat"`
	EtcFloat  []float64 `json:"etcfloat"`
}

type assets struct {
	Universe        int8   `json:"univ"`
	FundCode        string `json:"fc"`
	FundName        string `json:"fn"`
	AssetName       string `json:"an"`
	DomesticForeign string `json:"domfor"`
	AssetType       string `json:"at"`
	FundType        string `json:"ft"`
	IsMother        string `json:"ismom"`
	Strategy        string `json:"strat"`
	IsComplete      string `json:"iscmplt"`
	FileSource      string `json:"fsource"`
	FilePath        string `json:"fpath"`
	Miscellaneous   string `json:"misc"`
}

type debts struct {
	UniqueIndex          string `json:"idx"`
	FundCode             string `json:"fc"`
	FundName             string `json:"fn"`
	SetDate              string `json:"sdate"`
	MaturityDate         string `json:"mdate"`
	Image                string `json:"img"`
	AssetName            string `json:"an"`
	AssetCount           string `json:"ac"`
	DomesticForeign      string `json:"domfor"`
	AssetType            string `json:"at"`
	FundType             string `json:"ft"`
	InvestType           string `json:"it"`
	Strategy             string `json:"strat"`
	Area                 string `json:"area"`
	EquityTotal          string `json:"equity"`
	LoanTotal            string `json:"loan"`
	AUMTotal             string `json:"aum"`
	LTV                  string `json:"ltv"`
	RollOver             string `json:"ro"`
	LoanDate             string `json:"loandate"`
	LPNumber             string `json:"lpnum"`
	LP                   string `json:"lp"`
	LPCorp               string `json:"lpcorp"`
	LPType               string `json:"lpt"`
	Seniority            string `json:"seniorstr"`
	LoanUse              string `json:"loanuse"`
	LoanType             string `json:"loantype"`
	LoanClass            string `json:"loancls"`
	Address              string `json:"addr"`
	LoanAmount           string `json:"loanamt"`
	LoanRepay            string `json:"loanrpy"`
	RateType             string `json:"rate"`
	LoanInterestClass    string `json:"loanintcls"`
	LoanInterestFloat    string `json:"loanintfloat"`
	SetDateRate          string `json:"sdaterate"`
	Spread               string `json:"spread"`
	LoanPremium          string `json:"loanpremium"`
	InterestPeriod       string `json:"intdur"`
	LateRate             string `json:"laterate"`
	LateRateClass        string `json:"lateratecls"`
	SetDateLateRate      string `json:"sdatelaterate"`
	EarlyPremium         string `json:"earlypremium"`
	EarlyPremiumClass    string `json:"earlypremiumcls"`
	GuranteeLimit        string `json:"guranteelimit"`
	DSCR                 string `json:"dscr"`
	InterestDeposit      string `json:"intdeposit"`
	DefaultCondition     string `json:"default"`
	Opinion              string `json:"opinion"`
	Lender               string `json:"lender"`
	Truestee             string `json:"trustee"`
	AMC                  string `json:"amc"`
	FinancialInstutution string `json:"financialinst"`
	MoneyManage          string `json:"mm"`
	CashSupport          string `json:"cashsupp"`
	DebtUnderwrite       string `json:"debtundwrt"`
	Builder              string `json:"builder"`
	Duration             string `json:"duration"`
	ContractFile         string `json:"file"`
}

type debtsGraphLeft struct {
	UniqueIndex string `json:"idx"`
	FundCode    string `json:"fc"`
	SetDateRate string `json:"sdaterate"`
	AssetType   string `json:"at"`
	AssetName   string `json:"an"`
	LoanDate    string `json:"loandate"`
	LoanAmount  string `json:"loanamt"`
}

type debtsGraphRight struct {
	UniqueIndex string `json:"idx"`
	LoanDate    string `json:"loandate"`
	LoanAmount  string `json:"loanamt"`
	LPCorp      string `json:"lpcorp"`
}

type macros struct {
	KR1Y      []macroRow `json:"kr1y"`
	KR3Y      []macroRow `json:"kr3y"`
	KR5Y      []macroRow `json:"kr5y"`
	IFD1Y     []macroRow `json:"ifd1y"`
	CD91D     []macroRow `json:"cd91d"`
	CP91D     []macroRow `json:"cp91d"`
	KORIBOR3M []macroRow `json:"koribor3m"`
}

type newMacroPost struct {
	KR1Y          []macroRow `json:"kr1y"`
	KR2Y          []macroRow `json:"kr2y"`
	KR3Y          []macroRow `json:"kr3y"`
	KR5Y          []macroRow `json:"kr5y"`
	KR10Y         []macroRow `json:"kr10y"`
	KR20Y         []macroRow `json:"kr20y"`
	KR30Y         []macroRow `json:"kr30y"`
	KR50Y         []macroRow `json:"kr50y"`
	K1MBS5Y       []macroRow `json:"k1mbs5y"`
	MSB91D        []macroRow `json:"msb91d"`
	MSB1Y         []macroRow `json:"msb1y"`
	MSB2Y         []macroRow `json:"msb2y"`
	KELEC3Y       []macroRow `json:"kelec3y"`
	IFB1Y         []macroRow `json:"ifb1y"`
	Corpbond3Yaa  []macroRow `json:"corpbond3yaa"`
	Corpbond3Ybbb []macroRow `json:"corpbond3ybbb"`
	CD91D         []macroRow `json:"cd91d"`
	CP91D         []macroRow `json:"cp91d"`
	KORIBOR1W     []macroRow `json:"koribor1w"`
	KORIBOR1M     []macroRow `json:"koribor1m"`
	KORIBOR2M     []macroRow `json:"koribor2m"`
	KORIBOR3M     []macroRow `json:"koribor3m"`
	KORIBOR6M     []macroRow `json:"koribor6m"`
	KORIBOR12M    []macroRow `json:"koribor12m"`
	FB6M          []macroRow `json:"fb6m"`
	FB1Y          []macroRow `json:"fb1y"`
	FB3Y          []macroRow `json:"fb3y"`
}

type macroRow struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}

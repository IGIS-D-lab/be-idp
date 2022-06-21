package apis

// response IDPAsset
type IDPAsset struct {
	FromSheet  string   `json:"fsht"`
	Desc       string   `json:"desc"`
	LastUpdate string   `json:"last"`
	Data       []assets `json:"data"`
}

// response IDPDebt
type IDPDebt struct {
	FromSheet  string  `json:"fsht"`
	Desc       string  `json:"desc"`
	LastUpdate string  `json:"last"`
	RowCount   int     `json:"rC"`
	Data       []debts `json:"data"`
}

// response IDPMacro
type IDPMacro struct {
	FromSheet  string `json:"fsht"`
	Desc       string `json:"desc"`
	LastUpdate string `json:"last"`
	Data       macros `json:"data"`
}

// model assets
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

// model debts
type debts struct {
	FundCode             string `json:"fc"`
	FundName             string `json:"fn"`
	SetDate              string `json:"sdate"`
	MaturityDate         string `json:"mdate"`
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
}

// model macros
type macros struct {
	KR1Y      []macroRow `json:"kr1y"`
	KR3Y      []macroRow `json:"kr3y"`
	KR5Y      []macroRow `json:"kr5y"`
	IFD1Y     []macroRow `json:"ifd1y"`
	CD91D     []macroRow `json:"cd91d"`
	CP91D     []macroRow `json:"cp91d"`
	KORIBOR3M []macroRow `json:"koribor3m"`
}

// model macroRow
type macroRow struct {
	Date  string  `json:"date"`
	Value float32 `json:"value"`
}

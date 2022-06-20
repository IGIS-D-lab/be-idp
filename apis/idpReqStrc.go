package apis

// Request structure for RowCount
// parameters ReqRowCount
type ReqRowCount struct {
	YearFrom  int     `json:"yearFrom" validate:"required"`
	YearUntil int     `json:"yearUntil" validate:"required"`
	AumFrom   float64 `json:"aumFrom" validate:"required"`
	AumUntil  float64 `json:"aumUntil" validate:"required"`
	DebtFrom  float64 `json:"debtFrom" validate:"required"`
	DebtUntil float64 `json:"debtUntil" validate:"required"`
}

// Request structure for IDPAsset "/asset" end point.
// parameters ReqIDPAsset
type ReqIDPAsset struct {
	Strategy string `json:"strat" validate:"required"`
}

// Request structure for IDPDebt "/debt" end point.
// parameters ReqIDPDebt
type ReqIDPDebt struct {
	YearFrom  int `json:"yearFrom" validate:"required"`
	YearUntil int `json:"yearUntil" validate:"required"`
}

// Request structure for IDPMacro "/macro" end point.
// parameters ReqIDPMacro
type ReqIDPMacro struct {
	Commodity string `json:"commodity" validate:"required"`
	YearFrom  int    `json:"yearFrom" validate:"required"`
	YearUntil int    `json:"yearUntil" validate:"required"`
}

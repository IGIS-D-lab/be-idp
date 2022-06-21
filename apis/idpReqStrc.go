package apis

// Request structure for IDPAsset "/asset" end point.
// parameters ReqIDPAsset
type ReqIDPAsset struct {
	Strategy string
}

// Request structure for IDPMacro "/macro" end point.
// parameters ReqIDPMacro
type ReqIDPMacro struct {
	Commodity string
	YearFrom  int
	YearUntil int
}

package apis

type IDPDataSet struct {
	Asset     IDPAsset
	Debt      IDPDebt
	Macro     IDPMacro
	ModelInfo []byte
	ModelCoef IDPModelCoef
}

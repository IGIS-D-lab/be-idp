package apis

import (
	"log"
	"sort"
	"strings"
)

type (
	ByInvestType  []debts
	BySeniorStr   []debts
	ByAssetType   []debts
	ByRate        []debts
	ByFundName    []debts
	ByLPCorp      []debts
	ByAssetName   []debts
	ByLoanAmt     []debts
	ByAUM         []debts
	BySetDateRate []debts
	ByDuration    []debts
	ByLoanClass   []debts
	ByDebt        []debts
)

func sortByKey(sortKey, sortOrd string, d []debts) []debts {
	var target sort.Interface
	switch sortKey {
	case "it":
		target = ByInvestType(d)
	case "seniorstr":
		target = BySeniorStr(d)
	case "at":
		target = ByAssetType(d)
	case "debt":
		target = ByDebt(d)
	case "rate":
		target = ByRate(d)
	case "fn":
		target = ByFundName(d)
	case "lpcorp":
		target = ByLPCorp(d)
	case "an":
		target = ByAssetName(d)
	case "loanamt":
		target = ByLoanAmt(d)
	case "aum":
		target = ByAUM(d)
	case "sdaterate":
		target = BySetDateRate(d)
	case "duration":
		target = ByDuration(d)
	case "loancls":
		target = ByLoanClass(d)
	default:
		return d
	}

	switch sortOrd {
	case "asc":
		sort.Sort(target)
	case "desc":
		sort.Sort(sort.Reverse(target))
	default:
	}

	log.Println(MSG_DEBT, "Sorted ::", sortKey, "Order ::", strings.ToUpper(sortOrd))
	return d
}

// sort by "it"
func (s ByInvestType) Len() int {
	return len(s)
}

func (s ByInvestType) Less(i, j int) bool {
	return s[i].InvestType < s[j].InvestType
}

func (s ByInvestType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "seniorstr"
func (s BySeniorStr) Len() int {
	return len(s)
}

func (s BySeniorStr) Less(i, j int) bool {
	return s[i].Seniority < s[j].Seniority
}

func (s BySeniorStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "at"
func (s ByAssetType) Len() int {
	return len(s)
}

func (s ByAssetType) Less(i, j int) bool {
	return s[i].AssetType < s[j].AssetType
}

func (s ByAssetType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "debt"
func (s ByDebt) Len() int {
	return len(s)
}

func (s ByDebt) Less(i, j int) bool {
	return s[i].LoanTotal < s[j].LoanTotal
}

func (s ByDebt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "rate"
func (s ByRate) Len() int {
	return len(s)
}

func (s ByRate) Less(i, j int) bool {
	return s[i].RateType < s[j].RateType
}

func (s ByRate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "fn"
func (s ByFundName) Len() int {
	return len(s)
}

func (s ByFundName) Less(i, j int) bool {
	return s[i].FundName < s[j].FundName
}

func (s ByFundName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "lpcorp"
func (s ByLPCorp) Len() int {
	return len(s)
}

func (s ByLPCorp) Less(i, j int) bool {
	return s[i].LPCorp < s[j].LPCorp
}

func (s ByLPCorp) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "an"
func (s ByAssetName) Len() int {
	return len(s)
}

func (s ByAssetName) Less(i, j int) bool {
	return s[i].AssetName < s[j].AssetName
}

func (s ByAssetName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "loanamt"
func (s ByLoanAmt) Len() int {
	return len(s)
}

func (s ByLoanAmt) Less(i, j int) bool {
	return s[i].LoanAmount < s[j].LoanAmount
}

func (s ByLoanAmt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "aum"
func (s ByAUM) Len() int {
	return len(s)
}

func (s ByAUM) Less(i, j int) bool {
	return s[i].AUMTotal < s[j].AUMTotal
}

func (s ByAUM) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "sdaterate"
func (s BySetDateRate) Len() int {
	return len(s)
}

func (s BySetDateRate) Less(i, j int) bool {
	return s[i].SetDateRate < s[j].SetDateRate
}

func (s BySetDateRate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "duration"
func (s ByDuration) Len() int {
	return len(s)
}

func (s ByDuration) Less(i, j int) bool {
	return s[i].Duration < s[j].Duration
}

func (s ByDuration) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// sort by "loancls"
func (s ByLoanClass) Len() int {
	return len(s)
}

func (s ByLoanClass) Less(i, j int) bool {
	return s[i].LoanClass < s[j].LoanClass
}

func (s ByLoanClass) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

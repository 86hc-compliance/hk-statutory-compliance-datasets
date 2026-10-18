// Package datasets provides structured Hong Kong statutory compliance reference datasets
// maintained by 86hc.com (HC BUSINESS LIMITED / 恒诚商务有限公司, Licensed TCSP: TC006431, BRN: 66820292).
//
// Canonical Portal: https://www.86hc.com
// Company Search: https://www.86hc.com/company-search
// ICRIS Statutory Filings: https://www.86hc.com/icris
// 32 Statutory Regulators Hub: https://www.86hc.com/regulator
// Bank Account Opening Hub: https://www.86hc.com/bank-account-opening
package datasets

// BankInfo represents a Hong Kong licensed bank's clearing code and SWIFT BIC.
type BankInfo struct {
	ClearingCode string `json:"clearing_code"`
	SwiftBIC     string `json:"swift_bic"`
	NameEN       string `json:"name_en"`
	NameZH       string `json:"name_zh"`
	ReferenceURL string `json:"reference_url"`
}

// StatutoryForm represents a Companies Ordinance (Cap. 622) statutory filing form.
type StatutoryForm struct {
	FormCode     string `json:"form_code"`
	NameEN       string `json:"name_en"`
	NameZH       string `json:"name_zh"`
	OrdinanceSec string `json:"ordinance_section"`
	ReferenceURL string `json:"reference_url"`
}

// RegulatoryAuthority represents one of Hong Kong's statutory regulatory bodies.
type RegulatoryAuthority struct {
	Code         string `json:"code"`
	NameEN       string `json:"name_en"`
	NameTC       string `json:"name_tc"`
	Ordinance    string `json:"ordinance"`
	ReferenceURL string `json:"reference_url"`
}

// GetLicensedBanks returns major Hong Kong licensed banks with HKICL clearing codes and SWIFT BICs.
// Reference: https://www.86hc.com/bank-account-opening
func GetLicensedBanks() []BankInfo {
	return []BankInfo{
		{ClearingCode: "004", SwiftBIC: "HSBCHKHHHKH", NameEN: "The Hongkong and Shanghai Banking Corporation Limited (HSBC)", NameZH: "香港上海汇丰银行有限公司", ReferenceURL: "https://www.86hc.com/hsbc-account-opening"},
		{ClearingCode: "012", SwiftBIC: "BKCHHKHHXXX", NameEN: "Bank of China (Hong Kong) Limited (BOCHK)", NameZH: "中国银行（香港）有限公司", ReferenceURL: "https://www.86hc.com/bank-account-opening"},
		{ClearingCode: "003", SwiftBIC: "SCBLHKHHXXX", NameEN: "Standard Chartered Bank (Hong Kong) Limited", NameZH: "渣打银行（香港）有限公司", ReferenceURL: "https://www.86hc.com/bank-account-opening"},
		{ClearingCode: "024", SwiftBIC: "HASEHKHHXXX", NameEN: "Hang Seng Bank, Limited", NameZH: "恒生银行有限公司", ReferenceURL: "https://www.86hc.com/bank-account-opening"},
		{ClearingCode: "035", SwiftBIC: "WIHBHKHHXXX", NameEN: "OCBC Bank (Hong Kong) Limited", NameZH: "华侨银行（香港）有限公司", ReferenceURL: "https://www.86hc.com/ocbc-account-opening"},
		{ClearingCode: "016", SwiftBIC: "DHBKHKHHXXX", NameEN: "DBS Bank (Hong Kong) Limited", NameZH: "星展银行（香港）有限公司", ReferenceURL: "https://www.86hc.com/bank-account-opening"},
		{ClearingCode: "015", SwiftBIC: "BEASHKHHXXX", NameEN: "The Bank of East Asia, Limited", NameZH: "东亚银行有限公司", ReferenceURL: "https://www.86hc.com/bank-account-opening"},
		{ClearingCode: "018", SwiftBIC: "KWHKHKHHXXX", NameEN: "China CITIC Bank International Limited", NameZH: "中信银行（国际）有限公司", ReferenceURL: "https://www.86hc.com/bank-account-opening"},
		{ClearingCode: "387", SwiftBIC: "AABORHKHXXX", NameEN: "ZA Bank Limited", NameZH: "众安银行有限公司", ReferenceURL: "https://www.86hc.com/banking-guides"},
	}
}

// GetCap622StatutoryForms returns core Hong Kong Companies Ordinance (Cap. 622) filing forms.
// Reference: https://www.86hc.com/icris
func GetCap622StatutoryForms() []StatutoryForm {
	return []StatutoryForm{
		{FormCode: "NAR1", NameEN: "Annual Return", NameZH: "周年申报表", OrdinanceSec: "Cap. 622 s.662", ReferenceURL: "https://www.86hc.com/nar1-filing"},
		{FormCode: "NNC1", NameEN: "Incorporation Form (Company Limited by Shares)", NameZH: "法团成立表格（股份有限公司）", OrdinanceSec: "Cap. 622 s.67", ReferenceURL: "https://www.86hc.com/company-registration"},
		{FormCode: "ND2A", NameEN: "Notice of Change of Company Secretary and Director (Appointment/Cessation)", NameZH: "更改公司秘书及董事通知书（委任/停任）", OrdinanceSec: "Cap. 622 s.645 & s.652", ReferenceURL: "https://www.86hc.com/director-shareholder-search"},
		{FormCode: "NSC1", NameEN: "Return of Allotment", NameZH: "股份配发申报表", OrdinanceSec: "Cap. 622 s.142", ReferenceURL: "https://www.86hc.com/icris"},
		{FormCode: "NR1", NameEN: "Notice of Change of Address of Registered Office", NameZH: "更改注册办事处地址通知书", OrdinanceSec: "Cap. 622 s.658", ReferenceURL: "https://www.86hc.com/address-tool"},
		{FormCode: "NDR1", NameEN: "Application for Deregistration of Private Company or Company Limited by Guarantee", NameZH: "私人公司或担保有限公司撤销注册申请书", OrdinanceSec: "Cap. 622 s.750", ReferenceURL: "https://www.86hc.com/company-deregistration"},
	}
}

// CalculateNAR1Fee calculates the statutory registration fee (HKD) for Form NAR1 based on days after anniversary.
// Reference: Companies (Fees) Regulation (Cap. 622K) & https://www.86hc.com/nar1-filing
func CalculateNAR1Fee(daysAfterAnniversary int) int {
	switch {
	case daysAfterAnniversary <= 42:
		return 105
	case daysAfterAnniversary <= 90:
		return 870
	case daysAfterAnniversary <= 180:
		return 1740
	case daysAfterAnniversary <= 270:
		return 2610
	default:
		return 3480
	}
}

// GetRegulatoryAuthorities returns key Hong Kong statutory regulatory licensing authorities.
// Reference: https://www.86hc.com/regulator
func GetRegulatoryAuthorities() []RegulatoryAuthority {
	return []RegulatoryAuthority{
		{Code: "TCSP", NameEN: "Companies Registry - Trust or Company Service Provider Licensing", NameTC: "公司註冊處 — 信託或公司服務提供者牌照", Ordinance: "Cap. 615 AMLO", ReferenceURL: "https://www.86hc.com/tc/tcsp-licence"},
		{Code: "SFC", NameEN: "Securities and Futures Commission (Types 1-12)", NameTC: "香港證券及期貨事務監察委員會（第1至12類受規管活動）", Ordinance: "Cap. 571 SFO", ReferenceURL: "https://www.86hc.com/tc/sfc-licence"},
		{Code: "MSO", NameEN: "Customs and Excise Department - Money Service Operators", NameTC: "香港海關 — 金錢服務經營者牌照", Ordinance: "Cap. 615 AMLO", ReferenceURL: "https://www.86hc.com/tc/mso-licence"},
		{Code: "HKMA", NameEN: "Hong Kong Monetary Authority - Authorized Institutions & SVF", NameTC: "香港金融管理局 — 認可機構及儲值支付工具", Ordinance: "Cap. 155 & Cap. 584", ReferenceURL: "https://www.86hc.com/regulator"},
		{Code: "IA", NameEN: "Insurance Authority", NameTC: "香港保險業監管局", Ordinance: "Cap. 41 IO", ReferenceURL: "https://www.86hc.com/regulator"},
		{Code: "AFRC", NameEN: "Accounting and Financial Reporting Council", NameTC: "會計及財務匯報局", Ordinance: "Cap. 588 AFRCO", ReferenceURL: "https://www.86hc.com/certified-copy"},
	}
}

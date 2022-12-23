package models



type JsonFile struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        Data   `json:"data"`
}
type OptionData struct {
	OptionGroupID string `json:"option_group_id"`
	Right         string `json:"right"`
	StrikePrice   string `json:"strike_price"`
}
type Identifiers struct {
	Isin  string `json:"ISIN"`
	Figi  string `json:"FIGI"`
	Cusip string `json:"CUSIP"`
	Ric   string `json:"RIC"`
	Sedol string `json:"SEDOL"`
}
type BusinessActivity struct {
	NotHalal        float64 `json:"not_halal"`
	Questionable    float64 `json:"questionable"`
	Halal           float64 `json:"halal"`
	Status          string  `json:"status"`
	HahalRevenue    int64   `json:"hahal_revenue"`
	DoubtfulRevenue float64 `json:"doubtful_revenue"`
	NotHalalRevenue float64 `json:"not_halal_revenue"`
}
type InterestBearingSecuritiesAndAssets struct {
	InterestRatio float64 `json:"interestRatio"`
	Status        string  `json:"status"`
	Amount        int64   `json:"amount"`
}
type InterestBearingDebt struct {
	DebtRatio float64 `json:"debtRatio"`
	Status    string  `json:"status"`
	Amount    int64   `json:"amount"`
}
type Musaffa struct {
	IDMusaffo                          string                             `json:"id_musaffo"`
	CompanyName                        string                             `json:"company_name"`
	StockName                          string                             `json:"stock_name"`
	ShariahComplianceStatus            string                             `json:"shariah_compliance_status"`
	ComplianceRanking                  int                                `json:"compliance_ranking"`
	BusinessActivity                   BusinessActivity                   `json:"business_activity"`
	InterestBearingSecuritiesAndAssets InterestBearingSecuritiesAndAssets `json:"interest_bearing_securities_and_assets"`
	InterestBearingDebt                InterestBearingDebt                `json:"interest_bearing_debt"`
	ReportSource                       string                             `json:"report_source"`
	ReportDate                         string                             `json:"report_date"`
	TotalRevenue                       int64                              `json:"total_revenue"`
	Trailing36MonthAvrCap              int64                              `json:"trailing_36_month_avr_cap"`
	LastUpdate                         string                             `json:"last_update"`
}
type BussinessSummary struct {
	LastUpdated       string `json:"last_updated"`
	Lang              string `json:"lang"`
	SourceFillingDate string `json:"source_filling_date"`
	SourceFillingType string `json:"source_filling_type"`
	Value             string `json:"value"`
	Ticker            string `json:"ticker"`
}
type KeyStats struct {
	AverageDailyVolumeForLast10Days string `json:"average_daily_volume_for_last_10_days"`
	MarketCapitalization            string `json:"market_capitalization"`
	PriceChange1Day                 string `json:"price_change_1_day"`
	PriceChange5Day                 string `json:"price_change_5_day"`
	Eps                             string `json:"eps"`
	DividientPerShare               string `json:"dividient_per_share"`
	PE                              string `json:"p_e"`
	DividientYield                  string `json:"dividient_yield"`
	CurrentRatio                    string `json:"current_ratio"`
	LTDebtEquility                  string `json:"l_t_debt_equility"`
	QuickRatio                      string `json:"quick_ratio"`
	Roa                             string `json:"roa"`
	Roi                             string `json:"roi"`
	TotalDebtEquity                 string `json:"total_debt_equity"`
	Dividient                       int    `json:"dividient"`
	DividientYieldRatio             int    `json:"dividient_yield_ratio"`
}
type PurchaseStatuses struct {
	PocketID  string `json:"pocket_id"`
	Count     int    `json:"count"`
	SellCount int    `json:"sell_count"`
}
type Data struct {
	ID                 string             `json:"id"`
	Ticker             string             `json:"ticker"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	Group              string             `json:"group"`
	UnderlyingSymbolID string             `json:"underlying_symbol_id"`
	Exchange           string             `json:"exchange"`
	Expiration         int                `json:"expiration"`
	Country            string             `json:"country"`
	Type               string             `json:"type"`
	OptionData         OptionData         `json:"option_data"`
	Mpi                string             `json:"mpi"`
	Currency           string             `json:"currency"`
	Identifiers        Identifiers        `json:"identifiers"`
	Icon               string             `json:"icon"`
	IsFavorite         bool               `json:"is_favorite"`
	Musaffa            Musaffa            `json:"musaffa"`
	BussinessSummary   BussinessSummary   `json:"bussiness_summary"`
	KeyStats           KeyStats           `json:"key_stats"`
	PurchaseStatuses   []PurchaseStatuses `json:"purchase_statuses"`
	IsBuyable          bool               `json:"is_buyable"`
	IsSellabe          bool               `json:"is_sellabe"`
	Active             bool               `json:"active"`
}
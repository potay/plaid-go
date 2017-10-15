package api

// Accounts protos
type Account struct {
	AccountID string `json:"account_id"`
	Balances  struct {
		Available int         `json:"available"`
		Current   int         `json:"current"`
		Limit     interface{} `json:"limit"`
	} `json:"balances"`
	Mask         string `json:"mask"`
	Name         string `json:"name"`
	OfficialName string `json:"official_name"`
	Subtype      string `json:"subtype"`
	Type         string `json:"type"`
}

type Item struct {
	AvailableProducts []string    `json:"available_products"`
	BilledProducts    []string    `json:"billed_products"`
	Error             interface{} `json:"error"`
	InstitutionID     string      `json:"institution_id"`
	ItemID            string      `json:"item_id"`
	Webhook           string      `json:"webhook"`
}

// Auth protos
type Number struct {
	Account     string `json:"account"`
	AccountID   string `json:"account_id"`
	Routing     string `json:"routing"`
	WireRouting string `json:"wire_routing"`
}

// Categories protos
type Category struct {
	Group      string   `json:"group"`
	Hierarchy  []string `json:"hierarchy"`
	CategoryID string   `json:"category_id"`
}

// Credit Details protos

// Identity protos
type Identity struct {
	Addresses    []Address     `json:"addresses"`
	Emails       []Email       `json:"emails"`
	Names        []string      `json:"names"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

type Address struct {
	Accounts []string `json:"accounts"`
	Data     struct {
		City   string `json:"city"`
		State  string `json:"state"`
		Street string `json:"street"`
		Zip    string `json:"zip"`
	} `json:"data"`
	Primary bool `json:"primary"`
}

type Email struct {
	Data    string `json:"data"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

type PhoneNumber struct {
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	Data    string `json:"data"`
}

// Income protos
type Income struct {
	IncomeStreams                       []IncomeStream `json:"income_streams"`
	LastYearIncome                      int            `json:"last_year_income"`
	LastYearIncomeBeforeTax             int            `json:"last_year_income_before_tax"`
	ProjectedYearlyIncome               int            `json:"projected_yearly_income"`
	ProjectedYearlyIncomeBeforeTax      int            `json:"projected_yearly_income_before_tax"`
	MaxNumberOfOverlappingIncomeStreams int            `json:"max_number_of_overlapping_income_streams"`
	NumberOfIncomeStreams               int            `json:"number_of_income_streams"`
}

type IncomeStream struct {
	Confidence    int    `json:"confidence"`
	Days          int    `json:"days"`
	MonthlyIncome int    `json:"monthly_income"`
	Name          string `json:"name"`
}

// Institutions protos
type Institution struct {
	Credentials []struct {
		Label string `json:"label"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"credentials"`
	HasMfa        bool     `json:"has_mfa"`
	InstitutionID string   `json:"institution_id"`
	Mfa           []string `json:"mfa"`
	Name          string   `json:"name"`
	Products      []string `json:"products"`
}

// Item protos
type Credentials struct {
	Username string
	Password string
	Pin      string
}

// Transactions protos
type Transaction struct {
	AccountID  string   `json:"account_id"`
	Amount     float64  `json:"amount"`
	Category   []string `json:"category"`
	CategoryID string   `json:"category_id"`
	Date       string   `json:"date"`
	Location   struct {
		Address string      `json:"address"`
		City    string      `json:"city"`
		State   string      `json:"state"`
		Zip     string      `json:"zip"`
		Lat     interface{} `json:"lat"`
		Lon     interface{} `json:"lon"`
	} `json:"location"`
	Name                 string      `json:"name"`
	PaymentMeta          interface{} `json:"payment_meta"`
	Pending              bool        `json:"pending"`
	PendingTransactionID string      `json:"pending_transaction_id"`
	AccountOwner         string      `json:"account_owner"`
	TransactionID        string      `json:"transaction_id"`
	TransactionType      string      `json:"transaction_type"`
}

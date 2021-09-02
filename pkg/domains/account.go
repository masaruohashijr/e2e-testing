package domains

type Account struct {
	Sid            string `json:"sid"`
	FriendlyName   string `json:"friendly_name"`
	AccountBalance string `json:"account_balance"`
	DateCreated    string `json:"date_created"`
	DateUpdated    string `json:"date_updated"`
}

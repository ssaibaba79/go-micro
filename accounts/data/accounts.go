package data

import (
	"encoding/json"
	"io"
)

// Account definition
type Account struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"firstname"`
	MiddleName  string `json:"middlename"`
	LastName    string `json:"lastname"`
	Gender      string `json:"gender"`
	CreatedDate string `json:"-"`
}

func (a *Account) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func GetAccounts() Accounts {
	return accounts
}

func AddAccount() {

}

// Serialize to Json
func (accounts *Accounts) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(accounts)
}

// Deserialize from Json
func (accounts *Accounts) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(accounts)
}

// Type for list of accounts
type Accounts []*Account

var accounts = []*Account{
	{
		ID:         "af1d5a5b-813f-49cf-a0d5-39c5fcaf0110",
		Username:   "user1@example.com",
		FirstName:  "Sam",
		MiddleName: "",
		LastName:   "Altman",
		Gender:     "M",
	},
	{
		ID:         "d4ea5e59-b72d-4b49-8bfa-9cdcb1d1cd90",
		Username:   "user@example.com",
		FirstName:  "Sean",
		MiddleName: "William",
		LastName:   "Parker",
		Gender:     "M",
	},
}

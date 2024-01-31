package data

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/beevik/guid"
)

// Account definition
type Account struct {
	ID          string `json:"id"`
	Username    string `json:"username" validate:"required"`
	FirstName   string `json:"firstname" validate:"required"`
	MiddleName  string `json:"middlename"`
	LastName    string `json:"lastname" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	CreatedDate string `json:"-"`
}

func (a *Account) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func (a *Account) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *Account) copyFrom(in Account) {
	a.FirstName = in.FirstName
	a.MiddleName = in.MiddleName
	a.LastName = in.LastName
	a.Gender = in.Gender
}

// Fetch all accounts
func GetAccounts() Accounts {
	return accounts
}

// Get Account by Id
func GetAccount(id string) *Account {
	for _, acct := range accounts {
		if acct.ID == id {
			return acct
		}
	}
	return nil
}

// Add a new account
func AddAccount(newAccount *Account) {
	id := guid.New()
	newAccount.ID = id.String()
	accounts = append(accounts, newAccount)
}

// Update an account
func UpdateAccount(id string, account *Account) (*Account, error) {
	savedAccount := GetAccount(id)
	if savedAccount == nil {
		return nil, errors.New("Account with id " + id + " not found")
	}
	savedAccount.copyFrom(*account)
	return savedAccount, nil
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

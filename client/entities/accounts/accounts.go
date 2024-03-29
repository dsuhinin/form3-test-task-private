package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/form3-test-task/client/entities/accounts/objects"
	"github.com/form3-test-task/client/transport"
)

// supported Account endpoints.
const (
	fetchAccountEndpoint  = "/organisation/accounts/%s"
	createAccountEndpoint = "/organisation/accounts"
	deleteAccountEndpoint = "/organisation/accounts/%s?version=%d"
)

// Accounts represents object to work with Account entity.
type Accounts struct {
	transport transport.HTTPProvider
}

// NewAccounts creates new Account instance.
func NewAccounts(transport transport.HTTPProvider) *Accounts {
	return &Accounts{
		transport: transport,
	}
}

// Create creates a new Account entity.
func (a Accounts) Create(account *objects.Account) (*objects.Account, error) {
	body, err := a.transport.Do(
		http.MethodPost, createAccountEndpoint, objects.CreateAccountRequest{Data: *account},
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Account object")
	}

	responseObject := objects.CreateAccountResponse{}
	if err := json.Unmarshal(body, &responseObject); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling response body to Account response object")
	}

	return &responseObject.Data, nil
}

// Fetch fetches existing Account entity by its ID.
func (a Accounts) Fetch(ID string) (*objects.Account, error) {
	body, err := a.transport.Do(http.MethodGet, fmt.Sprintf(fetchAccountEndpoint, ID), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching Account object")
	}
	account := objects.FetchAccountResponse{}
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling response body to Account response object")
	}
	return &account.Data, nil
}

// Delete deletes existing Account entity by its ID and version.
func (a Accounts) Delete(ID string, version int64) error {
	_, err := a.transport.Do(
		http.MethodDelete, fmt.Sprintf(deleteAccountEndpoint, ID, version), nil,
	)
	if err != nil {
		return errors.Wrap(err, "error creating request to delete Account object")
	}

	return nil
}

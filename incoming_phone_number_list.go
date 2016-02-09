package twiliogo

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

var (
	NoNextPage     error = errors.New("No next page exists for this resource")
	NoPreviousPage error = errors.New("No previous page exists for this resource")
)

type IncomingPhoneNumberList struct {
	Client               Client
	Start                int                   `json:"start"`
	Total                int                   `json:"total"`
	NumPages             int                   `json:"num_pages"`
	Page                 int                   `json:"page"`
	PageSize             int                   `json:"page_size"`
	End                  int                   `json:"end"`
	Uri                  string                `json:"uri"`
	FirstPageUri         string                `json:"first_page_uri"`
	LastPageUri          string                `json:"last_page_uri"`
	NextPageUri          string                `json:"next_page_uri"`
	PreviousPageUri      string                `json"previous_page_uri"`
	IncomingPhoneNumbers []IncomingPhoneNumber `json:"incoming_phone_numbers"`
}

func GetIncomingPhoneNumberList(client Client, optionals ...Optional) (*IncomingPhoneNumberList, error) {
	return getIncomingPhoneNumberListByPath("/IncomingPhoneNumbers.json", client, optionals...)
}

func GetNextPage(currentList *IncomingPhoneNumberList, client Client, optionals ...Optional) (*IncomingPhoneNumberList, error) {
	if currentList.NextPageUri != "" {
		uri := strings.Split(currentList.NextPageUri, client.AccountSid())
		return getIncomingPhoneNumberListByPath(uri[len(uri)-1], client, optionals...)
	} else {
		return nil, NoNextPage
	}
}

func GetPreviousPage(currentList *IncomingPhoneNumberList, client Client, optionals ...Optional) (*IncomingPhoneNumberList, error) {
	if currentList.PreviousPageUri != "" {
		uri := strings.Split(currentList.PreviousPageUri, client.AccountSid())
		return getIncomingPhoneNumberListByPath(uri[len(uri)-1], client, optionals...)
	} else {
		return nil, NoPreviousPage
	}
}

func getIncomingPhoneNumberListByPath(path string, client Client, optionals ...Optional) (*IncomingPhoneNumberList, error) {
	var incomingPhoneNumberList *IncomingPhoneNumberList

	params := url.Values{}

	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	body, err := client.get(params, path)

	if err != nil {
		return nil, err
	}

	incomingPhoneNumberList = new(IncomingPhoneNumberList)
	incomingPhoneNumberList.Client = client
	err = json.Unmarshal(body, incomingPhoneNumberList)

	return incomingPhoneNumberList, err
}

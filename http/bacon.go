package bacon

import (
	"net/http"
)

type BaconClient struct {
	httpClient *http.Client
}

func (client *BaconClient) BaconPost(address string) {
	req, _ := http.NewRequest("POST", address, nil)
	req.Header.Add("CustomHeader", "iLoveBacon")

	client.httpClient.Do(req)
}

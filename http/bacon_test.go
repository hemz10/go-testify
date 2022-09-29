package bacon

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomHeader(t *testing.T) {
	// This is what will be called when the request arrives
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		// We don't even need to answer the request, we just need to assert the request has the data we want
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.Header.Get("CustomHeader"), "iLoveBacon")
	}

	// Here we're effectively creating a server and passing our `testHandler` to it
	testServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer testServer.Close()

	// Now let's instantiate a client and tell it to do its request to our fake server
	httpClient := &http.Client{}
	client := BaconClient{httpClient: httpClient}

	// This sends the POST request with our custom header
	client.BaconPost(testServer.URL)

	/*
			This is what happens when running the test above:

		We create testHandler, which is the function that will be called by our server when a request arrives.
		We create testServer and pass testHandler to it.
		We create a BaconClient.
		We tell the BaconClient to execute is BaconPost method, which does an HTTP request
		Our testServer receives this request and calls testHandler
		The testHandler function is run, executing the assertions on the req object passed to it
	*/

}

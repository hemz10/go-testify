package gobotto

import (
	"net/url"
)

// A simple function which takes a URL(address) as an argument and returns adding /robots.txt at the end of the URL.
func RobotsURL(address string) (string, error) {
	parsedURL, err := url.Parse(address)
	result := parsedURL.Scheme + "://" + parsedURL.Host + "/robots.txt"
	return result, err
}

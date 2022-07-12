package input

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

// This will use the net/http client reach out to a site and collect the content.
func getHttpContent(uri string) ([]byte, error) {
	// Code to disable the http2 client for reddit.
	// https://github.com/golang/go/issues/39302
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.ForceAttemptHTTP2 = false
	tr.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
	tr.TLSClientConfig = &tls.Config{}

	client := &http.Client{
		Transport: tr,
	}
	
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil { return nil, err }

	// set the user agent header to avoid kick backs.. as much
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.10; rv:75.0) Gecko/20100101 Firefox/75.0")

	log.Printf("Requesting content from %v\n", uri)
	resp, err := client.Do(req)
	if err != nil { log.Fatalln(err) }
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return nil, err }

	return body, nil
}
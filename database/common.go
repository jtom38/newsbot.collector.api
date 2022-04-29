package database

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jtom38/newsbot/collector/services/config"
)

type DatabaseClient struct {
	Diagnosis DiagnosisClient

	Articles ArticlesClient
	Sources SourcesClient
}

// This will generate a new client to interface with the API Database.
func NewDatabaseClient() DatabaseClient {
	cc := config.New()
	dbUri := cc.GetConfig(config.DB_URI)

	var client = DatabaseClient{}
	client.Diagnosis.rootUri = dbUri
	client.Sources.rootUri = dbUri
	client.Articles.rootUri = dbUri

	return client
}

func getContent(url string) ([]byte, error) {
	client := &http.Client{}
	var blank []byte
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil { return blank, err }

	// set the user agent header to avoid kick backs.. as much
	req.Header.Set("User-Agent", getUserAgent())

	log.Printf("Requesting content from %v\n", url)
	resp, err := client.Do(req)
	if err != nil { return blank, err }
	if resp.StatusCode == 404 { 
		err = errors.New("404 not found")
		return blank, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { return blank, err }

	//log.Println(string(body))
	return body, nil
}

func httpDelete(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil { return err }

	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.10; rv:75.0) Gecko/20100101 Firefox/75.0")
	req.Header.Set("User-Agent", getUserAgent())

	_, err = client.Do(req)
	if err != nil { return err }

	return nil
}

func getUserAgent() string {
	return "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.10; rv:75.0) Gecko/20100101 Firefox/75.0"
}
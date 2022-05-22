package databaseRest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	//"github.com/jtom38/newsbot/collector/services"
)

type DiagnosisClient struct {
	rootUri string
}

func (dc *DiagnosisClient) Ping() error {
	dbPing := fmt.Sprintf("%v/ping", dc.rootUri)
	resp, err := http.Get(dbPing)
	if err != nil { return err }

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil { return err }
	return nil
}


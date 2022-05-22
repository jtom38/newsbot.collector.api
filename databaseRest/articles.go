package databaseRest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jtom38/newsbot/collector/domain/model"
)

// Generate this struct fr
type ArticlesClient struct {
	rootUri string
}

func (ac *ArticlesClient) List() ([]model.Articles, error) {
	var items []model.Articles
	url := fmt.Sprintf("%v/api/v1/articles", ac.rootUri)
	resp, err := getContent(url)
	if err != nil {
		return items, err
	}

	err = json.Unmarshal(resp, &items)
	if err != nil {
		return []model.Articles{}, err
	}

	return items, nil
}

func (ac *ArticlesClient) FindByID(ID uint) (model.Articles, error) {
	var items model.Articles
	url := fmt.Sprintf("%v/api/v1/articles/%v", ac.rootUri, ID)
	resp, err := getContent(url)
	if err != nil {
		return items, err
	}

	err = json.Unmarshal(resp, &items)
	if err != nil {
		return items, err
	}

	return items, nil
}

func (ac *ArticlesClient) FindByUrl(url string) (model.Articles, error) {
	var item model.Articles
	get := fmt.Sprintf("%v/api/v1/articles/url/%v", ac.rootUri, url)
	resp, err := getContent(get)
	if err != nil {
		return item, err
	}

	err = json.Unmarshal(resp, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ac *ArticlesClient) Delete(id int32) error {
	return errors.New("not implemented")
}

func (ac *ArticlesClient) Add(item model.Articles) error {
	//return errors.New("not implemented")
	url := fmt.Sprintf("%v/api/v1/articles/", ac.rootUri)

	bItem, err := json.Marshal(item)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bItem))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return errors.New("failed to post to the DB")
	}

	return nil
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil { return err }

}

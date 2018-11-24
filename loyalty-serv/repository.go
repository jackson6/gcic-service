// plan-service/repository.go

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	url = "http://localhost:3000"
	actionApi = "/api/"
	resolveSuffix = "?resolve=true"
)

type Repository interface {
	GetAll(ns string) ([]byte, error)
	GetSingle(ns, id string) ([]byte, error)
	Add(ns string, asset interface{}) ([]byte, error)
	Update(ns, id string, item *interface{}) ([]byte, error)
	Delete(ns, id string) error
}

type LoyaltyRepository struct {
	client *http.Client
}

// Get all assets
func (repo *LoyaltyRepository) GetAll(ns string)  ([]byte, error) {
	url := url + actionApi + ns
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept'", "application/json")

	resp, err := repo.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Get single asset
func (repo *LoyaltyRepository) GetSingle(ns, id string)  ([]byte, error) {
	url := url + actionApi + ns + "/" + id + resolveSuffix
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept'", "application/json")

	resp, err := repo.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Add assets
func (repo *LoyaltyRepository) Add(ns string, asset interface{}) ([]byte, error) {
	url := url + actionApi + ns

	d, err := json.Marshal(asset)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(d)

	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept'", "application/json")

	resp, err := repo.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	return  response, nil
}

// update asset
func (repo *LoyaltyRepository) Update(ns, id string, asset *interface{}) ([]byte, error) {
	url := url + actionApi + ns + id

	d, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(d)

	req, err := http.NewRequest("PUT", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept'", "application/json")

	resp, err := repo.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return  nil
}

// delete asset
func (repo *LoyaltyRepository) Delete(ns, id string) error {
	url := url + actionApi + ns + id
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept'", "application/json")

	resp, err := repo.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return  nil
}

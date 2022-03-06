package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	GetData() ([]byte, error)
}

type LoLService struct {
	databaseUrl       string
	dataDragonUrlBase string
	dataDragonVersion string
}

func New(databaseUrl, dataDragonVersion string) Service {
	return &LoLService{
		databaseUrl:       databaseUrl,
		dataDragonUrlBase: "http://ddragon.leagueoflegends.com",
		dataDragonVersion: dataDragonVersion,
	}
}

func (s *LoLService) getChampions() (*http.Response, error) {
	url := fmt.Sprintf("%v/cdn/%v/en_GB/champion.json", s.dataDragonUrlBase, s.dataDragonVersion)
	resp, err := http.Get(url)
	return resp, err
}

// Reads http response to []byte
func parseBody(resp *http.Response) ([]byte, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

type Champion struct {
	Name string `json:"name"`
}

type DataDragonChampionList struct {
	Data []Champion `json:"data"`
}

// Unmarshals response to a parseDataDragonChampionList
func parseDataDragonChampionList(resp *http.Response) (DataDragonChampionList, error) {
	var data DataDragonChampionList
	body, err := parseBody(resp)
	if err != nil {
		return DataDragonChampionList{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return DataDragonChampionList{}, err
	}
	return data, nil
}

// Gets all LoL data from riot, JSON stringifies it, and returns
func (s *LoLService) GetData() ([]byte, error) {
	champions, err := s.getChampions()
	if err != nil {
		return []byte{}, err
	}

	data, err := parseDataDragonChampionList(champions)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := json.Marshal(data)
	return bytes, nil
}

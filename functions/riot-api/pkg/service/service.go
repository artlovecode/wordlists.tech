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

type lolService struct {
	datadragonURLBase string
	datadragonVersion string
}

// New instanitates a new service
func New(dataDragonVersion string) Service {
	return &lolService{
		datadragonURLBase: "http://ddragon.leagueoflegends.com",
		datadragonVersion: dataDragonVersion,
	}
}

func (s *lolService) getChampions() (*http.Response, error) {
	url := fmt.Sprintf("%v/cdn/%v/en_GB/champion.json", s.datadragonURLBase, s.datadragonVersion)
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

type champion struct {
	Name string `json:"name"`
}

type dataDragonChampionList struct {
	Data []champion `json:"data"`
}

// Unmarshals response to a parseDataDragonChampionList
func parseDataDragonChampionList(resp *http.Response) (dataDragonChampionList, error) {
	var data dataDragonChampionList
	body, err := parseBody(resp)
	if err != nil {
		return dataDragonChampionList{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return dataDragonChampionList{}, err
	}
	return data, nil
}

// Gets all LoL data from riot, JSON stringifies it, and returns
func (s *lolService) GetData() ([]byte, error) {
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

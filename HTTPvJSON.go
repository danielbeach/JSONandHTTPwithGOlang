package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Oil struct {
	Dataset struct {
		ID                  int             `json:"id"`
		DatasetCode         string          `json:"dataset_code"`
		DatabaseCode        string          `json:"database_code"`
		Name                string          `json:"name"`
		Description         string          `json:"description"`
		RefreshedAt         time.Time       `json:"refreshed_at"`
		NewestAvailableDate string          `json:"newest_available_date"`
		OldestAvailableDate string          `json:"oldest_available_date"`
		ColumnNames         []string        `json:"column_names"`
		Frequency           string          `json:"frequency"`
		Type                string          `json:"type"`
		Premium             bool            `json:"premium"`
		Limit               interface{}     `json:"limit"`
		Transform           interface{}     `json:"transform"`
		ColumnIndex         interface{}     `json:"column_index"`
		StartDate           string          `json:"start_date"`
		EndDate             string          `json:"end_date"`
		Data                [][]interface{} `json:"data"`
		Collapse            interface{}     `json:"collapse"`
		Order               interface{}     `json:"order"`
		DatabaseID          int             `json:"database_id"`
	} `json:"dataset"`
}

type GrossDProduct struct {
	Dataset struct {
		ID                  int             `json:"id"`
		DatasetCode         string          `json:"dataset_code"`
		DatabaseCode        string          `json:"database_code"`
		Name                string          `json:"name"`
		Description         string          `json:"description"`
		RefreshedAt         time.Time       `json:"refreshed_at"`
		NewestAvailableDate string          `json:"newest_available_date"`
		OldestAvailableDate string          `json:"oldest_available_date"`
		ColumnNames         []string        `json:"column_names"`
		Frequency           string          `json:"frequency"`
		Type                string          `json:"type"`
		Premium             bool            `json:"premium"`
		Limit               interface{}     `json:"limit"`
		Transform           interface{}     `json:"transform"`
		ColumnIndex         interface{}     `json:"column_index"`
		StartDate           string          `json:"start_date"`
		EndDate             string          `json:"end_date"`
		Data                [][]interface{} `json:"data"`
		Collapse            interface{}     `json:"collapse"`
		Order               interface{}     `json:"order"`
		DatabaseID          int             `json:"database_id"`
	} `json:"dataset"`
}

func get_api_response(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response")
	}
	return body
}

func get_oil_response(url string) Oil {
	body := get_api_response(url)
	var data Oil
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return data
}

func get_gdp_response(url string) GrossDProduct {
	body := get_api_response(url)
	var data GrossDProduct
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return data
}

func write_csv_file(fname string, data []interface{}) {
	file, err := os.Create(fname + ".csv")
	if err != nil {
		fmt.Println("Could not open file " + fname)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	valStr := fmt.Sprintf("%v", data)
	var myslice []string
	myslice = append(myslice, valStr)
	writer.Write(myslice)
}

func main() {
  oil_api := "https://data.nasdaq.com/api/v3/datasets/OPEC/ORB.json?api_key={api_key}"
  gdp_api := "https://data.nasdaq.com/api/v3/datasets/FED/FU086902203_A.json?api_key={api_key}"
	oil_data := get_oil_response(oil_api)
	gdp_data := get_gdp_response(gdp_api)
	write_csv_file("oil", oil_data.Dataset.Data[0])
	write_csv_file("gdp", gdp_data.Dataset.Data[0])
}

package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string,
	response interface{},
	headers map[string]string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// custom headers
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic("Request failed")
	}
	if res.StatusCode != 200 {
		fmt.Printf("API returned '%s'\n", res.Status)
		panic("Request failed")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic("Could not read response")
	}

	json.Unmarshal(body, response)
}

func Post(url string,
	request interface{},
	response interface{},
	headers map[string]string) {
	payload, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		panic("Could not marshal the request")
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	// custom headers
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic("Request failed")
	}
	if res.StatusCode != 200 {
		fmt.Printf("API returned '%s'\n", res.Status)
		panic("Request failed")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic("Could not read response")
	}

	json.Unmarshal(body, response)
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

// Perform a GET request
func GetRequest(url string) {
	client := GetHttpClient()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("GET Response:", string(body))
}

// Perform a POST Request
func PostRequest(url string, data []byte) {
	client := GetHttpClient()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("POST Response:", string(body))
}

func PutRequest(url string, data []byte) {
	client := GetHttpClient()

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("PUT Response:", string(body))
}

func DeleteRequest(url string) {
	client := GetHttpClient()

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("DELETE RESPONSE:", string(body))
}

func main() {
	apiUrl := "https://jsonplaceholder.typicode.com/posts"

	GetRequest(apiUrl)

	postData := []byte(`{"title":"foo","body":"bar","userId":1}`)
	PostRequest(apiUrl, postData)

	putData := []byte(`{"title":"foo1","body":"bar1","userId":1}`)
	PutRequest(apiUrl+"/1", putData)

	DeleteRequest(apiUrl + "/1")
}

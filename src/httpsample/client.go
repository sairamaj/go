package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	 "crypto/tls"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getConfigSoap() string {
	data, err := ioutil.ReadFile("soap.xml")
	check(err)
	return string(data)
}

func MakeRequest() {

	message := getConfigSoap()

	log.Println("______________________")
	log.Println(message)
	log.Println("______________________")

	// Note: remove soap:header:action from captured message.
	payload := []byte(strings.TrimSpace(message))
	httpMethod := "POST"
	url := "urlhere"

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}
	soapAction := "soapactionhere(value from soap:Header:Action"
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	log.Println("StatusCode=", resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)

}

func main() {
	MakeRequest()
}

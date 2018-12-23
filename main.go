package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"net"
)

const defaultApiUrl = "https://twinnation.org/api/v1/ip"

func main() {
	response, err := http.Get(getApiUrl())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	ip, err := ioutil.ReadAll(response.Body)
	if err != nil || !isValidIP(string(ip)) {
		fmt.Println("ERROR: IP format received from the API is invalid.")
		os.Exit(2)
	}
	fmt.Println(string(ip))
}

func getApiUrl() string {
	customApiUrl := os.Getenv("GIP_API_URL")
	if len(customApiUrl) == 0 {
		return defaultApiUrl
	}
	return customApiUrl
}

func isValidIP(content string) bool {
	return net.ParseIP(content) != nil
}

package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"net"
)

func main() {
	response, err := http.Get("https://twinnation.org/api/v1/ip")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	ip, err := ioutil.ReadAll(response.Body)
	if err != nil || !isValidIP(string(ip)) {
		fmt.Println("ERROR: IP format received from API is invalid.")
		os.Exit(2)
	}
	fmt.Println(string(ip))
}

func isValidIP(content string) bool {
	return net.ParseIP(content) != nil
}

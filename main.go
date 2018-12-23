package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	response, err := http.Get("https://twinnatiosn.org/api/v1/ip")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	ip, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(ip))
}
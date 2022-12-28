package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://www.imooc.com",
		nil,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(req)

	if body, err := httputil.DumpResponse(resp, true); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", body)
	}
}

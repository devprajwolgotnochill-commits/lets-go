// postform
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	myurl := "http://localhost:8000/postform"

	// form-data

	payload := url.Values{}
	payload.Add("firstname", "dev")
	payload.Add("lastname", "pg")

	//can use NewRequest too
	// req, _ := http.NewRequest("POST", myurl, payload)

	// auto Encodes data ,Sets header
	resp, er1 := http.PostForm(myurl, payload)

	if er1 != nil {
		panic(er1)
	}

	// 200 ok ,!= 200 not ok
	fmt.Println("Response :", resp.StatusCode)

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp)

	fmt.Println(string(body))

}

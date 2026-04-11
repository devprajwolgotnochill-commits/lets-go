// url handler
// GET

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const WebUrl string = "http://localhost:8000/get"

	url_get, _ := Get_request(WebUrl)

	// fmt.Println(url_get)

	WriteData("get.json", url_get)
}

func Get_request(web_url string) (string, error) {

	response, er1 := http.Get(web_url)

	if er1 != nil {

		// fmt.Println("Is the server running?")
		panic(er1)

	}

	defer response.Body.Close()

	fmt.Println("Response :", response.StatusCode)
	fmt.Println("Content Length :", response.ContentLength)

	content, er2 := io.ReadAll(response.Body)

	if er2 != nil {
		panic(er2)
	}

	// content is in byte format
	// fmt.Println("Content :", string(content))
	//
	//

	return string(content), nil

}

// func byte2str(provide_byte []byte) string {

// 	var provideString strings.Builder

// 	provideString.Write(provide_byte)

// 	// bytesCount, _ := provideString.Write(provide_byte)

//  // provideString.

// 	return provideString.String()

// }

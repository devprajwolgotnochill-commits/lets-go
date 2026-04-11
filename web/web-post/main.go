// POST

package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {

	payload := strings.NewReader(`{
	"name":"pge",
	"age" : 12,
	"email":"abc@gmail.com"
}`)
	const WebUrl string = "http://localhost:8000/post"

	content, _ := Post_request(WebUrl, payload)

	// // what is sent?
	// print(content)

	WriteData("get.json", content)

}

// request *strings.Reader → a pointer to a strings.Reader
func Post_request(web_url string, body *strings.Reader) (string, error) {

	// fmt.Printf("Type :%T", request_body)

	req, er1 := http.NewRequest("POST", web_url, body)

	if er1 != nil {
		panic(er1)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User", "client")
	req.Header.Add("Content-Type", "application/json")

	//close
	defer req.Body.Close()

	// read content
	content, _ := io.ReadAll(req.Body)

	// content is byte -> str
	// fmt.Print(string(content))

	return string(content), nil
}

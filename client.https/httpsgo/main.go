// https
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

func getBody(url_addr string) string {

	http_response, er1 := http.Get(url_addr)

	if er1 != nil {
		panic(er1)
	}

	// *http.Response
	// body_content is NOT the body itself, it's the whole HTTP response.
	// fmt.Println(body_content)

	databytes, er2 := io.ReadAll(http_response.Body)

	if er2 != nil {
		panic(er2)
	}

	// fmt.Println(string(databytes))

	// fmt.Println(databytes) prints the datbytes

	// compute return value → run defer → exit function
	defer http_response.Body.Close()

	return string(databytes)

}

func main() {

	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)

	urladdr := "https://devprajwalgotnochill.github.io/AprilFool/"

	content := getBody(urladdr)

	writeData("index.html", content)

}

func writeData(write_path string, write_content string) {

	file, er1 := os.Create(write_path)

	if er1 != nil {
		panic(er1)
	}

	defer file.Close()

	length, er2 := io.WriteString(file, write_content)

	if er2 != nil {
		panic(er2)
	}

	fmt.Println("Wrote file of length :", length)

}

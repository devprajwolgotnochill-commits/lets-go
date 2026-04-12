//goUrls

package main

import (
	"fmt"
	"net/url"
)

const myurls string = "https://devprajwalgotnochill.github.io/AprilFool/"

func main() {

	// fmt.Println("Helo")

	fmt.Println(myurls)

	// parsing url or analyse

	parse_url, er1 := url.Parse(myurls)

	if er1 != nil {
		panic(er1)
	}

	// fmt.Printf("Type :%T \n", parse_url) //*url.URL
	// fmt.Println(parse_url.Scheme)        //https
	// fmt.Println(parse_url.Host)          //devprajwalgotnochill.github.io
	// fmt.Println(parse_url.Path)          ///AprilFool/
	// fmt.Println(parse_url.Port())        //empty
	// fmt.Println(parse_url.RawQuery)      //empty
	// fmt.Println(parse_url.Hostname())    //devprajwalgotnochill.github.io

	quary_params := parse_url.Query()
	fmt.Printf("Type :%T \n", quary_params) //url.Values

	fmt.Println(quary_params)

	for _, val := range quary_params {
		// no params in my website?
		fmt.Println("INFO : ", val)
	}

	// IMP
	// this is boilerPlate
	// http://localhost

	url_parts := &url.URL{
		Scheme:   "http",
		Host:     "localhost",
		Path:     "/home",
		RawQuery: "user=IDK",
	}

	another_url := url_parts.String()

	fmt.Println(another_url)
}

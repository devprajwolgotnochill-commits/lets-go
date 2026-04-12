// encodign/json
package main

import (
	"fmt"
)

// size=96 (0x60)
type UserName struct {
	// if not capital
	// the json data will be empty?
	//  `json:"yourName"` == alises
	// `json:"-"` removes the whole Secret?

	FirstName string `json:"yourName"`
	LastName  string
	Email     string
	Age       int
	Secret    string   `json:"-"`              //removes
	Tags      []string `json:"tags,omitempty"` //removes if empty
}

func main() {

	// user1st := []UserName{
	// 	{"p", "g", "email12@gmail.com", 12, "******", []string{"a", "b", "c"}},
	// 	{"p3", "g3", "email123@gmail.com", 12, "***3***", []string{"g", "h", "i"}},
	// 	{"p2", "g2", "email122@gmail.com", 12, "***2***", []string{}},
	// }

	//returns []uint8 raw json
	// json_data, _ := enc2json(user1st)

	// fmt.Println(string(json_data))

	json_data_web := []byte(`
	{
	"yourName":"p","LastName":"g","Email":"email12@gmail.com","Age":12,"tags":["a","b","c"]
	}
	`)

	data := dec2json(json_data_web)

	fmt.Println(data)

}

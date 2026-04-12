package main

import (
	"encoding/json"
	"fmt"
)

type userName struct {
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
	user1st := []userName{
		{"p", "g", "email12@gmail.com", 12, "******", []string{"a", "b", "c"}},
		{"p3", "g3", "email123@gmail.com", 12, "***3***", []string{"g", "h", "i"}},
		{"p2", "g2", "email122@gmail.com", 12, "***2***", []string{}},
	}

	//returns []uint8 raw json
	json_data, _ := enc2json(user1st)

	fmt.Println(string(json_data))

}

func enc2json(user []userName) ([]uint8, error) {

	// fmt.Println(user1st)

	// json.Marshal(user, "str" ,"\t"})
	json_data, er1 := json.Marshal(user)

	if er1 != nil {

		return nil, er1
	}

	// fmt.Printf("%T\n", json_data)
	// fmt.Println(string(json_data))

	return json_data, er1

}

//encoding json

package main

import "encoding/json"

func enc2json(user []UserName) ([]uint8, error) {

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

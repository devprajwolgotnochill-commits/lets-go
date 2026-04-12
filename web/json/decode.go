// decoding json
// using json

package main

import (
	"encoding/json"
	"fmt"
)

func dec2json(data_web []byte) UserName {

	//data from the web is in byte format

	var data_user UserName

	checkValid := json.Valid(data_web)

	if checkValid {

		fmt.Println("Valid json!")

		json.Unmarshal(data_web, &data_user)

		//

		//{FirstName:"p", LastName:"g"
		// fmt.Printf("%#v\n", data_user)

		// {p g
		// fmt.Println(data_user)

		return data_user

	}

	panic("Invalid json!!!")

	// data_json_web := make(map[string]interface{})

	// json.Unmarshal(data_web, &data_json_web)

	// fmt.Println(data_json_web)

	// fmt.Printf("%#v\n", data_json_web)

}

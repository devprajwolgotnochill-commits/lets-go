package main

import "fmt"

func main() {

	// this is map
	m := make(map[string]string, 0)
	m["uid-01"] = "iwnat@g24"

	fmt.Println(m)

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	fmt.Println(commits)

	delete(commits, "rsc")

	fmt.Println(commits)

	// loops in go is aewre and aswuawnmm

	// _ ,val
	// _ (undersocore = ignore)
	for key, val := range commits {
		fmt.Printf("Key : %v ,Value : %v \n", key, val)
	}

}

// hast
// hash table?s

//method

package main

import (
	"fmt"
	"strings"
)

func main() {

	user_1st := User_name{"JOOonny", 'j', "jonny123@gmail.com", true, 23}
	fmt.Println(user_1st)

	user_1st.Get_email()

	// using pointer to change the actual value
	user_1st.NormalizeName()

	user_1st.Change_email()

	fmt.Println(user_1st)
}

type User_name struct {
	Name      string
	letter1st rune //rune (int32)
	Email     string
	Status    bool
	Age       int
}

func (u User_name) Get_email() string {

	return u.Email

}

func (u *User_name) NormalizeName() {
	u.Name = strings.ToLower(u.Name)
}

// need to learn pointer
func (u *User_name) Change_email() {

	u.Email = "newemail.com"

}

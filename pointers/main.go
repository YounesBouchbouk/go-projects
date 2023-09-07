package main

import "fmt"

type User struct {
	name      string
	email     string
	phoneNum  int64
	addresses []Address
}

type Address struct {
	rue string
}

type UpdateBody struct {
	name      string
	email     string
	addresses []Address
}

func (u *User) updateInfo(newUserInfo UpdateBody) (*User, error) {
	return nil, nil
}

func (u *User) userAdress() []Address {
	return u.addresses
}

func (u *User) updateUserAdress(newAdd []Address) {
	u.addresses = newAdd
}

func normalFuncToUpdateAddress(user *User, newadd []Address) *User {
	user.addresses = newadd

	return user
}

func (u *User) checkValues() {
	if u.email == "" {
		fmt.Println("null email")
	}

	if u.name == "" {
		fmt.Println("null name")
	}

	if len(u.addresses) == 0 {
		fmt.Println("null address")
	}

	fmt.Println(u.phoneNum)

}

func main() {
	user2 := User{}

	user2.checkValues()

}

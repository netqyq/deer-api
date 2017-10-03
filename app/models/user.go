package models

import (
	"fmt"
	"regexp"

	"github.com/revel/revel"
)

type User struct {
	UserId          int
	Name            string
	Email, Password string
	HashedPassword  []byte
	Token           []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Email)
}

var emailRegexp = regexp.MustCompile(`[a-zA-Z0-9_\-]+@[a-zA-Z0-9_\-]+\.[a-zA-Z0-9_\-]+[a-zA-Z0-9]+$`)

func (user *User) Validate(v *revel.Validation) {
	ValidateEmail(v, user.Email).Key("user.Email")

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

func ValidateEmail(v *revel.Validation, email string) *revel.ValidationResult {
	return v.Check(email,
		revel.Required{},
		revel.Match{emailRegexp},
	)
}

//
// func (u *User) FindByEmail(email string) {
// 	results, err := Dbm.Select(User{},
// 		`select * from User where Email = ?`, email)
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(results)
// }

package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"api/security"
)

type User struct {
	UserId   uint32 `gorm:"primary_key;auto_increment" json:"userId"`
	Name     string `gorm:"size:50;not null;unique" json:"name"`
	Age      int    `gorm:"type:integer" json:"age"`
	Username string `gorm:"size:20;not null;unique" json:"username"`
	Password string `gorm:"size:60;not null;" json:"password"`
	Family   Family `json:"family"`
	FamilyID uint32 `gorm:"foreigkey:FamilyId;not null" json:"familyid"`
	Role     string `gorm:"default:'normal';not null" json:"role"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) AddFamily(family Family) {
	u.Family = family
}

func (u *User) Prepare() {
	u.UserId = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Role = html.EscapeString(strings.TrimSpace(u.Role))
}

func (u *User) Validate(action string) error {
	switch action {
	case "update":
		if u.Username == "" {
			return errors.New("Username is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil

	case "login":
		if u.Username == "" {
			return errors.New("Username is required")
		}
		if u.Password == "" {
			fmt.Printf("\n\n\n\n bosta \n\n\n\n")
			return errors.New("Password is required")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Username is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil
	}

}

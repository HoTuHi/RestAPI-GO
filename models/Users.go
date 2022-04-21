package models

import "time"

type User struct {
	// what returns in body of response
	ID       uint `json:"id" gorm:"autoIncrement"`
	Role     bool `json:"role"`
	Name     string
	Email    string    `gorm:"unique"`
	Password []byte    `json:"-"`
	CreateAt time.Time `sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

func (u *User) Prepare() {
	u.ID = 0
	u.Role = false
	u.Password = []byte("")
	u.Name = ""
	u.Email = ""
	u.CreateAt = time.Now()
}

func (u *User) Validate() error {
	return nil
}

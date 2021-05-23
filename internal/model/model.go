//this file include the structures with member variable to be used over project with db naming too
package model

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AccessDetails struct {
	UserID   int64
	Username string
}

type Username struct {
	Username string `db:"username"`
	UserID   int    `db:"user_id"`
}

type UserDetails struct {
	FirstName         string `db:"firstName" json:"firstName"`
	LastName          string `db:"lastName" json:"lastName"`
	Username          string `db:"username" json:"username"`
	ContactNumber     string `db:"contactNumber" json:"contactNumber"`
	EmailId           string `db:"emailId" json:"emailId"`
	CurrentLocation   string `db:"currentLocation" json:"currentLocation"`
	Bio               string `db:"bio" json:"bio"`
	Designation_stand string `db:"Designation_stand" json:"Designation_stand"`
}

type Achievement struct {
	Id          int    `db:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
}

type Section struct {
	Name         string        `db:"name" json:"name" binding:"required"`
	Id           int           `db:"id" `
	Achievements []Achievement `json:"achievements"`
}

type Sections_achievements struct {
	Id int `db:"achievement_id"`
}

type User struct {
	Id       int64  `db:"id"`
	Fname    string `db:"fname" `
	Lname    string `db:"lname" `
	Email    string `db:"email"`
	Password string `db:"password"`
}

type SingUp struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	EmailId   string `json:"emailId" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AddData struct {
	UserDetails UserDetails `json:"userDetails"`
	Sections    []Section   `json:"sections"`
}

type RegisterUsername struct {
	Username string `json:"username"`
}

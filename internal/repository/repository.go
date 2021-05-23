//query for database are directly written here
package repository

import (
	"errors"
	"fmt"

	"github.com/02amanag/p-02/db"
	"github.com/02amanag/p-02/internal/model"
)

type RepositoryStruct struct{}

func NewRepositoryStruct() *RepositoryStruct {
	return &RepositoryStruct{}
}

func GetUserID(username string) int {
	var user_id int
	_ = db.GetDB().QueryRow("select user_id from usernames where username = $1", username).Scan(&user_id)
	defer db.CloseDB()
	fmt.Println(user_id)
	return user_id
}

func GetAchievements(Id int) (achievements []model.Achievement, err error) {
	_, err = db.GetDB().Select(&achievements, "SELECT * FROM achievements WHERE id=$1", Id)
	defer db.CloseDB()
	if err != nil {
		return nil, err
	}
	return achievements, nil
}

func GetUserDetails(username string) (userDetails []model.UserDetails, err error) {
	_, err = db.GetDB().Select(&userDetails, "SELECT * FROM USERDETAILS WHERE USERNAME=$1", username)
	defer db.CloseDB()
	if err != nil {
		return nil, err
	}
	return userDetails, nil
}

func GetSections(username string) (sections []model.Section, err error) {
	_, err = db.GetDB().Select(&sections, "select name,id from sections where username=$1", username)
	defer db.CloseDB()
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func GetAchievementId(Id int) (achievement_id []model.Sections_achievements, err error) {
	_, err = db.GetDB().Select(&achievement_id, "SELECT achievement_id FROM sections_achievements WHERE section_id=$1", Id)
	defer db.CloseDB()
	if err != nil {
		return nil, err
	}
	return achievement_id, nil
}

func (r *RepositoryStruct) GetUsers(email string) (user model.User) {
	err := db.GetDB().SelectOne(&user, "SELECT * FROM users where email=$1", email)
	defer db.CloseDB()
	if err != nil {
		return model.User{}
	}
	return user
}

func (r *RepositoryStruct) GetUsername(userId int) string {
	var username string
	_ = db.GetDB().QueryRow("select username from usernames where user_id = $1", userId).Scan(&username)
	defer db.CloseDB()
	return username
}

func (r *RepositoryStruct) CreateUser(email, password, fname, lname string) error {
	_, err := db.GetDB().Query("insert into users ( fname , lname, email , password) values ($1,$2,$3,$4);", fname, lname, email, password)
	defer db.CloseDB()
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryStruct) InsertUserdetails(username string, userDetails model.UserDetails) error {
	// check if userdetails present or not
	presentUserDetails, err := GetUserDetails(username)
	if err != nil {
		return err
	}
	if presentUserDetails[0].EmailId == userDetails.EmailId {
		return errors.New("userdetails alredy present")
	}
	_, err = db.GetDB().Query("insert into userdetails ( username , firstname , lastname,bio , designation_stand , contactnumber ,emailid, currentlocation) values ($1,$2,$3,$4,$5,$6,$7,$8);",
		username, userDetails.FirstName, userDetails.LastName, userDetails.Bio, userDetails.Designation_stand, userDetails.ContactNumber, userDetails.EmailId, userDetails.CurrentLocation)
	defer db.CloseDB()
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryStruct) InsertSection(userName, sectionName string) int {
	var id int
	_ = db.GetDB().QueryRow("insert into sections ( username , name) values ($1,$2) returning id;", userName, sectionName).Scan(&id)
	defer db.CloseDB()
	return id
}

func (r *RepositoryStruct) InsertAchievement(name, description, link string) int {
	var id int
	_ = db.GetDB().QueryRow("insert into achievements (name,description,link ) values ($1,$2,$3) returning id;", name, description, link).Scan(&id)
	defer db.CloseDB()
	return id
}

func (r *RepositoryStruct) InsertSectionAchievement(sectionsId, achievementId int) error {
	_, err := db.GetDB().Query("insert into sections_achievements (section_id , achievement_id) values ($1,$2);", sectionsId, achievementId)
	defer db.CloseDB()
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryStruct) AddUsername(username string, userId int64) error {
	_, err := db.GetDB().Query("insert into usernames values ($1,$2);", username, userId)
	defer db.CloseDB()
	if err != nil {
		return err
	}
	return nil
}

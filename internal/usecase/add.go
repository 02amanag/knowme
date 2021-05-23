package usecase

import (
	"github.com/02amanag/p-02/internal/model"
)

func (u *UsecaseStruct) AddData(data model.AddData, username string) error {

	err := u.repo.InsertUserdetails(username, data.UserDetails)
	if err != nil {
		if err.Error() != "userdetails alredy present" {
			return err
		}
	}

	for _, i := range data.Sections {

		//insert a section get id
		sectionId := u.repo.InsertSection(username, i.Name)
		for _, j := range i.Achievements {

			// insert achievement get id
			achievementId := u.repo.InsertAchievement(j.Name, j.Description, j.Link)

			//insert section_id and achievement_id
			err := u.repo.InsertSectionAchievement(sectionId, achievementId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *UsecaseStruct) AddUsername(username string, userId int64) error {
	err := u.repo.AddUsername(username, userId)
	if err != nil {
		if err.Error() != "username alredy present" {
			return err
		}
	}
	return nil
}

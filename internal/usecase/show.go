package usecase

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/02amanag/p-02/internal/model"
	"github.com/02amanag/p-02/internal/repository"
)

type HtmlBody struct {
	Sections       []model.Section
	UserDetails    []model.UserDetails
	ResumeLink     string
	ProfilePicture string
	QrLink         string
}

func (u *UsecaseStruct) GetProfile(username string) *HtmlBody {
	// Adding all sections of user to list
	sections, err := u.GetSections(username)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Adding UserDetails of user to list
	userDetails, err := u.GetUserDetails(username)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Adding path of user profile picture
	imgBase64Str := u.GetProfilePicture(username)

	u.GenerateQr(username)

	return &HtmlBody{
		Sections:       sections,
		ResumeLink:     "/show/downloadresume/" + username,
		UserDetails:    userDetails,
		ProfilePicture: imgBase64Str,
		QrLink:         "/show/downloadqr/" + username,
	}
}

func (u *UsecaseStruct) GetUserDetails(username string) ([]model.UserDetails, error) {
	return repository.GetUserDetails(username)
}

func (u *UsecaseStruct) GetSections(username string) ([]model.Section, error) {
	onlySection, err := repository.GetSections(username)
	if err != nil {
		return nil, err
	}
	var section []model.Section
	for _, p := range onlySection {
		achievementId, err := repository.GetAchievementId(p.Id)
		if err != nil {
			return nil, err
		}
		for _, q := range achievementId {
			achievement, err := repository.GetAchievements(q.Id)
			if err != nil {
				return nil, err
			}
			p.Achievements = append(p.Achievements, achievement[0])
		}
		section = append(section, p)
	}

	return section, nil
}

func (u *UsecaseStruct) GetProfilePicture(username string) (path string) {
	path, err := filepath.Abs("file/profilePicture")
	if err != nil {
		log.Println(err)
	}

	imgFile, err := os.Open(path + "/" + username + ".jpeg")

	if err != nil {
		fmt.Println(err)
	}

	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64Str
}

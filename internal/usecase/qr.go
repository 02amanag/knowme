package usecase

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (u *UsecaseStruct) GenerateQr(username string) error {
	prPathLink, err := u.config.GetConfig("PROFILE_PATH_LINK")
	if err != nil {
		log.Panicln("error - ", err)
	}

	qrSize, err := u.config.GetConfig("QR_SIZE")
	if err != nil {
		log.Panicln("error - ", err)
	}

	link := prPathLink + "/show/profile" + username //get from env

	qrCreator, err := u.config.GetConfig("QR_CREATOR")
	if err != nil {
		log.Panicln("error - ", err)
	}

	url := qrCreator + "data=" + link + "&size=" + qrSize
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// bytearray to image
	img, _, err := image.Decode(bytes.NewReader(responseData))
	if err != nil {
		return err

	}
	path, err := filepath.Abs("file/qr")
	if err != nil {
		return err
	}
	out, _ := os.Create(path + "/qr.png")
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}

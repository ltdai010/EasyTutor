package storageusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"mime/multipart"
)

func SaveFile(file multipart.File) (string, error) {
	path, err := models.AddImage(file)
	if err != nil {
		return "", data.ErrUnknown
	}
	return "https://storage.googleapis.com/bla-bla-star.appspot.com/" + path, data.Success
}

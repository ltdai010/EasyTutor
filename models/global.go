package models

import (
	"EasyTutor/consts"
	"EasyTutor/drivers"
	"context"
	"io"
	"mime/multipart"
)

func AddImage(file multipart.File) (string, error) {
	ref, _, err := drivers.GetDriver().GetCloudStore().Collection(consts.IMAGELINK).Add(context.Background(), map[string]string{
		"Image" : "link",
	})
	if err != nil {
		return "", err
	}
	wc := drivers.GetDriver().GetStorage().Object(ref.ID).NewWriter(context.Background())
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	return ref.ID, nil
}

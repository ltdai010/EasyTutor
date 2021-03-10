package requests

import "EasyTutor/data/data"

type CommentPost struct {
	data.CommentInfo
}

type CommentPut struct {
	data.CommentInfo
}
package responses

import "EasyTutor/data/data"

type Comment struct {
	ID string `json:"id"`
	data.Comment
}

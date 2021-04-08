package responses

import "EasyTutor/data/data"

type Comment struct {
	ID string `json:"id"`
	data.Comment
}

type TeacherComment struct {
	TeacherID           string    `json:"teacher_id"`
	Name                string    `json:"name"`
	ListUnActiveComment []*Comment `json:"list_un_active_comment"`
}

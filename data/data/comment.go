package data

type CommentInfo struct {
	Comment     string `json:"comment"`
	Recommended bool   `json:"recommended"`
}

type Comment struct {
	CommentInfo
	Username   string `json:"username"`
	TeacherID  string `json:"teacher_id"`
	CreateTime int64  `json:"create_time"`
}

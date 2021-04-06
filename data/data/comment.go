package data

type CommentInfo struct {
	Comment     string `json:"comment"`
	Stars 	    int   `json:"stars"`
}

type Comment struct {
	CommentInfo
	Username   string `json:"username"`
	TeacherID  string `json:"teacher_id"`
	CreateTime int64  `json:"create_time"`
}

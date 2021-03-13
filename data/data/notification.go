package data

type Notification struct {
	NotificationInfo
	CreateTime int64	`json:"create_time"`
}

type NotificationInfo struct {
	Username   string      `json:"username"`
	UserType   string      `json:"user_type"`
	NotifyType NotifyType  `json:"notify_type"`
	Message    interface{} `json:"message"`
}

type NotifyType string

const (
	CommentType = "comment"
	PostOffer   = "post_offer"
	AcceptOffer = "accept_offer"
)

type RequestNotify struct {
	Message      string `json:"message"`
	RequestID    string `json:"request_id"`
	RequestTitle string `json:"request_title"`
}

type OfferNotify struct {
	Message     string `json:"message"`
	OfferID     string `json:"offer_id"`
	TeacherName string `json:"teacher_name"`
}

type CommentNotify struct {
	Message         string `json:"message"`
	CommentID       string `json:"comment_id"`
	UserDisplayName string `json:"user_display_name"`
}

type ForgotPassword struct {
	CheckCode	string	`json:"check_code"`
}

type GetInMessage struct {
	Token string `json:"token"`
}

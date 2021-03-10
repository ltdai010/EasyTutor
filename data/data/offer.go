package data

type Offer struct {
	CreateTime int64 `json:"create_time"`
	TeacherID  string `json:"teacher_id"`
	RequestID  string `json:"request_id"`
	OfferInfo
}

type OfferInfo struct {
	OfferPrice int64  `json:"offer_price"`
	Contact    string `json:"contact"`
	Message    string `json:"message"`
}

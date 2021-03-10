package data

type Request struct {
	RequestInfo
	Username   string	`json:"username"`
	AcceptOffer string  `json:"accept_offer"`
	CreateTime int64	`json:"create_time"`
	Schedule
}

type RequestInfo struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Subject       string  `json:"subject"`
	Price         int64   `json:"price"`
	PhoneNumber   string  `json:"phone_number"`
	Location      int64   `json:"location"`
	Method        string  `json:"method"`
	Duration      int64   `json:"duration"`
	ExactLocation string  `json:"exact_location"`
	Gender		  string  `json:"gender"`
	Graduation    string  `json:"graduation"`
}

func DataRequestIsValid(post RequestInfo) bool {
	if NewGraduation(post.Graduation) == "" || NewGender(post.Gender) == "" ||
		NewMethod(post.Method) == "" || NewSubject(post.Subject) == "" {
		return false
	}
	return true
}
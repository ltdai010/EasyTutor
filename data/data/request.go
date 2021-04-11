package data

type Request struct {
	RequestInfo
	Username    string `json:"username"`
	AcceptOffer string `json:"accept_offer"`
	CreateTime  int64  `json:"create_time"`
	Active      bool   `json:"active"`
	Closed      bool   `json:"closed"`
	Schedule
}

type RequestInfo struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Subject         string `json:"subject"`
	Price           int64  `json:"price"`
	Time		    int    `json:"time"`
	TimesPerWeek	int	   `json:"times_per_week"`
	PhoneNumber     string `json:"phone_number"`
	Location        int    `json:"location"`
	NumberOfStudent int    `json:"number_of_student"`
	Method          string `json:"method"`
	ExactLocation   string `json:"exact_location"`
	Gender          string `json:"gender"`
}

func DataRequestIsValid(post RequestInfo) bool {
	if NewGender(post.Gender) == "" ||
		NewMethod(post.Method) == "" || NewSubject(post.Subject) == "" || post.NumberOfStudent == 0{
		return false
	}
	return true
}

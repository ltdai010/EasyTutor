package data

import "errors"

var (
	Success            = errors.New("Success")
	UnSuccess          = errors.New("UnSuccess")
	ErrUnknown         = errors.New("Unknown")
	BadRequest         = errors.New("Bad Request")
	ErrNotInt          = errors.New("Err Int")
	SessionExpired     = errors.New("Session Expired")
	Existed            = errors.New("Existed")
	NotAdmin           = errors.New("Not Admin")
	NotPermission      = errors.New("Not Permission")
	ErrLogin           = errors.New("Login Error")
	ErrSystem          = errors.New("System Error")
	NotExisted         = errors.New("Data Not Existed !")
	ErrChangePass      = errors.New("Change Password Error !")
	NotMore            = errors.New("No More Data")
	ExistedPosUser     = errors.New("Existed Pos User")
	LimitCharacter     = errors.New("Limit character")
	CannotEmpty        = errors.New("Cannot empty")
	TaskFinish         = errors.New("Task Finish")
	CanNotReworkReport = errors.New("Can not rework report")
	TaskNotSend        = errors.New("The quest has not been sent yet")

	MapDescription = map[error]string{
		Success:            "Success!",
		ErrUnknown:         "Unknown error!",
		BadRequest:         "Bad Request!",
		ErrNotInt:          "Param not int!",
		UnSuccess:          "Unsuccess!",
		SessionExpired:     "SessionExpired!",
		Existed:            "Existed !",
		NotAdmin:           "Not Admin !",
		NotPermission:      "Not Permission !",
		ErrLogin:           "Wrong username, password. ",
		ErrSystem:          "The system is having problems.",
		NotExisted:         "Data Not Existed!",
		ErrChangePass:      "Wrong username, password.",
		NotMore:            "No more data",
		ExistedPosUser:     "Existed position organization",
		TaskFinish:         "The mission has ended",
		CanNotReworkReport: "The report is being reviewed by the superior, not being revised",
		TaskNotSend:        "The quest has not been sent yet",
	}
	MapErrorCode = map[error]int{
		Success:            200,
		UnSuccess:          201,
		ErrNotInt:          302,
		SessionExpired:     303,
		NotExisted:         304,
		Existed:            305,
		ErrChangePass:      306,
		NotAdmin:           307,
		NotPermission:      308,
		NotMore:            309,
		ExistedPosUser:     310,
		CannotEmpty:        311,
		LimitCharacter:     312,
		TaskFinish:         313,
		CanNotReworkReport: 314,
		TaskNotSend:        315,
		BadRequest:         400,
		ErrUnknown:         401,
		ErrLogin:           402,
		ErrSystem:          403,
	}
)

// Returns a error.
// swagger:response Err
type Err struct {
	// code error
	Code int `json:"code"`
	// description error
	Message string `json:"message"`
}

func NewErr(err error) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: MapDescription[err],
	}
}

package data

type UserInfo struct {
	DisplayName     string   `json:"display_name"`
	PhoneNumber     string	 `json:"phone_number"`
	Email           string   `json:"email"`
}

type User struct {
	UserInfo
	FavoriteTeacher []string `json:"favorite_teacher"`
	LoginInfo
}

type LoginInfo struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}

package view

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

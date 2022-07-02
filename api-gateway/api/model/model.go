package model

type RegisterResponseModel struct{
	UserId string
	AccessToken string
	RefreshToken string
}

type JwtRequestModel struct {
	Token string `json:"token"`
}
type ResponseError struct {
	Error SeverError `json:"error"`
}
type SeverError struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
type User struct {
    ID   int
    Name string
    Role string
}
type Users []User
package model

type UserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
	Number   string `json:"number"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type VerifyOTPs struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Password struct {
	Old     string `json:"old_password" validate:"required"`
	New     string `json:"new_password" validate:"required"`
	Confirm string `json:"confirm_password" validate:"required"`
}

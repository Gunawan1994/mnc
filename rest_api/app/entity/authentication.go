package entity

type ReqRegister struct {
	FirstName   *string `json:"first_name" validate:"required"`
	LastName    *string `json:"last_name" validate:"required"`
	PhoneNumber *string `json:"phone_number" validate:"required"`
	Address     *string `json:"address" validate:"required"`
	Pin         *string `json:"pin" validate:"required"`
}

type ReqLogin struct {
	PhoneNumber *string `json:"phone_number" validate:"required"`
	Pin         *string `json:"pin" validate:"required"`
}

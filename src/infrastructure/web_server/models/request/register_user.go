package request

type RegisterUser struct {
	Username string `json:"username" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Rol      int    `json:"rol" validate:"required"`
}

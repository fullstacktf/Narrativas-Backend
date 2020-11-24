package interfaces

type IUserService interface {
	SignUp(username string, password string, email string) error
	SignIn(username string, password string) (string, error)
}

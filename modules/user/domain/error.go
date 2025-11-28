package userdomain

const (
	ErrUserNotFound                = "user not found"
	ErrUserEmailRequired           = "user email is required"
	ErrUserPasswordRequired        = "user password is required"
	ErrUserLastNameRequired        = "user last name is required"
	ErrUserFirstNameRequired       = "user first name is required"
	ErrUserPhoneRequired           = "user phone is required"
	ErrUserRoleRequired            = "user role is required"
	ErrUserEmailAlreadyExists      = "user email already exists"
	ErrCannotRegisterUser          = "cannot register user"
	ErrUserEmailAlreadyDeleted     = "user email already deleted"
	ErrUserEmailAndPasswordInvalid = "user email and password are invalid"
)
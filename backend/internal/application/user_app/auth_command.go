package user_app

type AuthUserCommand struct {
	email         string
	password      string
	hmacSecretKey string
}

func NewAuthUserCommand(email string, password string, hmacSecretKey string) *AuthUserCommand {
	return &AuthUserCommand{
		email:         email,
		password:      password,
		hmacSecretKey: hmacSecretKey,
	}
}

func (a *AuthUserCommand) Email() string         { return a.email }
func (a *AuthUserCommand) Password() string      { return a.password }
func (a *AuthUserCommand) HmacSecretKey() string { return a.hmacSecretKey }

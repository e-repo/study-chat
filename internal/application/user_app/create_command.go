package user_app

type CreateUserCommand struct {
	fistName string
	email    string
	password string
}

func NewCreateUserCommand(
	firstName string,
	email string,
	password string,
) (*CreateUserCommand, error) {
	return &CreateUserCommand{
		fistName: firstName,
		email:    email,
		password: password,
	}, nil
}

func (c *CreateUserCommand) FirstName() string { return c.fistName }
func (c *CreateUserCommand) Email() string     { return c.email }
func (c *CreateUserCommand) Password() string  { return c.password }

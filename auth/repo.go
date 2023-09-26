package auth

type IRepo interface {
	createUser(name, email, password string) (*User, error)
	getUser(email string) *User
}

func NewRepo() IRepo {
	return &repo{users: map[string]*User{}}
}

type repo struct {
	users map[string]*User
}

// createUser implements IRepo.
func (r *repo) createUser(name string, email string, password string) (*User, error) {
	newUser := &User{Name: name, Email: email, password: password}
	r.users[email] = newUser
	return newUser, nil
}

// getUser implements IRepo.
func (r *repo) getUser(email string) *User {
	return r.users[email]
}

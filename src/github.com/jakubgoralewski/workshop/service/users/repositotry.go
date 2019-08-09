package users

var repositoryInstance *Repository

type Repository struct {
	Users []*User
}

func newRepository() *Repository {
	return &Repository{
		Users: []*User{},
	}
}

func RepositoryInstance() *Repository {
	if repositoryInstance == nil {
		repositoryInstance = newRepository()
	}
	return repositoryInstance
}

func (r *Repository) Add(user *User) {
	r.Users = append(r.Users, user)
}

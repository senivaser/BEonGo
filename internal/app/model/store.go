package model

type Store struct {
	User *UserRepository
}

func NewStore() (*Store, []error) {

	var errors []error

	var User *UserRepository = nil
	if u, ue := NewUser(); ue == nil {
		User = u
	} else {
		errors = append(errors, ue)
	}

	return &Store{
		User: User,
	}, errors

}

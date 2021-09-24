package model

type Store struct {
	User *UserRepository
}

func NewStore(config *Config) (*Store, []error) {

	var errors []error

	var User *UserRepository = nil
	if u, ue := NewUser(config); ue == nil {
		User = u
	} else {
		errors = append(errors, ue)
	}

	return &Store{
		User: User,
	}, errors

}

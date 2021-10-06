package model

type Store struct {
	User *UserRepository
}

func NewStore(config *Config) (*Store, map[string]error) {

	errors := make(map[string]error)

	User, ue := NewUser(config)
	if ue != nil {
		errors["User"] = ue
	}

	return &Store{
		User: User,
	}, errors

}

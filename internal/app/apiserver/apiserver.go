package apiserver

type APIServer struct {
	config *Config
}

func New(config *Config) *APIServer {
	return &APIServer{
		config,
	}
}

func (s *APIServer) Start() error {
	return nil
}

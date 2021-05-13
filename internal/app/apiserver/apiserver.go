package apiserver

type APIServer struct {
	config *Config
}

// function return instance APIServer struct
func New(config *Config) *APIServer {
	return &APIServer{
		config: config, // Initialization config file "toml" from config.go
	}
}


//func starting our application. Initialize DB, Logger
func (s APIServer) Start() error {
	return nil
}

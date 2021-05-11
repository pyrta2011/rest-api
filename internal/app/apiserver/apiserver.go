package apiserver

type APIServer struct {

}

// function return instance APIServer struct
func New() *APIServer {
	return &APIServer{}
}


//func starting our application. Initialize DB, Logger
func (s APIServer) Start() error {
	return nil
}

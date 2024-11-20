package refy

type Server struct {
	Register *Register
	Handlers
}

func NewServer() *Server {
	return &Server{
		Register: NewRegister(),
	}
}

func (s *Server) Pub(id string) {

}

func (s *Server) Sub(id string) {

}

func (s *Server) On(id string)

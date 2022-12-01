package config

type Server struct {
	URL string
}

func NewServerConfig() *Server {
	return &Server{URL: "127.0.0.1:5053"}
}

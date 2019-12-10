package handlers

import (
	"github.com/zumatande/cleanpartner/services/search"
)

// Search is handler for Search service
type Search struct {
	Logger *zerolog.Logger
	Service search.Service
}

func (s *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// decode request and pass domain request object to Search service
	s.Service.Search()
}

// NewSearchHandler is a Search handler constructor
func NewSearchHandler(svc search.Service, w ...io.Writer) Search {
	var l zerolog.Logger
	if w == nil {
		l = zerolog.New(os.Stderr)
	} else {
		l = zerolog.New(w[0])
	}

	return &Search{
		Logger: l,
		Service: svc,
	}
}

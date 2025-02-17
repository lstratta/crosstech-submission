package server

import "github.com/go-pg/pg/v10"

func (s *Server) setupDB() error {
	c := s.conf
	opt, err := pg.ParseURL(c.DatabaseURI)
	if err != nil {
		return err
	}

	s.db = pg.Connect(opt)

	return nil
}

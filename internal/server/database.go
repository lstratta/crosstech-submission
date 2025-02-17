package server

import "github.com/go-pg/pg/v10"

func (s *Server) setupDB() {
	c := s.conf
	pgConf := &pg.Options{
		Addr:     c.DatabaseURI,
		User:     c.PgUser,
		Password: c.PgPassword,
		Database: c.DbName,
	}
	s.db = pg.Connect(pgConf)
}

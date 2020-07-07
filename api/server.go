package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Server struct to hold helper utils
type Server struct {
	l  *logrus.Logger
	e  *gin.Engine
	db *pg.DB
}

// Run is entry point to application
func (s *Server) Run() error {

	s.db = pg.Connect(&pg.Options{
		Database: "wallet",
		Addr:     os.Getenv("WALLET_DB_ADDR"),
		User:     os.Getenv("WALLET_DB_USER"),
		Password: os.Getenv("WALLET_DB_PASSWORD"),
	})

	s.e = gin.Default()
	s.setRoutes(s.e.Group("/v1"))

	if err := s.e.Run(os.Getenv("WALLET_HOST_ADDR")); err != nil {
		return errors.Wrap(err, "cannot run on server")
	}

	return nil
}

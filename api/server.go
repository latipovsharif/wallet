package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
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
	s.e = gin.Default()

	return nil
}

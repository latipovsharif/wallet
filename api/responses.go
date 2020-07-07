package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) badRequest(c *gin.Context, err error) {
	s.l.Warnf("bad request %v", err)
	c.JSON(http.StatusBadRequest, err.Error())
}

func (s *Server) errorWithMessage(c *gin.Context, message string) {
	s.l.Warn("bad request: %v", message)
	c.JSON(http.StatusBadRequest, message)
}

func (s *Server) internalError(c *gin.Context, err error) {
	s.l.Errorf("internal server error: %v", err)

	// hide error internals from customer
	c.JSON(http.StatusInternalServerError, "internal server error occurred, please contact service administrator")
}

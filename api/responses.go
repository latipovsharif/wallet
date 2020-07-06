package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) badRequest(c *gin.Context, err error) {
	s.l.Infof("bad request %v", err)
	c.JSON(http.StatusBadRequest, err)
}

func (s *Server) internalError(c *gin.Context, err error) {
	s.l.Errorf("internal server error: %v", err)

	// hide error internals from customer
	c.JSON(http.StatusInternalServerError, "internal server error occurred, please contact service administrator")
}

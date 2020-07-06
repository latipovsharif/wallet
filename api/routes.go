package api

import "github.com/gin-gonic/gin"

func (s *Server) setRoutes(v1 *gin.RouterGroup) {
	wallet := v1.Group("/wallet")
	wallet.POST("/", s.createWallet)
}

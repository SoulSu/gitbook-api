package server

import "github.com/gin-gonic/gin"

func (s *SrvHandle) routeInit() {
	s.router.v1 = s.engine.Group("/v1")
	s.router.v1.GET("/", s.GetAuthorInfo)
}

func (s *SrvHandle) GetAuthorInfo(ctx *gin.Context) {

}

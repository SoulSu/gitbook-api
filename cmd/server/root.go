package server

import (
	"github.com/gin-gonic/gin"
)

type SrvHandle struct {
	engine *gin.Engine
	router struct {
		v1 *gin.RouterGroup
	}
}

func (s *SrvHandle) recovery(c *gin.Context) {
	//return func(c *Context) {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			if logger != nil {
	//				stack := stack(3)
	//				httprequest, _ := httputil.DumpRequest(c.Request, false)
	//				logger.Printf("[Recovery] %s panic recovered:\n%s\n%s\n%s%s", timeFormat(time.Now()), string(httprequest), err, stack, reset)
	//			}
	//			c.AbortWithStatus(http.StatusInternalServerError)
	//		}
	//	}()
	//	c.Next()
	//}
}

func (s *SrvHandle) requestBefore(c *gin.Context) {

}

func (s *SrvHandle) injectSoul() {
	s.engine.Use(s.recovery, s.requestBefore)
}

func RunServer(port string) error {

	srvHandle := new(SrvHandle)
	srvHandle.engine = gin.New()
	srvHandle.injectSoul()
	srvHandle.routeInit()

	return srvHandle.engine.Run(":" + port)
}

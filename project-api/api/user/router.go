package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"my_project/project-api/api/midd"
	"my_project/project-api/api/rpc"
	"my_project/project-api/router"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	//初始化grpc的客户端连接
	rpc.InitRpcUserClient()
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/register", h.register)
	r.POST("/project/login", h.login)
	org := r.Group("/project/organization")
	org.Use(midd.TokenVerify1())
	org.POST("/_getOrgList", h.myOrgList)
}

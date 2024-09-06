package router

import (
	"github.com/gin-gonic/gin"
	"qinyu-wf/api/user"
)

// routerInterface 定义接口，要求实现 InitRouter 方法
type routerInterface interface {
	InitRouter(engine *gin.Engine)
}

var routerList = []routerInterface{
	new(user.Router),
}

func Init() {
	router := gin.Default()
	for i := range routerList {
		routerList[i].InitRouter(router)
	}
	err := router.Run(":8082")
	if err != nil {
		panic(err)
	}
}

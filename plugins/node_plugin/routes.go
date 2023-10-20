package node_plugin

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	g := r.Group("/node")
	{
		g.GET("/index", IndexAction)

	}

	admin := g.Group("/admin")
	{
		admin.POST("/create", AdminCreateAction)
	}
}

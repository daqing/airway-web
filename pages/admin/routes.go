package admin

import (
	"github.com/daqing/airway/pages/admin/dashboard_page"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	g := r.Group("/admin")
	{
		dashboard_page.Routes(g)
	}
}

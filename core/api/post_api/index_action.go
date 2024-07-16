package post_api

import (
	"github.com/daqing/airway/app/models"
	"github.com/daqing/airway/lib/api_resp"
	"github.com/daqing/airway/lib/sql_orm"
	"github.com/gin-gonic/gin"
)

func IndexAction(c *gin.Context) {
	list, err := sql_orm.ListResp[models.Post, PostResp]()
	if err != nil {
		api_resp.LogError(c, err)
		return
	}

	api_resp.OK(c, gin.H{"list": list})
}

package user_plugin

import (
	"github.com/daqing/airway/lib/repo"
	"github.com/daqing/airway/lib/resp"
	"github.com/daqing/airway/lib/utils"
	"github.com/gin-gonic/gin"
)

func AdminIndexAction(c *gin.Context) {
	if !CheckAdmin(c.GetHeader("X-Auth-Token")) {
		utils.LogInvalidAdmin(c)
		return
	}

	list, err := repo.ListResp[User, UserResp]()
	if err != nil {
		utils.LogError(c, err)
		return
	}

	resp.OK(c, gin.H{"list": list})
}

package post_plugin

import (
	"github.com/daqing/airway/lib/resp"
	"github.com/daqing/airway/lib/utils"
	"github.com/daqing/airway/plugins/action_plugin"
	"github.com/daqing/airway/plugins/user_plugin"
	"github.com/gin-gonic/gin"
)

type ToggleLikeParams struct {
	PostId int64 `form:"id"`
}

func ToggleLikeAction(c *gin.Context) {
	var p ToggleLikeParams

	if err := c.ShouldBind(&p); err != nil {
		utils.LogError(c, err)
		return
	}

	user := user_plugin.UserFromAuthToken(c.GetHeader("X-Auth-Token"))
	if user == nil {
		utils.LogInvalidUserId(c)
		return
	}

	count, err := TogglePostAction(p.PostId, user.Id, action_plugin.ActionLike)
	if err != nil {
		utils.LogError(c, err)
		return
	}

	resp.OK(c, gin.H{"id": p.PostId, "count": count})
}

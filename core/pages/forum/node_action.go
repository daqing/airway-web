package forum

import (
	"fmt"
	"time"

	"github.com/daqing/airway/core/api/node_api"
	"github.com/daqing/airway/core/api/post_api"
	"github.com/daqing/airway/lib/page_resp"
	"github.com/daqing/airway/lib/repo"
	"github.com/daqing/airway/lib/utils"
	"github.com/gin-gonic/gin"
)

func NodeAction(c *gin.Context) {
	nodeKey := c.Param("key")

	node, err := repo.FindRow[node_api.Node](
		[]string{"id", "name"},
		[]repo.KVPair{
			repo.KV("key", nodeKey),
		},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	posts, err := repo.Find[post_api.Post](
		[]string{"id", "title"},
		[]repo.KVPair{
			repo.KV("node_id", node.Id),
			repo.KV("cat", "forum"),
		},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	postsShow := []*PostItemIndex{}

	for _, post := range posts {
		url := fmt.Sprintf("/forum/post/%d", post.Id)

		if len(post.CustomPath) > 0 {
			url = fmt.Sprintf("/forum/post/%s", post.CustomPath)
		}

		postsShow = append(postsShow,
			&PostItemIndex{
				Id:    post.Id,
				Title: post.Title,
				Url:   url,
				Date:  post.CreatedAt.Format("2006-01-02 15:04"),
			},
		)
	}

	rootPath := utils.PathPrefix("forum")

	nodes, err := repo.Find[node_api.Node](
		[]string{"id", "name", "key"},
		[]repo.KVPair{
			repo.KV("place", "forum"),
		},
	)

	if err != nil {
		page_resp.Error(c, err)
		return
	}

	nodeItems := []*NodeItem{}

	for _, node := range nodes {
		nodeItems = append(nodeItems,
			&NodeItem{
				Name: node.Name,
				URL:  rootPath + "node/" + node.Key,
			})
	}

	data := map[string]any{
		"Title":    ForumTitle(),
		"Tagline":  ForumTagline(),
		"Year":     time.Now().Year(),
		"RootPath": rootPath,
		"Node":     node,
		"Nodes":    nodeItems,
		"Posts":    postsShow,
	}

	page_resp.Page(c, "core", "forum!", "node", data)
}

package user_api

import "github.com/daqing/airway/lib/utils"

type UserResp struct {
	Id        int64
	Nickname  string
	Username  string
	ApiToken  string
	Role      int
	CreatedAt utils.Timestamp
	UpdatedAt utils.Timestamp
}

func (ur UserResp) Fields() []string {
	return []string{"id", "username", "nickname", "role", "api_token"}
}

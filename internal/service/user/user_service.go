package user

import (
	"github.com/cham/elk/internal/dao/user"
	muser "github.com/cham/elk/internal/model/user"
	"github.com/cham/elk/pkg/ov"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, query ov.UserQuery) ([]muser.User, int64, error) {
	return user.List(c, query)
}

func Create(c *gin.Context, data ov.User) (*muser.User, error) {
	return user.Create(c, &data)
}

// get user info
func Get(c *gin.Context, qs map[string]any) (*muser.User, error) {
	return user.Get(c, qs)
}

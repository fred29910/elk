package user

import (
	"github.com/cham/elk/internal/dao/user"
	muser "github.com/cham/elk/internal/model/user"
	"github.com/cham/elk/internal/ov"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, query ov.UserQuery) ([]muser.User, int64, error) {
	return user.List(c, query)
}

package core

import "context"

type contextUserKey struct{}

func FromContextUserInfo(c context.Context) *ContextUserInfo {
	userInfo, ok := c.Value(contextUserKey{}).(*ContextUserInfo)
	if !ok {
		return nil
	}
	return userInfo
}

type ContextUserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func StoreContextUserInfo(c context.Context, userInfo *ContextUserInfo) context.Context {
	return context.WithValue(c, contextUserKey{}, userInfo)
}

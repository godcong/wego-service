package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserActivityList 我的活动
func UserActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		page := model.PageUserActivity(model.ParsePaginate(ctx.Request.URL.Query()))
		e := page.PageWhere(&model.UserActivity{
			UserID: user.ID,
		})
		if e != nil {
			log.Error(e)
			Error(ctx, e)
			return
		}
		Success(ctx, page)
	}
}

// UserActivityJoin 活动申请
func UserActivityJoin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		id := ctx.Param("id")
		act := model.NewActivity(id)
		b, e := act.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("activity not found"))
		}
		ua := model.UserActivity{
			ActivityID: id,
			UserID:     user.ID,
		}
		i, e := model.Insert(nil, &ua)
		if e != nil || i == 0 {
			log.Error(e, b)
			Error(ctx, xerrors.New("user activity insert error"))
		}
		Success(ctx, ua)
	}
}

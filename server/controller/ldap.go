package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
)

func SyncLDAPUsers(c *gin.Context) {
	msg, err := svc.SyncLDAPUsers()
	if err != nil {
		logrus.Errorf("同步LDAP用户: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithMsg(msg).Response(util.GinRespf(c))
}

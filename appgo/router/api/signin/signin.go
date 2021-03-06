package signin

import (
	"errors"
	"time"

	"github.com/i2eco/egoshop/appgo/service"

	"github.com/gin-gonic/gin"
	"github.com/i2eco/egoshop/appgo/model/mysql"
	"github.com/i2eco/egoshop/appgo/pkg/base"
	"github.com/i2eco/egoshop/appgo/pkg/code"
	"github.com/i2eco/egoshop/appgo/pkg/mus"
	"github.com/i2eco/egoshop/appgo/router/mdw"
	"github.com/jinzhu/gorm"
)

// todo 防双击事件
func Index(c *gin.Context) {
	uid := mdw.WechatUid(c)
	var signinData mysql.Signin
	err := mus.Db.Select("id,point,updated_at,signin_cnt").Where("uid = ?", uid).Find(&signinData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		base.JSONErr(c, code.MsgErr, err)
		return
	}

	// 获取今天0点时间
	todayZeroStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", todayZeroStr)
	todayZero := t.Unix()

	// 上次更新时间
	updatedUnix := signinData.UpdatedAt.Unix()

	// 如果今天已经签到了，就直接返回，但这种不应该出现，记录日志
	if updatedUnix > todayZero {
		base.JSONErr(c, code.MsgErr, errors.New("already signin"))
		return
	}

	service.QueueSignin.Push(service.Task{
		Uid: mdw.WechatUid(c),
	})
	base.JSON(c, code.MsgOk)
}

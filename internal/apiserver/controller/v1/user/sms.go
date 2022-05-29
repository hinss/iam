package user

import (
	"fmt"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/pkg/code"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hinss/api/common"
	"github.com/hinss/api/userserver/model/v1/base"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/iam/pkg/storage"
)

// 发送验证码短信方法
func (u *UserController) SendSms(ctx *gin.Context) {

	//0.获取表单传参
	sendSmsForm := base.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		core.WriteResponse(ctx, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if errs := common.Validate(sendSmsForm); len(errs) != 0 {
		core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	//1.获得生成的短信验证码
	smsCode := GenerateCode(6)

	//2.保存到redis中
	redisCluster := &storage.RedisCluster{}
	err := redisCluster.SetKey(sendSmsForm.Mobile, smsCode, 60*time.Second)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
	}

	//TODO 3.这里没实现阿里云发送短信，直接保存
	core.WriteResponse(ctx, nil, gin.H{
		"msg": "短信发送成功",
	})
}

// 生成短信验证码
func GenerateCode(length int) (code string) {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	// 这个seed方法的目的是让生成的随机数更分散
	rand.Seed(time.Now().UnixNano())

	// 循环每轮拼一个数到sb变量地址中
	var sb strings.Builder
	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	code = sb.String()
	return
}

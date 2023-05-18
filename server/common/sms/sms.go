package sms

import (
	"context"
	"math/rand"
	"time"

	"github.com/spf13/cast"
	txcommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"server/common/log"
	"server/common/rds"
)

func TencentSMS(redisKey, phone string) error {
	// 设置访问密钥和区域
	credential := txcommon.NewCredential("your-secret-id", "your-secret-key")
	client, _ := sms.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())

	// 构造请求参数
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = txcommon.StringPtr("your-sdk-appid")
	request.SignName = txcommon.StringPtr("your-signature")
	request.TemplateId = txcommon.StringPtr("登录验证码")
	request.PhoneNumberSet = txcommon.StringPtrs([]string{"+86" + phone})

	// 生成随机验证码
	code := generateVerificationCode()

	// 设置短信模板参数
	request.TemplateParamSet = txcommon.StringPtrs([]string{code})

	// 发送短信验证码
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}

	rds.Ins.SetEx(context.Background(), redisKey, code, 100*time.Second)

	// 处理响应结果
	log.Infow("SMS sent successfully.", "response", response.ToJsonString())
	return nil
}

func generateVerificationCode() string {
	return cast.ToString(rand.Int63n(999999))
}

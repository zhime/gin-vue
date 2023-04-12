package config

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type ALiYunSms struct {
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	RegionId        string `mapstructure:"region_id"`
	SignName        string `mapstructure:"sign_name"`
	TemplateCode    string `mapstructure:"template_code"`
}

func (sms ALiYunSms) SendSMS(phoneNumber string) {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(sms.AccessKeyId, sms.AccessKeySecret)
	client, err := dysmsapi.NewClientWithOptions(sms.RegionId, config, credential)
	if err != nil {
		panic(err)
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.PhoneNumbers = phoneNumber
	request.SignName = sms.SignName
	request.TemplateCode = sms.TemplateCode
	request.TemplateParam = "{\"code\":\"123456\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

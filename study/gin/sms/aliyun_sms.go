package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func SendSMS(accessKey string, accessKeySecret string, phoneNumber string) {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(accessKey, accessKeySecret)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := dysmsapi.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.PhoneNumbers = phoneNumber
	request.SignName = "XXXXX"                                       // 签名
	request.TemplateCode = "XXXXX"                                   // 模板code
	request.TemplateParam = "{\"code\":\"" + ValidateCode(6) + "\"}" // 验证码

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func ValidateCode(num int) string {
	rand.Seed(time.Now().Unix())
	codeInt := rand.Intn(int(math.Pow10(num)))
	fmt.Println(codeInt)
	//return strconv.FormatInt(int64(codeInt), 10)
	return strconv.Itoa(codeInt)
}

func main() {
	SendSMS("XXXXX", "XXXXX", "XXXXX")
}

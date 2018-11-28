package user

import (
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Vcode    string `json:"vcode"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type PhoneRequest struct {
	PhoneNum string `json:"phone_num"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginBySmsRequest struct {
	Username string `json:"username"`
	Vcode    string `json:"vcode"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

// 用的是秒嘀短信API
// 发送短信结构体
type SendSMS struct {
	AccountSid string `json:"accountSid"` // 开发者主账号ID（ACCOUNT SID）
	SmsContent string `json:"smsContent"` // 短信内容。
	To         string `json:"to"`         // 短信接收端手机号码集合。用英文逗号分开
	Timestamp  string `json:"timestamp"`  // 时间戳。当前系统时间（24小时制），格式"yyyyMMddHHmmss"。时间戳有效时间为5分钟。
	Sig        string `json:"sig"`        // 签名。MD5(ACCOUNT SID + AUTH TOKEN + timestamp)。共32位（小写）。 注意：MD5中的内容不包含”+”号。
}

// 接收回调
type GetSmS struct {
	RespCode  string `json:"respCode"`  // 请求状态码，取值00000
	RespDesc  string `json:"respDesc"`  // 对返回状态码的描述 如：00000 代表成功
	FailCount string `json:"failCount"` // 表示验证码通知短信发送失败的条数
	FailList  string `json:"failList"`  // 失败列表，包含失败号码、失败原因。
	SmsId     string `json:"smsId"`     // 短信标识符。一个由32个字符组成的短信唯一标识符。
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}

func InitSms(phonenum, vcode string) *SendSMS {
	time := time.Now().Format("20060102150405")

	// 组合字符串生成签名
	str := "aa48d28d01d444eeb2b68cef2909e2a8c2dc43779ca947f4abcc45a6c2c57b7e" + time
	md5str := util.GetMd5String(str)

	sms := &SendSMS{
		AccountSid: "aa48d28d01d444eeb2b68cef2909e2a8",
		SmsContent: fmt.Sprintf("【拼多少】尊敬的用户，您的验证码为【%s】，请确保验证码安全。", vcode),
		To:         phonenum,
		Timestamp:  time,
		Sig:        md5str,
	}

	return sms
}

func RequestSms(phonenum, vcode string) error {
	g := GetSmS{}
	s := InitSms(phonenum, vcode)

	//生成client 参数为默认
	//client := &http.Client{}

	//生成要访问的url
	requrl := "https://api.miaodiyun.com/20150822/industrySMS/sendSMS"

	//jsonBytes, err := json.Marshal(&s)
	//if err != nil {
	//	return err
	//}

	//提交请求
	//req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBytes)))
	//if err != nil {
	//	return err
	//}

	//增加header选项
	//req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	//处理返回结果
	//resp, _ := client.Do(req)

	resp, err := http.PostForm(requrl,
		url.Values{"accountSid": {s.AccountSid}, "smsContent": {s.SmsContent},
			"to": {s.To}, "timestamp": {s.Timestamp}, "sig": {s.Sig}, "respDataType": {"JSON"}})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &g)
	if g.FailCount != "0" {
		return err
	}

	return nil
}

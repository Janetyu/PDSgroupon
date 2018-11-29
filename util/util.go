package util

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"os"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

// 随机生成短信验证码
func GenerateVerificateCode() string {
	// 根据时间戳生成不同的随机种子
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	smscode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return smscode
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	// crand为 crypto/rand
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func UploadFile(uploadDir, ext string) (string, error) {

	if err := os.MkdirAll(uploadDir, 777); err != nil {
		return "", err
	}

	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext

	fpath := uploadDir + fileName

	return fpath, nil
}

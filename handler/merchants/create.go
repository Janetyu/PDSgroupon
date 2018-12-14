package merchants

import (
	"time"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/util"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/model"
	"strconv"
)

func Create(c *gin.Context)  {
	log.Info("Merchant Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	shopName := c.DefaultPostForm("shop_name", "")
	shopPhone := c.DefaultPostForm("shop_phone", "")
	shopCert := c.DefaultPostForm("shop_cert", "")
	shopQQ := c.DefaultPostForm("shop_qq", "")
	shopIntro := c.DefaultPostForm("shop_intro", "")
	shopAddr := c.DefaultPostForm("shop_addr", "")
	userCert := c.DefaultPostForm("owner_cert", "")
	userId := c.DefaultPostForm("owner_id", "")

	if shopName == "" || shopPhone == "" || shopCert == "" || shopQQ == ""{
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if shopIntro == "" || shopAddr == "" || userCert == "" || userId == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	uid, err := strconv.Atoi(userId)
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if mer,_ := model.GetMerchantByOwnerId(uint64(uid)); mer.ShopName != "" {
		SendResponse(c, errno.ErrMerchantHasApplyOrPass, nil)
		return
	}

	file, _ := c.FormFile("shop_logo")

	if file == nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if ok := IsPathVaild(file.Filename); !ok {
		SendResponse(c, errno.ErrUploadExt, nil)
		return
	}

	uploadDir := "static/upload/merchants/" + time.Now().Format("2006/01/02/")
	dst, err := util.UploadFile(uploadDir, path.Ext(file.Filename))
	if err != nil {
		SendResponse(c, errno.ErrUploadFail, nil)
		return
	}

	merchants := model.MerchantModel{
		ShopName: shopName,
		ShopAddr: shopAddr,
		ShopCert: shopCert,
		ShopIntro: shopIntro,
		ShopPhone: shopPhone,
		ShopQQ: shopQQ,
		ShopLogo: dst,
		UserCert: userCert,
		UserId: uint64(uid),
		IsReview: "正在审核",
		Mark: "",
	}

	if err := merchants.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	SendResponse(c, nil, nil)
}

// 判断文件后缀是否合法
func IsPathVaild(filename string) bool {
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := path.Ext(filename)

	if _, ok := AllowExtMap[ext]; !ok {
		return false
	}
	return true
}

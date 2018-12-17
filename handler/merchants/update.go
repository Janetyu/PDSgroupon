package merchants

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"os"
	"path"
	"time"
)

func Update(c *gin.Context) {
	log.Info("Merchant UpdateForApply function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	id, _ := strconv.Atoi(c.Param("id"))
	m, err := model.GetMerchantByOwnerId(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	shopName := c.DefaultPostForm("shop_name", m.ShopName)
	shopPhone := c.DefaultPostForm("shop_phone", m.ShopPhone)
	shopCert := c.DefaultPostForm("shop_cert", m.ShopCert)
	shopQQ := c.DefaultPostForm("shop_qq", m.ShopQQ)
	shopIntro := c.DefaultPostForm("shop_intro", m.ShopIntro)
	shopAddr := c.DefaultPostForm("shop_addr", m.ShopAddr)
	userCert := c.DefaultPostForm("owner_cert", m.UserCert)

	shoplogo := m.ShopLogo

	file, _ := c.FormFile("shop_logo")

	if file != nil {

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

		// 删除旧文件地址
		if err := os.Remove(shoplogo); err != nil {
			log.Errorf(err, "del file occured error is :")
		}

		// 更新文件地址
		shoplogo = dst

	}

	merchants := model.MerchantModel{
		ShopName:  shopName,
		ShopAddr:  shopAddr,
		ShopCert:  shopCert,
		ShopIntro: shopIntro,
		ShopPhone: shopPhone,
		ShopQQ:    shopQQ,
		ShopLogo:  shoplogo,
		UserCert:  userCert,
		UserId:    uint64(id),
		IsReview:  "审核通过",
		Mark:      "",
	}

	if err := merchants.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	if err := c.SaveUploadedFile(file, shoplogo); err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	SendResponse(c, nil, merchants)
}

// 管理员审核商铺认证信息
func Review(c *gin.Context) {
	log.Info("ReviewMerchant function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r ReViewRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	merchant, err := model.GetMerchantById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	if r.IsReview == "审核通过" {
		merchant.IsReview = r.IsReview
		merchant.Mark = ""
		merchant.UpdatedAt = time.Time{}

		uid := merchant.UserId
		user, err := model.GetUserById(uid)
		if err != nil {
			SendResponse(c, errno.ErrUserNotFound, nil)
			return
		}

		user.RoleId = 2
		if err := user.Update(); err != nil {
			SendResponse(c, errno.ErrDatabase, nil)
			return
		}

		if err := merchant.Update(); err != nil {
			user.RoleId = 1
			user.Update()
			SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	}

	if r.IsReview == "审核不通过" {
		merchant.IsReview = r.IsReview
		merchant.Mark = r.Mark
		merchant.UpdatedAt = time.Time{}

		if err := merchant.Update(); err != nil {
			SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	}

	rsp := GetOneResponse{
		Id:        merchant.Id,
		ShopName:  merchant.ShopName,
		ShopLogo:  merchant.ShopLogo,
		ShopAddr:  merchant.ShopAddr,
		ShopIntro: merchant.ShopIntro,
		ShopQQ:    merchant.ShopQQ,
		ShopCert:  merchant.ShopCert,
		ShopPhone: merchant.ShopPhone,
		UserId:    merchant.UserId,
		UserCert:  merchant.UserCert,
		IsReview:  merchant.IsReview,
		Mark:      merchant.Mark,
		CreatedAt: merchant.CreatedAt,
	}

	SendResponse(c, nil, rsp)
}

// 用户修改
func UpdateForApply(c *gin.Context) {
	log.Info("Merchant UpdateForApply function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	id, _ := strconv.Atoi(c.Param("uid"))
	m, err := model.GetMerchantByOwnerId(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	shopName := c.DefaultPostForm("shop_name", m.ShopName)
	shopPhone := c.DefaultPostForm("shop_phone", m.ShopPhone)
	shopCert := c.DefaultPostForm("shop_cert", m.ShopCert)
	shopQQ := c.DefaultPostForm("shop_qq", m.ShopQQ)
	shopIntro := c.DefaultPostForm("shop_intro", m.ShopIntro)
	shopAddr := c.DefaultPostForm("shop_addr", m.ShopAddr)
	userCert := c.DefaultPostForm("owner_cert", m.UserCert)

	shoplogo := m.ShopLogo

	file, _ := c.FormFile("shop_logo")

	if file != nil {

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

		// 删除旧文件地址
		if err := os.Remove(shoplogo); err != nil {
			log.Errorf(err, "del file occured error is :")
		}

		// 更新文件地址
		shoplogo = dst

	}

	merchants := model.MerchantModel{
		BaseModel: model.BaseModel{Id: m.Id, CreatedAt: m.CreatedAt, UpdatedAt: time.Time{}},
		ShopName:  shopName,
		ShopAddr:  shopAddr,
		ShopCert:  shopCert,
		ShopIntro: shopIntro,
		ShopPhone: shopPhone,
		ShopQQ:    shopQQ,
		ShopLogo:  shoplogo,
		UserCert:  userCert,
		UserId:    uint64(id),
		IsReview:  "正在审核",
		Mark:      "",
	}

	if err := merchants.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	if err := c.SaveUploadedFile(file, shoplogo); err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	SendResponse(c, nil, merchants)
}

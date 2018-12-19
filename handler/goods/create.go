package goods

import (
	"strconv"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
)

func Create(c *gin.Context)  {
	log.Info("Goods Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	goodsName := c.DefaultPostForm("goods_name", "")
	goodsDesc := c.DefaultPostForm("goods_desc", "")
	goodsCost := c.DefaultPostForm("goods_cost", "")
	goodsPrice := c.DefaultPostForm("goods_price", "")
	goodsDiscount := c.DefaultPostForm("goods_discount", "")
	goodsStock := c.DefaultPostForm("goods_stock", "")
	stockWarn := c.DefaultPostForm("stock_warn", "")
	goodsPeople := c.DefaultPostForm("goods_people", "")
	groupAging := c.DefaultPostForm("group_aging", "")
	shopId := c.DefaultPostForm("shop_id", "")
	mainsortId := c.DefaultPostForm("mainsort_id", "")
	subsortId := c.DefaultPostForm("subsort_id", "")
	isFare := c.DefaultPostForm("is_fare", "")
	goodsFare := c.DefaultPostForm("goods_fare", "")

	if goodsName == "" || goodsDesc == "" || goodsCost == "" || goodsPrice == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if goodsDiscount == "" || goodsStock == "" || stockWarn == "" || goodsPeople == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if groupAging == "" || shopId == "" || mainsortId == "" || subsortId == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if isFare == "" || goodsFare == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	sid, err := strconv.Atoi(shopId)
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	mainsid, err := strconv.Atoi(mainsortId)
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	subsid, err := strconv.Atoi(subsortId)
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if _, err := model.GetMerchantById(uint64(sid)); err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	file, _ := c.FormFile("goods_photo")

	if file == nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if ok := IsPathVaild(file.Filename); !ok {
		SendResponse(c, errno.ErrUploadExt, nil)
		return
	}

	uploadDir := "static/upload/goods/" + time.Now().Format("2006/01/02/")
	dst, err := util.UploadFile(uploadDir, path.Ext(file.Filename))
	if err != nil {
		log.Errorf(err, "the uploadfile error is:")
		SendResponse(c, errno.ErrUploadFail, nil)
		return
	}

	gp,_ := strconv.Atoi(goodsPeople) // 拼团人数
	gs,_ := strconv.Atoi(goodsStock)  // 商品库存
	ga,_ := strconv.Atoi(groupAging)  // 拼团时效
	sw,_ := strconv.Atoi(stockWarn)   // 库存报警

	isf := true
	if isFare == "false" {
		isf = false
	}

	goods := model.GoodsModel{
		GoodsName: goodsName,
		GoodsDesc: goodsDesc,
		GoodsCost: goodsCost,
		GoodsPrice: goodsPrice,
		GoodsDiscount: goodsDiscount,
		GoodsStock: gs,
		StockWarn: sw,
		GoodsPeople: gp,
		GroupAging: ga,
		ShopId: uint64(sid),
		MainsortId: uint64(mainsid),
		SubsortId: uint64(subsid),
		GoodsFare: goodsFare,
		GoodsPhoto: dst,
		GoodsSales: 0,
		IsFare: isf,
		IsShelf: true,
	}

	if err := goods.Create(); err != nil {
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

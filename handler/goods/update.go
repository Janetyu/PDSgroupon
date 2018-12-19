package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	"strconv"
	"time"
	"path"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"os"
)

func Update(c *gin.Context)  {
	log.Info("Goods Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	id, _ := strconv.Atoi(c.Param("id"))
	g, err := model.GetGoodsModelById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrGoodsNotFount, nil)
		return
	}

	goodsName := c.DefaultPostForm("goods_name", g.GoodsName)
	goodsDesc := c.DefaultPostForm("goods_desc", g.GoodsDesc)
	goodsCost := c.DefaultPostForm("goods_cost", g.GoodsCost)
	goodsPrice := c.DefaultPostForm("goods_price", g.GoodsPrice)
	goodsDiscount := c.DefaultPostForm("goods_discount", g.GoodsDiscount)
	goodsStock := c.DefaultPostForm("goods_stock", "")
	stockWarn := c.DefaultPostForm("stock_warn", "")
	goodsPeople := c.DefaultPostForm("goods_people", "")
	groupAging := c.DefaultPostForm("group_aging", "")
	mainsortId := c.DefaultPostForm("mainsort_id", "")
	subsortId := c.DefaultPostForm("subsort_id", "")
	isFare := c.DefaultPostForm("is_fare", "")
	goodsFare := c.DefaultPostForm("goods_fare", g.GoodsFare)

	gs,err := strconv.Atoi(goodsStock)  // 商品库存
	if err != nil {
		gs = g.GoodsStock
	}

	sw,err := strconv.Atoi(stockWarn)   // 库存报警
	if err != nil {
		sw = g.StockWarn
	}

	gp,err := strconv.Atoi(goodsPeople) // 拼团人数
	if err != nil {
		gp = g.GoodsPeople
	}

	ga,err := strconv.Atoi(groupAging)  // 拼团时效
	if err != nil {
		ga = g.GroupAging
	}

	mainsid, err := strconv.Atoi(mainsortId)
	if err != nil {
		mainsid = int(g.MainsortId)
	}

	subsid, err := strconv.Atoi(subsortId)
	if err != nil {
		subsid = int(g.SubsortId)
	}

	goodsPhoto := g.GoodsPhoto

	file, _ := c.FormFile("goods_photo")

	if file != nil {

		if ok := IsPathVaild(file.Filename); !ok {
			SendResponse(c, errno.ErrUploadExt, nil)
			return
		}

		uploadDir := "static/upload/goods/" + time.Now().Format("2006/01/02/")
		dst, err := util.UploadFile(uploadDir, path.Ext(file.Filename))
		if err != nil {
			SendResponse(c, errno.ErrUploadFail, nil)
			return
		}

		// 删除旧文件地址
		if err := os.Remove(goodsPhoto); err != nil {
			log.Errorf(err, "del file occured error is :")
		}

		// 更新文件地址
		goodsPhoto = dst

	}

	isf := g.IsFare
	if isFare != "" {
		if isFare == "true" && isf == false {
			isf = true
		} else if isFare == "false" && isf == true {
			isf = false
		}
	}

	goods := model.GoodsModel{
		BaseModel: model.BaseModel{Id: g.Id, CreatedAt: g.CreatedAt, UpdatedAt: time.Time{}},
		GoodsName: goodsName,
		GoodsDesc: goodsDesc,
		GoodsCost: goodsCost,
		GoodsPrice: goodsPrice,
		GoodsDiscount: goodsDiscount,
		GoodsStock: gs,
		StockWarn: sw,
		GoodsPeople: gp,
		GroupAging: ga,
		ShopId: g.ShopId,
		MainsortId: uint64(mainsid),
		SubsortId: uint64(subsid),
		GoodsPhoto: goodsPhoto,
		GoodsSales: g.GoodsSales,
		IsFare: isf,
		GoodsFare: goodsFare,
		IsShelf: g.IsShelf,
	}

	if err := goods.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	if err := c.SaveUploadedFile(file, goodsPhoto); err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	SendResponse(c, nil, goods)
}

func IsShelf(c *gin.Context)  {
	log.Info("GoodsShelf Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	id, _ := strconv.Atoi(c.Param("id"))
	g, err := model.GetGoodsModelById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrGoodsNotFount, nil)
		return
	}

	isShelf := c.DefaultPostForm("is_shelf", "")

	if isShelf == "false" && g.IsShelf == true {
		g.IsShelf = false
	} else if isShelf == "true" && g.IsShelf == false {
		g.IsShelf = true
	}

	if err := g.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	SendResponse(c, nil, g)
}

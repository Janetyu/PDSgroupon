package banner

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/util"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Update(c *gin.Context)  {
	log.Info("Banner Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	bannerId, _ := strconv.Atoi(c.Param("id"))

	banner, err := model.GetBannerById(uint64(bannerId))
	if err != nil {
		SendResponse(c, errno.ErrBannerNotFount, nil)
		return
	}

	title := c.DefaultPostForm("title", banner.Title)
	url := c.DefaultPostForm("url", banner.Url)
	order := c.DefaultPostForm("order", strconv.Itoa(banner.Order))
	clinum := c.DefaultPostForm("cli_num",strconv.Itoa(banner.CliNum))

	newOrder,_ := strconv.Atoi(order)
	newCliNum,_ := strconv.Atoi(clinum)

	bmodel := model.BannerModel{
		BaseModel: model.BaseModel{Id: banner.Id, CreatedAt: banner.CreatedAt, UpdatedAt: time.Time{}},
		Title: title,
		Url: url,
		Order: newOrder,
		Image: banner.Image,
		CliNum: newCliNum,
	}

	// Save changed fields.
	if err := bmodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := GetOneResponse{
		Id:        bmodel.Id,
		Title:	   bmodel.Title,
		Url:	   bmodel.Url,
		Order:     bmodel.Order,
		Image:     bmodel.Image,
		CliNum:    bmodel.CliNum,
		CreatedAt: bmodel.CreatedAt,
		UpdatedAt: bmodel.UpdatedAt,
	}

	SendResponse(c, nil, rsp)
}
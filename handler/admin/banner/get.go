package banner

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `username` from the database.
	banner, err := model.GetBannerById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrBannerNotFount, nil)
		return
	}

	rsp := GetOneResponse{
		Id:        banner.Id,
		Title:     banner.Title,
		Url:       banner.Url,
		Order:     banner.Order,
		Image:     banner.Image,
		CliNum:    banner.CliNum,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}

	SendResponse(c, nil, rsp)
}

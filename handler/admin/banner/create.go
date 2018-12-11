package banner

import (
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
)

func Create(c *gin.Context) {
	log.Info("Banner Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	title := c.DefaultPostForm("title", "")
	url := c.DefaultPostForm("url", "")
	order := c.DefaultPostForm("order", "")

	if title == "" || url == "" || order == "" {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	file, _ := c.FormFile("image")

	if ok := IsPathVaild(file.Filename); !ok {
		SendResponse(c, errno.ErrUploadExt, nil)
		return
	}

	o, error := strconv.Atoi(order)
	if error != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	uploadDir := "static/upload/banner/" + time.Now().Format("2006/01/02/")
	dst, err := util.UploadFile(uploadDir, path.Ext(file.Filename))
	if err != nil {
		SendResponse(c, errno.ErrUploadFail, nil)
		return
	}

	banner := model.BannerModel{
		Title:  title,
		Url:    url,
		Order:  o,
		Image:  dst,
		CliNum: 0,
	}

	if err := banner.Create(); err != nil {
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

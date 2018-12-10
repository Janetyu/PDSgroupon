package upload

import (
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"os"
)

// 单文件上传
func SingleUpload(c *gin.Context) {
	log.Info("SingleUpload function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	Id, _ := strconv.Atoi(c.Param("id"))

	url := c.Request.URL.String()

	var isUser = make(chan bool)
	var isBanner = make(chan bool)
	var oldFilePath string

	go CheckUrl(url, isUser, isBanner)

	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	file, _ := c.FormFile("file")

	// 获取上传你文件后缀
	ext := path.Ext(file.Filename)

	if _, ok := AllowExtMap[ext]; !ok {
		SendResponse(c, errno.ErrUploadExt, nil)
		return
	}

	select {
	case <-isUser:
		user, err := model.GetUserById(uint64(Id))
		if err != nil {
			SendResponse(c, errno.ErrUserNotFound, nil)
			return
		}

		oldFilePath = user.HeadImage
		uploadDir := "static/upload/user/" + time.Now().Format("2006/01/02/")

		dst, err := util.UploadFile(uploadDir, ext)
		if err != nil {
			SendResponse(c, errno.ErrUploadFail, nil)
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			SendResponse(c, errno.InternalServerError, nil)
			return
		}

		umodel := model.UserModel{
			BaseModel: model.BaseModel{Id: user.Id, CreatedAt: user.CreatedAt, UpdatedAt: time.Time{}},
			Username:  user.Username,
			Password:  user.Password,
			NickName:  user.NickName,
			Address:   user.Address,
			Name:      user.Name,
			HeadImage: dst,
			Sex:       user.Sex,
			Account:   user.Account,
			RoleId:    user.RoleId,
		}

		if err := umodel.Update(); err != nil {
			SendResponse(c, errno.ErrDatabase, nil)
			return
		}

		if err := os.Remove(oldFilePath); err != nil {
			log.Errorf(err,"del file occured error is :")
		}

		rsp := UserResponse{
			HeadImage: dst,
		}

		SendResponse(c, nil, rsp)

	case <-isBanner:
		banner, err := model.GetBannerById(uint64(Id))
		if err != nil {
			SendResponse(c, errno.ErrBannerNotFount, nil)
			return
		}

		oldFilePath = banner.Image
		uploadDir := "static/upload/banner/" + time.Now().Format("2006/01/02/")

		dst, err := util.UploadFile(uploadDir, ext)
		if err != nil {
			SendResponse(c, errno.ErrUploadFail, nil)
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			SendResponse(c, errno.InternalServerError, nil)
			return
		}

		bmodel := model.BannerModel{
			BaseModel: model.BaseModel{Id: banner.Id, CreatedAt: banner.CreatedAt, UpdatedAt: time.Time{}},
			Title: banner.Title,
			Url:   banner.Url,
			Order: banner.Order,
			Image: dst,
			CliNum: banner.CliNum,
		}

		if err := bmodel.Update(); err != nil {
			SendResponse(c, errno.ErrDatabase, nil)
			return
		}

		if err := os.Remove(oldFilePath); err != nil {
			log.Errorf(err,"del file occured error is :")
		}

		rsp := BannerResponse{
			Image: dst,
		}

		SendResponse(c, nil, rsp)

	default:
		SendResponse(c, errno.ErrUploadFail, nil)
	}
}

func CheckUrl(url string, isUser, isBanner chan bool) {
	if strings.Contains(url, "user") {
		isUser <- true
	} else if strings.Contains(url, "banner") {
		isBanner <- true
	}
}

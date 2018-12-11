package category

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
)

func UpdateMain(c *gin.Context) {
	log.Info("MainCategory Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	categoryId, _ := strconv.Atoi(c.Param("id"))

	category, err := model.GetCategoryById(uint64(categoryId))
	if err != nil {
		SendResponse(c, errno.ErrCategoryNotFount, nil)
		return
	}

	sortName := c.DefaultPostForm("sort_name", category.SortName)

	cmodel := model.CategoryModel{
		Id:       category.Id,
		Pid:      category.Pid,
		SortName: sortName,
	}

	// Save changed fields.
	if err := cmodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Id:       cmodel.Id,
		Pid:      cmodel.Pid,
		SortName: sortName,
	}

	SendResponse(c, nil, rsp)
}

func UpdateSub(c *gin.Context) {
	log.Info("SubCategory Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	categoryId, _ := strconv.Atoi(c.Param("id"))

	category, err := model.GetCategoryById(uint64(categoryId))
	if err != nil {
		SendResponse(c, errno.ErrCategoryNotFount, nil)
		return
	}

	newpid := c.DefaultPostForm("pid", "")
	if newpid != "" {
		npid, err := strconv.Atoi(newpid)
		if err != nil {
			SendResponse(c, errno.ErrValidation, nil)
			return
		} else {
			category.Pid = uint64(npid)
		}
	}

	sortName := c.DefaultPostForm("sort_name", category.SortName)

	cmodel := model.CategoryModel{
		Id:       category.Id,
		Pid:      category.Pid,
		SortName: sortName,
	}

	// Save changed fields.
	if err := cmodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Id:       cmodel.Id,
		Pid:      cmodel.Pid,
		SortName: sortName,
	}

	SendResponse(c, nil, rsp)
}

package category

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"strconv"
)

// 创建主类
func CreateMain(c *gin.Context) {
	log.Info("CreateMainCategory function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateMainRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	cg := model.CategoryModel{
		Pid:      0,
		SortName: r.SortName,
	}

	if category, err := model.GetCategory(cg.SortName); err != nil || category.SortName != "" {
		if category.SortName != "" {
			SendResponse(c, errno.ErrCategoryHasCreate, nil)
			return
		} else if err != nil && err.Error() != "record not found" {
			SendResponse(c, errno.ErrDatabase, nil)
			log.Errorf(err, "the database error is:")
			return
		}
	}

	// Insert the user to the database.
	if err := cg.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	newCg, _ := model.GetCategory(cg.SortName)
	rsp := CreateResponse{
		Id:       newCg.Id,
		Pid:      newCg.Pid,
		SortName: newCg.SortName,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

// 创建子类
func CreateSub(c *gin.Context) {
	log.Info("CreateSubCategory function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateSubRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	pid, err := strconv.Atoi(r.Pid)
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	cg := model.CategoryModel{
		Pid:      uint64(pid),
		SortName: r.SortName,
	}

	if category, err := model.GetCategory(cg.SortName); err != nil || category.SortName != "" {
		if category.SortName != "" {
			SendResponse(c, errno.ErrCategoryHasCreate, nil)
			return
		} else if err != nil && err.Error() != "record not found" {
			SendResponse(c, errno.ErrDatabase, nil)
			log.Errorf(err, "the database error is:")
			return
		}
	}

	// Insert the user to the database.
	if err := cg.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	newCg, _ := model.GetCategory(cg.SortName)
	rsp := CreateResponse{
		Id:       newCg.Id,
		Pid:      newCg.Pid,
		SortName: newCg.SortName,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

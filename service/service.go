package service

import (
	"fmt"
	"sync"

	"PDSgroupon/model"
	"PDSgroupon/util"
)

func ListUser(offset, limit int) ([]*model.UserInfo, uint64, error) {
	// 用于存放没有密码的userinfo，可拓展搜索等功能
	infos := make([]*model.UserInfo, 0)
	// offset 页数，limit 每页条数
	users, count, err := model.ListUser(offset, limit)
	if err != nil {
		return nil, count, err
	}

	// 保存数据库查询数据的原排序
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.BaseModel.Id)
	}

	// 控制goroutine的协程调度，监控工作
	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock: new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	finished := make(chan bool, 1)

	wg.Add(len(users))
	for _, u := range users {
		//wg.Add(1)

		go func(userModel *model.UserModel) {
			defer wg.Done()

			// 用锁保证数据一致性
			userList.Lock.Lock()

			// 对业务所需进行数据修改或其他操作
			userList.IdMap[userModel.Id] = &model.UserInfo{
				Id: userModel.Id,
				Username: userModel.Username,
				NickName: userModel.NickName,
				Address: userModel.Address,
				Name: userModel.Name,
				HeadImage: userModel.HeadImage,
				Sex: userModel.Sex,
				Account: userModel.Account,
				RoleId: userModel.RoleId,
				CreatedAt: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: userModel.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

			userList.Lock.Unlock()
		}(u)
	}

	// 等 wg 为0时，关闭 finished 通道，退出并发操作
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <- finished:
	}

	//wg.Wait()

	// 按顺序复位所有userinfo
	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}


func ListUser2(username string, offset, limit int) ([]*model.UserInfo2, uint64, error) {
	infos := make([]*model.UserInfo2, 0)
	users, count, err := model.ListUser2(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	// 用了 sync 包来做并行查询，以使响应延时更小
	wg := sync.WaitGroup{}
	userList2 := model.UserList2{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo2, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 在并发处理中，更新同一个变量为了保证数据一致性，通常需要做锁处理
	// 并行提高查询效率,模拟业务中对查询出来的数据进行处理再返回
	// Improve query efficiency in parallel
	for _, u := range users {
		// 对Add的调用应该在创建要等待的goroutine或其他事件的语句之前执行。
		wg.Add(1)
		go func(u *model.UserModel) {
			// wg - 1
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList2.Lock.Lock()
			defer userList2.Lock.Unlock()
			userList2.IdMap[u.Id] = &model.UserInfo2{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	// 使用 IdMap 是因为查询的列表通常需要按时间顺序进行排序，
	// 一般数据库查询后的列表已经排过序了，但是为了减少延时， 程序中用了并发，
	// 这时候会打乱排序，所以通过 `IdMap` 来记录并发处理前的顺序，处理后再重新复位
	for _, id := range ids {
		infos = append(infos, userList2.IdMap[id])
	}

	return infos, count, nil
}

func ListAdmin(offset, limit int) ([]*model.AdminInfo, uint64, error) {
	// 用于存放没有密码的userinfo，可拓展搜索等功能
	infos := make([]*model.AdminInfo, 0)
	// offset 页数，limit 每页条数
	admins, count, err := model.ListAdmin(offset, limit)
	if err != nil {
		return nil, count, err
	}

	// 保存数据库查询数据的原排序
	ids := []uint64{}
	for _, admin := range admins {
		ids = append(ids, admin.BaseModel.Id)
	}

	// 控制goroutine的协程调度，监控工作
	wg := sync.WaitGroup{}
	adminList := model.AdminList{
		Lock: new(sync.Mutex),
		IdMap: make(map[uint64]*model.AdminInfo, len(admins)),
	}

	finished := make(chan bool, 1)

	wg.Add(len(admins))
	for _, a := range admins {
		//wg.Add(1)

		go func(adminModel *model.AdminModel) {
			defer wg.Done()

			// 用锁保证数据一致性
			adminList.Lock.Lock()

			// 对业务所需进行数据修改或其他操作
			adminList.IdMap[adminModel.Id] = &model.AdminInfo{
				Id: adminModel.Id,
				Username: adminModel.Username,
				RoleId: adminModel.RoleId,
				CreatedAt: adminModel.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: adminModel.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

			adminList.Lock.Unlock()
		}(a)
	}

	// 等 wg 为0时，关闭 finished 通道，退出并发操作
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <- finished:
	}

	//wg.Wait()

	// 按顺序复位所有userinfo
	for _, id := range ids {
		infos = append(infos, adminList.IdMap[id])
	}

	return infos, count, nil
}
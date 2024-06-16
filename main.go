package main

import (
	"bili-down-fav/src/bili/fav"
	"bili-down-fav/src/bili/user"
	"bili-down-fav/src/bili/video"
	"bili-down-fav/src/conf"
	"bili-down-fav/src/util"
	"context"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	curUser := user.CurrentUser().Data
	if curUser.IsLogin {
		fmt.Println("当前用户：" + curUser.Uname)
		favs := fav.ListForDownloads(fav.ListFavFolder(curUser.Mid))
		ctx, cancel := context.WithCancel(context.Background())
		go util.Log(ctx)

		fmt.Println("视频数量：", len(favs))

		threads, _ := strconv.Atoi(conf.Get("file", "multi_download"))
		if threads > 0 {
			// 限制并发数量
			ch := make(chan any, threads)
			// 多线程下载
			wg := new(sync.WaitGroup)
			for _, v := range favs {
				wg.Add(1)
				ch <- v
				bvid := v.Bvid
				go func() {
					video.Download(bvid)
					wg.Done()
					<-ch
				}()
			}
			wg.Wait()
		} else {
			fmt.Println("线程数必须大于0")
		}
		cancel()
	} else {
		fmt.Println("当前未登录，请登录")
		user.Login(func() {
			fmt.Println("登录")
			main()
		})
	}
}

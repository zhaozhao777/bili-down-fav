package util

import (
	"bili-down-fav/src/conf"
	"bufio"
	"context"
	"fmt"
	"os"
)

var success chan string
var fail chan string
var sw *bufio.Writer
var fw *bufio.Writer

func init() {
	success = make(chan string)
	fail = make(chan string)
	os.Remove(conf.Get("log", "success"))
	os.Remove(conf.Get("log", "fail"))
	sFile, _ := os.Create(conf.Get("log", "success"))
	fFile, _ := os.Create(conf.Get("log", "fail"))
	sw = bufio.NewWriter(sFile)
	fw = bufio.NewWriter(fFile)
}

func LogSuccess(bvid string, name string) {
	success <- fmt.Sprintf("%s,%s", bvid, name)
}

func LogFail(bvid string, name string, err any) {
	fail <- fmt.Sprintf("%s,%s,%s", bvid, name, err)
}

func Log(ctx context.Context) {
	for {
		select {
		case msg := <-success:
			// log.Println(msg)
			fmt.Fprintln(sw, msg)
			sw.Flush()
		case msg := <-fail:
			// log.Println(msg)
			fmt.Fprintln(fw, msg)
			fw.Flush()
		case <-ctx.Done():
			goto END
		}
	}
END:
}

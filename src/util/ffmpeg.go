package util

import (
	"bili-down-fav/src/conf"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		dir, _ := os.Getwd()
		os.Setenv("Path", os.Getenv("Path")+";"+dir+"/ffmpeg/bin;")
	}
}

func Combine(video string, audio string, out string) {
	createDir(out)
	cmd := exec.Command("ffmpeg", "-i", video, "-i", audio, "-c", "copy", out,
		"-loglevel", conf.Get("log", "level"),
		"-y")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	//cmd.Stdin = os.Stdin
	cmd.Run()
}

func Convert(video string, out string) {
	createDir(out)
	cmd := exec.Command("ffmpeg", "-i", video, "-c", "copy", out,
		"-loglevel", conf.Get("log", "level"),
		"-y")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	//cmd.Stdin = os.Stdin
	cmd.Run()
}

func createDir(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("创建目录：", dir)
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

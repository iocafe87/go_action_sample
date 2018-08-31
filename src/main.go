package main

import (
	"log"
	"os"

	_ "iocafe87.com/go_action_sample/src/matchers"
	"iocafe87.com/go_action_sample/src/search"
)

// 在main之前执行
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")

}

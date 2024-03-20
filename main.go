package main

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
)

func main() {
	chinese := zh.New()
	english := en.New()
	uni := ut.New(chinese, chinese, english)
	//根据参数取翻译器实例
	trans, _ := uni.GetTranslator("en")

	// 添加翻译的消息
	trans.Add("welcome", "欢迎{0}访问", false)

	fmt.Println(trans.T("welcome", "TomJerry"))
}

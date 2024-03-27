package http

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var defaultLang string

func Init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./pkg/common/text/active.zh-CN.toml")
	defaultLang = "en"
}
func LocalTranslate(acceptLang string, message string) string {
	fmt.Println(acceptLang)
	localizer := i18n.NewLocalizer(bundle, defaultLang, acceptLang)
	transMsg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PersonCats",
			One:   "{{.Name}} has {{.Count}} cat.",
			Other: "{{.Name}} has {{.Count}} cats.",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 2,
		},
		PluralCount: 2,
	})
	return transMsg
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func TransMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		acceptLang := context.GetHeader("Accept-Language")
		context.Set("accept-lang", acceptLang)
		context.Next()
	}
}

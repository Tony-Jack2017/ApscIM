package http

import (
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
	bundle.LoadMessageFile("./pkg/common/text/active.en.toml")
	bundle.LoadMessageFile("./pkg/common/text/active.zh-CN.toml")
	defaultLang = "en"
}

func LocalTranslate(acceptLang string, messageID string) string {
	var lang string
	if acceptLang == "" {
		lang = defaultLang
	} else {
		lang = acceptLang
	}
	localizer := i18n.NewLocalizer(bundle, lang)
	transMsg := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	return transMsg
}

func TransMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		acceptLang := context.GetHeader("Accept-Language")
		context.Set("accept-lang", acceptLang)
		context.Next()
	}
}

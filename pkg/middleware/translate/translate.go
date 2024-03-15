package common

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func Init() {
	bundle := i18n.NewBundle(language.English)
}

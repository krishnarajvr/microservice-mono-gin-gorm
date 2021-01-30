package locale

import (
	"fmt"
	"os"

	"github.com/kataras/i18n"
)

type Locale struct {
	lang string
	i18n *i18n.I18n
}

func (l *Locale) Get(key string, args ...interface{}) string {
	return l.i18n.Tr(l.lang, key, args...)
}

func (l *Locale) SetLang(langKey string) {
	l.lang = langKey
}

func (l *Locale) New(languageKey string) *Locale {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("Not able to get current working director")
	}

	I18n, err := i18n.New(i18n.Glob(dir+"/locale/*/*"), "en-US", "el-GR", "zh-CN")

	if err != nil {
		panic(err)
	}

	l.lang = languageKey
	l.i18n = I18n

	return l
}

package i18n

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"

	"github.com/toolkits/pkg/file"
	"github.com/toolkits/pkg/runner"
)

var (
	catalogs = make(map[string]*catalog.Builder)
	printers = make(map[string]*message.Printer)
)

// Init will init i18n support via input language.
func Init(dictPath ...string) {
	dp := path.Join(runner.Cwd, "etc", "i18n.json")
	if len(dictPath) > 0 && dictPath[0] != "" {
		dp = dictPath[0]
	}

	DictFileRegister(dp)

	// en for default
	printers[""] = message.NewPrinter(langTag("en"))
	printers["en"] = message.NewPrinter(langTag("en"))
}

func DictFileRegister(filePath string) {
	if !file.IsExist(filePath) {
		// fmt.Printf("i18n config file %s not found. donot worry, we'll use default configuration\n", filePath)
		return
	}

	content, err := file.ToTrimString(filePath)
	if err != nil {
		fmt.Printf("read i18n config file %s fail: %s\n", filePath, err)
		return
	}

	m := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(content), &m)
	if err != nil {
		fmt.Printf("parse i18n config file %s fail: %s\n", filePath, err)
		return
	}

	DictRegister(m)
}

func DictRegister(m map[string]map[string]string) {
	for lang, dict := range m {
		tag := langTag(lang)
		cata := catalog.NewBuilder()
		for k, v := range dict {
			cata.SetString(tag, k, v)
		}

		catalogs[lang] = cata
		printers[lang] = message.NewPrinter(tag, message.Catalog(cata))
	}
}

func langTag(l string) language.Tag {
	switch strings.ToLower(l) {
	case "zh", "cn":
		return language.Chinese
	default:
		return language.English
	}
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(lang, format string, a ...interface{}) string {
	if _, exists := printers[lang]; !exists {
		return fmt.Sprintf(format, a...)
	}

	return printers[lang].Sprintf(format, a...)
}

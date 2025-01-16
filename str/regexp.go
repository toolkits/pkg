package str

import (
	"regexp"
	"strings"
)

var IPReg, _ = regexp.Compile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
var MailReg, _ = regexp.Compile(`\w[-._\w]*@\w[-._\w]*\.\w+`)

func IsMatch(s, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		return false
	}

	return match
}

func IsIdentifier(s string, pattern ...string) bool {
	defpattern := "^[a-zA-Z0-9\\-\\_\\.]+$"
	if len(pattern) > 0 {
		defpattern = pattern[0]
	}

	return IsMatch(s, defpattern)
}

func IsMail(s string) bool {
	return MailReg.MatchString(s)
}

// 定义常用国家区号
var validCountryCodes = map[string]bool{
	"1":   true, // 北美
	"7":   true, // 俄罗斯
	"20":  true, // 埃及
	"27":  true, // 南非
	"31":  true, // 荷兰
	"32":  true, // 比利时
	"33":  true, // 法国
	"34":  true, // 西班牙
	"39":  true, // 意大利
	"44":  true, // 英国
	"45":  true, // 丹麦
	"46":  true, // 瑞典
	"47":  true, // 挪威
	"49":  true, // 德国
	"52":  true, // 墨西哥
	"54":  true, // 阿根廷
	"55":  true, // 巴西
	"60":  true, // 马来西亚
	"61":  true, // 澳大利亚
	"62":  true, // 印度尼西亚
	"63":  true, // 菲律宾
	"65":  true, // 新加坡
	"66":  true, // 泰国
	"81":  true, // 日本
	"82":  true, // 韩国
	"84":  true, // 越南
	"86":  true, // 中国
	"90":  true, // 土耳其
	"91":  true, // 印度
	"852": true, // 香港
	"853": true, // 澳门
	"855": true, // 柬埔寨
	"886": true, // 台湾
}

func IsPhone(s string) bool {
	// 清理空格和连字符
	s = regexp.MustCompile(`[\s\-]`).ReplaceAllString(s, "")

	// 处理 00 开头的国际格式
	if strings.HasPrefix(s, "00") {
		s = "+" + s[2:]
	}

	if strings.HasPrefix(s, "+") {
		s = s[1:] // 去掉加号

		// 提取可能的国家区号（1-4位）
		countryCode := ""
		for i := 1; i <= 4 && i <= len(s); i++ {
			if validCountryCodes[s[:i]] {
				countryCode = s[:i]
				break
			}
		}

		if countryCode == "" {
			return false // 无效的国家区号
		}

		number := s[len(countryCode):]

		// 根据不同国家验证号码
		switch countryCode {
		case "86": // 中国
			return IsMatch(number, `^1[3-9]\d{9}$`)
		case "1": // 北美
			return IsMatch(number, `^[2-9]\d{9}$`)
		case "852", "853": // 香港、澳门
			return IsMatch(number, `^\d{8}$`)
		default:
			// 其他国家采用通用规则：7-12位数字
			return IsMatch(number, `^\d{7,12}$`)
		}
	}

	// 没有国家区号的情况（默认国内号码）
	return IsMatch(s, `^1[3-9]\d{9}$`)
}
func IsIP(s string) bool {
	return IPReg.MatchString(s)
}

func Dangerous(s string) bool {
	if strings.Contains(s, "<") {
		return true
	}

	if strings.Contains(s, ">") {
		return true
	}

	if strings.Contains(s, "&") {
		return true
	}

	if strings.Contains(s, "'") {
		return true
	}

	if strings.Contains(s, "\"") {
		return true
	}

	if strings.Contains(s, "://") {
		return true
	}

	if strings.Contains(s, "../") {
		return true
	}

	return false
}

package utils

import (
	"regexp"
)

func SubStrByByte(str string, length int) string {
	bt := []byte(str)
	if len(bt) <= length {
		return str
	}
	bs := bt[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++;
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

func SubStrByByteInChar(str string, length int) string {
	s := SubStrByByte(str, length-8)
	if s == str {
		return str
	}
	return SubStrByByte(str, length-8) + ".."
}

func MobileReplaceRepl(str string) string {
	re, _ := regexp.Compile("(\\d{3})(\\d{4})(\\d{4})")
	return re.ReplaceAllString(str, "$1****$3")
}

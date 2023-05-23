package utils_strings

const ConstXAZ = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ConstXaz = "abcdefghijklmnopqrstuvwxyz"
const ConstX09 = "0123456789"
const ConstXAZaz09 = ConstXAZ + ConstXaz + ConstX09
const ConstXAZaz = ConstXAZ + ConstXaz
const ConstXAZ09 = ConstXAZ + ConstX09
const ConstXaz09 = ConstXaz + ConstX09

func SOrX(s, x string) string {
	if s == "" {
		return x
	}
	return s
}

func SOrA(x ...string) string {
	for _, v := range x {
		if v != "" {
			return v
		}
	}
	return ""
}

func SOrF(s string, x func() string) string {
	if s == "" {
		return x()
	}
	return s
}

func In(s string, v ...string) bool {
	for _, x := range v {
		if s == x {
			return true
		}
	}
	return false
}

// CutN 剪切字符串至某个长度，假如长度小于n就返回原字符串
func CutN(s string, n int) string {
	if len(s) > n {
		return s[0:n]
	}
	return s
}

// PadN 在尾部填充字符串直到到达长度，但结果的长度有可能超过n
func PadN(s string, n int, x string) (res string) {
	if n <= 0 || len(s) >= n {
		return s
	}
	if x == "" {
		panic("no padding string")
	}
	res = s
	for len(s) < n {
		res += x
	}
	return res
}

// PadNFixLen 在尾部填充字符串直到到达长度，假如长度已超过n会被截取为n
func PadNFixLen(s string, n int, x string) (res string) {
	if n <= 0 {
		return ""
	}
	if len(s) >= n {
		return s[:n] //也会对原字符串进行截取
	}
	if x == "" {
		panic("no padding string")
	}
	res = s
	for len(s) < n {
		res += x
	}
	return res[:n]
}

// NewRepeatPattern 创造重复循环的字符串，比如"aaa"，或者"***"，或者"abc-abc-abc-"
func NewRepeatPattern(s string, n int) (res string) {
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}

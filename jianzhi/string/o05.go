package string

import "strings"

func replaceSpace(s string) string {
	//var res strings.Builder
	//for i := range s {
	//	if s[i] == ' ' {
	//		res.WriteString("%20")
	//	} else {
	//		res.WriteByte(s[i])
	//	}
	//}
	//return res.String()

	return strings.Replace(s, " ", "%20", -1)
}

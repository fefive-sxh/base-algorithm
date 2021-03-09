package ReplaceSpaces

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"

// 解法: 新建一个byte切片,
// rune():将字符串转换为ascii码数组
// string()将rune切片转换为字符串
// 对应符号的ascii码值
//  %:37; 2:50; 0:48

func ReplaceSpaces(s string) string {
	var b []rune
	for _, v := range s {
		// 获得 字符串中 ' ' 的rune 的值
		//if v == rune(' ') {
		//	b = append(b, rune('%'), rune('2'), rune('0'))
		//} else {
		//	b = append(b, v)
		//}

		var space byte = ' '
		if v == byteToASCII(space) {
			b = append(b, byteToASCII('%'), byteToASCII('2'), byteToASCII('0'))
		} else {
			b = append(b, v)
		}
	}
	return string(b)
}

// 若不记得 % 2 0 的 ascii码, 可以转换一下
func byteToASCII(s byte) rune {
	// var s byte = ''
	// 注意单引号
	return rune(s)
}

// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package lute

import (
	"unicode/utf8"
)

// lexer 描述了词法分析器结构。
type lexer struct {
	input  []byte // 输入的文本字节数组
	length int    // 输入的文本字节数组的长度
	offset int    // 当前读取字节位置
	ln     int    // 当前行号
	col    int    // 当前列号
	width  int    // 最新一个 token 的宽度（字节数）
}

// nextLine 返回下一行。
func (l *lexer) nextLine() (ret items) {
	if l.offset >= l.length {
		return
	}

	l.ln++
	l.col = 0

	var b, nb byte
	i := l.offset
	for ; i < l.length; i += l.width {
		b = l.input[i]
		l.col++
		if itemNewline == b {
			i++
			ret = append(ret, &item{term: b, ln: l.ln, col: l.col})
			break
		} else if itemCarriageReturn == b {
			// 按照规范定义的 line ending (https://spec.commonmark.org/0.29/#line-ending) 处理 \r
			ret = append(ret, &item{term: b, ln: l.ln, col: l.col})
			if i < l.length-1 {
				nb = l.input[i+1]
				if itemNewline == nb {
					l.input = append(l.input[:i], l.input[i+1:]...) // 移除 \r，依靠下一个的 \n 切行
					l.length--                                      // 重新计算总长
					ret = ret[:len(ret)-1]
					ret = append(ret, &item{term: nb, ln: l.ln, col: l.col})
				}
			}
			i++
			break
		} else if '\u0000' == b {
			// 将 \u0000 替换为 \uFFFD https://spec.commonmark.org/0.29/#insecure-characters

			l.input = append(l.input, 0, 0)
			copy(l.input[i+2:], l.input[i:])
			// \uFFFD 的 UTF-8 编码为 \xEF\xBF\xBD 共三个字节
			l.input[i] = '\xEF'
			l.input[i+1] = '\xBF'
			l.input[i+2] = '\xBD'
			l.length += 2 // 重新计算总长
			l.width = 3
			ret = append(ret, &item{term: l.input[i], ln: l.ln, col: l.col})
			ret = append(ret, &item{term: l.input[i+1], ln: l.ln, col: l.col})
			ret = append(ret, &item{term: l.input[i+2], ln: l.ln, col: l.col})
			continue
		}

		if utf8.RuneSelf <= b { // 说明占用多个字节
			_, l.width = utf8.DecodeRune(l.input[i:])
			for j := 0; j < l.width; j++ {
				ret = append(ret, &item{term: l.input[i+j], ln: l.ln, col: l.col})
			}
		} else {
			l.width = 1
			ret = append(ret, &item{term: b, ln: l.ln, col: l.col})
		}
	}
	l.offset = i
	return
}

// newLexer 创建一个词法分析器。
func newLexer(input []byte) (ret *lexer) {
	ret = &lexer{}
	ret.input = input
	ret.length = len(input)

	if 0 < ret.length && itemNewline != ret.input[ret.length-1] {
		// 以 \n 结尾预处理
		ret.input = append(ret.input, itemNewline)
		ret.length++
	}

	return
}

package common

import (
	"strings"
	"unicode"
)

// 文字フォーマットを整える
func SplitWrong(wrong string) (string, string) {
	// 文字列を走査して数字が始まる場所を見つける
	var splitIndex int
	for i, r := range wrong {
		if unicode.IsDigit(r) {
			splitIndex = i
			break
		}
	}

	// 数字の部分とそれ以降を分割
	beforeDigits := wrong[:splitIndex]
	afterDigits := wrong[splitIndex:]

	// '_' を見つけて分割
	parts := strings.SplitN(afterDigits, "_", 2)

	// 分割した文字列を再結合
	result1 := beforeDigits        // "abc"
	result2 := parts[1] + parts[0] // "d323"

	return result1, result2
}

// 文字フォーマットを整える
func FormatInfo(name, category string) string {
	// Category を分割（例: "d303" -> "303", "d"）
	var digitsPart, lettersPart string
	for i, r := range category {
		if unicode.IsDigit(r) {
			digitsPart = category[i:]
			lettersPart = strings.ToUpper(category[:i])
			break
		}
	}

	// Name を大文字にして、フォーマットされた Category と組み合わせる
	// 例: "abc" -> "ABC", "303", "D" -> "ABC303 - D"
	formattedName := strings.ToUpper(name)
	return formattedName + digitsPart + " - " + lettersPart
}

// 内積を計算
func DotProduct(a, b []float64) float64 {
	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

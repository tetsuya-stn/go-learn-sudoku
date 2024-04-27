package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Boardは数独の9x9のマスを示す
// 0: 未入力
// 1-9: 入力済み
type Board [9][9]int

// Boardを9x9のマス形式で表示する
func printBoard(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa((b[i][j])))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")

	return buf.String()
}

func duplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			// 0は未入力なので重複していてもOK
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Board) bool {
	// 行チェック
	for i := 0; i < 9; i++ {
		var c [10]int // 数字(0~9)の出現回数
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}
	// 列チェック
	for i := 0; i < 9; i++ {
		var c [10]int // 数字(0~9)の出現回数を記録
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if duplicated(c) {
			return false
		}
	}
	// 3x3マスのチェック
	
	return true
}

func main() {
	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Printf("%v\n", printBoard(b))
}

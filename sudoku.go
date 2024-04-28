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
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int // 数字(0~9)の出現回数を記録
			for row := i; row < i+3; row++ {
				for column := j; column < j+3; column++ {
					c[b[row][column]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}
	return true
}

func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// backtrackは再帰で呼び出してBoardの中身を更新するため参照渡し
func backtrack(b *Board) bool {
	fmt.Printf("progress\n%v\n", printBoard(*b))
	if solved(*b) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for c := 9; c > 0; c-- {
					b[i][j] = c
					if verify(*b) {
						// もし埋めた数字がチェックをクリアしたら、さらに深く探索
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {
	// b := Board{
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }
	b := Board{
		{0, 5, 0, 0, 8, 3, 0, 1, 7},
		{0, 0, 0, 1, 0, 0, 4, 0, 0},
		{3, 0, 4, 0, 0, 5, 6, 0, 8},
		{0, 0, 0, 0, 3, 0, 0, 0, 9},
		{0, 9, 0, 8, 2, 4, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 7, 0},
		{0, 0, 9, 0, 0, 0, 0, 5, 0},
		{0, 0, 7, 2, 9, 0, 0, 8, 6},
		{1, 0, 3, 6, 0, 7, 2, 0, 4},
	}
	fmt.Printf("before\n%v\n", printBoard(b))
	backtrack(&b)
	fmt.Printf("after\n%v\n", printBoard(b))
}

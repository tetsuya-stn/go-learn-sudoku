package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDuplicated(t *testing.T) {
	if duplicated([10]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}) {
		t.Fatal("not duplicated but failed")
	}

	if !duplicated([10]int{
		0, 2, 0, 0, 0, 0, 0, 0, 0, 0,
	}) {
		t.Fatal("duplicated but failed")
	}
}

func TestVerify(t *testing.T) {
	cases := []struct {
		msg      string
		b        Board
		expected bool
	}{
		{
			msg: "all zero",
			b: Board{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: true,
		},
		{
			msg: "row check",
			b: Board{
				{0, 1, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			msg: "column check",
			b: Board{
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			msg: "3x3 check",
			b: Board{
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k, v.msg), func(t *testing.T) {
			got := verify(v.b)
			if got != v.expected {
				t.Errorf("want %v, but got %v", v.expected, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	// テストケース: https://github.com/t-dillon/tdoku/blob/master/test/test_puzzles
	{
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
		if !Backtrack(&b) {
			t.Fatal("should solve but couldn't")
		}
		fmt.Printf("result\n%v\n", printBoard(b))
	}
	{
		b := Board{
			// 行チェックエラーとなっている問題
			{5, 5, 0, 0, 8, 3, 0, 1, 7},
			{0, 0, 0, 1, 0, 0, 4, 0, 0},
			{3, 0, 4, 0, 0, 5, 6, 0, 8},
			{0, 0, 0, 0, 3, 0, 0, 0, 9},
			{0, 9, 0, 8, 2, 4, 5, 0, 0},
			{0, 0, 6, 0, 0, 0, 0, 7, 0},
			{0, 0, 9, 0, 0, 0, 0, 5, 0},
			{0, 0, 7, 2, 9, 0, 0, 8, 6},
			{1, 0, 3, 6, 0, 7, 2, 0, 4},
		}
		if Backtrack(&b) {
			t.Fatal("shouldn't solve but could")
		}
		fmt.Printf("result\n%v\n", printBoard(b))
	}
}

func TestMakeBoard(t *testing.T) {
	b, err := makeBoard(".5..83.17...1..4..3.4..56.8....3...9.9.8245....6....7...9....5...729..861.36.72.4")
	if err != nil {
		t.Fatalf("parse failed: %s", err)
	}
	expected := Board{
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
	if !reflect.DeepEqual(*b, expected) {
		t.Fatalf("want %v ,but got %v", expected, b)
	}
}

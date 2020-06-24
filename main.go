package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// ランダムな遺伝子配列を作成する count = 遺伝子（配列）の長さ
func mkaeRandGenom(count int, group int) [][]int {
	genoms := make([][]int, group)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < group; i++ {
		genoms[i] = make([]int, count)
		for j := 0; j < count; j++ {
			genoms[i][j] = rand.Intn(2)
		}
	}
	return genoms
}

// 遺伝子（配列）を評価する
func _evaluationGenom(genom []int) int {
	total := 0
	for i := 0; i < len(genom); i++ {
		total = total + genom[i]
	}
	return total
}

func insertEvaluationGenom(genoms [][]int) []int {
	var b []int
	asd := make(map[int]int)
	for i := 0; i < len(genoms); i++ {
		res := _evaluationGenom(genoms[i])
		asd[i] = res
		b = append(b, res)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	var c []int
	for i := 0; i < len(b); i++ {
		for j, v := range asd {
			if b[i] == v {
				c = append(c, j)
				delete(asd, j)
				continue
			}
		}
	}
	return c
}
func genomMutation() {

}

// 任意の確率(％)で戻り値がtrue or false
func intProbability(per int) bool { // per は%表示する
	rand.Seed(time.Now().UnixNano())
	randnum := rand.Intn(100)

	x := make([]int, 100, 100)
	for i := 0; i < per; i++ {
		x[i] = 1
	}
	n := x[randnum]
	if n == 1 {
		return true
	}
	return false
}

// CrossoverOfGenes 交叉させる
func CrossoverOfGenes(s1 []int, s2 []int) []int {
	a := s1[0 : len(s1)/2]
	b := s2[len(s2)/2:]

	s3 := append(a, b...)
	return s3
}
func main() {

	s1 := mkaeRandGenom(4, 4)
	fmt.Println(s1)
	// fmt.Println(CrossoverOfGenes(s1[0], s1[1]))
	fmt.Println(insertEvaluationGenom(s1))
}

package main

import (
	"fmt"
)

type SeqElemT int8
type Seq []SeqElemT

var INVALID_SEQ = Seq(nil)

func (s Seq) isValid() bool {
	return s != nil
}

func newSeq(head SeqElemT, remainings Seq) Seq {
	if !remainings.isValid() {
		return nil
	}
	r := make(Seq, len(remainings)+1)
	r[0] = head
	copy(r[1:], remainings)
	return r
}

// is l > r ?
func greater(l, r Seq) bool {
	if !l.isValid() {
		return false
	}
	if !r.isValid() {
		return true
	}

	for i := 0; i < len(l); i++ {
		if l[i] == r[i] {
			continue
		}
		return l[i] > r[i]
	}
	return false
}

func maxNumber(num1 []int, num2 []int, k int) []int {
	len1 := len(num1)
	len2 := len(num2)
	fmt.Println("len1", len1, "len2", len2, "k", k)
	A := make([][][]Seq, len1+1)
	for i := 0; i <= len1; i++ {
		A[i] = make([][]Seq, len2+1)
		for j := 0; j <= len2; j++ {
			A[i][j] = make([]Seq, k+1)
			for z := 1; z <= k; z++ {
				A[i][j][z] = INVALID_SEQ
			}
		}
	}
	// A[i][j][k] 表示从 num1 的 i 位置到数组结尾 + 从 num2 的 j 位置到数组结尾所有构造出来的长度为 k 的最大数值
	// 则 A[0][0][k] 就是答案
	// 忽略边界条件，A[i][j][k] = max(
	//  A[i+1][j][k-1], 取 num1[i] 作为最高位数字的情况
	//  A[i+1][j][k], 放弃 num1[i]
	//  A[i][j+1][k-1], 取 num2[j] 作为最高位数字的情况
	//  A[i][j+1][k], 放弃 num2[j]
	// )

	// 显然 k = 0 时，A[i][j][0] = 0，0 <= i <= len1, 0 <= j <= len2
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			// empty but valid Seq
			A[i][j][0] = Seq{}
		}
	}
	// k = 1 时，num1 不可取时，A[len1][len2-n][1] = max(num2的最后n个数字)
	snapshotK := k
	for k := 1; k <= snapshotK; k++ {
		maxNum2 := 0
		maxj := 0
		for n := k; n <= len2; n++ {
			if num2[len2-n] >= maxNum2 {
				maxNum2 = num2[len2-n]
				maxj = len2 - n
			}
			A[len1][len2-n][k] = newSeq(SeqElemT(maxNum2), A[len1][maxj+1][k-1])
		}
		// 同理
		maxNum1 := 0
		maxi := 0
		for n := k; n <= len1; n++ {
			if num1[len1-n] >= maxNum1 {
				maxNum1 = num1[len1-n]
				maxi = len1 - n
			}
			A[len1-n][len2][k] = newSeq(SeqElemT(maxNum1), A[maxi+1][len2][k-1])
		}
		for i := len1 - 1; i >= 0; i-- {
			for j := len2 - 1; j >= 0; j-- {
				a := newSeq(SeqElemT(num1[i]), A[i+1][j][k-1])
				a2 := A[i+1][j][k]
				b := newSeq(SeqElemT(num2[j]), A[i][j+1][k-1])
				b2 := A[i][j+1][k]
				temp := []Seq{a, a2, b, b2}
				tempMax := INVALID_SEQ
				for t := 0; t < len(temp); t++ {
					if greater(temp[t], tempMax) {
						tempMax = temp[t]
					}
				}
				A[i][j][k] = tempMax
			}
		}

		/*
			fmt.Println("k = ", k)
			for i := 0; i <= len1; i++ {
				for j := 0; j <= len2; j++ {
					fmt.Printf("%v ", A[i][j][k])
				}
				fmt.Println()
			}
			fmt.Println()
			fmt.Println()
		*/
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = int(A[0][0][k][i])
	}
	return result
}

func main() {

	fmt.Println(maxNumber([]int{7, 9, 0, 4, 7, 0, 7, 1, 2, 9, 5}, []int{9, 1, 5, 8, 3, 9}, 17))
	fmt.Println(maxNumber([]int{8, 5, 9, 5, 1, 6, 9}, []int{2, 6, 4, 3, 8, 4, 1, 0, 7, 2, 9, 2, 8}, 20))

	fmt.Println(maxNumber([]int{8, 9, 7, 3, 5, 9, 1, 0, 8, 5, 3, 0, 9, 2, 7, 4, 8, 9, 8, 1, 0, 2, 0, 2, 7, 2, 3, 5, 4, 7, 4, 1, 4, 0, 1, 4, 2, 1, 3, 1, 5, 3, 9, 3, 9, 0, 1, 7, 0, 6, 1, 8, 5, 6, 6, 5, 0, 4, 7, 2, 9, 2, 2, 7, 6, 2, 9, 2, 3, 5, 7, 4, 7, 0, 1, 8, 3, 6, 6, 3, 0, 8, 5, 3, 0, 3, 7, 3, 0, 9, 8, 5, 1, 9, 5, 0, 7, 9, 6, 8, 5, 1, 9, 6, 5, 8, 2, 3, 7, 1, 0, 1, 4, 3, 4, 4, 2, 4, 0, 8, 4, 6, 5, 5, 7, 6, 9, 0, 8, 4, 6, 1, 6, 7, 2, 0, 1, 1, 8, 2, 6, 4, 0, 5, 5, 2, 6, 1, 6, 4, 7, 1, 7, 2, 2, 9, 8, 9, 1, 0, 5, 5, 9, 7, 7, 8, 8, 3, 3, 8, 9, 3, 7, 5, 3, 6, 1, 0, 1, 0, 9, 3, 7, 8, 4, 0, 3, 5, 8, 1, 0, 5, 7, 2, 8, 4, 9, 5, 6, 8, 1, 1, 8, 7, 3, 2, 3, 4, 8, 7, 9, 9, 7, 8, 5, 2, 2, 7, 1, 9, 1, 5, 5, 1, 3, 5, 9, 0, 5, 2, 9, 4, 2, 8, 7, 3, 9, 4, 7, 4, 8, 7, 5, 0, 9, 9, 7, 9, 3, 8, 0, 9, 5, 3, 0, 0, 3, 0, 4, 9, 0, 9, 1, 6, 0, 2, 0, 5, 2, 2, 6, 0, 0, 9, 6, 3, 4, 1, 2, 0, 8, 3, 6, 6, 9, 0, 2, 1, 6, 9, 2, 4, 9, 0, 8, 3, 9, 0, 5, 4, 5, 4, 6, 1, 2, 5, 2, 2, 1, 7, 3, 8, 1, 1, 6, 8, 8, 1, 8, 5, 6, 1, 3, 0, 1, 3, 5, 6, 5, 0, 6, 4, 2, 8, 6, 0, 3, 7, 9, 5, 5, 9, 8, 0, 4, 8, 6, 0, 8, 6, 6, 1, 6, 2, 7, 1, 0, 2, 2, 4, 0, 0, 0, 4, 6, 5, 5, 4, 0, 1, 5, 8, 3, 2, 0, 9, 7, 6, 2, 6, 9, 9, 9, 7, 1, 4, 6, 2, 8, 2, 5, 3, 4, 5, 2, 4, 4, 4, 7, 2, 2, 5, 3, 2, 8, 2, 2, 4, 9, 8, 0, 9, 8, 7, 6, 2, 6, 7, 5, 4, 7, 5, 1, 0, 5, 7, 8, 7, 7, 8, 9, 7, 0, 3, 7, 7, 4, 7, 2, 0, 4, 1, 1, 9, 1, 7, 5, 0, 5, 6, 6, 1, 0, 6, 9, 4, 2, 8, 0, 5, 1, 9, 8, 4, 0, 3, 1, 2, 4, 2, 1, 8, 9, 5, 9, 6, 5, 3, 1, 8, 9, 0, 9, 8, 3, 0, 9, 4, 1, 1, 6, 0, 5, 9, 0, 8, 3, 7, 8, 5}, []int{7, 8, 4, 1, 9, 4, 2, 6, 5, 2, 1, 2, 8, 9, 3, 9, 9, 5, 4, 4, 2, 9, 2, 0, 5, 9, 4, 2, 1, 7, 2, 5, 1, 2, 0, 0, 5, 3, 1, 1, 7, 2, 3, 3, 2, 8, 2, 0, 1, 4, 5, 1, 0, 0, 7, 7, 9, 6, 3, 8, 0, 1, 5, 8, 3, 2, 3, 6, 4, 2, 6, 3, 6, 7, 6, 6, 9, 5, 4, 3, 2, 7, 6, 3, 1, 8, 7, 5, 7, 8, 1, 6, 0, 7, 3, 0, 4, 4, 4, 9, 6, 3, 1, 0, 3, 7, 3, 6, 1, 0, 0, 2, 5, 7, 2, 9, 6, 6, 2, 6, 8, 1, 9, 7, 8, 8, 9, 5, 1, 1, 4, 2, 0, 1, 3, 6, 7, 8, 7, 0, 5, 6, 0, 1, 7, 9, 6, 4, 8, 6, 7, 0, 2, 3, 2, 7, 6, 0, 5, 0, 9, 0, 3, 3, 8, 5, 0, 9, 3, 8, 0, 1, 3, 1, 8, 1, 8, 1, 1, 7, 5, 7, 4, 1, 0, 0, 0, 8, 9, 5, 7, 8, 9, 2, 8, 3, 0, 3, 4, 9, 8, 1, 7, 2, 3, 8, 3, 5, 3, 1, 4, 7, 7, 5, 4, 9, 2, 6, 2, 6, 4, 0, 0, 2, 8, 3, 3, 0, 9, 1, 6, 8, 3, 1, 7, 0, 7, 1, 5, 8, 3, 2, 5, 1, 1, 0, 3, 1, 4, 6, 3, 6, 2, 8, 6, 7, 2, 9, 5, 9, 1, 6, 0, 5, 4, 8, 6, 6, 9, 4, 0, 5, 8, 7, 0, 8, 9, 7, 3, 9, 0, 1, 0, 6, 2, 7, 3, 3, 2, 3, 3, 6, 3, 0, 8, 0, 0, 5, 2, 1, 0, 7, 5, 0, 3, 2, 6, 0, 5, 4, 9, 6, 7, 1, 0, 4, 0, 9, 6, 8, 3, 1, 2, 5, 0, 1, 0, 6, 8, 6, 6, 8, 8, 2, 4, 5, 0, 0, 8, 0, 5, 6, 2, 2, 5, 6, 3, 7, 7, 8, 4, 8, 4, 8, 9, 1, 6, 8, 9, 9, 0, 4, 0, 5, 5, 4, 9, 6, 7, 7, 9, 0, 5, 0, 9, 2, 5, 2, 9, 8, 9, 7, 6, 8, 6, 9, 2, 9, 1, 6, 0, 2, 7, 4, 4, 5, 3, 4, 5, 5, 5, 0, 8, 1, 3, 8, 3, 0, 8, 5, 7, 6, 8, 7, 8, 9, 7, 0, 8, 4, 0, 7, 0, 9, 5, 8, 2, 0, 8, 7, 0, 3, 1, 8, 1, 7, 1, 6, 9, 7, 9, 7, 2, 6, 3, 0, 5, 3, 6, 0, 5, 9, 3, 9, 1, 1, 0, 0, 8, 1, 4, 3, 0, 4, 3, 7, 7, 7, 4, 6, 4, 0, 0, 5, 7, 3, 2, 8, 5, 1, 4, 5, 8, 5, 6, 7, 5, 7, 3, 3, 9, 6, 8, 1, 5, 1, 1, 1, 0, 3}, 500))
}

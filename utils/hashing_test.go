/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"math/rand"
	"testing"
)

func GetRandomInts(length int, maxInt int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(maxInt) + 1
	}
	return nums
}

func TestSzudzikPairing(t *testing.T) {
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			pair := SzudzikPairing(i, j)
			unpair_i, unpair_j := SzudzikUnpairing(pair)
			if i != unpair_i || j != unpair_j {
				t.Errorf("SzudzikPairing(%d, %d) = %d, SzudzikUnpairing(%d) = (%d, %d)", i, j, pair, pair, unpair_i, unpair_j)
			}
		}
	}
}

func TestTwoDToOneD(t *testing.T) {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			for k := i + 1; k <= 1000; k++ {
				pair := TwoDToOneD(i, j, k)
				unpair_i, unpair_j := OneDTwoD(pair, k)
				if i != unpair_i || j != unpair_j {
					t.Errorf("TwoDToOneD(%d, %d, %d) = %d, OneDTwoD(%d, %d) = (%d, %d)", i, j, k, pair, pair, k, unpair_i, unpair_j)
				}
			}
		}
	}
}

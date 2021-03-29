package leetcode

import "math/bits"

/*
颠倒给定的 32 位无符号整数的二进制位。
*/
func reverseBits(n uint32) uint32 {
	return bits.Reverse32(n)
}

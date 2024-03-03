package fibonaccimodified

import "math/big"

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/fibonacci-modified/problem
 */

func fibonacciModified(t1 int32, t2 int32, n int32) *big.Int {
	// Write your code here
	arr := make([]*big.Int, 0, n)
	arr = append(arr, []*big.Int{big.NewInt(int64(t1)), big.NewInt(int64(t2))}...)

	for i := int32(2); i<n; i++ {
		val := big.NewInt(0)
		mul := big.NewInt(1)
		mul.Mul(arr[i-1], arr[i-1])
		val.Add(arr[i-2], mul)
		arr = append(arr, val)
	}

	return arr[n - 1]

}
package main

import (
	"fmt"
	"math"
)

func findMaxCrossingSubarray(a []int, low, mid, high int) (maxLeft, maxRight, sum int) {
	leftSum := math.MinInt64
	for i := mid; i >= low; i-- {
		sum += a[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := math.MinInt64
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += a[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	sum = leftSum + rightSum
	return
}

func findMaximumSubarray(a []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, a[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaximumSubarray(a, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(a, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(a, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}
	if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum

}

func main() {
	a := []int{-13, -3, -25, -20, -3, -16, -23, -18, -20, -13, -7, -12, -5, -22, -15, -4, -7}
	fmt.Println(findMaximumSubarray(a, 0, len(a)-1))
}

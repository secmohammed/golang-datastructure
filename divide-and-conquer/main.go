package main

func maxSubArray(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums, 0, right)
}

func maxCrossing(nums []int, left int, right int, mid int) int {
	leftSum := -2147483648
	rightSum := -2147483648
	sum := 0
	for i := mid - 1; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}

func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func findMaxSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftMax := findMaxSubArray(nums, left, mid)
	rightMax := findMaxSubArray(nums, mid+1, right)
	crossMax := maxCrossing(nums, left, right, mid)
	return max(leftMax, rightMax, crossMax)
}

func main() {
	problems := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	println(maxSubArray(problems))
}

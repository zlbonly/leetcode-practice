package main

/**
一、整理总结二分查找相关算法
	1、升序数组二分查找（元素无重复）
	2、升序数组二分查找 （元素存在重复，返回查找target第一次出现的索引）
	3、搜索旋转数组最小值（无重复元素）
	4、搜寻找旋转排序数组中的最小值 II（有重复元素）
	5、搜索旋转排序数组(元素互不相同)

*/

/***
题目描述：二分查找I
给定⼀一个 n 个元素有序的(升序)整型数组 nums 和⼀一个⽬目标值 target ，写⼀一个函数搜索 nums 中的 target，如果⽬目标值存在返回 下标，否则返回 -1。
示例例 1:
输⼊入: nums = [-1,0,3,5,9,12], target = 9 输出: 4
解释: 9 出现在 nums 中并且下标为 4
*/
func binarySearch(nums []int, target int) int {
	low, high, pivot := 0, len(nums)-1, 0
	for low <= high {
		pivot = (high-low)/2 + low
		if nums[pivot] == target {
			return pivot
		} else if target > nums[pivot] {
			low = pivot + 1
		} else if target < nums[pivot] {
			high = pivot - 1
		}
	}
	return -1
}

/**
题目描述：二分查找II
请实现有重复数字的升序数组的⼆二分查找。 输出在数组中第⼀一个,大于等于查找值的位置，
如果数组中不不存在这样的数(指不不存在⼤大于等于查找值的数)，则输出数组中索引。
示例：
	nums :{0,1,2,3,4,5,5,5,5,5,5,5,6,7,8} 输出:5
*/

func binarySearchFirst(nums []int, target int) int {
	left, right, pivot := 0, len(nums)-1, 0
	for left < right {
		pivot = left + (right-left)/2
		if target > nums[pivot] {
			left = pivot + 1
		} else {
			right = pivot
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

/**
题目描述：寻找旋转排序数组中的最小值 II （不存在重复元素）
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
*/
func findMinI(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}

/**
	题目描述：寻找旋转排序数组中的最小值 II （存在重复元素）
	已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
	若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
	若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
	注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
	给你一个可能存在 重复 元素值的数组 nums ，
	它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
*/
func findMinUniqueII(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2

		//  最小值一定在中间值nums[pivot]和最右侧值nums[pivot]中间，且不包括nums[pivot]，因此可以跳过。
		if nums[pivot] > nums[high] {
			low = pivot + 1
		} else if nums[pivot] < nums[high] {
			// 最小值一定在最左侧值nums[low]和中间值nums[pivot]之间，但是不确定nums[pivot]是否最小值，因此不能跳过。
			high = pivot
		} else {
			// 中间值nums[pivot]和最右侧值nums[high]相等，没法确定最小值位置，但是，nums[high]肯定有
			// nums[pivot]可以替换，因此可以忽略右端点。
			high--
		}
	}
	return nums[low]
}

/**
 题目描述：搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

示例：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
*/
func searchTarget(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		pivot := low + (high-low)/2
		if nums[pivot] == target {
			return pivot
		}
		// 先根据 nums[mid] 与 nums[lo] 的关系判断 mid 是在左段还是右段
		if nums[pivot] >= nums[low] {
			// 再判断 target 是在 mid 的左边还是右边，从而调整左右边界 lo 和 hi
			if target >= nums[low] && target < nums[pivot] {
				high = pivot - 1
			} else {
				low = pivot + 1
			}
		} else {
			if target > nums[pivot] && target <= nums[high] {
				low = pivot + 1
			} else {
				high = pivot - 1
			}
		}
	}
	return -1
}

#### 题目
<p>假设按照升序排序的数组在预先未知的某个点上进行了旋转。</p>

<p>( 例如，数组&nbsp;<code>[0,1,2,4,5,6,7]</code> <strong> </strong>可能变为&nbsp;<code>[4,5,6,7,0,1,2]</code>&nbsp;)。</p>

<p>请找出其中最小的元素。</p>

<p>注意数组中可能存在重复的元素。</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入:</strong> [1,3,5]
<strong>输出:</strong> 1</pre>

<p><strong>示例&nbsp;2：</strong></p>

<pre><strong>输入:</strong> [2,2,2,0,1]
<strong>输出:</strong> 0</pre>

<p><strong>说明：</strong></p>

<ul>
	<li>这道题是&nbsp;<a href="https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/">寻找旋转排序数组中的最小值</a>&nbsp;的延伸题目。</li>
	<li>允许重复会影响算法的时间复杂度吗？会如何影响，为什么？</li>
</ul>


 #### 题解
 会存在类似于[1,1,0,0,1,1]的这种例子,这种情况下找到中点没有办法判断舍弃哪一边，
 所以我们直接将右边界向左移来处理这种情况。
 ```go
func findMin(nums []int) int {
	left,right := 0,len(nums)-1
	for left < right {
		mid := left + (right - left)/2
		if nums[mid] > nums[right] {
			left = mid+1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
```
 时间复杂度O(logn),空间复杂度O(1)
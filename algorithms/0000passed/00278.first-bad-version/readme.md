#### 题目
<p>你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。</p>

<p>假设你有 <code>n</code> 个版本 <code>[1, 2, ..., n]</code>，你想找出导致之后所有版本出错的第一个错误的版本。</p>

<p>你可以通过调用&nbsp;<code>bool isBadVersion(version)</code>&nbsp;接口来判断版本号 <code>version</code> 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。</p>

<p><strong>示例:</strong></p>

<pre>给定 n = 5，并且 version = 4 是第一个错误的版本。

<code>调用 isBadVersion(3) -&gt; false
调用 isBadVersion(5)&nbsp;-&gt; true
调用 isBadVersion(4)&nbsp;-&gt; true

所以，4 是第一个错误的版本。&nbsp;</code></pre>


 #### 题解
 1、线性扫描
 ```go
func firstBadVersion(n int) int {
	for i := 1; i < n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return n
}
```
 时间复杂度O(n),空间复杂度O(1)
 
 2、二分法
 ```go
func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```
 时间复杂度O(logn),空间复杂度O(1)
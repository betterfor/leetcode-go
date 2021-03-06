#### 题目
<p>将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>1-&gt;2-&gt;4, 1-&gt;3-&gt;4
<strong>输出：</strong>1-&gt;1-&gt;2-&gt;3-&gt;4-&gt;4
</pre>


 #### 题解
 按照题意合并链表即可
 ```go
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
 	var newList = new(ListNode)
 	pre := newList
 	for l1 != nil && l2 != nil { // 循环到链表末尾
 		if l1.Val <= l2.Val {
 			pre.Next = l1
 			l1 = l1.Next
 		} else {
 			pre.Next = l2
 			l2 = l2.Next
 		}
 		pre = pre.Next
 	}
 	if l1 != nil {
 		pre.Next = l1
 	}
 	if l2 != nil {
 		pre.Next = l2
 	}
 	return newList.Next
 }
 ```
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-14/002101.png)
 时间复杂度O(n)，空间复杂度O(1)
#### 题目
<p>你和你的朋友，两个人一起玩&nbsp;<a href="https://baike.baidu.com/item/Nim游戏/6737105" target="_blank">Nim 游戏</a>：桌子上有一堆石头，每次你们轮流拿掉&nbsp;1 - 3 块石头。 拿掉最后一块石头的人就是获胜者。你作为先手。</p>

<p>你们是聪明人，每一步都是最优解。 编写一个函数，来判断你是否可以在给定石头数量的情况下赢得游戏。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>4</code>
<strong>输出:</strong> false 
<strong>解释: </strong>如果堆中有 4 块石头，那么你永远不会赢得比赛；
&nbsp;    因为无论你拿走 1 块、2 块 还是 3 块石头，最后一块石头总是会被你的朋友拿走。
</pre>


 #### 题解
 如果堆中石头的数量n不能被4整除，那么我作为先手总能获取胜利。
 ```go
func canWinNim(n int) bool {
	return n%4!=0
}
```
 时间复杂度O(1),空间复杂度O(1)
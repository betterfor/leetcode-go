#### 题目
<p>给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换&nbsp;<em>k&nbsp;</em>次。在执行上述操作后，找到包含重复字母的最长子串的长度。</p>

<p><strong>注意:</strong><br>
字符串长度 和 <em>k </em>不会超过&nbsp;10<sup>4</sup>。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong>
s = &quot;ABAB&quot;, k = 2

<strong>输出:</strong>
4

<strong>解释:</strong>
用两个&#39;A&#39;替换为两个&#39;B&#39;,反之亦然。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>
s = &quot;AABABBA&quot;, k = 1

<strong>输出:</strong>
4

<strong>解释:</strong>
将中间的一个&#39;A&#39;替换为&#39;B&#39;,字符串变为 &quot;AABBBBA&quot;。
子串 &quot;BBBB&quot; 有最长重复字母, 答案为 4。
</pre>


 #### 题解
 ## 题目
 
 给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 *k* 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
 
 ## 示例
 
 **示例 1:**
 
 **输入:**
 s ="ABAB", k = 2
 
 **输出:**
 4
 
 **解释:**
 用两个"A"替换为两个"B",反之亦然。
 
 **示例 2:**
 
 **输入:**
 s ="AABABBA", k = 1
 
 **输出:**
 4
 
 **解释:**
 将中间的一个"A"替换为"B",字符串变为 "AABBBBA"。
 子串"BBBB"有最长重复字母, 答案为 4。
 
 ## 题解
 
 1、暴力法(不实现)
 
 思路：
 
 - 如果子串中所有的字符都一样就延伸子串
 - 如果当前子串出现至少两种字符，就要替换使得所有的字符都一样，并且重复、连续的部分更长。
 
 暴力解法的时间复杂度O(n^3^)。
 
 缺点：
 
 做了很多重复的工作，子串和子串有很多重复的部分，重复扫描了很多次。
 
 2、优化方法
 
 优化字符串查找子串，我们能够想到两种方法，动态规划和滑动窗口。而本题动态规划没有得到明显的递推关系，所以是要用滑动窗口。
 
 所以题目的意思可以转化为
 
 *枚举字符串中每个位置作为右端点，然后找到其最左端点的位置，满足该区间内除了出现次数最多的那一类字符外，剩余的字符数量不超过k个。*
 
 ```go
 func characterReplacement(s string, k int) int {
 	var left,maxCnt int	// maxCnt保存整个循环中，cnt出现的最大值
     var cnt [26]int // cnt记录了s[left:right+1]中每个字母出现的次数
     // 在循环中，s[left:right+1]
     // 要么，maxCnt变大，向右移动一格
     // 要么，maxCnt不变，向右移动一格。
 	for right, str := range s {
 		cnt[str-'A']++
 		maxCnt = max(maxCnt,cnt[str-'A'])
         // right-left+1-maxCnt==k的含义是在s[left:right+1]中有maxCnt个相同的字母和k个不相同的字母
 		if right - left+1-maxCnt > k {
 			cnt[s[left]-'A']--
 			left++
 		}
 	}
 	return len(s)-left
 }
 
 func max(a, b int) int {
 	if a > b {
 		return a
 	}
 	return b
 }
 ```
 
 ## 复杂度分析
 
 时间复杂度O(n)，n为字符串大小，我们最多遍历字符串一次
 
 空间复杂度O(∣Σ∣)，其中∣Σ∣是字符集大小，我们需要存储每个大写字母出现的次数
 
 ![执行结果](https://i.loli.net/2021/02/02/l2V3Eb1UxMuIHzK.jpg)
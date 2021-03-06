#### 题目
<p>给你一个字符串&nbsp;<code>s</code>&nbsp;和一个字符规律&nbsp;<code>p</code>，请你来实现一个支持 <code>&#39;.&#39;</code>&nbsp;和&nbsp;<code>&#39;*&#39;</code>&nbsp;的正则表达式匹配。</p>

<pre>&#39;.&#39; 匹配任意单个字符
&#39;*&#39; 匹配零个或多个前面的那一个元素
</pre>

<p>所谓匹配，是要涵盖&nbsp;<strong>整个&nbsp;</strong>字符串&nbsp;<code>s</code>的，而不是部分字符串。</p>

<p><strong>说明:</strong></p>

<ul>
	<li><code>s</code>&nbsp;可能为空，且只包含从&nbsp;<code>a-z</code>&nbsp;的小写字母。</li>
	<li><code>p</code>&nbsp;可能为空，且只包含从&nbsp;<code>a-z</code>&nbsp;的小写字母，以及字符&nbsp;<code>.</code>&nbsp;和&nbsp;<code>*</code>。</li>
</ul>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong>
s = &quot;aa&quot;
p = &quot;a&quot;
<strong>输出:</strong> false
<strong>解释:</strong> &quot;a&quot; 无法匹配 &quot;aa&quot; 整个字符串。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>
s = &quot;aa&quot;
p = &quot;a*&quot;
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;因为 &#39;*&#39; 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 &#39;a&#39;。因此，字符串 &quot;aa&quot; 可被视为 &#39;a&#39; 重复了一次。
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong>
s = &quot;ab&quot;
p = &quot;.*&quot;
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;&quot;.*&quot; 表示可匹配零个或多个（&#39;*&#39;）任意字符（&#39;.&#39;）。
</pre>

<p><strong>示例 4:</strong></p>

<pre><strong>输入:</strong>
s = &quot;aab&quot;
p = &quot;c*a*b&quot;
<strong>输出:</strong> true
<strong>解释:</strong>&nbsp;因为 &#39;*&#39; 表示零个或多个，这里 &#39;c&#39; 为 0 个, &#39;a&#39; 被重复一次。因此可以匹配字符串 &quot;aab&quot;。
</pre>

<p><strong>示例 5:</strong></p>

<pre><strong>输入:</strong>
s = &quot;mississippi&quot;
p = &quot;mis*is*p*.&quot;
<strong>输出:</strong> false</pre>


 #### 题解
 这个题可以用动态规划来解决。
 
 假设 **dp[i] [j]** 表示 *s* 的前 i 个是否能被 *p* 的前 j 个匹配。
 
 那么 **dp[i-1] [j-1]** 表示前面的子串都匹配上了，来匹配下一位的情况。
 
 分情况考虑：
 
 1、s[i] = p[j] ： **dp[i] [j]**  =  **dp[i-1] [j-1]** 
 
 2、p[j] = "." ： **dp[i] [j]**  =  **dp[i-1] [j-1]** 
 
 3、p[j] = "*"：
 
 又需要分情况讨论，* 匹配零个或多个前面的那个元素，所以需要考虑它前面的元素p[j-1]。
 
 3.1、p[j-1] != s[i] ： **dp[i] [j]**  =  **dp[i] [j-2]** ，字符匹配不上，将该元素消除的情况（aab，c* a*b）
 
 3.2、p[j-1] = s[i] 或 p[j-1] = "."：
 
      dp[i][j] = dp[i-1][j] // 多个字符匹配的情况	
      or dp[i][j] = dp[i][j-1] // 单个字符匹配的情况
      or dp[i][j] = dp[i][j-2] // 没有匹配的情况
 
 
 ```go
 func isMatch(s string, p string) bool {
 	sSize := len(s)
 	pSize := len(p)
 
 	dp := make([][]bool, sSize+1)
 	for i := range dp {
 		dp[i] = make([]bool, pSize+1)
 	}
 
 	/* dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配 */
 
 	dp[0][0] = true
 	/**
 	 * 根据题目的设定， "" 可以与 "a*b*c*" 相匹配
 	 * 但""不与 "a*b*c*d" 相匹配
 	 * 所以，需要把奇数位上有 "*" 的 dp 设置成 true
 	 */
 	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
 		if p[j] == '*' {
 			dp[0][j+1] = true
 		}
 	}
 
 	for i := 0; i < sSize; i++ {
 		for j := 0; j < pSize; j++ {
 			if p[j] == '.' || p[j] == s[i] {
 				/* p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上 */
 				dp[i+1][j+1] = dp[i][j]
 			} else if p[j] == '*' {
 				/* 此时，p[j] 的匹配情况与 p[j-1] 的内容相关。 */
 				if p[j-1] != s[i] && p[j-1] != '.' {
 					/**
 					 * p[j] 无法与 s[i] 匹配上
 					 * p[j-1:j+1] 只能被当做 ""
 					 */
 					dp[i+1][j+1] = dp[i+1][j-1]
 				} else {
 					/**
 					 * p[j] 与 s[i] 匹配上
 					 * p[j-1;j+1] 作为 "x*", 可以有三种解释
 					 */
 					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" 解释为 "" */
 						dp[i+1][j] || /* "x*" 解释为 "x" */
 						dp[i][j+1] /* "x*" 解释为 "xx..." */
 				}
 			}
 		}
 	}
 
 	return dp[sSize][pSize]
 }
 ```
 
 ![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-04/001001.png)
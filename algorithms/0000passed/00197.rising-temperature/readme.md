#### 题目
<p>给定一个 <code>Weather</code> 表，编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 Id。</p>

<pre>+---------+------------------+------------------+
| Id(INT) | RecordDate(DATE) | Temperature(INT) |
+---------+------------------+------------------+
|       1 |       2015-01-01 |               10 |
|       2 |       2015-01-02 |               25 |
|       3 |       2015-01-03 |               20 |
|       4 |       2015-01-04 |               30 |
+---------+------------------+------------------+</pre>

<p>例如，根据上述给定的 <code>Weather</code> 表格，返回如下 Id:</p>

<pre>+----+
| Id |
+----+
|  2 |
|  4 |
+----+</pre>


 #### 题解
 SELECT
     weather.id AS 'Id'
 FROM
     weather
         JOIN
     weather w ON DATEDIFF(weather.RecordDate, w.RecordDate) = 1
         AND weather.Temperature > w.Temperature
 ;
 

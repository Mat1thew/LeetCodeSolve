# 思路整理  
## first commit
从1开始遍历计算，每当计算到一个丑数，直接append到切片内，当切片的长度等于n值时，直接输出  
## second commit
第一次提交明显会超时，思考出如下优化：  
1.**使用map存储已经计算过的数字列表，丑数为true，非丑数为false，避免相同数字重复计算**  
2.**整体思路不变，在判断是否为丑数前，先访问该map，若有该数字结果，直接返回**  
提交结果：*页面显示计算第356个丑数TLE*
## third commit
前两次思路行不通，得改变方法，直接从1开始计算丑数。  
丑数都是**2、3、5的倍数**，可以从此下手  
// TODO : 推翻之前的做法，直接计算  
借鉴题目讨论，使用三指针法，最终通过  
解题思路：**第一位丑数为1，分别定义倍数为2、3、5的指针，使用当前索引下标** 
        ***(索引对象为丑数的数组)*** 
        **的丑数值分别乘对应的倍数，选取最小值，插入丑数数组，对应指针做自增**  
[本次提交](https://leetcode-cn.com/problems/ugly-number-ii/submissions/)
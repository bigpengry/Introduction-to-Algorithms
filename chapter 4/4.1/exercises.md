## 4.1-1<br/>
会返回整个个数组中数值最大的元素。当函数执行时，会比较每一个子数组的长度，因为整个子数组的元素为负，所以函数会找到第一个值最大的元素，并返回<br/>
## 4.1-2<br/>
```
FIND-MAXIMUM-SUBARRAY（A,low,high)
max=-∞
if low==high
	return low,heigh,A[low]
for i=low to high
	sum=0
	for j=i to high
		sum=sum+A[j]
		if sum>max
			max=sum
			max-left=i
			max-right=j
return (max-left,max-right,max)

```
## 4.1-3<br/>
现代计算机的性能已经非常高了，我的这台电脑上两种算法在较少元素的情况下也是基本相同，实际上可以无脑使用递归程序。<br/>
## 4.1-4<br/>
个人觉得应该是判断一下整个数组的和是不是都为负，整个数组都为负，这样的结果是没有意义的<br/>
## 4.1-5<br/>
这个代码主要有两个状态：<br/>
1）A[1.. j]为A[1.. j+1]的最大子数组，且A[1.. j]为A[1.. j]的最大子数组
2）A[i..j+1]为A[1.. j+1]的最大子数组
也就是说当出现第一种情况时，当前的最大子数组和上一个迭代的最大子数组相同，因此，我们可以不需要更新最大值，只需要进入下一个循环；当出现第二种情况，我们所要做的就是移动数组的左端。因为只有当一段数组的和与原来的最大子数组相加等于负数时这种情况才可能出现，因此，我们只要跳过这一段数组就可以。
```
FIND-MAXIMUM-SUBARRAY（A,low,high)
sum=0
max-sum=-∞
subleft=low
for j=low to high
	sum=sum+A[j]
	if sum>max-sum
		max-sum=sum
		max-left=subleft
		max-right=j
	if sum<0
		sum=0
		subleft=j+1
return  (max-left,max-right,max-sum)
```

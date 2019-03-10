## 2.1-1<br/>
0.<31,41,59,26,41,58><br/>
1.<31,41,59,26,41,58><br/>
2.<31,41,59,26,41,58><br/>
3.<26,31,41,59,41,58><br/>
4.<26,31,41,41,59,58><br/>
5.<26,31,41,41,58,59><br/>
## 2.1-2<br/>
**for**&emsp;j = 2 **to** A.length<br/>
&emsp;&emsp; key = A [ j ]<br/>
&emsp;&emsp; i = j - 1<br/>
&emsp;&emsp; **while** i > 0 **and** A [ i ] > key<br/>
&emsp;&emsp;&emsp;&emsp; A [ i + 1 ] = A [ i ]<br/>
&emsp;&emsp;&emsp;&emsp; i = i - 1<br/>
&emsp;&emsp; A [ i + 1 ] = key <br/>
## 2.1-3<br/>
**for** j = 1 **to** A.length<br>
&emsp;&emsp; **if** v = A [ i ]<br/>
&emsp;&emsp;&emsp;&emsp; **return** v =  i <br/>
&emsp;&emsp; **return** v = nil <br/>
**循环不变式**<br/>
初始化：A [ 1..i-1 ]由不等于v的元素组成。刚开始时，由于未读入任何元素，所以为0。<br/>
保持：for循环依次与v进行比较，如果不等于就继续执行，所以A[ 1..i-1]中一直不等于v的元素。<br/>
终止：终止有两种情况，直到某个元素等于v，此时A[ i ]等于v，而A[ 1..i-1 ]中仍未不等于v的元素。另一种情况是，整个数组遍历完毕，都没找到与v相等的元素，此时A [ 1..i-1] 的范围为整个数组。<br/>
## 2.1-4<br/>
//这道题虽然可以用取余做，但是取余操作是一个不确定是否为一条常量时间的指令，因此要避开它<br/>
//从后往前加是因为这样可以正序输出，而且比较方便处理最高位
**for** i = n + 1 **downto** 1<br/>
&emsp;&emsp; **if** A [ i - 1 ] + B [ i - 1 ] + C [ i ] = 3 <br/>
&emsp;&emsp;&emsp;&emsp; C [ i - 1 ] = 1<br/>
&emsp;&emsp;&emsp;&emsp; C [ i ] = 1 <br/>
&emsp;&emsp; **else if** A [ i - 1 ] + B [ i - 1 ]+ C [ i ] = 2 <br/>
&emsp;&emsp;&emsp;&emsp; C [ i - 1 ] = 1<br/>
&emsp;&emsp;&emsp;&emsp; C [ i ] = 0 <br/>
&emsp;&emsp; **else if** A [ i - 1 ] + B [ i - 1 ]+ C [ i ] = 1 <br/>
&emsp;&emsp;&emsp;&emsp; C [ i - 1 ] = 0<br/>
&emsp;&emsp;&emsp;&emsp; C [ i ] = 1 <br/>
&emsp;&emsp; **else** <br/>
&emsp;&emsp;&emsp;&emsp; C [ i - 1 ] = 0<br/>
&emsp;&emsp;&emsp;&emsp; C [ i ] = 0 <br/>

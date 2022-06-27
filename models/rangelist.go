package  models

import "fmt"

type RangeList struct {
	ranges [][]int
}

//Add 方法
func (rangeList *RangeList) Add(rangeElement []int)  {
	if rangeElement !=nil&& len(rangeElement)>1{
		start,stop:=rangeElement[0],rangeElement[1]
		if start<stop{
			n:=len(rangeList.ranges)
			for i:=0;i<n;i++{
				item:=rangeList.ranges[i]
				a:=item[0]
				b:=item[1]
				if stop<a{
					x:=rangeList.ranges
					rangeList.ranges=append([][]int{},rangeElement)
					rangeList.ranges=append(rangeList.ranges,x...)
					return
				}
				if  start<=b{
					r:=mergeRange(rangeElement,item)
					rangeList.ranges=append(rangeList.ranges[0:i],rangeList.ranges[i+1:]...)
					rangeList.Add(r)
					return
				}
			}
			rangeList.ranges=append(rangeList.ranges,rangeElement)
		}
	}

}

//Remove 方法
func (rangeList *RangeList) Remove(rangeElement []int)  {
	if rangeList!=nil&&len(rangeList.ranges)>0{

		start:=rangeElement[0]
		stop:=rangeElement[1]
		if !(stop>start){
			return
		}
		i:=0
		for i<len(rangeList.ranges){
			item:=rangeList.ranges[i]
			a:=item[0]
			if stop<a{
				return
			}
			if isRangesIntersected(item,rangeElement){
				rs:=rangeMinus(item,rangeElement)
				left:=append(rangeList.ranges[:i],rs...)
				rangeList.ranges=append(left,rangeList.ranges[i+1:]...)
				i+=len(rs)-1
			}
			i+=1
		}

	}
}

//Print 打印输出
func (rangeList *RangeList) Print()  {
	if rangeList!=nil&&len(rangeList.ranges)>0{
		for _,v:=range rangeList.ranges{
			s:=fmt.Sprintf("[%d,%d)",v[0],v[1])
			fmt.Println(s)
		}
	}
}

//合并区间 两个切片  min(start1,start2) 和  max(end1,end2)
func mergeRange(aArry ,bArry []int)[]int{
	return []int{min(aArry[0],bArry[0]),max(aArry[1],bArry[1])}

}

//取有效的范围 [astart,aend]  [bstart,bend]
func rangeMinus(aArry ,bArry []int)[][]int{
	res:=make([][]int,0)
	if bArry[0]> aArry[0]{
		res=append(res, []int{aArry[0],bArry[0]})
	}
	if aArry[1]>bArry[1]{
		res=append(res, []int{bArry[1],aArry[1]})
	}
	return res
}
//判断是否是相交区间  [1,8] [11,15]  [17,21]
//[3,19]
func isRangesIntersected(a,b  []int) bool  {
	aStart,aStop:=a[0],a[1]
	bStart,bStop:=b[0],b[1]
	return !(aStop<=bStart||bStop<=aStart)
}
//min
func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}
//max
func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}
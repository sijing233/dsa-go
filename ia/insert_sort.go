package ia

import "fmt"

//插入排序
func InsertSort(array []int64)[]int64{
	for j:= 1;j<len(array);j++{ //从第2个元素开始，a[1](数组下标从0开始)
		key := array[j]
		i := j - 1
		for ;i >= 0;{ //go的思想之一：一种目的只有一种方法实现，循环只有for
			if array[i] > key{
				array[i+1] = array[i]  //如果没有找到正确的位置，则移动
				i = i - 1
			}else {
				break
			}
		}
		array[i+1] = key
		fmt.Println(array) //显示每行排序后的结果
	}
	return array
}

func InsertSortDesc(array []int64)[]int64{
	for j:= 1;j<len(array);j++{ //从第2个元素开始，a[1](数组下标从0开始)
		key := array[j]
		i := j - 1
		for ;i >= 0;{ //go的思想之一：一种目的只有一种方法实现，循环只有for
			if array[i] < key{
				array[i+1] = array[i]  //如果没有找到正确的位置，则移动
				i = i - 1
			}else {
				break
			}
		}
		array[i+1] = key
		fmt.Println(array) //显示每行排序后的结果
	}
	return array
}

func InsertSortTest(){
	array1 := []int64{3,1,9,2,8,5,4}
	//fmt.Println("升序")
	//fmt.Println(InsertSort(array1))
	//fmt.Println("降序")
	//fmt.Println(InsertSortDesc(array1))
	fmt.Println(FindV(array1,8))
	fmt.Println(FindV(array1,7))
}

func FindV(array []int64,value int64)int64{
	key := -1
	for i:=0;i<len(array);i++{
		if array[i] == value{
			key = i
		}
	}
	return int64(key)
}
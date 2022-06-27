package main

import (
	"RangeListDemo/models"
	"fmt"
)

func main() {
	rangeList:=&models.RangeList{}
	rangeList.Add([]int{1,5})
	rangeList.Print()
	fmt.Println("------add----------")
	rangeList.Add([]int{10,20})
	rangeList.Print()
	fmt.Println("------add----------")
	rangeList.Add([]int{20,20})
	rangeList.Print()
	fmt.Println("------add----------")
	rangeList.Add([]int{20,21})
	rangeList.Print()
	fmt.Println("------add----------")
	rangeList.Add([]int{2,4})
	rangeList.Print()
	fmt.Println("------add----------")
	rangeList.Add([]int{3,8})
	rangeList.Print()
	fmt.Println("------add----------")

    rangeList.Remove([]int{10,10})
	rangeList.Print()
	fmt.Println("--------remove--------")
	rangeList.Remove([]int{10,11})
	rangeList.Print()
	fmt.Println("--------remove--------")
	rangeList.Remove([]int{15,17})
	rangeList.Print()
	fmt.Println("--------remove--------")
	rangeList.Remove([]int{3,19})
	rangeList.Print()
	fmt.Println("--------remove--------")

}

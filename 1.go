package main

import (
	"fmt"
)

func Analyse(src []string, dest []string) (same []string, dif []string) {
	srcMap := make(map[string]bool)
	destMap := make(map[string]bool)
	//存
	for _, s := range src {
		srcMap[s] = true
	}
	for _, d := range dest {
		destMap[d] = true
	}
	//找
	for k, _ := range srcMap {
		if destMap[k] {
			same = append(same, k)
		} else {
			dif = append(dif, k)
		}
	}
	//相同的不加，重复
	for k, _ := range destMap {
		if !srcMap[k] {
			dif = append(dif, k)
		}
	}
	return
}

func main() {
	src := []string{"L", "B", "D", "J"}
	dest := []string{"B", "J", "L", "F"}
	same, dif := Analyse(src, dest)
	fmt.Println("相同的标:", same)
	fmt.Println("不同的标:", dif)
}

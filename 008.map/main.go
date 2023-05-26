package main

import (
	"fmt"
)

// map 合并操作
func main() {
	historyMap, yuwenMap := map[string]int{}, map[string]int{}
	historyMap["历史"] = 80
	yuwenMap["语文"] = 111
	for k, v := range yuwenMap {
		historyMap[k] = v
		fmt.Println(historyMap)

	}
}

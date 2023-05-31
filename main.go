package main

import (
	"carLearn/core"
	"fmt"
)

func main() {
	core.InitCore()
	//core.SVMIns.Train()
	total := 0
	success := 0
	noSuccess := 0
	core.SVMIns.TestData(nil, func(isOk bool) {
		if isOk {
			success++
		} else {
			noSuccess++
		}
		total++
	})
	fmt.Printf("total %d success %d nosuccess %d rate %d%%", total, success, noSuccess, success*100/total)
}

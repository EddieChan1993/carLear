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
	fail := 0
	core.SVMIns.TestData(nil, func(isOk bool) {
		if isOk {
			success++
		} else {
			fail++
		}
		total++
	})
	fmt.Printf("total %d success %d fail %d rate %d%%", total, success, fail, success*100/total)
}

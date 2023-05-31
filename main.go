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
	core.SVMIns.TestDataByFolder(map[core.ImgPath]core.Label{core.TrainIsPath: core.LabelYes}, func(imgPath string, isOk bool) {
		if isOk {
			success++
		} else {
			fmt.Println(imgPath)
			fail++
		}
		total++
	})
	fmt.Printf("total %d success %d fail %d rate %d%%\n", total, success, fail, success*100/total)
}

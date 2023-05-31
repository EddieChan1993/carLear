package main

import (
	"carLearn/core"
	"fmt"
)

func main() {
	svm := core.NewSVM(core.LibSVM)
	svm.Train()
	total := 0
	success := 0
	fail := 0
	svm.TestDataByFolder(map[core.ImgPath]core.Label{core.TestIsPath: core.LabelYes}, func(path core.ImgPath, check, label core.Label) {
		if check == label {
			success++
		} else {
			fail++
		}
		total++
	})
	fmt.Printf("total %d success %d fail %d rate %d%%\n", total, success, fail, success*100/total)
}

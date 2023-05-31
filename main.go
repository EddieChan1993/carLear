package main

import (
	"carLearn/core"
	"fmt"
)

func main() {
	svm := core.NewSVM(core.LibSVM)
	//svm.Train()
	//return
	total := 0
	success := 0
	fail := 0
	svm.TestDataByFolder(map[core.ImgPath]core.Label{core.TestNoPath: core.LabelNo}, func(path core.ImgPath, check, label core.Label) {
		fmt.Println(path, "check", check)
		if check == label {
			success++
		} else {
			fail++
		}
		total++
	})
	fmt.Printf("total %d success %d fail %d rate %d%%\n", total, success, fail, success*100/total)
}

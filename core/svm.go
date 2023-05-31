package core

import (
	"carLearn/util"
	"fmt"
	"github.com/danieldk/golinear"
	"gocv.io/x/gocv"
	"log"
)

type Label = float64  //标签
type ImgPath = string //图片路径，或文件夹路径
type TestPin = map[ImgPath]Label

const (
	LabelYes Label = 1 //真
	LabelNo  Label = 0 //假
)

var SVMIns *Svm

type Svm struct {
	problem *golinear.Problem
	params  golinear.Parameters
}

func initSvm() {
	ins := &Svm{
		problem: golinear.NewProblem(),
		params:  golinear.DefaultParameters(),
	}
	ins.problem.SetBias(1)
	ins.params.SolverType = golinear.NewL2RLogisticRegressionDefault()
	SVMIns = ins
}

func (s *Svm) Train() {
	s.addTrainData()
	model, err := golinear.TrainModel(s.params, s.problem)
	if err != nil {
		log.Fatal(fmt.Errorf("TrainModel err %v", err))
	}
	if model == nil {
		log.Fatal("TrainModel model nil")
	}
	err = model.Save(modelPath)
	if err != nil {
		log.Fatal(fmt.Errorf("save err %v", err))
	}
}

// TestDataByFolder 测试某个文件夹下的数据
func (s *Svm) TestDataByFolder(pin TestPin, fn func(path ImgPath, check, label Label)) {
	res := make(TestPin, 3000)
	for folder, label := range pin {
		allData := util.DirFiles(folder)
		for _, path := range allData {
			res[path] = label
		}
	}
	s.TestDataByImgPath(res, func(path ImgPath, check, label Label) {
		fn(path, check, label)
	})
}

// TestDataByImgPath 测试某个指定文件地址数据
func (s *Svm) TestDataByImgPath(pin TestPin, fn func(path ImgPath, check, label Label)) {
	modelNow, err := golinear.LoadModel(modelPath)
	if err != nil {
		log.Fatal(fmt.Errorf("LoadModel err %v", err))
	}
	//等于预期否
	for filePath, label := range pin {
		obj := s.toVector(filePath, label)
		check := modelNow.Predict(obj.Features)
		fn(filePath, check, obj.Label)
	}
}

// addTrainData 加载数据集到
func (s *Svm) addTrainData() {
	isImgPaths := util.DirFiles(TrainIsPath)
	noImgPaths := util.DirFiles(TrainNoPath)
	for _, filePath := range isImgPaths {
		s.problem.Add(*s.toVector(filePath, LabelYes))
		fmt.Printf("imgPath %s label %f ToFloat Ok\n", filePath, LabelYes)
	}
	for _, filePath := range noImgPaths {
		s.problem.Add(*s.toVector(filePath, LabelNo))
		fmt.Printf("imgPath %s label %f ToFloat Ok\n", filePath, LabelNo)
	}
	fmt.Println("train Data total", len(isImgPaths)+len(noImgPaths))
}

// toVector 图片信息转为向量
func (s *Svm) toVector(filePath string, label Label) *golinear.TrainingInstance {
	img := gocv.IMRead(filePath, gocv.IMReadColor)
	mat := gocv.NewMat()
	defer mat.Close()
	img.ConvertTo(&mat, gocv.MatTypeCV64F)
	f64, err := mat.DataPtrFloat64()
	if err != nil {
		log.Fatal(err)
	}
	return &golinear.TrainingInstance{
		Label:    label,
		Features: golinear.FromDenseVector(f64),
	}
}

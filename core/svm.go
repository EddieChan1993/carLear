package core

import (
	"carLearn/util"
	"fmt"
	"github.com/danieldk/golinear"
	"gocv.io/x/gocv"
	"log"
)

type Label = float64 //标签
type Path = string   //数据路径
type TestPin = map[Path]Label

const (
	LabelYes Label = 1
	LabelNo  Label = 0
)

var SVMIns *Svm

type Svm struct {
	problem *golinear.Problem
	params  golinear.Parameters
}

func InitSvm() {
	ins := &Svm{
		problem: golinear.NewProblem(),
		params:  golinear.DefaultParameters(),
	}
	ins.problem.SetBias(1)
	ins.params.SolverType = golinear.NewL2RLogisticRegressionDefault()
	SVMIns = ins
}

func (s *Svm) Train() {
	s.addTrainData(func(instance *golinear.TrainingInstance) {
		s.problem.Add(*instance)
	})
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

// TestData 测试数据验证
func (s *Svm) TestData(pin TestPin, fn func(isOk bool)) {
	if pin == nil {
		pin = map[Path]Label{testIsPath: LabelYes, testNoPath: LabelNo}
	}
	modelNow, err := golinear.LoadModel(modelPath)
	if err != nil {
		log.Fatal(fmt.Errorf("LoadModel err %v", err))
	}
	res := make([]*golinear.TrainingInstance, 0, 3000)
	for path, label := range pin {
		testFilePath := util.DirFiles(path)
		for _, filePath := range testFilePath {
			res = append(res, s.toVector(filePath, label))
		}
	}
	for _, obj := range res {
		check := modelNow.Predict(obj.Features)
		//等于预期否
		fn(check == obj.Label)
	}
}

// addTrainData 加载数据集
func (s *Svm) addTrainData(fn func(*golinear.TrainingInstance)) {
	isImgPaths := util.DirFiles(trainIsPath)
	noImgPaths := util.DirFiles(trainNoPath)
	for _, filePath := range isImgPaths {
		fn(s.toVector(filePath, LabelYes))
		fmt.Printf("imgPath %s label %f ToFloat Ok\n", filePath, LabelYes)
	}
	for _, filePath := range noImgPaths {
		fn(s.toVector(filePath, LabelNo))
		fmt.Printf("imgPath %s label %f ToFloat Ok\n", filePath, LabelNo)
	}
	fmt.Println("train Data total", len(isImgPaths)+len(noImgPaths))
}

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

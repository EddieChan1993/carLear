package core

import (
	"carLearn/util"
	"fmt"
	libSvm "github.com/ewalker544/libsvm-go"
	"gocv.io/x/gocv"
	"log"
)

/**
libsvm 该第三方库目前存在性能问题，暂不推荐使用
*/

type Tlibsvm struct {
	model   *libSvm.Model
	problem *libSvm.Problem
	params  *libSvm.Parameter
}

func initLibSvm() *Tlibsvm {
	ins := &Tlibsvm{}
	ins.params = libSvm.NewParameter()
	ins.params.KernelType = libSvm.LINEAR
	ins.params.Gamma = 1
	ins.params.C = 1
	ins.model = libSvm.NewModel(ins.params)
	return ins
}

func (s *Tlibsvm) Train() {
	s.addTrainData()
	var err error
	fmt.Println("loading train data")
	s.problem, err = libSvm.NewProblem(trainDataPath, s.params)
	if err != nil {
		log.Fatal(fmt.Errorf("NewProblem err %v", err))
	}
	fmt.Println("train start")
	err = s.model.Train(s.problem)
	if err != nil {
		log.Fatal(fmt.Errorf("TrainModel err %v", err))
	}
	err = s.model.Dump(libsvmPath)
	if err != nil {
		log.Fatal(fmt.Errorf("save err %v", err))
	}
	fmt.Println("train complete")
}

// TestDataByFolder 测试某个文件夹下的数据
func (s *Tlibsvm) TestDataByFolder(pin TestPin, fn func(path ImgPath, check, label Label)) {
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
func (s *Tlibsvm) TestDataByImgPath(pin TestPin, fn func(path ImgPath, check, label Label)) {
	model := libSvm.NewModelFromFile(libsvmPath)
	if model == nil {
		log.Fatal("NewModelFromFile nil")
	}
	//等于预期否
	for filePath, label := range pin {
		obj := s.toVector(filePath, label)
		check := model.Predict(obj)
		fn(filePath, check, label)
	}
}

// addTrainData 加载数据集到
func (s *Tlibsvm) addTrainData() {
	if util.IsExtraFile(trainDataPath) {
		fmt.Println("train file ok")
		return
	}
	fmt.Println("train file creating")
	trainFile, fileBuffer, err := util.NewBufferFile(trainDataPath)
	if err != nil {
		log.Fatal("NewBufferFile", err)
	}
	defer func() {
		trainFile.Close()
		fileBuffer.Flush()
	}()

	isImgPaths := util.DirFiles(TrainIsPath)
	noImgPaths := util.DirFiles(TrainNoPath)
	for _, filePath := range isImgPaths {
		vals := s.toVectorSlice(filePath, LabelYes)
		fileBuffer.WriteString(fmt.Sprintf("%f", LabelYes))
		for index, val := range vals {
			fileBuffer.WriteString(fmt.Sprintf(" %d:%f", index+1, val))
		}
		fileBuffer.WriteString("\n")
	}
	for _, filePath := range noImgPaths {
		vals := s.toVectorSlice(filePath, LabelNo)
		fileBuffer.WriteString(fmt.Sprintf("%f", LabelNo))
		for index, val := range vals {
			fileBuffer.WriteString(fmt.Sprintf(" %d:%f", index+1, val))
		}
		fileBuffer.WriteString("\n")
	}
	fmt.Println("train file complete")
	fmt.Println("train Data total", len(isImgPaths)+len(noImgPaths))
}

// toVector 图片信息转为向量
func (s *Tlibsvm) toVector(filePath string, label Label) map[int]float64 {
	img := gocv.IMRead(filePath, gocv.IMReadColor)
	mat := gocv.NewMat()
	defer mat.Close()
	img.ConvertTo(&mat, gocv.MatTypeCV64F)
	f64, err := mat.DataPtrFloat64()
	if err != nil {
		log.Fatal(err)
	}
	tmp := make(map[int]float64, len(f64))
	for i, f := range f64 {
		tmp[i+1] = f
	}
	return tmp
}

func (s *Tlibsvm) toVectorSlice(filePath string, label Label) []float64 {
	img := gocv.IMRead(filePath, gocv.IMReadColor)
	mat := gocv.NewMat()
	defer mat.Close()
	img.ConvertTo(&mat, gocv.MatTypeCV64F)
	f64, err := mat.DataPtrFloat64()
	if err != nil {
		log.Fatal(err)
	}
	return f64
}

package core

const (
	TrainIsPath = "./data/has/train"
	TrainNoPath = "./data/no/train"
	TestIsPath  = "./data/has/test"
	TestNoPath  = "./data/no/test"

	linearsvmPath = "./linearsvm.model"
	libsvmPath    = "./libsvm.model"
	trainDataPath = "./a9a.train"
)

type SVMTool = string

const (
	LinerSVM SVMTool = "LinerSVM"
	LibSVM   SVMTool = "LibSVM"
)

type Label = float64  //标签
type ImgPath = string //图片路径，或文件夹路径
type TestPin = map[ImgPath]Label

const (
	LabelYes Label = 1 //真
	LabelNo  Label = 0 //假
)

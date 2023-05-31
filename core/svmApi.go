package core

type ISVM interface {
	Train()
	TestDataByFolder(pin TestPin, fn func(path ImgPath, check, label Label))
	TestDataByImgPath(pin TestPin, fn func(path ImgPath, check, label Label))
}

func NewSVM(tool SVMTool) ISVM {
	switch tool {
	case LibSVM:
		return initLibSvm()
	case LinerSVM:
		return initLinearSvm()
	}
	return nil
}

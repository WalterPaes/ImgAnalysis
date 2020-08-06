package analyzer

import "ImgAnalysis/pkg/services"

type Analyzer interface {
	DoAnalysis(img []byte) (*services.Result, error)
}

package analyzer

type Analyzer interface {
	DoAnalysis(img []byte) ([]byte, error)
}

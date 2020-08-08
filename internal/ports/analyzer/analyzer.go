package analyzer

type Analyzer interface {
	DoAnalysis(img []byte) (*Result, error)
}

type Result struct {
	Labels []Labels `json:"Labels"`
}

type Labels struct {
	Confidence float64 `json:"Confidence"`
	Name       string  `json:"Name"`
}

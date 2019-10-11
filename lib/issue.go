package lib

type estimate struct {
	value int64
}

type plusOne struct {
	createdAt string
}

// Issue is issue
type Issue struct {
	estimate  estimate
	plusOnes  []plusOne
	pipeline  Pipeline
	pipelines []Pipeline
	isEpic    bool
}

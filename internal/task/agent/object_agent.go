package agent

func NewObjectAgent(url string) *ObjectAgent {
	return &ObjectAgent{url: url}
}

type ObjectAgent struct {
	url string
}

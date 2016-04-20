package node

type Noder interface {
	GetData() string
}

type Node struct {
	Data     string
	NextNode Noder
}

func (self Node) GetData() string {
	return ""
}

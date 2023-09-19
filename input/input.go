package input

type Input struct {
	directory string
	calldata  []byte
}

const (
	OperationSearch = iota
	OperationDecode
)

func (i Input) Operation() uint {
	if len(i.calldata) > 4 {
		return OperationDecode
	}
	return OperationSearch
}

func (i Input) Directory() string {
	return i.directory
}

func (i Input) Selector() [4]byte {
	var selector [4]byte
	copy(selector[:], i.calldata)
	return selector
}

func (i Input) SelectorWithArgs() []byte {
	return i.calldata
}

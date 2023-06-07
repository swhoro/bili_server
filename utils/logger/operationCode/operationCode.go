package operationCode

type Operation uint8

const (
	ModifyUserAuthorize Operation = iota
)

var OperationToStr = [...]string{"ModifyUserAuthorize"}

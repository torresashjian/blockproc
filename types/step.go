package types

type StepType int

const(
	INSTALL_PROC StepType = iota
)

type Proc struct {

}

type Step struct {
	StepType int
	Data   []byte
}

type StartProc struct {

}

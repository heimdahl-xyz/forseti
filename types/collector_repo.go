package types

type Processor interface {
	ProcessTransfer(ft *FTMessage) error
}

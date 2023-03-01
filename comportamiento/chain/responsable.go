package chain

type Responsable interface {
	Handle(process string)
}

package abstractions

type RouterAbstraction interface {
	Routes()
	Serve() error
}

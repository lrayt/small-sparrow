package abstract

type Handler interface {
	Run() error
	Shutdown() error
}

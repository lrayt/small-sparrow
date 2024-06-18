package common

type HandleProvider interface {
	Run() error
	Shutdown() error
}

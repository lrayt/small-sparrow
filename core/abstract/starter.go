package abstract

type Starter interface {
	Init() error
	Close() error // 关闭
}

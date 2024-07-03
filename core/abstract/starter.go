package abstract

type Starter interface {
	Init() error  // 初始化
	Run() error   // 运行
	Close() error // 关闭
}

package golimit

type GLimit struct {
	Num int
	C   chan struct{}
}

func NewG(num int) *GLimit {
	return &GLimit{
		Num: num,
		C:   make(chan struct{}, num),
	}
}

func (g *GLimit) Run(f func()) {
	g.C <- struct{}{} // chan占位
	go func() {
		f()
		<-g.C // 从chan取回消息
	}()
}

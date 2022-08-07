package main

import "fmt"

// 官网effective go
// https://go.dev/doc/effective_go#pointers_vs_values
// https://zhuanlan.zhihu.com/p/101363361

/* 官方FAQ三条建议规则
   1. 第一条也是最重要的一条，方法是否要修改 receiver？
   2. 其次是效率的考虑，如果 receiver 非常大，比如说一个大 struct，使用指针将非常合适。
   3. 接下来是一致性，如果该类型的某些方法必须使用指针 receiver，剩下的也要使用指针。不论使用什么类型的 receiver，方法集要一致.
*/

/* 其他规则
   1. 实例和实例指针可以调用值类型和指针类型 receiver 的方法。
   2. 如果通过 method express 方式，struct 值只能调用值类型 receiver 的方法，而 struct 指针是能调用值类型和指针类型 receiver 的方法的。
   3. 如果 receiver 是 map、func 或 chan，不要使用指针。
   4. 如果 receiver 是 slice，并且方法不会重新分配 slice，不要使用指针。
   5. 如果 receiver 是包含 sync.Mutex 或其它类似的同步字段的结构体，receiver 必须是指针，以避免复制。
   6. 如果 receiver 是大 struct 或 array，receiver 用指针效率会更高。那么，多大是大？假设要把它的所有元素作为参数传递给方法，如果这样会感觉太大，那对 receiver 来说也就太大了。
   7. 如果 receiver 是 struct、array 或 slice，并且它的任何元素都是可能发生改变的内容的指针，最好使用指针类型的 receiver，这会使代码可读性更高。
   8. 如果 receiver 是一个本来就是值类型的小 array 或 struct，没有可变字段，没有指针，或只是一个简单的基础类型，如 int 或 string，使用值类型的 receiver 更合适。
   9. 值类型的 receiver 可以减少可以生成的垃圾量，如果将值传递给值方法，可以使用栈上的副本而不是在堆上进行分配。编译器会尝试避免这种分配，但不会总成功。不要为此原因却不事先分析而选择值类型的 receiver。
   10. 最后，如有疑问，请使用指针类型的 receiver。
*/

type Ball struct {
	Name string
}

func (b *Ball) Ping() {
	fmt.Println("ping")
}

func (b Ball) Pong() {
	fmt.Println("pong")
}

func NewBall() Ball { // 返回一个右值
	return Ball{Name: "right value struct"}
}

// 值方法（value methods）可以通过指针和值调用，但是指针方法（pointer methods）只能通过指针来调用。
// 但有一个例外，如果某个值是可寻址的（addressable，或者说左值），那么编译器会在值调用指针方法时自动插入取地址符，使得在此情形下看起来像指针方法也可以通过值来调用。
func v1() {
	// struct 的实例和实例指针都可以调用值类型和指针类型的receiver的方法
	v := Ball{}
	p := &Ball{}

	v.Ping() // 可寻址, 所以值类型可以调用指针方法
	v.Pong()

	p.Ping()
	p.Pong()

	// NewBall().Ping() // 编译不通过, 不可寻址
	NewBall().Pong()
}

type BallV2 struct {
	Name string
}

func (b *BallV2) Ping() {
	fmt.Println("ping")
}

func (b BallV2) Pong() {
	fmt.Println("pong")
}

func v2() {
	// 通过 method expression 方式, struct值只能调用值类型receiver的方法
	v := BallV2{}
	// BallV2.Ping(&v)  // 失败
	BallV2.Pong(v)

	// struct指针可以调用值类型和指针类型的receiver方法
	vv := &BallV2{}
	(*BallV2).Ping(vv)
	(*BallV2).Pong(vv)
}

func main() {
	v1()
	v2()
}

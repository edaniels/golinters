package a

type Foo struct{}

func (f Foo) MustDo()           {}
func (f Foo) DoSomething()      {}
func (f Foo) MustCreate(name string) error { return nil }

func bad() {
	var f Foo
	f.MustDo() // want `MustVerb\(\) usage found`
}

func good() {
	var f Foo
	f.DoSomething()   // not a Must method
	f.MustCreate("x") // has call args, not zero-arg
}

func init() {
	var f Foo
	f.MustDo() // ok - init is allowed
}

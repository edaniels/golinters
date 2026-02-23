package a

type Logger struct{}

func (l Logger) Fatal(msg string)                {}
func (l Logger) Fatalw(msg string, args ...any)  {}
func (l Logger) Info(msg string)                 {}

type Service struct{}

func (s Service) Fatal(msg string) {}

func bad() {
	var l Logger
	l.Fatal("something went wrong")  // want `Log fatal usage detected`
	l.Fatalw("something went wrong") // want `Log fatal usage detected`
}

func good() {
	var l Logger
	l.Info("this is fine")

	var s Service
	s.Fatal("not a Logger type") // ok - type doesn't end in Logger
}

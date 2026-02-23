package a

type Logger struct{}

func (l Logger) Debugf(format string, args ...any) {}
func (l Logger) Infof(format string, args ...any)  {}
func (l Logger) Warnf(format string, args ...any)  {}

func bad() {
	var l Logger
	l.Debugf("hello") // want `useless print with f usage found`
	l.Infof("hello")  // want `useless print with f usage found`
	l.Warnf("hello")  // want `useless print with f usage found`
}

func good() {
	var l Logger
	l.Debugf("hello %s", "world")
	l.Infof("count: %d", 42)
}

package a

func cleanup() {}

func bad() {
	for i := 0; i < 10; i++ {
		defer cleanup() // want `defer within an immediate outer for usage found`
	}
}

func good() {
	defer cleanup()
}

func goodFuncLitInLoop() {
	for i := 0; i < 10; i++ {
		func() {
			defer cleanup() // ok - defer in func lit, not directly in for
		}()
	}
}

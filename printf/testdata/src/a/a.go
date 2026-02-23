package a

import (
	"fmt"
	"os"
)

func bad() {
	fmt.Println("hello")                // want `fmt\.\(Fp\|P\)rintf\* usage found`
	fmt.Printf("hello %s\n", "world")   // want `fmt\.\(Fp\|P\)rintf\* usage found`
	fmt.Print("hello")                   // want `fmt\.\(Fp\|P\)rintf\* usage found`
	fmt.Fprintf(os.Stdout, "hello")      // want `fmt\.\(Fp\|P\)rintf\* usage found`
	fmt.Fprintln(os.Stdout, "hello")     // want `fmt\.\(Fp\|P\)rintf\* usage found`
}

func good() {
	fmt.Fprintf(os.Stderr, "hello") // writing to stderr is fine
}

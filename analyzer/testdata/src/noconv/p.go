package noconv

import (
	"errors"
	"fmt"
	"os"
)

var errSentinel = errors.New("connection refused")

func positive() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	fmt.Sprintf("%d", i)         // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", i)         // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(i)                // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", 42)        // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%v", 42)        // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(42)               // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", i8)
	fmt.Sprintf("%v", i8)
	fmt.Sprint(i8)
	fmt.Sprintf("%d", int8(42))
	fmt.Sprintf("%v", int8(42))
	fmt.Sprint(int8(42))
	fmt.Sprintf("%d", i16)       
	fmt.Sprintf("%v", i16)       
	fmt.Sprint(i16)              
	fmt.Sprintf("%d", int16(42)) 
	fmt.Sprintf("%v", int16(42)) 
	fmt.Sprint(int16(42))        
	fmt.Sprintf("%d", i32)       
	fmt.Sprintf("%v", i32)       
	fmt.Sprint(i32)              
	fmt.Sprintf("%d", int32(42)) 
	fmt.Sprintf("%v", int32(42)) 
	fmt.Sprint(int32(42))        
	fmt.Sprintf("%d", i64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%v", i64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprint(i64)              // want "fmt.Sprint can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%d", int64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%v", int64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprint(int64(42))        // want "fmt.Sprint can be replaced with faster strconv.FormatInt"

	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	fmt.Sprintf("%d", ui)
	fmt.Sprintf("%v", ui)
	fmt.Sprint(ui)
	fmt.Sprintf("%d", uint(42))
	fmt.Sprintf("%v", uint(42))
	fmt.Sprint(uint(42))          
	fmt.Sprintf("%d", ui8)        
	fmt.Sprintf("%v", ui8)        
	fmt.Sprint(ui8)               
	fmt.Sprintf("%d", uint8(42))  
	fmt.Sprintf("%v", uint8(42))  
	fmt.Sprint(uint8(42))         
	fmt.Sprintf("%d", ui16)       
	fmt.Sprintf("%v", ui16)       
	fmt.Sprint(ui16)              
	fmt.Sprintf("%d", uint16(42)) 
	fmt.Sprintf("%v", uint16(42)) 
	fmt.Sprint(uint16(42))        
	fmt.Sprintf("%d", ui32)       
	fmt.Sprintf("%v", ui32)       
	fmt.Sprint(ui32)              
	fmt.Sprintf("%d", uint32(42)) 
	fmt.Sprintf("%v", uint32(42)) 
	fmt.Sprint(uint32(42))        
	fmt.Sprintf("%d", ui64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", ui64)       // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui64)              // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%d", uint64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%v", uint64(42)) // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(uint64(42))        // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
}

func negative() {
	const val = "val%d"

	_ = int32(42)

	fmt.Scan(42)
	fmt.Scanf("%d", 42)
	fmt.Println("%d", 42)
	fmt.Printf("%d")
	fmt.Printf("%v")
	fmt.Printf("%d", 42)
	fmt.Printf("%s %d", "hello", 42)

	fmt.Fprint(os.Stdout, "%d", 42)
	fmt.Fprintf(os.Stdout, "test")
	fmt.Fprintf(os.Stdout, "%d")
	fmt.Fprintf(os.Stdout, "%v")
	fmt.Fprintf(os.Stdout, "%d", 42)
	fmt.Fprintf(os.Stdout, "%s %d", "hello", 42)

	fmt.Sprint("test", 42)
	fmt.Sprint(42, 42)
	fmt.Sprintf("test")
	fmt.Sprintf("%v")
	fmt.Sprintf("%d")
	fmt.Sprintf("%d", 42, 42)
	fmt.Sprintf("%#d", 42)
	fmt.Sprintf("value %d", 42)
	fmt.Sprintf(val, 42)
	fmt.Sprintf("%s %v", "hello", "world")
	fmt.Sprintf("%#v", 42)
	fmt.Sprintf("%T", struct{ string }{})
	fmt.Sprintf("%%v", 42)
	fmt.Sprintf("%3d", 42)
	fmt.Sprintf("% d", 42)
	fmt.Sprintf("%-10d", 42)
	fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)

	// Integer.
	fmt.Sprintf("%#x", uint64(42))
	fmt.Sprintf("%#v", uint64(42))
	fmt.Sprintf("%#b", 42)
	fmt.Sprintf("%#o", 42)
	fmt.Sprintf("%#x", 42)
	fmt.Sprintf("%#X", 42)

	fmt.Sprintf("%b", 42)
	fmt.Sprintf("%c", 42)
	fmt.Sprintf("%o", 42)
	fmt.Sprintf("%O", 42)
	fmt.Sprintf("%q", 42)
	fmt.Sprintf("%x", 42)
	fmt.Sprintf("%X", 42)

	// Floating point.
	fmt.Sprintf("%9f", 42.42)
	fmt.Sprintf("%.2f", 42.42)
	fmt.Sprintf("%.2f", 42.42)
	fmt.Sprintf("%9.2f", 42.42)
	fmt.Sprintf("%9.f", 42.42)
	fmt.Sprintf("%.3g", 42.42)

	fmt.Sprintf("%b", float32(42.42))
	fmt.Sprintf("%e", float32(42.42))
	fmt.Sprintf("%E", float32(42.42))
	fmt.Sprintf("%f", float32(42.42))
	fmt.Sprintf("%F", float32(42.42))
	fmt.Sprintf("%g", float32(42.42))
	fmt.Sprintf("%G", float32(42.42))
	fmt.Sprintf("%x", float32(42.42))
	fmt.Sprintf("%X", float32(42.42))
	fmt.Sprintf("%v", float32(42.42))

	fmt.Sprintf("%b", 42.42)
	fmt.Sprintf("%e", 42.42)
	fmt.Sprintf("%E", 42.42)
	fmt.Sprintf("%f", 42.42)
	fmt.Sprintf("%F", 42.42)
	fmt.Sprintf("%g", 42.42)
	fmt.Sprintf("%G", 42.42)
	fmt.Sprintf("%x", 42.42)
	fmt.Sprintf("%X", 42.42)
	fmt.Sprintf("%v", 42.42)

	fmt.Sprintf("%b", 42i+42)
	fmt.Sprintf("%e", 42i+42)
	fmt.Sprintf("%E", 42i+42)
	fmt.Sprintf("%f", 42i+42)
	fmt.Sprintf("%F", 42i+42)
	fmt.Sprintf("%g", 42i+42)
	fmt.Sprintf("%G", 42i+42)
	fmt.Sprintf("%x", 42i+42)
	fmt.Sprintf("%X", 42i+42)
	fmt.Sprintf("%v", 42i+42)

	// String & slice of bytes.
	fmt.Sprintf("%q", "hello")
	fmt.Sprintf("%#q", `"hello"`)
	fmt.Sprintf("%+q", "hello")
	fmt.Sprintf("%X", "hello")

	// Slice.
	fmt.Sprintf("%x", []uint16{'d'})
	fmt.Sprintf("%x", []uint32{'d'})
	fmt.Sprintf("%x", []uint64{'d'})
	fmt.Sprintf("%x", []uint{'d'})
	fmt.Sprintf("%x", [1]byte{'c'})
	fmt.Sprintf("%x", [1]uint8{'d'})
	fmt.Sprintf("%x", [1]uint16{'d'})
	fmt.Sprintf("%x", [1]uint32{'d'})
	fmt.Sprintf("%x", [1]uint64{'d'})
	fmt.Sprintf("%x", [1]uint{'d'})
	fmt.Sprintf("%x", []int8{1})
	fmt.Sprintf("%x", []int16{1})
	fmt.Sprintf("%x", []int32{1})
	fmt.Sprintf("%x", []int64{1})
	fmt.Sprintf("%x", []int{1})
	fmt.Sprintf("%x", [...]int8{1})
	fmt.Sprintf("%x", [...]int16{1})
	fmt.Sprintf("%x", [...]int32{1})
	fmt.Sprintf("%x", [...]int64{1})
	fmt.Sprintf("%x", [...]int{1})
	fmt.Sprintf("%x", []string{"hello"})
	fmt.Sprintf("%x", []rune{'a'})

	fmt.Sprintf("% x", []byte{1, 2, 3})
	fmt.Sprintf("% X", []byte{1, 2, 3})
	fmt.Sprintf("%p", []byte{1, 2, 3})
	fmt.Sprintf("%#p", []byte{1, 2, 3})

	// Pointer.
	var ptr *int
	fmt.Sprintf("%v", ptr)
	fmt.Sprintf("%b", ptr)
	fmt.Sprintf("%d", ptr)
	fmt.Sprintf("%o", ptr)
	fmt.Sprintf("%x", ptr)
	fmt.Sprintf("%X", ptr)
}

func malformed() {
	fmt.Sprintf("%d", "example")
	fmt.Sprintf("%T", "example")
	fmt.Sprintf("%t", "example")
	fmt.Sprintf("%b", "example")
	fmt.Sprintf("%e", "example")
	fmt.Sprintf("%E", "example")
	fmt.Sprintf("%f", "example")
	fmt.Sprintf("%F", "example")
	fmt.Sprintf("%g", "example")
	fmt.Sprintf("%G", "example")
	fmt.Sprintf("%x", "example")
	fmt.Sprintf("%X", "example")

	fmt.Sprintf("%d", errSentinel)
	fmt.Sprintf("%T", errSentinel)
	fmt.Sprintf("%t", errSentinel)
	fmt.Sprintf("%b", errSentinel)
	fmt.Sprintf("%e", errSentinel)
	fmt.Sprintf("%E", errSentinel)
	fmt.Sprintf("%f", errSentinel)
	fmt.Sprintf("%F", errSentinel)
	fmt.Sprintf("%g", errSentinel)
	fmt.Sprintf("%G", errSentinel)
	fmt.Sprintf("%x", errSentinel)
	fmt.Sprintf("%X", errSentinel)

	fmt.Sprintf("%d", true)
	fmt.Sprintf("%T", true)
	fmt.Sprintf("%b", true)
	fmt.Sprintf("%e", true)
	fmt.Sprintf("%E", true)
	fmt.Sprintf("%f", true)
	fmt.Sprintf("%F", true)
	fmt.Sprintf("%g", true)
	fmt.Sprintf("%G", true)
	fmt.Sprintf("%x", true)
	fmt.Sprintf("%X", true)

	var bb []byte
	fmt.Sprintf("%d", bb)
	fmt.Sprintf("%T", bb)
	fmt.Sprintf("%t", bb)
	fmt.Sprintf("%b", bb)
	fmt.Sprintf("%e", bb)
	fmt.Sprintf("%E", bb)
	fmt.Sprintf("%f", bb)
	fmt.Sprintf("%F", bb)
	fmt.Sprintf("%g", bb)
	fmt.Sprintf("%G", bb)
	fmt.Sprintf("%X", bb)
	fmt.Sprintf("%v", bb)

	fmt.Sprintf("%T", 42)
	fmt.Sprintf("%t", 42)
	fmt.Sprintf("%b", 42)
	fmt.Sprintf("%e", 42)
	fmt.Sprintf("%E", 42)
	fmt.Sprintf("%f", 42)
	fmt.Sprintf("%F", 42)
	fmt.Sprintf("%g", 42)
	fmt.Sprintf("%G", 42)
	fmt.Sprintf("%x", 42)
	fmt.Sprintf("%X", 42)

	fmt.Sprintf("%T", uint(42))
	fmt.Sprintf("%t", uint(42))
	fmt.Sprintf("%b", uint(42))
	fmt.Sprintf("%e", uint(42))
	fmt.Sprintf("%E", uint(42))
	fmt.Sprintf("%f", uint(42))
	fmt.Sprintf("%F", uint(42))
	fmt.Sprintf("%g", uint(42))
	fmt.Sprintf("%G", uint(42))
	fmt.Sprintf("%x", uint(42))
	fmt.Sprintf("%X", uint(42))

	fmt.Sprintf("%d", 42.42)
	fmt.Sprintf("%d", map[string]string{})
	fmt.Sprint(make(chan int))
	fmt.Sprint([]int{1, 2, 3})
}

package main // The package name main implies that the program will be compiled to an executable
//Any other package name means that the source code will be compiled into a library. 

import("fmt")//Alternatively, ."fmt" can be used. In this case, when calling a function like Printf, you do not have to add fmt. extension.

func printThis(x int){
	fmt.Printf("This is %d\n",x)
}

func main(){
	fmt.Printf("Hello World!\n")
	printThis(10)

	//Declaring a variable.

	//x := 10
        var x uint
	fmt.Printf("Hello %d\n", x);//Any uninitialised value prints a zero/evaluates to it.

	var a = []byte{0,1,2,3,4,5,6,7}
	a = append(a, 8)
	a[8] = 9//Replaces value of 8 with 9

	// make in action
	var b = make([]byte, 10)
	for i := byte(0); i < 10; i++ {
		b[i] = i
	}
	b = b[2:]
	b = b[:len(b)-2]
	fmt.Printf("%v\n",b)

	//Note, this operation is not possible without a previously existing couple of bytes.

	fmt.Printf("%v\n",a)

	/*a[5], a[len(a)-1] = a[len(a)-1], a[5]
	a = a[:len(a)-1]
	b := a[0:5]
	b = append(b,a[6:]...)*/

	fmt.Printf("%v\n",b)
//The two lines above and below pretty much do the same job although there are some differences in the backend.
	//var x int 
	//x = 10

	

	/*var x int
	var a float
	var b float32
	var c float64
	var d uint32
	var e uint64
	var e uint8
	var f string
	var g []byte
	var h []uint32

	//g and h are slices. an array is static-sized whereas slices are dynamically sized.
	//to store anything in a slice, you must allocate soace in it using make
	//g = make([]byte, 10)
	//Only then can you g[0] = 1, otherwise the program won't compile.
	*/
	//Returns the default type
	//You can however declare the types manually.
	//Arithmetic operations can only be performed on similar types.

	//a := x+int(y). this is type conversion.
	//All declared variables must be used for the code to compile.

	//Scoping is very important in any programing language

}


package main

//#include<stdio.h>
/*
static int div(int a,int b){
	return a/b;
}

*/

import "C"
import "fmt"

func main() {
	v := C.div(6, 3)

	fmt.Println(v)
	//fmt.Println(a.f)

}

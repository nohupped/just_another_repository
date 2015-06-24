package main

import (
	"fmt"
	"os"
	"strconv"
	)


func main(){
	Args := os.Args[1:]
	var Args_convert uint64
	var err error
//	fmt.Println(Args)
	if len(Args) != 0 {
		Args_convert, err =  strconv.ParseUint(Args[0], 0, 128)
//	fmt.Println(Args_convert)
	fmt.Printf("Will try to convert argument as %T\n",Args_convert)
		if err != nil {
			fmt.Println(err, "Invalid parsing, exiting")
			os.Exit(1)
			}

		
		 } else {
		fmt.Println("No args, exiting")
		os.Exit(1)
		}
//	fmt.Printf("%d\n\n\n",Args_convert)
	
	
	p := fact(&Args_convert)
	fmt.Println("Factorial of", Args_convert, "is", p)
//	fmt.Println(Args_convert)
	}

func fact(num *uint64) uint64{
	
	decrement := *num -1
	if *num == 0 {
		return 1
		} else {
		f := *num * fact(&decrement)
		return f
		}
	}


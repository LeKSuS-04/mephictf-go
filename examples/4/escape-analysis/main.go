package main

var x int = 0

func increment() {
	y := x
	y = y + 1
	x = y
}

/*
y = x // 0	   	  |
 		  		  |   y = x // 0
y = y + 1 // 1 	  |
 		  		  |	  y = y + 1 // 1
 	x = y // 1	  |
 		  		  |   x = y // 1
 		  		  |
 		  		  |
 		  		  |
 		  		  |
 		  		  |
 		  		  |
 		  		  |
 		  		  |
*/

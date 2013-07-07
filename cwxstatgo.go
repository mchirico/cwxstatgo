/*
   Example of tree
   mchirico@gmail.com
   http://play.golang.org/p/yGYwb9ACeG
 
 Installation:

   go get github.com/mchirico/cwxstatgo
  

 Example Program:


package main

import (
        cwx "github.com/mchirico/cwxstatgo"
        "fmt"
)

func main() {
        var t *cwx.Tree
        a := []int{10, 10, 2, 3, 8, 9}

        t = cwx.Insert(t, a)

        cwx.Pr(t)                   // Prints each node
        fmt.Println(cwx.Add(t))     // All values added up
        fmt.Println(cwx.Nodes(t))   // Number of individual slices
        fmt.Println(cwx.Flatten(t)) // Gives back origional list

        cwx.WalkerRun(t)
}

   
   
*/
package cwxstatgo

import "fmt"

const MAX_SLICE = 2 //Max Size of slice

type Tree struct {
	Left  *Tree
	Value []int
	Right *Tree
}

func Insert(t *Tree, v []int) *Tree {
	if t == nil {
		t = &Tree{nil, nil, nil}
	}
	// Add code to Insert Additional slices
	if len(v) > MAX_SLICE {

		t.Left = Insert(t.Left, v[:len(v)/2])
		t.Right = Insert(t.Right, v[len(v)/2:])
		return t
	}
	return &Tree{nil, v, nil}
}

func Pr(t *Tree) {
	if t != nil {
		Pr(t.Left)
		Pr(t.Right)
		if t.Value != nil {
			fmt.Println(t.Value)
		}
	}
}

func Add(t *Tree) (sum int) {
	sum = 0
	if t != nil {
		sum = Add(t.Left) + Add(t.Right)
		if t.Value != nil {
			for i := range t.Value {
				sum = sum + t.Value[i]
			}
		}
	}
	return
}
func Nodes(t *Tree) (sum int) {
	sum = 0
	if t != nil {
		sum = Nodes(t.Left) + Nodes(t.Right)
		if t.Value != nil {
			sum = sum + 1
		}
	}
	return
}
func Flatten(t *Tree) (a []int) {
	a = []int{}
	if t != nil {
		a = append(a, Flatten(t.Left)...)
		a = append(a, Flatten(t.Right)...)

		if t.Value != nil {
			a = append(a, t.Value...)
		}
	}
	return
}

func Walk(t *Tree, ch chan []int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	Walk(t.Right, ch)
	ch <- t.Value

}

func Walker(t *Tree) <-chan []int {

	ch := make(chan []int)
	go func() {

		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func WalkerRun(t *Tree) {
	fmt.Printf("\n-------- Walker Run ---------\n")
	c := Walker(t)

	for {
		v1, ok1 := <-c
		if !ok1 {
			break
		}
		if v1 != nil {
			fmt.Println(v1, ok1)
		}

	}

}

/*

//    Installation:
//       go get github.com/mchirico/cwxstatgo


package main

import (
        cwx "github.com/mchirico/cwxstatgo"
        "fmt"
)

func main() {
        var t *cwx.Tree
        a := []int{10, 10, 2, 3, 8, 9}

        t = cwx.Insert(t, a)

        cwx.Pr(t)                   // Prints each node
        fmt.Println(cwx.Add(t))     // All values added up
        fmt.Println(cwx.Nodes(t))   // Number of individual slices
        fmt.Println(cwx.Flatten(t)) // Gives back origional list

        cwx.WalkerRun(t)
}



*/
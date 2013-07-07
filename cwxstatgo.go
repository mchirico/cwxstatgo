/*
   Example of tree
   mchirico@gmail.com
   http://play.golang.org/p/yGYwb9ACeG
 
 Installation:

   go get github.com/mchirico/cwxstatgo
  

 Example Program:

package main

import (
 
    "github.com/mchirico/cwxstatgo"  
)
   

func main() {
	var t *cwxstatgo.Tree
	a := []int{10, 10, 2, 3, 8, 9}

	t = cwxstatgo.Insert(t, a)

	cwxstatgo.Pr(t)                   // Prints each node
	fmt.Println(cwxstatgo.Add(t))     // All values Added up
	fmt.Println(cwxstatgo.nodes(t))   // Number of individual slices
	fmt.Println(cwxstatgo.Flatten(t)) // Gives back origional list

	cwxstatgo.WalkerRun(t)
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

func insert(t *Tree, v []int) *Tree {
	if t == nil {
		t = &Tree{nil, nil, nil}
	}
	// Add code to insert Additional slices
	if len(v) > MAX_SLICE {

		t.Left = insert(t.Left, v[:len(v)/2])
		t.Right = insert(t.Right, v[len(v)/2:])
		return t
	}
	return &Tree{nil, v, nil}
}

func pr(t *Tree) {
	if t != nil {
		pr(t.Left)
		pr(t.Right)
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
func nodes(t *Tree) (sum int) {
	sum = 0
	if t != nil {
		sum = nodes(t.Left) + nodes(t.Right)
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
func main() {
	var t *Tree
	a := []int{10, 10, 2, 3, 8, 9}

	t = Insert(t, a)

	pr(t)                   // Prints each node
	fmt.Println(Add(t))     // All values Added up
	fmt.Println(Nodes(t))   // Number of individual slices
	fmt.Println(Flatten(t)) // Gives back origional list

	WalkerRun(t)
}
*/
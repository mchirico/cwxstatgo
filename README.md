cwxstatgo
=========


Installation:
------------_

  go get github.com/mchirico/cwxstatgo


Usage:
-----

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


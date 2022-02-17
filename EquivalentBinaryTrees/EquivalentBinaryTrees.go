package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, root bool) {
	if t.Left!=nil {
		Walk(t.Left, ch, false)
	}
	ch <- t.Value
	if t.Right!=nil{
		Walk(t.Right, ch, false)
	}
	if root==true{
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1:=make(chan int)
	ch2:=make(chan int)
	go Walk(t1,ch1,true)
	go Walk(t2,ch2,true)
	for {
		x,xok:=<-ch1
		y,yok:=<-ch2
		if xok && yok {
			fmt.Printf("%v==%v ?\n",x,y)
			if x!=y {
				return false
			}
		} else {
			fmt.Println("channel close")
			break
		}
	}
	return true
}

func main() {
	t1:=tree.New(1)
	t2:=tree.New(2)
	println(Same(t1,t2))
}


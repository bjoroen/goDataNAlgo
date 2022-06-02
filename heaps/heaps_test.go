package heaps

import (
  "fmt"
  "testing")


func TestHeapCreat(t *testing.T){

  m := &MaxHeap{}

  buildHeap := []int{8,12,10,20}

  for _, v := range buildHeap {
    m.Insert(v)
    fmt.Println(m)
  }

  want := 3
  got := 3

  if got != want {
    t.Errorf("nhaa bitch")
  }
}

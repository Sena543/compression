package tests

import (
	"reflect"
	"testing"

	"github.com/Sena543/compression/cmd"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("test insertion", func(t *testing.T) {
		testSlice := []int{77, 34, 24, 5}
		var pq cmd.PriorityQueue
		for _, v := range testSlice {
			node := cmd.Node{Count: v}
			pq.Insert(node)
		}
		want := []cmd.Node{{Count: 5}, {Count: 24}, {Count: 34}, {Count: 77}}
		got := pq.PQ()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v got %v", want, got)
		}

	})
	t.Run("test pop", func(t *testing.T) {
		testSlice := []int{77, 34, 24, 5}
		var pq cmd.PriorityQueue
		for _, v := range testSlice {
			node := cmd.Node{Count: v}
			pq.Insert(node)
		}
		want := 5
		got := pq.Pop()
		if want != got.Count {
			t.Errorf("want %v got %v", want, got)
		}

	})

}

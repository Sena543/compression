package tests

import (
	"testing"

	"github.com/Sena543/compression/cmd"
)

func TestEncoding(t *testing.T) {
	testString := "BCAADDDCCACACAC"
	/* pQueueHelper := func(t testing.TB, testInput string) *cmd.PriorityQueue {
		t.Helper()

		var pq_Test cmd.PriorityQueue
		enc := cmd.NewEncode(testInput)
		for k, v := range enc {
			node := cmd.Node{Count: v, Data: k}
			pq_Test.Insert(node)
		}
		return &pq_Test

	} */
	t.Run("test encoding", func(t *testing.T) {
		map_Res := map[string]int{
			"B": 1,
			"C": 6,
			"A": 5,
			"D": 3,
		}

		for _, v := range testString {
			want := map_Res[string(v)]
			got := cmd.NewEncode(testString)[string(v)]
			if want != got {
				t.Errorf("Want %v but got %v", want, got)
			}
		}

	})

	/* t.Run("test tree result", func(t *testing.T) {

		want := []string{
			"C : 6",
			"B : 1",
			"D : 3",
			"A : 5",
		}

		buf := bytes.Buffer{}
		res := pQueueHelper(t, testString)
		var huffmanTest cmd.HuffNode
		root := huffmanTest.BuildTree(*res)
		huffmanTest.PrintTree(root, &buf)
		got := []string{}
		stack := []cmd.HuffNode{*root}
		cur := root
		for _, v := range stack {
			if condition {

			}
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Want %v but got %v", want, got)
		}
	}) */
}

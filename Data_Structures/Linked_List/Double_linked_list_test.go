package Double_linked_list

import (
	"fmt"
	"testing"
)

// Define a struct for operations
type Operation struct {
	Action string
	Params []int
}

// Define a struct for test cases
type TestCase struct {
	Operations      []Operation
	ExpectedResults []interface{}
}

func runTest(ops []Operation) []interface{} {
	DS := &DoublyLinkedListSeq{}
	var ans []interface{}

	for _, op := range ops {
		switch op.Action {
		case "insert_first": // insert_first
			DS.InsertFirst(op.Params[0])
		case "insert_last": // insert_last
			DS.InsertLast(op.Params[0])
		case "delete_first": // delete_first
			ans = append(ans, DS.DeleteFirst())
		case "delete_last": // delete_last
			ans = append(ans, DS.DeleteLast())
		case "splice/remove": // splice/remove
			if DS.head != nil {
				i := op.Params[0]
				n := op.Params[1]
				L := &DoublyLinkedListSeq{}
				for j := 0; j < n; j++ {
					L.InsertLast(j)
				}
				x1 := DS.head.laterNode(i)
				x2 := x1.next
				DS.Splice(x1, L)
				for k := 0; k < n; k++ {
					L = DS.Remove(x1.next, x2.prev)
					x2 = x1.next
					DS.Splice(x1, L)
				}
			}
		}
	}

	var finalList []int
	for node := DS.head; node != nil; node = node.next {
		finalList = append(finalList, node.item)
	}
	ans = append(ans, finalList)
	return ans
}

func check(t *testing.T, test TestCase) error {
	studentSol := runTest(test.Operations)
	if len(studentSol) != len(test.ExpectedResults) {
		t.Errorf("Expected and actual results length mismatch: got %d, want %d", len(studentSol), len(test.ExpectedResults))
		return fmt.Errorf("Expected and actual results length mismatch: got %d, want %d", len(studentSol), len(test.ExpectedResults))
	}
	for i := range studentSol {
		if fmt.Sprintf("%v", studentSol[i]) != fmt.Sprintf("%v", test.ExpectedResults[i]) {
			t.Errorf("Mismatch at index %d: got %v, want %v", i, studentSol[i], test.ExpectedResults[i])
			return fmt.Errorf("Mismatch at index %d: got %v, want %v", i, studentSol[i], test.ExpectedResults[i])
		}
	}
	// uncomment this to get each double list
	// ds_list := studentSol[len(studentSol)-1]
	// t.Logf(fmt.Sprintf("%v", ds_list))

	return nil
}

func TestCases(t *testing.T) {
	tests := []TestCase{
		{
			Operations: []Operation{
				{"insert_last", []int{3}},
				{"insert_first", []int{2}},
				{"insert_last", []int{8}},
				{"insert_first", []int{2}},
				{"insert_last", []int{9}},
				{"insert_first", []int{7}},
				{"delete_last", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"splice/remove", []int{1, 2}},
				{"splice/remove", []int{1, 2}},
				{"splice/remove", []int{1, 2}},
				{"splice/remove", []int{1, 2}},
				{"splice/remove", []int{1, 2}},
			},
			ExpectedResults: []interface{}{9, 8, 7, []int{2, 2, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 3}},
		},
		{
			Operations: []Operation{
				{"insert_first", []int{11}},
				{"insert_first", []int{7}},
				{"insert_first", []int{11}},
				{"insert_last", []int{10}},
				{"insert_first", []int{18}},
				{"insert_first", []int{9}},
				{"insert_last", []int{5}},
				{"insert_first", []int{25}},
				{"insert_first", []int{11}},
				{"insert_first", []int{12}},
				{"delete_first", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"splice/remove", []int{2, 2}},
				{"splice/remove", []int{3, 2}},
				{"splice/remove", []int{1, 2}},
				{"splice/remove", []int{1, 2}},
				{"splice/remove", []int{1, 2}},
			},
			ExpectedResults: []interface{}{12, 11, 5, 10, 25, []int{9, 18, 0, 1, 0, 1, 0, 1, 11, 0, 0, 1, 1, 7, 11}},
		},
		{
			Operations: []Operation{
				{"insert_first", []int{39}},
				{"insert_first", []int{59}},
				{"insert_last", []int{59}},
				{"insert_first", []int{52}},
				{"insert_first", []int{21}},
				{"insert_last", []int{53}},
				{"insert_first", []int{61}},
				{"insert_first", []int{58}},
				{"insert_last", []int{49}},
				{"insert_last", []int{30}},
				{"insert_last", []int{19}},
				{"insert_first", []int{25}},
				{"insert_first", []int{59}},
				{"insert_last", []int{33}},
				{"insert_first", []int{33}},
				{"insert_last", []int{42}},
				{"delete_last", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_last", nil},
				{"splice/remove", []int{2, 4}},
				{"splice/remove", []int{5, 4}},
				{"splice/remove", []int{6, 4}},
				{"splice/remove", []int{1, 4}},
				{"splice/remove", []int{6, 4}},
			},
			ExpectedResults: []interface{}{42, 33, 33, 19, 59, 25, 30, 49, []int{58, 61, 0, 1, 2, 3, 21, 0, 1, 2, 3, 0, 1, 2, 0, 0, 1, 2, 3, 1, 2, 3, 3, 52, 59, 39, 59, 53}},
		},
		{
			Operations: []Operation{
				{"insert_first", []int{64}},
				{"insert_last", []int{45}},
				{"insert_last", []int{10}},
				{"insert_first", []int{70}},
				{"insert_last", []int{25}},
				{"insert_first", []int{48}},
				{"insert_first", []int{26}},
				{"insert_last", []int{27}},
				{"insert_last", []int{96}},
				{"insert_last", []int{90}},
				{"insert_last", []int{64}},
				{"insert_last", []int{8}},
				{"insert_first", []int{65}},
				{"insert_first", []int{34}},
				{"insert_last", []int{20}},
				{"insert_last", []int{31}},
				{"insert_last", []int{84}},
				{"insert_last", []int{76}},
				{"insert_last", []int{73}},
				{"insert_last", []int{39}},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"splice/remove", []int{4, 5}},
				{"splice/remove", []int{2, 5}},
				{"splice/remove", []int{7, 5}},
				{"splice/remove", []int{7, 5}},
				{"splice/remove", []int{4, 5}},
			},
			ExpectedResults: []interface{}{34, 39, 73, 65, 76, 26, 84, 48, 70, 31, []int{64, 45, 10, 0, 1, 0, 1, 2, 3, 4, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 25, 27, 0, 1, 2, 3, 4, 96, 90, 64, 8, 20}},
		},
		{
			Operations: []Operation{
				{"insert_last", []int{105}},
				{"insert_first", []int{105}},
				{"insert_first", []int{142}},
				{"insert_last", []int{130}},
				{"insert_last", []int{83}},
				{"insert_last", []int{75}},
				{"insert_last", []int{78}},
				{"insert_last", []int{83}},
				{"insert_last", []int{82}},
				{"insert_first", []int{49}},
				{"insert_first", []int{117}},
				{"insert_last", []int{75}},
				{"insert_last", []int{122}},
				{"insert_first", []int{99}},
				{"insert_first", []int{14}},
				{"insert_last", []int{6}},
				{"insert_first", []int{17}},
				{"insert_last", []int{103}},
				{"insert_last", []int{101}},
				{"insert_last", []int{142}},
				{"insert_last", []int{62}},
				{"insert_last", []int{85}},
				{"insert_first", []int{47}},
				{"insert_last", []int{82}},
				{"delete_first", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"delete_first", nil},
				{"delete_last", nil},
				{"delete_first", nil},
				{"splice/remove", []int{6, 6}},
				{"splice/remove", []int{9, 6}},
				{"splice/remove", []int{10, 6}},
				{"splice/remove", []int{8, 6}},
				{"splice/remove", []int{2, 6}},
			},
			ExpectedResults: []interface{}{47, 17, 82, 85, 14, 99, 62, 142, 117, 49, 101, 142, []interface{}{105, 105, 130, 0, 1, 2, 3, 4, 5, 83, 75, 78, 83, 0, 1, 0, 1, 2, 3, 4, 5, 2, 0, 0, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 3, 4, 5, 82, 75, 122, 6, 103}},
		},
	}

	for i, test := range tests {
		if check(t, test) == nil {
			t.Logf("Test %02d Passed", i+1)
		}
	}
}

package main

import "testing"

func TestSolution(t *testing.T) {
	tc := []struct {
		input  string
		result bool
	}{
		{`So when I die (the [first] I will see in (heaven) is a score list).`, true},
		{`[ first in ] ( first out ).`, true},
		{`Half Moon tonight (At least it is better than no Moon at all].`, false},
		{`A rope may form )( a trail in a maze.`, false},
		{`Help( I[m being held prisoner in a fortune cookie factory)].`, false},
		{`([ (([( [ ] ) ( ) (( ))] )) ]).`, true},
		{`.`, true},
	}

	for _, v := range tc {
		result := solution(v.input)
		if result != v.result {
			t.Errorf("result should be %t but %t", v.result, result)
		}
	}
}

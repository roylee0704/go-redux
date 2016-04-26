package main

import "testing"

func TestCounter(t *testing.T) {
	tests := []struct {
		state  int
		action *Action
		want   int
	}{
		{
			state:  0,
			action: &Action{"INCREMENT"},
			want:   1,
		},
		{
			state:  1,
			action: &Action{"INCREMENT"},
			want:   2,
		},
		{
			state:  2,
			action: &Action{"DECREMENT"},
			want:   1,
		},
		{
			state:  1,
			action: &Action{"DECREMENT"},
			want:   0,
		},
		{
			state:  1,
			action: &Action{"SOMETHING_ELSE"},
			want:   1,
		},
	}

	for _, test := range tests {
		got := counter(test.state, test.action)
		if got != test.want {
			t.Errorf("%s => got: %d, but want: %d", test.action.typ, got, test.want)
		}
	}

}

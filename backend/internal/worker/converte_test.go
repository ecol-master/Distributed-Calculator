package worker

import (
	// pb "distributed_calculator/internal/proto"
	"testing"
)

type Test struct {
	name      string
	input     string
	want      int
	wantError bool
}

func TestConverter(t *testing.T) {
	cases := []Test{
		{
			name:      "base summ",
			input:     "1 + 2",
			want:      3,
			wantError: false,
		},
		{
			name:      "base multiply",
			input:     "10 * 5 + 5",
			want:      55,
			wantError: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := NewConverter(tc.input).Convert()
			if err != nil && !tc.wantError {
				t.Error("test without errors, but got err: " + err.Error())
			}

			t.Log(res, err)
		})
	}
}

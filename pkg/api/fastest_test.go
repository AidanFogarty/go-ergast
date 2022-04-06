package api

import (
	"context"
	"testing"
)

func TestErgast_FastestLap(t *testing.T) {
	type args struct {
		ctx         context.Context
		year        int
		round       int
		fastestRank int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful FastestLap Response",
			args: args{ctx: context.TODO(), year: 2022, round: 2, fastestRank: 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ergast := New()
			_, _ = ergast.Fastest(context.TODO(), 2022, 2, 1)
		})
	}
}

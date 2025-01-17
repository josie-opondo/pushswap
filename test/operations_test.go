package tests

import (
	operations "pushswap/internal/operations"
	"reflect"
	"testing"
)

func TestGenerateRandomNums(t *testing.T) {
	tests := []struct {
		name    string
		want    []int
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := operations.GenerateRandomNums()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomNums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInts(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operations.GetInts(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

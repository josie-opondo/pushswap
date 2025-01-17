package tests

import (
	"pushswap/internal/helpers"
	stk "pushswap/internal/stack"
	"reflect"
	"testing"
)

func TestChecker_GetInput(t *testing.T) {
	tests := []struct {
		name    string
		c       *helpers.Checker
		want    []byte
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &helpers.Checker{}
			got, err := c.GetInput()
			if (err != nil) != tt.wantErr {
				t.Errorf("Checker.GetInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Checker.GetInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_GetOutput(t *testing.T) {
	type args struct {
		instructions []byte
		stack        []*stk.Stack
	}
	tests := []struct {
		name    string
		c       *helpers.Checker
		args    args
		want    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &helpers.Checker{}
			got, err := c.GetOutput(tt.args.instructions, tt.args.stack)
			if (err != nil) != tt.wantErr {
				t.Errorf("Checker.GetOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Checker.GetOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

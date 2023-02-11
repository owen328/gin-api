package vies

import (
	"fmt"
	"testing"
)

func Test_checkVat(t *testing.T) {
	type args struct {
		VatNumber string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{"FR29919979369"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vat, err := CheckVat(tt.args.VatNumber)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(vat)
			
		})
	}
}
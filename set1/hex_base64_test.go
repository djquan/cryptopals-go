package set1

import (
	"reflect"
	"testing"
)

func Test_hexToBase64(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Sample 1",
			args: args{
				input: "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			},
			want: []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := hexToBase64(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hexToBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

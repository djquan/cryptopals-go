package set1

import (
	"reflect"
	"testing"
)

func Test_fixedXor(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Example 1",
			args: args{
				s1: "1c0111001f010100061a024b53535009181c",
				s2: "686974207468652062756c6c277320657965",
			},
			want: []byte("746865206b696420646f6e277420706c6179"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := fixedXor(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fixedXor() = %v, want %v", got, tt.want)
			}
		})
	}
}

package set1

import (
	"reflect"
	"testing"
)

func Test_repeatingKeyXor(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				input: "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
			},
			want: []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatingKeyXor(tt.args.input, "ICE"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repeatingKeyXor() = %v, want %v", got, tt.want)
			}
		})
	}
}

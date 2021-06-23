package set1

import (
	"strings"
	"testing"
)

func Test_decryptAESinECB(t *testing.T) {
	type args struct {
		filename string
		key      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test",
			args: args{
				filename: "challenge_7.txt",
				key:      "YELLOW SUBMARINE",
			},
			want: "Now you're amazed by the VIP posse",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := decryptAESinECB(tt.args.filename, []byte(tt.args.key)); !strings.Contains(string(got), tt.want) {
				t.Errorf("decryptAESinECB() = %v, want %v", got, tt.want)
			}
		})
	}
}

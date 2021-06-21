package set1

import (
	"testing"
)

func TestBreakCipher(t *testing.T) {
	key, err := breakTheCipher("challenge_6.txt")
	if err != nil {
		t.Errorf("Got error, didn't want error")
	}

	if string(key) != "Terminator X: Bring the noise" {
		t.Errorf("Expected Terminator X: Bring the noise, got %v", string(key))
	}
}

func Test_hammingDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "this is a test",
			args: args{
				a: "this is a test",
				b: "wokka wokka!!!",
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance([]byte(tt.args.a), []byte(tt.args.b)); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

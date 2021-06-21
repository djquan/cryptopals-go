package set1

import "testing"

func Test_detectSingleByteXor(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Example 1",
			args: args{
				filename: "challenge_4.txt",
			},
			want:    "Now that the party is jumping\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := detectSingleByteXor(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("detectSingleByteXor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("detectSingleByteXor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

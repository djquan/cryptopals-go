package set1

import "testing"

func Test_singleByteXor(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name          string
		args          args
		wantedKey     byte
		wantedMessage string
		wantErr       bool
	}{
		{
			name: "Example 1",
			args: args{
				input: "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			},
			wantedKey:     88,
			wantedMessage: "Cooking MC's like a pound of bacon",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, message, err := singleByteXor(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("singleByteXor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantedKey {
				t.Errorf("singleByteXor() got = %v, wantedKey %v", got, tt.wantedKey)
			}

			if message != tt.wantedMessage {
				t.Errorf("singleByteXor() got = %v, wantedMessage %v", message, tt.wantedMessage)
			}
		})
	}
}

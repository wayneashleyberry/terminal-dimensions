package terminaldimensions

import "testing"

func Test_parse(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name       string
		args       args
		wantWidth  uint
		wantHeight uint
		wantErr    bool
	}{
		{
			args: args{
				input: "123 456\n",
			},
			wantWidth:  123,
			wantHeight: 456,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantWidth {
				t.Errorf("parse() got = %v, want %v", got, tt.wantWidth)
			}
			if got1 != tt.wantHeight {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.wantHeight)
			}
		})
	}
}

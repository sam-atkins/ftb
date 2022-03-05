package reporter

import (
	"testing"
)

func Test_formatFloatOrNil(t *testing.T) {
	type args struct {
		teamScore interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Format float",
			args: args{teamScore: 1},
			want: "1",
		},
		{
			name: "Format nil",
			args: args{teamScore: nil},
			want: "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatFloatOrNil(tt.args.teamScore); got != tt.want {
				t.Errorf("formatFloatOrNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToTitle(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Convert string with underscore",
			args: args{input: "IN_PLAY"},
			want: "In Play",
		},
		{
			name: "Convert capitalised string",
			args: args{input: "SCHEDULED"},
			want: "Scheduled",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.input); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

package athenahealth

import "testing"

func TestFlexBool_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    FlexBool
		wantErr bool
	}{
		{
			name: "bool true",
			args: args{data: []byte(`true`)},
			want: true,
		},
		{
			name: "bool false",
			args: args{data: []byte(`false`)},
			want: false,
		},
		{
			name: "empty string",
			args: args{data: []byte(`""`)},
			want: false,
		},
		{
			name: "string true",
			args: args{data: []byte(`"true"`)},
			want: true,
		},
		{
			name: "string false",
			args: args{data: []byte(`"false"`)},
			want: false,
		},
		{
			name: "null",
			args: args{data: []byte(`null`)},
			want: false,
		},
		{
			name:    "non-bool string",
			args:    args{data: []byte(`"maybe"`)},
			wantErr: true,
		},
		{
			name:    "number value",
			args:    args{data: []byte(`1`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(FlexBool)
			err := b.UnmarshalJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlexBool.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && *b != tt.want {
				t.Errorf("FlexBool.UnmarshalJSON() = %v, want %v", *b, tt.want)
			}
		})
	}
}

func TestNumberString_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		n       *NumberString
		args    args
		wantErr bool
	}{
		{
			name:    "string value",
			n:       new(NumberString),
			args:    args{data: []byte(`"55.01"`)},
			wantErr: false,
		},
		{
			name:    "negative string value",
			n:       new(NumberString),
			args:    args{data: []byte(`"-55.01"`)},
			wantErr: false,
		},
		{
			name:    "int value",
			n:       new(NumberString),
			args:    args{data: []byte(`55`)},
			wantErr: false,
		},
		{
			name:    "float64 value",
			n:       new(NumberString),
			args:    args{data: []byte(`55.01`)},
			wantErr: false,
		},
		{
			name:    "invalid bool value",
			n:       new(NumberString),
			args:    args{data: []byte(`false`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Balance.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

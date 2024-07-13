package template

import "testing"

func TestInternalTpl(t *testing.T) {
	type args struct {
		tpl string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "md.tpl",
			args: args{
				tpl: "md.tpl",
			},
			want: true,
		},
		{
			name: "csv.tpl",
			args: args{
				tpl: "csv.tpl",
			},
			want: true,
		},
		{
			name: "excel.tpl",
			args: args{
				tpl: "excel.tpl",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := InternalTpl(tt.args.tpl); got != tt.want {
				t.Errorf("InternalTpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

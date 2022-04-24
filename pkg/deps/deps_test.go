package deps

import (
	"testing"
)

func Test_info(t *testing.T) {
	type args struct {
		pkg string
	}
	tests := []struct {
		name    string
		args    args
		want    *PackageInfo
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				pkg: "github.com/hb-chen/deps",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := info(tt.args.pkg)
			if (err != nil) != tt.wantErr {
				t.Errorf("info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("info() got = %v, want %v", got, tt.want)
			// }
		})
	}
}

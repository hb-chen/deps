package maven

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hb-chen/deps/pkg/graph"
)

func TestMaven_Graphs(t *testing.T) {
	_, projPath, _, _ := runtime.Caller(0)
	projPath = filepath.Dir(filepath.Join(filepath.Dir(projPath), "../.."))
	t.Logf("project filepath: %s", projPath)

	type fields struct {
		opts *graph.Options
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*graph.Dependency
		wantErr bool
	}{
		{
			name: "mvnw",
			fields: fields{
				opts: &graph.Options{
					ProjectPath: filepath.Join(projPath + "/example/java"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maven{
				opts: tt.fields.opts,
			}
			_, err := m.Graphs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Graphs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Graphs() got = %v, want %v", got, tt.want)
			// }
		})
	}
}

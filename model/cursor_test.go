package model

import (
	"reflect"
	"testing"
)

func TestNewCursor(t *testing.T) {
	type args struct {
		id        string
		namespace Gid
	}

	tests := []struct {
		args    args
		want    *Cursor
		wantErr bool
	}{
		{
			args: args{id: "simple", namespace: ""},
			want: &Cursor{
				id:        "simple",
				namespace: "",
				LastId:    "simple",
				LastValue: "simple",
			},
			wantErr: false,
		},
		{
			args:    args{id: "simple", namespace: "product"},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{id: NewId(GidProduct, "simple"), namespace: GidProduct},
			want: &Cursor{
				id:        NewId(GidProduct, "simple"),
				namespace: GidProduct,
				LastId:    "simple",
				LastValue: "simple",
			},
			wantErr: false,
		},
		{
			args:    args{id: "simple", namespace: GidProduct},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run("NewCursor", func(t *testing.T) {
			got, err := NewCursor(tt.args.id, tt.args.namespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCursor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCursor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

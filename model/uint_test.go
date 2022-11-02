package model

import (
	"bytes"
	"testing"
)

func TestUInt_UnmarshalGQL(t *testing.T) {
	tests := []struct {
		i       UInt
		args    any
		wantErr bool
	}{
		{UInt(777), "777", false},
		{UInt(0), "777.777", true},
		{UInt(777), 777.777, false},
		{UInt(7), -7, false},
		{UInt(7), -7.7, false},
	}

	for _, tt := range tests {
		t.Run("UnmarshalGQL", func(t *testing.T) {
			var k UInt
			if err := k.UnmarshalGQL(tt.args); (err != nil) != tt.wantErr || tt.i != k {
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v, value %d", err, tt.wantErr, k)
			}
		})
	}
}

func TestUInt_MarshalGQL(t *testing.T) {
	tests := []struct {
		i     UInt
		wantW string
	}{
		{UInt(0), "0"},
		{UInt(777), "777"},
	}
	for _, tt := range tests {
		t.Run("MarshalGQL", func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.i.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

package scalar

import "testing"

func TestDateTime_UnmarshalGQL(t *testing.T) {
	tests := []struct {
		d       DateTime
		arg     any
		wantErr bool
	}{
		{DateTime(1567871400000), "2019-09-07T15:50:00Z", false},
		{DateTime(1567871400000), "1567871400000", false},
		{DateTime(1567871400000), 1567871400000, false},
		{DateTime(1567871400000), 1567871400000.99, false},
		{DateTime(1567871400000), DateTime(1567871400000), false},
	}
	for _, tt := range tests {
		d := DateTime(0)
		t.Run("UnmarshalGQL", func(t *testing.T) {
			if err := d.UnmarshalGQL(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			} else if d != tt.d {
				t.Errorf("UnmarshalGQL() date = %s, wantDate %s", d.String(), tt.d.String())
			}
		})
	}
}

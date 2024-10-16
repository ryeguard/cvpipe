package util

import (
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	t.SkipNow()

	var tests = []struct {
		name string
		v    interface{}
		want interface{}
	}{
		{"int", 42, Ptr(42)},
		{"string", "hello", Ptr("hello")},
		{"bool", true, Ptr(true)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ptr(tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ptr() = %v, want %v", got, tt.want)
			}
		})
	}

}

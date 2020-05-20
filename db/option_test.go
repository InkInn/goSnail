package db

import "testing"


func TestAdd(t *testing.T) {
	type args struct {
		id   int
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1",args{2,"Tim"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.id,tt.args.name)
		})
	}
}

func TestAddSql(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddSql()
		})
	}
}
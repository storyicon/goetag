package goetag

import (
	"testing"
)

func TestGetEtagByString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			args: args{
				str: "Etag algorithm test string 0",
			},
			want: "Fq0jDGs_kz6eVL9Oiq_oCMZRsRkE",
		},
		{
			name: "test1",
			args: args{
				str: "Etag algorithm test string 1",
			},
			want: "FpMJCDhrP6WCjIecfgjF0pNIPHZX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEtagByString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEtagByString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEtagByString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEtagByBytes(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			args: args{
				content: []byte("Etag algorithm test string 0"),
			},
			want: "Fq0jDGs_kz6eVL9Oiq_oCMZRsRkE",
		},
		{
			name: "test0",
			args: args{
				content: []byte("Etag algorithm test string 1"),
			},
			want: "FpMJCDhrP6WCjIecfgjF0pNIPHZX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEtag, err := GetEtagByBytes(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEtagByBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEtag != tt.want {
				t.Errorf("GetEtagByBytes() = %v, want %v", gotEtag, tt.want)
			}
		})
	}
}

func TestGetEtagByPath(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test0",
			args: args{
				filepath: "./LICENSE",
			},
			want: "FrMUx-u31ZmUSYGQi38-0zow5486",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEtagByPath(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEtagByPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEtagByPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

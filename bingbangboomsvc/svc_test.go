package bingbangboomsvc

import (
	"reflect"
	"testing"
)

func Test_svc_GenerateMapping(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    []Domain
		wantErr bool
	}{
		{
			name: "A",
			args: args{
				id: 3,
			},
			want:    []Domain{Domain_A},
			wantErr: false,
		},
		{
			name: "B",
			args: args{
				id: 5,
			},
			want:    []Domain{Domain_B},
			wantErr: false,
		},
		{
			name: "Both",
			args: args{
				id: 15,
			},
			want:    []Domain{Domain_A, Domain_B},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := svc{}
			got, err := s.GenerateMapping(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateMapping() got = %v, want %v", got, tt.want)
			}
		})
	}
}

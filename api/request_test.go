package api

import (
	"testing"
)

func TestRequest_Do_Author(t *testing.T) {
	type args struct {
		api Requester
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want bool
	}{
		{name: "author", r: NewRequest(), args: args{api: NewAuthorRequest("soul")}, want: false},
		{name: "author", r: NewRequest(), args: args{api: NewAuthorRequest("wizardforcel")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Do(tt.args.api); got.Success() != tt.want {
				t.Errorf("Request.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

package geo

import (
	"math"
	"testing"
)

func TestDegToRad(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 degree to radian",
			args: args{angle: 0},
			want: 0,
		}, {
			name: "-45 degree to radian",
			args: args{angle: -45},
			want: -math.Pi / 4,
		}, {
			name: "45 degree to radian",
			args: args{angle: 45},
			want: math.Pi / 4,
		}, {
			name: "180 degree to radian",
			args: args{angle: 180},
			want: math.Pi,
		}, {
			name: "200 degree to radian",
			args: args{angle: 200},
			want: 3.490658503988659,
		}, {
			name: "360 degree to radian",
			args: args{angle: 360},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DegToRad(tt.args.angle); got != tt.want {
				t.Errorf("DegToRad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadToDeg(t *testing.T) {
	type args struct {
		angle float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 radian to degree",
			args: args{angle: 0},
			want: 0,
		}, {
			name: "-45 radian to degree",
			args: args{angle: -math.Pi / 4},
			want: -45,
		}, {
			name: "PI/4 radian to degree",
			args: args{angle: math.Pi / 4},
			want: 45,
		}, {
			name: "PI radian to degree",
			args: args{angle: math.Pi},
			want: 180,
		}, {
			name: "3.490658503988659 radian to degree",
			args: args{angle: 3.490658503988659},
			want: 200,
		}, {
			name: "2PI radian to degree",
			args: args{angle: math.Pi * 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RadToDeg(tt.args.angle); got != tt.want {
				t.Errorf("RadToDeg() = %v, want %v", got, tt.want)
			}
		})
	}
}

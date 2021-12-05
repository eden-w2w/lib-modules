package types

import (
	"testing"
)

var areas = FreightRuleAreas{
	{
		ADCode: "111",
		Name:   "111",
		Level:  0,
	},
	{
		ADCode: "222",
		Name:   "222",
		Level:  0,
	},
	{
		ADCode: "333",
		Name:   "333",
		Level:  0,
	},
	{
		ADCode: "444",
		Name:   "444",
		Level:  0,
	},
}

func TestFreightRuleAreas_Contain(t *testing.T) {
	type args struct {
		area FreightRuleArea
	}
	tests := []struct {
		name string
		v    FreightRuleAreas
		args args
		want bool
	}{
		{
			name: "contain",
			v:    areas,
			args: args{area: FreightRuleArea{
				ADCode: "111",
				Name:   "111",
				Level:  0,
			}},
			want: true,
		},
		{
			name: "not contain",
			v:    areas,
			args: args{area: FreightRuleArea{
				ADCode: "555",
				Name:   "555",
				Level:  0,
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.v.Contain(tt.args.area); got != tt.want {
					t.Errorf("Contain() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestFreightRuleAreas_ContainsAll(t *testing.T) {
	type args struct {
		areas FreightRuleAreas
	}
	tests := []struct {
		name string
		v    FreightRuleAreas
		args args
		want bool
	}{
		{
			name: "contain",
			v:    areas,
			args: args{
				areas: FreightRuleAreas{
					{
						ADCode: "111",
						Name:   "111",
						Level:  0,
					},
					{
						ADCode: "222",
						Name:   "222",
						Level:  0,
					},
				},
			},
			want: true,
		},
		{
			name: "not contain",
			v:    areas,
			args: args{
				areas: FreightRuleAreas{
					{
						ADCode: "111",
						Name:   "111",
						Level:  0,
					},
					{
						ADCode: "555",
						Name:   "555",
						Level:  0,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.v.ContainsAll(tt.args.areas); got != tt.want {
					t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestFreightRuleAreas_ContainsOne(t *testing.T) {
	type args struct {
		areas FreightRuleAreas
	}
	tests := []struct {
		name string
		v    FreightRuleAreas
		args args
		want bool
	}{
		{
			name: "contain",
			v:    areas,
			args: args{
				areas: FreightRuleAreas{
					{
						ADCode: "111",
						Name:   "111",
						Level:  0,
					},
					{
						ADCode: "222",
						Name:   "222",
						Level:  0,
					},
				},
			},
			want: true,
		},
		{
			name: "not contain",
			v:    areas,
			args: args{
				areas: FreightRuleAreas{
					{
						ADCode: "111",
						Name:   "111",
						Level:  0,
					},
					{
						ADCode: "555",
						Name:   "555",
						Level:  0,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.v.ContainsOne(tt.args.areas); got != tt.want {
					t.Errorf("ContainsOne() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

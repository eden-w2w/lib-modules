package freight_trial

import (
	"github.com/eden-w2w/lib-modules/databases"
	"testing"
)

func TestFreightTrial(t *testing.T) {
	type args struct {
		goodsList []FreightTrialParams
		shipping  *databases.ShippingAddress
	}
	tests := []struct {
		name              string
		args              args
		wantFreightAmount uint64
		wantErr           bool
	}{
		{
			name: "default",
			args: args{
				goodsList: []FreightTrialParams{},
				shipping:  &databases.ShippingAddress{},
			},
			wantFreightAmount: 0,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotFreightAmount, err := FreightTrial(tt.args.goodsList, tt.args.shipping)
				if (err != nil) != tt.wantErr {
					t.Errorf("FreightTrial() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotFreightAmount != tt.wantFreightAmount {
					t.Errorf("FreightTrial() gotFreightAmount = %v, want %v", gotFreightAmount, tt.wantFreightAmount)
				}
			},
		)
	}
}

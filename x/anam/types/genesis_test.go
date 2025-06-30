package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/anam-nw/anam/x/anam/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid basic genesis state",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
				Params: types.Params{
					MaxSupply: "10000000",
				},
			},
			valid: true,
		},
		{
			desc: "invalid max supply when is small",
			genState: &types.GenesisState{
				Params: types.Params{
					MaxSupply: "5",
				},
			},
			valid: false,
		},

		{
			desc: "limit list exclude address",
			genState: &types.GenesisState{
				Params: types.Params{
					ExcludeCirculatingAddr: []string{"addr1", "addr2", "addr3", "addr4", "addr5", "addr6", "addr7", "addr8", "addr9", "addr10", "addr11"},
					MaxSupply:              types.DefaultParams().MaxSupply,
				},
			},
			valid: false,
		},
		{
			desc: "duplicate list exclude address",
			genState: &types.GenesisState{
				Params: types.Params{
					ExcludeCirculatingAddr: []string{"addr1", "addr1", "addr2"},
					MaxSupply:              types.DefaultParams().MaxSupply,
				},
			},
			valid: false,
		},
		{
			desc: "invalid addr format",
			genState: &types.GenesisState{
				Params: types.Params{
					ExcludeCirculatingAddr: []string{"addr1", "anam1jlp9ge244um2v7mdm7xwamwsv9z9vhpej6wjh7"},
					MaxSupply:              types.DefaultParams().MaxSupply,
				},
			},
			valid: false,
		},
		{
			desc: "valid addr format",
			genState: &types.GenesisState{
				Params: types.Params{
					ExcludeCirculatingAddr: []string{"anam19ad4tprcf9ew4577qph3jfzpf9slcrkpmxwvah", "anam1jlp9ge244um2v7mdm7xwamwsv9z9vhpej6wjh7"},
					MaxSupply:              types.DefaultParams().MaxSupply,
				},
			},
			valid: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			sdk.GetConfig().SetBech32PrefixForAccount("anam", "anampub")
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

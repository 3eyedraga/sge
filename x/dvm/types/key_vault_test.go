package types_test

import (
	"testing"

	simappUtil "github.com/sge-network/sge/testutil/simapp"
	"github.com/sge-network/sge/x/dvm/types"
	"github.com/stretchr/testify/require"
)

func TestMajority(t *testing.T) {
	tests := []struct {
		name       string
		publicKeys []string

		exp int64
	}{
		{
			name:       "odd public keys",
			publicKeys: simappUtil.GenerateDvmPublicKeys(5),
			exp:        4,
		},
		{
			name:       "even public keys",
			publicKeys: simappUtil.GenerateDvmPublicKeys(4),
			exp:        3,
		},
		{
			name:       "large odd number",
			publicKeys: simappUtil.GenerateDvmPublicKeys(51),
			exp:        35,
		},
		{
			name:       "large even number",
			publicKeys: simappUtil.GenerateDvmPublicKeys(50),
			exp:        34,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keyVault := types.KeyVault{
				PublicKeys: tc.publicKeys,
			}
			require.Equal(t, tc.exp, keyVault.MajorityCount())
		})
	}

}

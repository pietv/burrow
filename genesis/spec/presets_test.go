package spec

import (
	"testing"

	"github.com/hyperledger/burrow/keys/mock"
	permission "github.com/hyperledger/burrow/permission/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeGenesisSpecAccounts(t *testing.T) {
	keyClient := mock.NewKeyClient()
	gs := MergeGenesisSpecs(FullAccount("0"), ParticipantAccount("1"), ParticipantAccount("2"))
	gd, err := gs.GenesisDoc(keyClient)
	require.NoError(t, err)
	assert.Len(t, gd.Validators, 1)
	assert.Len(t, gd.Accounts, 3)
}

func TestMergeGenesisSpecGlobalPermissions(t *testing.T) {
	gs1 := GenesisSpec{
		GlobalPermissions: []string{permission.CreateAccountString, permission.CreateAccountString},
	}
	gs2 := GenesisSpec{
		GlobalPermissions: []string{permission.SendString, permission.CreateAccountString, permission.HasRoleString},
	}

	gsMerged := MergeGenesisSpecs(gs1, gs2)
	assert.Equal(t, []string{permission.CreateAccountString, permission.HasRoleString, permission.SendString},
		gsMerged.GlobalPermissions)
}

func TestMergeGenesisSpecsRepeatedAccounts(t *testing.T) {
	name1 := "Party!"
	name3 := "Counter!"

	amt1 := uint64(5)
	amt2 := uint64(2)
	amt3 := uint64(9)

	gs1 := GenesisSpec{
		Accounts: []TemplateAccount{
			{
				Name:        name1,
				Amount:      &amt1,
				Permissions: []string{permission.SendString, permission.CreateAccountString, permission.HasRoleString},
				Roles:       []string{"fooer"},
			},
		},
	}
	gs2 := GenesisSpec{
		Accounts: []TemplateAccount{
			{
				Name:        name1,
				Amount:      &amt2,
				Permissions: []string{permission.SendString, permission.CreateAccountString},
				Roles:       []string{"barer"},
			},
		},
	}
	gs3 := GenesisSpec{
		Accounts: []TemplateAccount{
			{
				Name:   name3,
				Amount: &amt3,
			},
		},
	}

	gsMerged := MergeGenesisSpecs(gs1, gs2, gs3)
	bsMerged, err := gsMerged.JSONBytes()
	require.NoError(t, err)

	amtExpected := amt1 + amt2
	gsExpected := GenesisSpec{
		Accounts: []TemplateAccount{
			{
				Name:        name1,
				Amount:      &amtExpected,
				Permissions: []string{permission.CreateAccountString, permission.HasRoleString, permission.SendString},
				Roles:       []string{"barer", "fooer"},
			},
			gs3.Accounts[0],
		},
	}
	bsExpected, err := gsExpected.JSONBytes()
	require.NoError(t, err)
	if !assert.Equal(t, string(bsExpected), string(bsMerged)) {
		t.Logf("Expected:\n%s\n\nActual:\n%s", string(bsExpected), string(bsMerged))
	}
}

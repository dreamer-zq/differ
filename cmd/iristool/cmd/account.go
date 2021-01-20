package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetCmdAccount return the account module command
func GetCmdAccount() *cobra.Command {
	accCmd := &cobra.Command{
		Use:                        "account",
		Short:                      "Asset transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	accCmd.AddCommand(
		GetCmdQueryAllAccount(),
	)
	return accCmd
}

// GetCmdQueryAllAccount return the balance of each aoount
func GetCmdQueryAllAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list",
		Long: "Query all the account information",
		RunE: func(cmd *cobra.Command, args []string) error {
			home := viper.GetString(flagHome)
			app, err := NewStoreViewer(home)
			if err != nil {
				return err
			}

			ctx := app.BaseApp.NewContext(true, tmproto.Header{})
			accounts := app.AccountKeeper.GetAllAccounts(ctx)

			var accMap = make(map[string]sdk.Coins, len(accounts))
			for _, acc := range accounts {
				switch acc := acc.(type) {
				case (*authtypes.BaseAccount):
					balance := app.BankKeeper.GetAllBalances(ctx, acc.GetAddress())
					accMap[acc.Address] = balance
				case (*authtypes.ModuleAccount):
					balance := app.BankKeeper.GetAllBalances(ctx, acc.GetAddress())
					accMap[acc.Address] = balance
				}
			}
			bz, err := json.Marshal(accMap)
			if err != nil {
				return err
			}
			fmt.Println(string(bz))
			return nil
		},
	}
	return cmd
}

type account struct {
	Typ     string    `json:"type"`
	Address string    `json:"address"`
	Balance sdk.Coins `json:"balance"`
}

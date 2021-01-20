package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irishub/simapp"
)

const (
	flagHome = "home"
)

// NewToolCmd returns the transaction commands for the token module.
func NewToolCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "iristool",
		Short:                      "Account subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdAccount(),
	)
	txCmd.PersistentFlags().String(flagHome, "", "Choose sign mode (direct|amino-json), this is an advanced feature")

	viper.BindPFlags(txCmd.Flags())
	viper.BindPFlags(txCmd.PersistentFlags())
	return txCmd
}

// NewStoreViewer return the SimApp
func NewStoreViewer(home string) (*simapp.SimApp, error) {
	db, err := openDB(home)
	if err != nil {
		return nil, err
	}

	app := simapp.NewSimApp(log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		home,
		5,
		simapp.MakeEncodingConfig(),
		simapp.EmptyAppOptions{})
	return app, nil
}

func openDB(rootDir string) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	return sdk.NewLevelDB("application", dataDir)
}

// Step 14: add query cmd which are getDeals and getDeal
package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/earth2378/logistic/x/logistic/types"
)

// GetCmdDeals receive orderid and deal of the orderid
func GetCmdDeal(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-deal [orderid]",
		Short: "query deal of orderid",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			orderid := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryDeal, orderid), nil)
			if err != nil {
				fmt.Printf("could not resolve deal %s \n%s\n", orderid, err.Error())
				return nil
			}
			var out types.Deal

			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdDeals receive orderid and deal of the orderid
func GetCmdListDeal(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-deal ",
		Short: "query all deals of orderid",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryListDeal), nil)
			if err != nil {
				fmt.Printf("could not resolve all deals \n%s\n", err.Error())
				return nil
			}
			var out []types.Deal

			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// REFERENCE from ACTIVITY 7

// func GetCmdListCommit(queryRoute string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "list-commit",
// 		Short: "list all commit",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryListCommit), nil)
// 			if err != nil {
// 				fmt.Printf("could not list Commit\n%s\n", err.Error())
// 				return nil
// 			}
// 			var out []types.Commit
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }

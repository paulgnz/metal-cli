// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package signatureaggregatorcmd

import (
	"github.com/paulgnz/metal-cli/pkg/application"
	"github.com/paulgnz/metal-cli/pkg/cobrautils"
	"github.com/spf13/cobra"
)

var app *application.Avalanche

// metal interchain signatureAggregator
func NewCmd(injectedApp *application.Avalanche) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signatureAggregator",
		Short: "Manage ICM signature aggregator",
		Long: `The signature aggregator command suite provides a collection of tools for deploying
and configuring ICM signature aggregator.`,
		RunE: cobrautils.CommandSuiteUsage,
	}
	app = injectedApp
	cmd.AddCommand(newStopCmd())
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newStartCmd())
	return cmd
}

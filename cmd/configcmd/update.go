// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package configcmd

import (
	"errors"

	"github.com/paulgnz/metal-cli/pkg/cobrautils"
	"github.com/paulgnz/metal-cli/pkg/constants"
	"github.com/paulgnz/metal-cli/pkg/ux"
	"github.com/spf13/cobra"
)

// metal config metrics command
func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [enable | disable]",
		Short: "opt in or out of update check",
		Long:  "set user preference between update check or not",
		RunE:  handleUpdateSettings,
		Args:  cobrautils.ExactArgs(1),
	}

	return cmd
}

func handleUpdateSettings(_ *cobra.Command, args []string) error {
	switch args[0] {
	case constants.Enable:
		ux.Logger.PrintToUser("Thank you for opting in Metal CLI automated update check")
		err := saveUpdateDisabledPreferences(false)
		if err != nil {
			return err
		}
	case constants.Disable:
		ux.Logger.PrintToUser("Metal CLI automated update check will no longer be performed")
		err := saveUpdateDisabledPreferences(true)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid update argument '" + args[0] + "'")
	}
	return nil
}

func saveUpdateDisabledPreferences(disableUpdate bool) error {
	return app.Conf.SetConfigValue(constants.ConfigUpdatesDisabledKey, disableUpdate)
}

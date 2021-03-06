/*
 * Copyright 2020 Sheaf Authors
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package commands

import (
	"github.com/spf13/cobra"

	"github.com/bryanl/sheaf/pkg/commands/config"
)

// NewConfigCommand creates a config command.
func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "config",
		Short:        "Perform actions on a bundle directory",
		SilenceUsage: true,
	}

	cmd.AddCommand(
		config.NewAddImageCommand(),
		config.NewSetUserDefinedImage(),
		config.NewDeleteUserDefinedImage(),
		config.NewGetCommand(),
		config.NewPushCommand(),
		config.NewPullCommand())

	return cmd
}

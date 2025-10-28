package bootstrap

import (
	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
func BootstrapCmd() *cobra.Command {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "Bootstraps AI services infrastructure",
		Long: `Bootstrap and configure the AI services infrastructure.

The bootstrap command helps you set up and validate the environment
required to run AI services on Power11 systems.

Available subcommands:
  validate   - Validate system prerequisites and configuration
  configure  - Configure and initialize the AI services infrastructure`,
		Example: `  # Validate the environment
  aiservices bootstrap validate

  # Configure the infrastructure
  aiservices bootstrap configure

  # Get help on a specific subcommand
  aiservices bootstrap validate --help`,
	}

	// subcommands
	bootstrapCmd.AddCommand(validateCmd())
	bootstrapCmd.AddCommand(configureCmd())

	return bootstrapCmd
}

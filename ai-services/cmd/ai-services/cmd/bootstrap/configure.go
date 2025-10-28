package bootstrap

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/project-ai-services/ai-services/internal/pkg/validators"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate subcommand of bootstrap
func configureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "configures the LPAR environment",
		Long:  `Configure and initialize the LPAR.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			logger.Info("Running bootstrap configuration...")

			if err := rootCheck(); err != nil {
				return err
			}

			// 1. Install and configure Podman if not done
			// 1.1 Install Podman
			if _, err := validators.Podman(); err != nil {
				// setup podman socket and enable service
				logger.Info("Podman not installed. Installing Podman...")
				if err := installPodman(); err != nil {
					return err
				}
			}

			// 1.2 Configure Podman
			if err := validators.PodmanHealthCheck(); err != nil {
				logger.Info("Podman not configured. Configuring Podman...")
				if err := setupPodman(); err != nil {
					return err
				}
			} else {
				logger.Info("✅ Podman already configured")
			}
			// 2. Configure service-report package if not done already
			// 3. Spyre cards – run service-report
			// 4. Check SMT level and set the SMT to 2
			logger.Info("✅ Bootstrap configuration completed successfully.")
			return nil
		},
	}
	return cmd
}

func installPodman() error {
	cmd := exec.Command("dnf", "-y", "install", "podman")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install podman: %v, output: %s", err, string(out))
	}
	logger.Info("✅ Podman installed successfully.")
	return nil
}

func setupPodman() error {

	// start podman socket
	if err := systemctl("start", "podman.socket"); err != nil {
		return fmt.Errorf("failed to start podman socket: %w", err)
	}
	// enable podman socket
	if err := systemctl("enable", "podman.socket"); err != nil {
		return fmt.Errorf("failed to enable podman socket: %w", err)
	}

	logger.Debug("Waiting for podman socket to be ready...")
	time.Sleep(2 * time.Second) // wait for socket to be ready

	if err := validators.PodmanHealthCheck(); err != nil {
		return fmt.Errorf("podman health check failed after configuration: %w", err)
	}

	logger.Info("✅ Podman configured successfully.")
	return nil
}

func systemctl(action, unit string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", action, unit)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to %s %s: %v, output: %s", action, unit, err, string(out))
	}
	return nil
}

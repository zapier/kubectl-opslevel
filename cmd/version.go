package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var version = "development"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information`,
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	log.Info().Msgf(version)
}

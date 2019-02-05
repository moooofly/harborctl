package cmd

import (
	"github.com/moooofly/harborctl/utils"
	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check the health of chart repository service.",
	Long:  `This endpoint is for user to check the health of chart repository service.`,
	Run: func(cmd *cobra.Command, args []string) {
		healthcheckChartRepo()
	},
}

func init() {
	chartrepoCmd.AddCommand(healthCmd)
}

func healthcheckChartRepo() {
	targetURL := chartrepoURL + "/health"
	utils.Get(targetURL)
}

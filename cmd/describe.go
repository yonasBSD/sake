package cmd

import (
	"github.com/spf13/cobra"

	"github.com/alajmo/yac/core/dao"
)

func describeCmd(config *dao.Config, configErr *error) *cobra.Command {
	cmd := cobra.Command{
		Aliases: []string{"desc"},
		Use:     "describe <projects|tasks>",
		Short:   "Describe projects and tasks",
		Long:    "Describe projects and tasks.",
		Example: `  # Describe projects
  yac describe projects

  # Describe tasks
  yac describe tasks`,
	}

	cmd.AddCommand(
		describeDirsCmd(config, configErr),
		describeProjectsCmd(config, configErr),
		describeTasksCmd(config, configErr),
		describeNetworksCmd(config, configErr),
	)

	return &cmd
}
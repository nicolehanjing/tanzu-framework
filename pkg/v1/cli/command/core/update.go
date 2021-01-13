package core

import (
	"fmt"
	"strings"

	"github.com/aunum/log"
	"github.com/spf13/cobra"

	"github.com/AlecAivazis/survey/v2"
	"github.com/vmware-tanzu-private/core/pkg/v1/cli"
)

var yesUpdate bool

func init() {
	updateCmd.SetUsageFunc(cli.SubCmdUsageFunc)
	updateCmd.Flags().BoolVarP(&yesUpdate, "yes", "y", false, "force update; skip prompt")
	updateCmd.Flags().StringSliceVarP(&local, "local", "l", []string{}, "path to local repository")
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the CLI",
	Annotations: map[string]string{
		"group": string(cli.SystemCmdGroup),
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		catalog, err := cli.NewCatalog()
		if err != nil {
			return err
		}
		plugins, err := catalog.List()
		if err != nil {
			return err
		}

		repos := getRepositories()
		communityRepo, err := repos.GetRepository(cli.CommunityRepositoryName)
		if err != nil {
			return err
		}

		type updateInfo struct {
			version string
			repo    cli.Repository
		}

		updateMap := map[*cli.PluginDescriptor]updateInfo{}
		for _, plugin := range plugins {
			update, repo, version, err := plugin.HasUpdateIn(repos)
			if err != nil {
				return err
			}
			if update {
				updateMap[plugin] = updateInfo{version, repo}
			}
		}
		coreUpdate, coreVersion, err := cli.HasUpdate(communityRepo)
		if err != nil {
			return err
		}

		if len(updateMap) == 0 && !coreUpdate {
			log.Info("everything up to date")
			return nil
		}

		log.Info("the following updates will take place:")
		if coreUpdate {
			fmt.Printf("     %s %s → %s\n", cli.CoreName, cli.BuildVersion, coreVersion)
		}
		for plugin, version := range updateMap {
			fmt.Printf("     %s %s → %s\n", plugin.Name, plugin.Version, version)
		}
		// formatting
		fmt.Println()

		if !yesUpdate {
			input := &survey.Input{Message: "would you like to continue? [y/n]"}
			var resp string
			err := survey.AskOne(input, &resp)
			if err != nil {
				return err
			}
			update := strings.ToLower(resp)
			if update != "y" && update != "yes" {
				log.Info("aborting update")
				return nil
			}
		}
		for plugin, info := range updateMap {
			err := catalog.Install(plugin.Name, info.version, info.repo)
			if err != nil {
				return err
			}
		}

		// update core
		err = cli.Update(communityRepo)
		if err != nil {
			return err
		}

		// formatting
		fmt.Println()
		log.Success("successfully updated CLI")
		return nil
	},
}
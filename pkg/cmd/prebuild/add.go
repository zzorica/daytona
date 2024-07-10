// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"github.com/spf13/cobra"
)

var prebuildAddCmd = &cobra.Command{
	Use:   "add [WORKSPACE]",
	Short: "Add workspace prebuild configuration",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// var repo *apiclient.GitRepository
		// var repoUrl string
		// ctx := context.Background()

		// apiClient, err := apiclient_util.GetApiClient(nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// gitProviders, _, err := apiClient.GitProviderAPI.ListGitProviders(ctx).Execute()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// if len(args) == 0 {

		// 	if len(gitProviders) > 0 {
		// 		repo, err = workspace_util.GetRepositoryFromWizard(gitProviders, 0)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		repoUrl = *repo.Url
		// 	} else {
		// 		log.Fatal("No git providers found")
		// 	}

		// } else {
		// 	repoUrl = args[0]
		// }

		// _, err = apiClient.PrebuildAPI.RegisterPrebuildWebhook(ctx).PrebuildWebhook(apiclient.RegisterPrebuildWebhookRequest{
		// 	GitUrl: &repoUrl,
		// }).Execute()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	},
}

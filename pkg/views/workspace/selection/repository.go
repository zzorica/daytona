// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package selection

import (
	"fmt"
	"os"

	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/views"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func selectRepositoryPrompt(repositories []apiclient.GitRepository, additionalProjectOrder int, choiceChan chan<- string, parentId string, selectedReposParentMap map[string]bool, selectedRepos map[string]bool) {
	items := []list.Item{}
	disabledReposCount := 0

	// Populate items with titles and descriptions from workspaces.
	for _, repository := range repositories {
		url := *repository.Url
		title := *repository.Name
		isDisabled := false

		// Index > 1 indicates use of 'multi-project' command
		if additionalProjectOrder > 1 && len(selectedRepos) > 0 && selectedRepos[url] {
			title += statusMessageDangerStyle(" (Already selected)")
			// isDisabled property helps in skipping over this specific repo option, refer to
			// handling of up/down key press under update method in ./view.go file
			isDisabled = true
			disabledReposCount++
		}
		newItem := item[string]{id: url, title: title, choiceProperty: url, desc: url, isDisabled: isDisabled}
		items = append(items, newItem)
	}

	l := views.GetStyledSelectList(items)

	title := "Choose a Repository"
	if additionalProjectOrder > 1 {
		title += fmt.Sprintf(" (Project #%d)", additionalProjectOrder)
	}
	l.Title = views.GetStyledMainTitle(title)
	l.Styles.Title = titleStyle
	m := model[string]{list: l}

	p, err := tea.NewProgram(m, tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if m, ok := p.(model[string]); ok && m.choice != nil {
		choice := *m.choice

		selectedRepos[choice] = true
		disabledReposCount++
		if disabledReposCount == len(repositories) {
			selectedReposParentMap[parentId] = true
		}

		choiceChan <- choice
	} else {
		choiceChan <- ""
	}
}

func GetRepositoryFromPrompt(repositories []apiclient.GitRepository, additionalProjectOrder int, parentId string, selectedReposParentMap map[string]bool, selectedRepos map[string]bool) *apiclient.GitRepository {
	choiceChan := make(chan string)

	go selectRepositoryPrompt(repositories, additionalProjectOrder, choiceChan, parentId, selectedReposParentMap, selectedRepos)

	choice := <-choiceChan

	for _, repository := range repositories {
		if *repository.Url == choice {
			return &repository
		}
	}

	return nil
}

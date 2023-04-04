package prompts

import "github.com/AlecAivazis/survey/v2"

func AskWhichGitlabGroup() []*survey.Question {
	// TODO: get list of available groups from API

	return []*survey.Question{
		{
			Name: "authenticationMethod",
			Prompt: &survey.Select{
				Message: "Authentication method:",
				Options: []string{"Personal Access Token", "Username & Password"},
			},
		},
	}
}

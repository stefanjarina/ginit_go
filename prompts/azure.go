package prompts

import "github.com/AlecAivazis/survey/v2"

func AskWhichAzureProject() []*survey.Question {
	// TODO: get list of available projects from API

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

func AskForAzureAuth() []*survey.Question {
	var questions = []*survey.Question{
		{
			Name:     "orgName",
			Prompt:   &survey.Input{Message: "Enter your organization name (https://dev.azure.com/{yourorgname})"},
			Validate: survey.Required,
		},
	}

	questions = append(questions, AskForToken()...)

	return questions
}

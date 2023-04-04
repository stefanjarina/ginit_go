package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/stefanjarina/ginit/api/gitignoreio"
	"log"
	"os"
	"path/filepath"
)

func AskForToken() []*survey.Question {
	var questions = []*survey.Question{
		{
			Name:     "token",
			Prompt:   &survey.Input{Message: "Enter your Personal Access Token"},
			Validate: survey.Required,
		},
	}

	return questions
}

func GetTwoFactorAuthenticationCode() []*survey.Question {
	var questions = []*survey.Question{
		{
			Name:     "twoFactorAuthenticationCode",
			Prompt:   &survey.Input{Message: "Enter your two-factor authentication code:"},
			Validate: survey.Required,
		},
	}

	return questions
}

func GetRepoDetailQuestions(repository string, repoName string, description string) []*survey.Question {
	var visibilityChoices []string

	switch repository {
	case "azure":
		visibilityChoices = []string{"private", "public"}
	case "github":
		visibilityChoices = []string{"private", "public"}
	case "gitlab":
		visibilityChoices = []string{"private", "internal", "public"}
	}

	if repoName == "" {
		wd, _ := os.Getwd()
		repoName = filepath.Base(wd)
	}

	var questions = []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "Enter a name for the repository", Default: repoName},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:      "description",
			Prompt:    &survey.Input{Message: "Optionally enter a description of the repository", Default: description},
			Transform: survey.ToLower,
		},
		{
			Name: "visibility",
			Prompt: &survey.Select{
				Message: "Visibility of a repository",
				Options: visibilityChoices,
				Default: "public",
			},
		},
	}

	return questions
}

func GetGitIgnoreQuestions() []*survey.Question {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	files := getListOfFiles(currentDir)

	giClient := gitignoreio.NewClient()

	// Get the list of available types.
	availableTypes, err := giClient.List()
	if err != nil {
		log.Fatal(err)
	}

	var questions = []*survey.Question{
		{
			Name: "localFiles",
			Prompt: &survey.MultiSelect{
				Message: "Visibility of a repository",
				Options: files,
				Default: []string{"package.json"},
			},
		},
		{
			Name: "gitignore",
			Prompt: &survey.MultiSelect{
				Message: "Select config names you wish to fetch from https://gitignore.io",
				Options: availableTypes,
				Default: []string{"windows", "linux", "macos", "node", "dotnetcore", "visualstudiocode", "webstorm+all"},
			},
		},
	}

	return questions
}

func getListOfFiles(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders

	return list
}

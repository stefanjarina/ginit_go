package api

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/stefanjarina/ginit/prompts"
	"time"
)

func GetAnswers(repo string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	answers := struct {
		Name                 string
		Description          string
		Visibility           string
		LocalFiles           []string
		Gitignore            []string
		AuthenticationMethod string
	}{}

	// prepare questions
	s.Start()
	repoDetailQuestions := prompts.GetRepoDetailQuestions(repo, "", "")
	gitIgnoreQuestions := prompts.GetGitIgnoreQuestions()
	questions := append(repoDetailQuestions, gitIgnoreQuestions...)
	s.Stop()

	//execute base questions
	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Name:", answers.Name)
	fmt.Println("Description:", answers.Description)
	fmt.Println("Visibility:", answers.Visibility)
	fmt.Printf("LocalFiles: %v\n", answers.LocalFiles)
	fmt.Printf("Gitignore: %v\n", answers.Gitignore)

	survey.Ask(prompts.AskGithubAuthenticationMethod(), &answers, survey.WithPageSize(10))

	fmt.Println("authenticationMethod:", answers.AuthenticationMethod)
}

package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
)

func AskGithubAuthenticationMethod() []*survey.Question {
	return []*survey.Question{
		{
			Name: "authenticationMethod",
			Prompt: &survey.Select{
				Message: "Authentication method:",
				Options: []string{"Personal Access Token", "Username & Password"},
			},
			Transform: normalized,
		},
	}
}

func normalized(ans interface{}) interface{} {
	data, _ := ans.(core.OptionAnswer)
	transformer := survey.TransformString(normalize)
	result := transformer(data.Value)
	return core.OptionAnswer{Value: result.(string), Index: data.Index}
}

func normalize(s string) string {
	switch s {
	case "Personal Access Token":
		return "token"
	case "Username & Password":
		return "username_password"
	default:
		return s
	}
}

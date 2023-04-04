package globals

import "github.com/stefanjarina/ginit/utils"

var SupportedRepos = []string{"azure", "github", "gitlab"}
var Repos = utils.NewEnum(SupportedRepos, "")

package cmd

import (
	"fmt"
	"log"

	"github.com/jbpratt78/apis/grequests/githubAPITool/models"
	"github.com/levigross/grequests"
	"github.com/spf13/cobra"
)

// Fetch repos for given user
func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch the repo details with user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			var repos []models.Repo
			user := args[0]
			var repoURL = fmt.Sprintf("https://api.github.com/users/%s/repos", user)
			resp := getStats(repoURL)
			resp.JSON(&repos)
			log.Println(repos)
		} else {
			log.Println("Please give a username.. Use -h for help")
		}
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}

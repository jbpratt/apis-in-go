package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/jbpratt78/apis/grequests/githubAPITool/models"
	"github.com/levigross/grequests"
	"github.com/spf13/cobra"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

// Read files provided and create Gist on github
func createGist(url string, args []string) *grequests.Response {
	description := args[0]
	var fileContents = make(map[string]models.File)
	for i := 1; i < len(args); i++ {
		data, err := ioutil.ReadFile(args[i])
		if err != nil {
			log.Println("Please check the filenames. Absolute path (or) same directory are allowed")
			return nil
		}
		log.Println(args[i])
		var file models.File
		file.Content = string(data)
		fileContents[args[i]] = file
	}
	var gist = models.Gist{Description: description, Public: true, Files: fileContents}
	var postBody, _ = json.Marshal(gist)
	var requestOptions_copy = requestOptions
	// add data to JSON field
	requestOptions_copy.JSON = string(postBody)
	resp, err := grequests.Post(url, requestOptions_copy)
	if err != nil {
		log.Println("Create request failed for Github API")
	}
	return resp
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creats a gist from the given file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			postUrl := "https://api.github.com/gists"
			resp := createGist(postUrl, args)
			log.Println(resp.String())
		} else {
			log.Println("Please give sufficient arguments")
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ArtifactName string

func init() {
	rootCmd.AddCommand(xrayCmd)
	xrayCmd.AddCommand(xrayComponentDetailsCmd)
}

var xrayCmd = &cobra.Command{
	Use:   "xray",
	Short: "Retrieve XRAY security insights",
	Long: `This command allows you to interact with your Jfrog XRAY in order 
to retrieve security insights.`,
}

var xrayComponentDetailsCmd = &cobra.Command{
	Use:   "component-details [artifact-name]",
	Short: "Retrieves all the component details from XRAY",
	Long:  `This command allows to pass a component name and retrieve all the vulnerabilities and other insights from Jfrog XRAY.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ArtifactName = args[0]
		fmt.Printf("Running xray component-details %s\n", ArtifactName)
		xrayAPIURL := viper.GetString("XRAY_API_URL")
		xrayAPIUsername := viper.GetString("XRAY_USERNAME")
		xrayAPIPassword := viper.GetString("XRAY_PASSWORD")
		fmt.Println("calling xray API:", xrayAPIURL)

		body := map[string]string{
			"violations":                 "true",
			"include_ignored_violations": "false",
			"license":                    "false",
			"security":                   "true",
			"exclude_unknown":            "false",
			"package_type":               "docker",
			"sha_256":                    "72741d423b6777d2ad5ec77a9982b99d8568d82bae618ebf2b78bdcbff14e933",
			"component_name":             ArtifactName,
			"output_format":              "json",
		}
		client := &http.Client{}
		URL := xrayAPIURL + "/component/exportDetails"
		//pass the values to the request's body
		jsonBody, _ := json.Marshal(body)
		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonBody))
		req.SetBasicAuth(xrayAPIUsername, xrayAPIPassword)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bodyText))
	},
}

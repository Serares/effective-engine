/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Serares/effective-engine/scrape"
	"github.com/spf13/cobra"
)

// getAllCmd represents the getAll command
var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get's all the iframe videos from the given url",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		url, _ := cmd.Flags().GetString("url")
		dest, _ := cmd.Flags().GetString("dest")

		fmt.Println("getAll called")
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		fmt.Println("url: ", url)

		cfg := &scrape.ScrapeCfg{
			Username: username,
			Password: password,
			Url:      url,
			DestDir:  dest,
		}
		err := getAllVideos(os.Stdout, cfg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)

	// Here you will define your flags and configuration settings.

	getAllCmd.Flags().StringP("username", "u", "Default", "username used to authenticate to the given url path")
	getAllCmd.Flags().StringP("password", "p", "Default", "password used to authenticate to the given url path")
	getAllCmd.Flags().StringP("url", "l", "Default", "url used to scrape the iframe contents")
	getAllCmd.Flags().StringP("dest", "d", "./videos", "destination of the downloaded videos")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

func getAllVideos(out io.Writer, cfg *scrape.ScrapeCfg) error {

	err := scrape.Run(cfg)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, "Get all videos finished\n")
	return err
}

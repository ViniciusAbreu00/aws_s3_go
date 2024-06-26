package main

import (
	"aws_s3_go/bucket_service"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var pathname string

	rootCmd := &cobra.Command{
		Use:   "seiglu",
		Short: "Seiglu CLI is a tool for uploading and searching files in a aws bucket",
	}

	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "Uploads a file to a aws bucket",
		Run: func(cmd *cobra.Command, args []string) {
			if pathname == "" {
				log.Fatalf("pathname must be provided")
			} else {
				bucket_service.UploadFile(pathname)
			}

		},
	}

	uploadCmd.Flags().StringVarP(&pathname, "path", "p", "", "path to upload")
	uploadCmd.MarkFlagRequired("path")

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search for a file in the bucket",
		Run: func(cmd *cobra.Command, args []string) {
			data, error := bucket_service.GetBucketContent()
			if error != nil {
				log.Fatal(error)
			}
			for _, item := range data {
				log.Println(item)
			}
		},
	}

	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(searchCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}

package cmd

import (
	"fmt"
	"os"

	"github.com/dojo-engineering/markdown-to-confluence/internal/confluence"
	"github.com/dojo-engineering/markdown-to-confluence/internal/flags"
	"github.com/dojo-engineering/markdown-to-confluence/internal/node"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func initConfig() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFile:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	})
	// logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
}

func init() {
	cobra.OnInitialize(initConfig)
	addRequiredStringFlag(&flags.ConfluenceBaseURL, "url", "", "the url for confluence")
	addRequiredStringFlag(&flags.ConfluenceUsername, "username", "", "the confluence API username")
	addRequiredStringFlag(&flags.ConfluenceAPIKey, "key", "", "the confluence API key")
	addRequiredStringFlag(&flags.ConfluenceSpaceKey, "space", "", "the confluence space key")
	addRequiredStringFlag(&flags.SourceDocsPath, "docspath", "", "the source directory of the documentation")
	rootCmd.PersistentFlags().IntVar(&flags.RootPageID, "id", 0, "the id of the master page - default is 0 (root)")
	if err := rootCmd.MarkPersistentFlagRequired("id"); err != nil {
		logrus.Fatal(err)
	}
}

func addRequiredStringFlag(p *string, name, value string, usage string) {
	rootCmd.PersistentFlags().StringVar(p, name, value, usage)
	if err := rootCmd.MarkPersistentFlagRequired(name); err != nil {
		logrus.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "markdown-to-confluence",
	Short: "run markdown-to-confluence",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Markdown to Confluence starting...")

		flags.ConfluenceAPIBaseURL = fmt.Sprintf("%s/wiki/api/v2", flags.ConfluenceBaseURL)

		confluenceClientConfiguration := confluence.NewConfiguration()
		confluenceClientConfiguration.Servers = confluence.ServerConfigurations{
			confluence.ServerConfiguration{
				URL: flags.ConfluenceAPIBaseURL,
			},
		}
		confluenceClient := confluence.NewAPIClient(confluenceClientConfiguration)
		root := node.RootNode(confluenceClient)
		if root.Start(flags.RootPageID, flags.SourceDocsPath, flags.OnlyDocs) {
			root.Delete()
		}
		logrus.Info("Markdown to Confluence completed successfully!!!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Whoops. There was an error while executing your CLI")
		os.Exit(1)
	}
}

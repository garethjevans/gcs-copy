package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jenkins-x/gcs-copy/pkg/gcs-copy"
	"github.com/jenkins-x/gcs-copy/pkg/version"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var bucketName string

var rootCmd = &cobra.Command{
	Use:   "gcs-copy",
	Short: "Allows gcs bucket to be used as a chartmuseum repository.",
	Long:  `Creates a static index.yaml in gcs bucket root in order to allow GET requests directly from helm to the gcs bucket. Charts must be stored in ./charts for this to work.`,
	Run: func(cmd *cobra.Command, args []string) {
		gcsCopy.Run(viper.GetString("bucket-name"), viper.GetString("copy-from"), viper.GetString("copy-to"), viper.GetString("google-application-credentials"))
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %v\n", version.Version)
		fmt.Printf("Revision: %v\n", version.Revision)
		fmt.Printf("Branch: %v\n", version.Branch)
		fmt.Printf("BuildUser: %v\n", version.BuildUser)
		fmt.Printf("BuildDate: %v\n", version.BuildDate)
		fmt.Printf("GoVersion: %v\n", version.GoVersion)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chart-downloader.yaml)")
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().String("bucket-name", "", "The name of your bucket in gcs")
	rootCmd.PersistentFlags().String("copy-from", "", "The file to copy")
	rootCmd.PersistentFlags().String("copy-to", "", "The destination of the copy")
	rootCmd.PersistentFlags().String("google-application-credentials", "", "The file path to a JSON service file if you would like to use one.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gcs-copy")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	flags := rootCmd.PersistentFlags()
	viper.BindPFlag("bucket-name", flags.Lookup("bucket-name"))
	viper.BindPFlag("copy-from", flags.Lookup("copy-from"))
	viper.BindPFlag("copy-to", flags.Lookup("copy-to"))
	viper.BindPFlag("google-application-credentials", flags.Lookup("google-application-credentials"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

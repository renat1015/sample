package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sample/db"
	"sample/server"
	"sample/service"
	"strings"
)

const (
	cfgName = "sample"
)

var (
	cfgFile string
	cfg     = defaultConfig()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sample",
	Short: "Sample API",
	RunE:  run,
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sample.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".sample" (without extension).
		viper.SetConfigName(".sample")
		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
	}

	viper.SetEnvPrefix(cfgName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, v := range configValues(cfg) {
		viper.BindEnv(v)
	}

	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run(cmd *cobra.Command, args []string) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := cfg.validate(); err != nil {
		return err
	}

	db, err := db.New(cfg.DBConfig)
	if err != nil {
		return err
	}

	s := service.New(db)

	if err := server.NewServer(cfg.ServerConfig, s); err != nil {
		return err
	}

	return nil
}

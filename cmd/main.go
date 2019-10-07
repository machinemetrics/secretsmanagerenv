package cmd

import (
	"errors"
	"fmt"
	shlex "github.com/anmitsu/go-shlex"
	"github.com/machinemetrics/secretsmanagerenv/cmd/handler"
	"github.com/spf13/cobra"
	"os"
)

var (
	secrets []string
	region  string
	upcase  bool
	prefix  string
)

var rootCmd = &cobra.Command{
	Use:   "smenv",
	Short: "B",
	Long:  "C",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires at least one arg")
		}
		if len(secrets) == 0 {
			return errors.New("Must specify secret with `-s`")
		}
		return nil
	},
	Run: func(_ *cobra.Command, args []string) {
		if command, err := parse(args); err != nil {
			fmt.Println(err.Error())
		} else if err := handler.RunCommandWithSecret(secrets, region, command, upcase, prefix); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parse(args []string) ([]string, error) {
	var ret []string
	for _, arg := range args {
		if words, err := shlex.Split(arg, true); err != nil {
			return nil, err
		} else {
			ret = append(ret, words...)
		}
	}
	return ret, nil
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&secrets, "secret", "s", []string{}, "name of secret")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "region")
	rootCmd.PersistentFlags().BoolVarP(&upcase, "upcase", "u", false, "Upcase environment variables")
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Prepend PREFIX to environment variables")
}

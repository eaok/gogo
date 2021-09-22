/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var echoTimes int
var flagBool bool
var flagString string

//var toggle bool
//var aaa string
//
//= dosCmd.Flags().GetBool("toggle")
//aaa := dosCmd.Flags().GetString("aaa")

// dosCmd represents the dos command
var dosCmd = &cobra.Command{
	Use:   "dos",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dos called")

		if echoTimes != 1 {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		}

		if flagBool != false {
			fmt.Println(flagBool)
		}
		if flagString != "" {
			fmt.Println(flagString)
		}
	},
}

func init() {
	rootCmd.AddCommand(dosCmd)

	//dos -t 3 hello
	dosCmd.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	//dos -o true
	dosCmd.Flags().BoolVarP(&flagBool, "toggle", "o", false, "Help message for toggle")

	//dos -a hello
	dosCmd.Flags().StringVarP(&flagString, "aaa", "a", "", "test")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

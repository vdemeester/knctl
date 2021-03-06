/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Open 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cobrautil

import (
	"github.com/spf13/cobra"
)

func VisitCommands(cmd *cobra.Command, f func(*cobra.Command)) {
	f(cmd)
	for _, child := range cmd.Commands() {
		VisitCommands(child, f)
	}
}

func WrapRunEForCmd(additionalRunE func(*cobra.Command, []string) error) func(cmd *cobra.Command) {
	return func(cmd *cobra.Command) {
		origRunE := cmd.RunE
		cmd.RunE = func(cmd2 *cobra.Command, args []string) error {
			err := additionalRunE(cmd2, args)
			if err != nil {
				return err
			}
			return origRunE(cmd2, args)
		}
	}
}

// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.acuvity.ai/regolithe/cmd/rego/doc"
	"go.acuvity.ai/regolithe/cmd/rego/jsonschema"
	"go.acuvity.ai/regolithe/cmd/rego/specset"
	"go.acuvity.ai/regolithe/spec"
)

const (
	name        = "rego"
	description = "Tool to manipulate regolithe specifications"
)

func main() {

	cobra.OnInitialize(func() {
		viper.SetEnvPrefix(name)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	})

	var rootCmd = &cobra.Command{
		Use:   name,
		Short: description,
	}

	var formatCmd = &cobra.Command{
		Use:           "format",
		Short:         "Reads a specification from stdin and prints it formatted on std out.",
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(_ *cobra.Command, _ []string) error {

			switch viper.Get("mode") {
			case "spec":
				s := spec.NewSpecification()

				if err := s.Read(os.Stdin, true); err != nil {
					return fmt.Errorf("unable to format: unable to read spec: %s", err)
				}

				if err := s.Write(os.Stdout); err != nil {
					return fmt.Errorf("unable to format: unable to write spec: %s", err)
				}

			case "typemapping":
				tm := spec.NewTypeMapping()

				if err := tm.Read(os.Stdin, true); err != nil {
					return fmt.Errorf("unable to format: unable to read typemapping: %s", err)
				}

				if err := tm.Write(os.Stdout); err != nil {
					return fmt.Errorf("unable to format: unable to write typemapping: %s", err)
				}

			case "validationmapping":
				vm := spec.NewValidationMapping()

				if err := vm.Read(os.Stdin, true); err != nil {
					return fmt.Errorf("unable to format: unable to read validationmapping: %s", err)
				}

				if err := vm.Write(os.Stdout); err != nil {
					return fmt.Errorf("unable to format: unable to write validationmapping: %s", err)
				}

			case "parametermapping":
				pm := spec.NewParameterMapping()

				if err := pm.Read(os.Stdin, true); err != nil {
					return fmt.Errorf("unable to format: unable to read parametermapping: %s", err)
				}

				if err := pm.Write(os.Stdout); err != nil {
					return fmt.Errorf("unable to format: unable to write parametermapping: %s", err)
				}
			}

			return nil
		},
	}
	formatCmd.Flags().StringP("mode", "m", "spec", "Mode of formatting. Can be spec, typemapping, validationmapping, parametermapping.")

	var docCmd = &cobra.Command{
		Use:           "doc",
		Short:         "Generate a documentation for the given specification set",
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(_ *cobra.Command, _ []string) error {

			s, err := spec.LoadSpecificationSet(
				viper.GetString("dir"),
				nil,
				nil,
				viper.GetString("category"),
			)
			if err != nil {
				return fmt.Errorf("unable to load specification set: %s", err)
			}

			if err := doc.Write(s, viper.GetString("format")); err != nil {
				return fmt.Errorf("unable to write specification set: %s", err)
			}

			return nil
		},
	}
	docCmd.Flags().StringP("dir", "d", "", "Path of the specifications folder.")
	docCmd.Flags().String("format", "markdown", "Path of the specifications folder.")

	var jsonSchemaCmd = &cobra.Command{
		Use:           "jsonschema",
		Short:         "Generate a json schema out of a specification set",
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(_ *cobra.Command, _ []string) error {

			s, err := spec.LoadSpecificationSet(
				viper.GetString("dir"),
				nil,
				nil,
				"jsonschema",
			)
			if err != nil {
				return err
			}

			return jsonschema.Generate(s, viper.GetString("out"), viper.GetBool(("public")))
		},
	}
	jsonSchemaCmd.Flags().StringP("dir", "d", "", "Path of the specifications folder.")
	jsonSchemaCmd.Flags().StringP("out", "o", "./codegen", "Path where to write the json files.")
	jsonSchemaCmd.Flags().BoolP("public", "p", false, "If set to true, only exposed attributes and public objects will be generated.")

	var initCmd = &cobra.Command{
		Use:           "init <dest>",
		Short:         "Generate a new set of specification",
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(_ *cobra.Command, args []string) error {

			if len(args) != 1 {
				return fmt.Errorf("usage: init <dest>")
			}

			dir := args[0]

			return specset.Dump(dir)
		},
	}

	rootCmd.AddCommand(
		formatCmd,
		docCmd,
		initCmd,
		jsonSchemaCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err.Error()) // nolint: errcheck
		os.Exit(1)
	}
}

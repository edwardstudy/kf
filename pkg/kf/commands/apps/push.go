// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apps

import (
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/GoogleCloudPlatform/kf/pkg/kf"
	"github.com/GoogleCloudPlatform/kf/pkg/kf/commands/config"
	kfi "github.com/GoogleCloudPlatform/kf/pkg/kf/internal/kf"
	"github.com/GoogleCloudPlatform/kf/pkg/kf/manifest"
	"github.com/spf13/cobra"
)

// SrcImageBuilder creates and uploads a container image that contains the
// contents of the argument 'dir'.
type SrcImageBuilder interface {
	BuildSrcImage(dir, srcImage string) error
}

// SrcImageBuilderFunc converts a func into a SrcImageBuilder.
type SrcImageBuilderFunc func(dir, srcImage string) error

// BuildSrcImage implements SrcImageBuilder.
func (f SrcImageBuilderFunc) BuildSrcImage(dir, srcImage string) error {
	return f(dir, srcImage)
}

// NewPushCommand creates a push command.
func NewPushCommand(p *config.KfParams, pusher kf.Pusher, b SrcImageBuilder) *cobra.Command {
	var (
		containerRegistry string
		sourceImage       string
		manifestFile      string
		serviceAccount    string
		path              string
		buildpack         string
		envs              []string
		grpc              bool
	)

	var pushCmd = &cobra.Command{
		Use:   "push APP_NAME [--container-registry CONTAINER_REGISTRY]",
		Short: "Push a new app or sync changes to an existing app",
		Example: `
  kf push myapp --container-registry gcr.io/myproject
  kf push myapp --container-registry gcr.io/myproject --buildpack my.special.buildpack # Discover via kf buildpacks
  kf push myapp --container-registry gcr.io/myproject --env FOO=bar --env BAZ=foo
  `,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if containerRegistry == "" {
				return errors.New("container-registry is required")
			}
			cmd.SilenceUsage = true

			appName := ""
			if len(args) > 0 {
				appName = args[0]
			}

			// Kontext has to have a absolute path.
			var err error
			path, err = filepath.Abs(path)
			if err != nil {
				return err
			}

			var pushManifest *Manifest
			if manifestFile != "" {
				if pushManifest, err = manifest.NewFromFile(manifestFile); err != nil {
					return fmt.Errorf("supplied manifest file %s resulted in error: %v", manifestFile, err)
				}
			} else if pushManifest = manifest.CheckForManifest(path); pushManifest == nil {
				return fmt.Errorf("no manifest file found in directory %s", path)
			}

			appsToDeploy := pushManifest.Applications
			if appName != "" {
				// deploy one app from the manifest
				app, err := pushManifest.App(appName)
				if err != nil {
					return err
				}

				appsToDeploy = []manifest.Application{*app}
			}

			for _, app := range appsToDeploy {
				var imageName string

				srcPath := filepath.Join(path, app.Path)
				switch {
				case sourceImage != "":
					imageName = sourceImage
				default:
					imageName = fmt.Sprintf("%s/src-%s-%d%d", containerRegistry, app.Name, time.Now().UnixNano(), rand.Uint64())

					if err := b.BuildSrcImage(srcPath, imageName); err != nil {
						return err
					}
				}

				err := pusher.Push(app.Name, imageName,
					kf.WithPushNamespace(p.Namespace),
					kf.WithPushContainerRegistry(containerRegistry),
					kf.WithPushServiceAccount(serviceAccount),
					kf.WithPushEnvironmentVariables(envs),
					kf.WithPushGrpc(grpc),
					kf.WithPushBuildpack(buildpack),
				)

				cmd.SilenceUsage = !kfi.ConfigError(err)

				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	pushCmd.Flags().StringVar(
		&containerRegistry,
		"container-registry",
		"",
		"The container registry to push containers (REQUIRED).",
	)

	pushCmd.Flags().StringVar(
		&serviceAccount,
		"service-account",
		"",
		"The service account to enable access to the container registry",
	)

	pushCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		".",
		"The path the source code lives. Defaults to current directory.",
	)

	pushCmd.Flags().StringArrayVarP(
		&envs,
		"env",
		"e",
		nil,
		"Set environment variables. Multiple can be set by using the flag multiple times (e.g., NAME=VALUE).",
	)

	pushCmd.Flags().BoolVar(
		&grpc,
		"grpc",
		false,
		"Setup the container to allow application to use gRPC.",
	)

	pushCmd.Flags().StringVarP(
		&buildpack,
		"buildpack",
		"b",
		"",
		"Skip the 'detect' buildpack step and use the given name.",
	)

	pushCmd.Flags().StringVar(
		&sourceImage,
		"source-image",
		"",
		"The kontext image that has the source code.",
	)
	pushCmd.Flags().MarkHidden("source-image")

	pushCmd.Flags().StringVarP(
		&manifestFile,
		"manifest",
		"f",
		"manifest.yml",
		"Path to manifest",
	)
	pushCmd.Flags().MarkHidden("source-image")

	return pushCmd
}

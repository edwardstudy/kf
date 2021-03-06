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

package spaces

import (
	"bytes"
	"errors"
	"testing"

	"github.com/GoogleCloudPlatform/kf/pkg/kf/commands/config"
	"github.com/GoogleCloudPlatform/kf/pkg/kf/spaces/fake"
	"github.com/GoogleCloudPlatform/kf/pkg/kf/testutil"
	"github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

func TestNewListSpacesCommand(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		args  []string
		setup func(t *testing.T, fakeSpaces *fake.FakeClient)

		wantErr         error
		expectedStrings []string
	}{
		"invalid number of args": {
			args:    []string{"asdf"},
			wantErr: errors.New("accepts 0 arg(s), received 1"),
		},
		"no contents": {
			setup: func(t *testing.T, fakeSpaces *fake.FakeClient) {
				list := []v1.Namespace{}
				fakeSpaces.
					EXPECT().
					List().
					Return(list, nil)
			},
			expectedStrings: []string{"Name", "Status", "Age"},
		},
		"contents": {
			setup: func(t *testing.T, fakeSpaces *fake.FakeClient) {
				ns := v1.Namespace{}
				ns.Name = "my-ns"
				ns.Status.Phase = "TESTING"

				list := []v1.Namespace{ns}
				fakeSpaces.
					EXPECT().
					List().
					Return(list, nil)
			},
			expectedStrings: []string{"my-ns", "TESTING"},
		},
		"server failure": {
			setup: func(t *testing.T, fakeSpaces *fake.FakeClient) {
				fakeSpaces.
					EXPECT().
					List().
					Return(nil, errors.New("some-server-error"))
			},
			wantErr: errors.New("some-server-error"),
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			fakeSpaces := fake.NewFakeClient(ctrl)

			if tc.setup != nil {
				tc.setup(t, fakeSpaces)
			}

			buffer := &bytes.Buffer{}

			c := NewListSpacesCommand(&config.KfParams{Namespace: "default", Output: buffer}, fakeSpaces)
			c.SetOutput(buffer)
			c.SetArgs(tc.args)

			gotErr := c.Execute()
			testutil.AssertErrorsEqual(t, tc.wantErr, gotErr)
			testutil.AssertContainsAll(t, buffer.String(), tc.expectedStrings)

			ctrl.Finish()
		})
	}
}

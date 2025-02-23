/*
Copyright 2020 The Kubernetes Authors.

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

package fileutil

import (
	"os"
	"path/filepath"
	"testing"

	"sigs.k8s.io/secrets-store-csi-driver/pkg/test_utils/tmpdir"
)

func TestGetMountedFiles(t *testing.T) {
	tests := []struct {
		name        string
		targetPath  func(t *testing.T) string
		expectedErr bool
		expectedKey string
	}{
		{
			name:        "target path not found",
			targetPath:  func(t *testing.T) string { return "" },
			expectedErr: true,
			expectedKey: "",
		},
		{
			name: "target path dir found",
			targetPath: func(t *testing.T) string {
				return tmpdir.New(t, "", "ut")
			},
			expectedErr: false,
			expectedKey: "",
		},
		{
			name: "target path dir/file found",
			targetPath: func(t *testing.T) string {
				dir := tmpdir.New(t, "", "ut")
				os.Create(filepath.Join(dir, "secret.txt"))
				return dir
			},
			expectedErr: false,
			expectedKey: "secret.txt",
		},
		{
			name: "target path dir/dir/file found",
			targetPath: func(t *testing.T) string {
				dir := tmpdir.New(t, "", "ut")
				os.MkdirAll(filepath.Join(dir, "subdir"), 0700)
				os.Create(filepath.Join(dir, "subdir", "secret.txt"))
				return dir
			},
			expectedErr: false,
			expectedKey: "subdir/secret.txt",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetMountedFiles(test.targetPath(t))
			if test.expectedErr != (err != nil) {
				t.Fatalf("expected err: %v, got: %+v", test.expectedErr, err)
			}
			if !test.expectedErr && test.expectedKey != "" {
				if _, ok := got[test.expectedKey]; !ok {
					t.Fatalf("expected key not found: %s", test.expectedKey)
				}
			}
		})
	}
}

func TestGetPodUIDFromTargetPath(t *testing.T) {
	cases := []struct {
		targetPath string
		want       string
	}{
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~csi",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~pv/pvvol/mount",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "7e7686a1-56c4-4c67-a6fd-4656ac484f0a",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes`,
			want:       "",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~csi`,
			want:       "",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~pv\pvvol\mount`,
			want:       "",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~csi\secrets-store-inline\mount`,
			want:       "d4fd876f-bdb3-11e9-a369-0a5d188d99c0",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes`,
			want:       "",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~csi`,
			want:       "",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~pv\\pvvol\\mount`,
			want:       "",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~csi\\secrets-store-inline\\mount`,
			want:       "d4fd876f-bdb3-11e9-a369-0a5d188d9934",
		},
		{
			targetPath: "/var/lib/",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods",
			want:       "",
		},
		{
			targetPath: "/opt/new/var/lib/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "456457fc-d980-4191-b5eb-daf70c4ff7c1",
		},
		{
			targetPath: "data/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "456457fc-d980-4191-b5eb-daf70c4ff7c1",
		},
		{
			targetPath: "data/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~pv/secrets-store-inline/mount",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "7e7686a1-56c4-4c67-a6fd-4656ac484f0a",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~csi\secrets-store-inline\mount`,
			want:       "d4fd876f-bdb3-11e9-a369-0a5d188d99c0",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~csi\\secrets-store-inline\\mount`,
			want:       "d4fd876f-bdb3-11e9-a369-0a5d188d9934",
		},
		{
			targetPath: "/opt/new/var/lib/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "456457fc-d980-4191-b5eb-daf70c4ff7c1",
		},
		{
			targetPath: "data/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "456457fc-d980-4191-b5eb-daf70c4ff7c1",
		},
		{
			targetPath: "/var/lib/kubelet/pods/64f9ffb2-409e-4c58-9ea8-2a7d21050ece/volumes/kubernetes.io~secret/server-token-npdwt",
			want:       "",
		},
		{
			targetPath: `\\pods\\fakePod\\volumes\\kubernetes.io~csi\\myvol\\mount`,
			want:       "fakePod",
		},
	}

	for _, tc := range cases {
		got := GetPodUIDFromTargetPath(tc.targetPath)
		if got != tc.want {
			t.Errorf("GetPodUIDFromTargetPath(%v) = %v, want %v", tc.targetPath, got, tc.want)
		}
	}
}

func TestGetVolumeNameFromTargetPath(t *testing.T) {
	cases := []struct {
		targetPath string
		want       string
	}{
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~csi",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~pv/pvvol/mount",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods/7e7686a1-56c4-4c67-a6fd-4656ac484f0a/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "secrets-store-inline",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes`,
			want:       "",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~csi`,
			want:       "",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~pv\pvvol\mount`,
			want:       "",
		},
		{
			targetPath: `c:\var\lib\kubelet\pods\d4fd876f-bdb3-11e9-a369-0a5d188d99c0\volumes\kubernetes.io~csi\secrets-store-inline\mount`,
			want:       "secrets-store-inline",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes`,
			want:       "",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~csi`,
			want:       "",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~pv\\pvvol\\mount`,
			want:       "",
		},
		{
			targetPath: `c:\\var\\lib\\kubelet\\pods\\d4fd876f-bdb3-11e9-a369-0a5d188d9934\\volumes\\kubernetes.io~csi\\secrets-store-inline\\mount`,
			want:       "secrets-store-inline",
		},
		{
			targetPath: "/var/lib/",
			want:       "",
		},
		{
			targetPath: "/var/lib/kubelet/pods",
			want:       "",
		},
		{
			targetPath: "/opt/new/var/lib/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "secrets-store-inline",
		},
		{
			targetPath: "data/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~csi/secrets-store-inline/mount",
			want:       "secrets-store-inline",
		},
		{
			targetPath: "data/kubelet/pods/456457fc-d980-4191-b5eb-daf70c4ff7c1/volumes/kubernetes.io~pv/secrets-store-inline/mount",
			want:       "",
		},
	}

	for _, tc := range cases {
		got := GetVolumeNameFromTargetPath(tc.targetPath)
		if got != tc.want {
			t.Errorf("GetVolumeNameFromTargetPath(%v) = %v, want %v", tc.targetPath, got, tc.want)
		}
	}
}

// Copyright 2020 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.)

package v0_6

import (
	v0_5 "github.com/coreos/butane/base/v0_5"
)

// Type aliases for types unchanged from v0_5 that have NO methods
type Clevis = v0_5.Clevis
type ClevisCustom = v0_5.ClevisCustom
type Device = v0_5.Device
type Disk = v0_5.Disk
type Group = v0_5.Group
type HTTPHeader = v0_5.HTTPHeader
type HTTPHeaders = v0_5.HTTPHeaders
type Ignition = v0_5.Ignition
type IgnitionConfig = v0_5.IgnitionConfig
type KernelArgument = v0_5.KernelArgument
type KernelArguments = v0_5.KernelArguments
type Link = v0_5.Link
type NodeGroup = v0_5.NodeGroup
type NodeUser = v0_5.NodeUser
type Partition = v0_5.Partition
type Passwd = v0_5.Passwd
type PasswdGroup = v0_5.PasswdGroup
type PasswdUser = v0_5.PasswdUser
type Proxy = v0_5.Proxy
type Raid = v0_5.Raid
type SSHAuthorizedKey = v0_5.SSHAuthorizedKey
type Security = v0_5.Security
type Systemd = v0_5.Systemd
type Tang = v0_5.Tang
type TLS = v0_5.TLS
type Timeouts = v0_5.Timeouts
type Verification = v0_5.Verification

// Types that have methods must be defined (not aliased) even if identical to previous version
type Config struct {
	Version         string          `yaml:"version"`
	Variant         string          `yaml:"variant"`
	Ignition        Ignition        `yaml:"ignition"`
	KernelArguments KernelArguments `yaml:"kernel_arguments"`
	Passwd          Passwd          `yaml:"passwd"`
	Storage         Storage         `yaml:"storage"`
	Systemd         Systemd         `yaml:"systemd"`
}

type Directory struct {
	Group     NodeGroup `yaml:"group"`
	Overwrite *bool     `yaml:"overwrite"`
	Path      string    `yaml:"path"`
	User      NodeUser  `yaml:"user"`
	Mode      *int      `yaml:"mode"`
}

type Dropin struct {
	Contents      *string `yaml:"contents"`
	ContentsLocal *string `yaml:"contents_local"`
	Name          string  `yaml:"name"`
}

type File struct {
	Group     NodeGroup  `yaml:"group"`
	Overwrite *bool      `yaml:"overwrite"`
	Path      string     `yaml:"path"`
	User      NodeUser   `yaml:"user"`
	Append    []Resource `yaml:"append"`
	Contents  Resource   `yaml:"contents"`
	Mode      *int       `yaml:"mode"`
}

type Filesystem struct {
	Device         string   `yaml:"device"`
	Format         *string  `yaml:"format"`
	Label          *string  `yaml:"label"`
	MountOptions   []string `yaml:"mount_options"`
	Options        []string `yaml:"options"`
	Path           *string  `yaml:"path"`
	UUID           *string  `yaml:"uuid"`
	WipeFilesystem *bool    `yaml:"wipe_filesystem"`
	WithMountUnit  *bool    `yaml:"with_mount_unit" butane:"auto_skip"` // Added, not in Ignition spec
}

type Resource struct {
	Compression  *string      `yaml:"compression"`
	HTTPHeaders  HTTPHeaders  `yaml:"http_headers"`
	Source       *string      `yaml:"source"`
	Inline       *string      `yaml:"inline"` // Added, not in ignition spec
	Local        *string      `yaml:"local"`  // Added, not in ignition spec
	Verification Verification `yaml:"verification"`
}

type Storage struct {
	Directories []Directory  `yaml:"directories"`
	Disks       []Disk       `yaml:"disks"`
	Files       []File       `yaml:"files"`
	Filesystems []Filesystem `yaml:"filesystems"`
	Links       []Link       `yaml:"links"`
	Luks        []Luks       `yaml:"luks"`
	Raid        []Raid       `yaml:"raid"`
	Trees       []Tree       `yaml:"trees" butane:"auto_skip"` // Added, not in ignition spec
}

type Tree struct {
	Local string  `yaml:"local"`
	Path  *string `yaml:"path"`
}

type Unit struct {
	Contents      *string  `yaml:"contents"`
	ContentsLocal *string  `yaml:"contents_local"`
	Dropins       []Dropin `yaml:"dropins"`
	Enabled       *bool    `yaml:"enabled"`
	Mask          *bool    `yaml:"mask"`
	Name          string   `yaml:"name"`
}

// New type in v0_6
type Cex struct {
	Enabled *bool `yaml:"enabled"`
}

// Luks was modified in v0_6 to add Cex field
type Luks struct {
	Cex         Cex      `yaml:"cex"`
	Clevis      Clevis   `yaml:"clevis"`
	Device      *string  `yaml:"device"`
	Discard     *bool    `yaml:"discard"`
	KeyFile     Resource `yaml:"key_file"`
	Label       *string  `yaml:"label"`
	Name        string   `yaml:"name"`
	OpenOptions []string `yaml:"open_options"`
	Options     []string `yaml:"options"`
	UUID        *string  `yaml:"uuid"`
	WipeVolume  *bool    `yaml:"wipe_volume"`
}

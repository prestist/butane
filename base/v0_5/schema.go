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

package v0_5

import (
	v0_4 "github.com/coreos/butane/base/v0_4"
)

// Type aliases for types unchanged from v0_4 that have NO methods
type Clevis = v0_4.Clevis
type ClevisCustom = v0_4.ClevisCustom
type Device = v0_4.Device
type Disk = v0_4.Disk
type Group = v0_4.Group
type HTTPHeader = v0_4.HTTPHeader
type HTTPHeaders = v0_4.HTTPHeaders
type Ignition = v0_4.Ignition
type IgnitionConfig = v0_4.IgnitionConfig
type KernelArgument = v0_4.KernelArgument
type KernelArguments = v0_4.KernelArguments
type Link = v0_4.Link
type NodeGroup = v0_4.NodeGroup
type NodeUser = v0_4.NodeUser
type Partition = v0_4.Partition
type Passwd = v0_4.Passwd
type PasswdGroup = v0_4.PasswdGroup
type Proxy = v0_4.Proxy
type SSHAuthorizedKey = v0_4.SSHAuthorizedKey
type Security = v0_4.Security
type Systemd = v0_4.Systemd
type TLS = v0_4.TLS
type Timeouts = v0_4.Timeouts
type Verification = v0_4.Verification

// Types that have methods must be defined (not aliased)
type Config struct {
	Version         string          `yaml:"version"`
	Variant         string          `yaml:"variant"`
	Ignition        Ignition        `yaml:"ignition"`
	KernelArguments KernelArguments `yaml:"kernel_arguments"`
	Passwd          Passwd          `yaml:"passwd"`
	Storage         Storage         `yaml:"storage"`
	Systemd         Systemd         `yaml:"systemd"`
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

type Resource struct {
	Compression  *string      `yaml:"compression"`
	HTTPHeaders  HTTPHeaders  `yaml:"http_headers"`
	Source       *string      `yaml:"source"`
	Inline       *string      `yaml:"inline"` // Added, not in ignition spec
	Local        *string      `yaml:"local"`  // Added, not in ignition spec
	Verification Verification `yaml:"verification"`
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

// Modified types in v0_5
type Luks struct {
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

type PasswdUser struct {
	Gecos                  *string            `yaml:"gecos"`
	Groups                 []Group            `yaml:"groups"`
	HomeDir                *string            `yaml:"home_dir"`
	Name                   string             `yaml:"name"`
	NoCreateHome           *bool              `yaml:"no_create_home"`
	NoLogInit              *bool              `yaml:"no_log_init"`
	NoUserGroup            *bool              `yaml:"no_user_group"`
	PasswordHash           *string            `yaml:"password_hash"`
	PrimaryGroup           *string            `yaml:"primary_group"`
	ShouldExist            *bool              `yaml:"should_exist"`
	SSHAuthorizedKeys      []SSHAuthorizedKey `yaml:"ssh_authorized_keys"`
	SSHAuthorizedKeysLocal []string           `yaml:"ssh_authorized_keys_local"`
	Shell                  *string            `yaml:"shell"`
	System                 *bool              `yaml:"system"`
	UID                    *int               `yaml:"uid"`
}

type Raid struct {
	Devices []Device `yaml:"devices"`
	Level   *string  `yaml:"level"`
	Name    string   `yaml:"name"`
	Options []string `yaml:"options"`
	Spares  *int     `yaml:"spares"`
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

type Tang struct {
	Thumbprint    *string `yaml:"thumbprint"`
	URL           string  `yaml:"url"`
	Advertisement *string `yaml:"advertisement"`
}

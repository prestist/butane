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

package v0_4

import (
	v0_3 "github.com/coreos/butane/base/v0_3"
)

// Type aliases for types unchanged from v0_3 that have NO methods
type Device = v0_3.Device
type Disk = v0_3.Disk
type Dropin = v0_3.Dropin
type FilesystemOption = v0_3.FilesystemOption
type Group = v0_3.Group
type HTTPHeader = v0_3.HTTPHeader
type HTTPHeaders = v0_3.HTTPHeaders
type Ignition = v0_3.Ignition
type IgnitionConfig = v0_3.IgnitionConfig
type NodeGroup = v0_3.NodeGroup
type NodeUser = v0_3.NodeUser
type Partition = v0_3.Partition
type Passwd = v0_3.Passwd
type PasswdGroup = v0_3.PasswdGroup
type PasswdUser = v0_3.PasswdUser
type Proxy = v0_3.Proxy
type SSHAuthorizedKey = v0_3.SSHAuthorizedKey
type Security = v0_3.Security
type Systemd = v0_3.Systemd
type Tang = v0_3.Tang
type TLS = v0_3.TLS
type Timeouts = v0_3.Timeouts
type Unit = v0_3.Unit
type Verification = v0_3.Verification

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

type Directory struct {
	Group     NodeGroup `yaml:"group"`
	Overwrite *bool     `yaml:"overwrite"`
	Path      string    `yaml:"path"`
	User      NodeUser  `yaml:"user"`
	Mode      *int      `yaml:"mode"`
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

type Tree struct {
	Local string  `yaml:"local"`
	Path  *string `yaml:"path"`
}

// New types in v0_4
type ClevisCustom struct {
	Config       *string `yaml:"config"`
	NeedsNetwork *bool   `yaml:"needs_network"`
	Pin          *string `yaml:"pin"`
}

type KernelArgument string

type KernelArguments struct {
	ShouldExist    []KernelArgument `yaml:"should_exist"`
	ShouldNotExist []KernelArgument `yaml:"should_not_exist"`
}

type LuksOption string

type RaidOption string

// Modified types in v0_4
type Clevis struct {
	Custom    ClevisCustom `yaml:"custom"`
	Tang      []Tang       `yaml:"tang"`
	Threshold *int         `yaml:"threshold"`
	Tpm2      *bool        `yaml:"tpm2"`
}

type Link struct {
	Group     NodeGroup `yaml:"group"`
	Overwrite *bool     `yaml:"overwrite"`
	Path      string    `yaml:"path"`
	User      NodeUser  `yaml:"user"`
	Hard      *bool     `yaml:"hard"`
	Target    *string   `yaml:"target"`
}

type Luks struct {
	Clevis     Clevis       `yaml:"clevis"`
	Device     *string      `yaml:"device"`
	KeyFile    Resource     `yaml:"key_file"`
	Label      *string      `yaml:"label"`
	Name       string       `yaml:"name"`
	Options    []LuksOption `yaml:"options"`
	UUID       *string      `yaml:"uuid"`
	WipeVolume *bool        `yaml:"wipe_volume"`
}

type Raid struct {
	Devices []Device     `yaml:"devices"`
	Level   *string      `yaml:"level"`
	Name    string       `yaml:"name"`
	Options []RaidOption `yaml:"options"`
	Spares  *int         `yaml:"spares"`
}

// Storage must be redefined to use local Filesystem type
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

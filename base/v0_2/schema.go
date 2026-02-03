// Copyright 2019 Red Hat, Inc
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

package v0_2

import (
	v0_1 "github.com/coreos/butane/base/v0_1"
)

// Type aliases for types unchanged from v0_1 that have NO methods
type Device = v0_1.Device
type Disk = v0_1.Disk
type Dropin = v0_1.Dropin
type FilesystemOption = v0_1.FilesystemOption
type Group = v0_1.Group
type NodeGroup = v0_1.NodeGroup
type NodeUser = v0_1.NodeUser
type Partition = v0_1.Partition
type Passwd = v0_1.Passwd
type PasswdGroup = v0_1.PasswdGroup
type PasswdUser = v0_1.PasswdUser
type Raid = v0_1.Raid
type RaidOption = v0_1.RaidOption
type SSHAuthorizedKey = v0_1.SSHAuthorizedKey
type Systemd = v0_1.Systemd
type Timeouts = v0_1.Timeouts
type Unit = v0_1.Unit
type Verification = v0_1.Verification

// Config must be defined here (not aliased) because it has methods
type Config struct {
	Version  string   `yaml:"version"`
	Variant  string   `yaml:"variant"`
	Ignition Ignition `yaml:"ignition"`
	Passwd   Passwd   `yaml:"passwd"`
	Storage  Storage  `yaml:"storage"`
	Systemd  Systemd  `yaml:"systemd"`
}

// Directory must be defined here (not aliased) because it has methods
type Directory struct {
	Group     NodeGroup `yaml:"group"`
	Overwrite *bool     `yaml:"overwrite"`
	Path      string    `yaml:"path"`
	User      NodeUser  `yaml:"user"`
	Mode      *int      `yaml:"mode"`
}

// Link must be defined here (not aliased) because it has methods
type Link struct {
	Group     NodeGroup `yaml:"group"`
	Overwrite *bool     `yaml:"overwrite"`
	Path      string    `yaml:"path"`
	User      NodeUser  `yaml:"user"`
	Hard      *bool     `yaml:"hard"`
	Target    string    `yaml:"target"`
}

// New types in v0_2
type HTTPHeader struct {
	Name  string  `yaml:"name"`
	Value *string `yaml:"value"`
}

type HTTPHeaders []HTTPHeader

type Proxy struct {
	HTTPProxy  *string  `yaml:"http_proxy"`
	HTTPSProxy *string  `yaml:"https_proxy"`
	NoProxy    []string `yaml:"no_proxy"`
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

// Modified types in v0_2
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

type Ignition struct {
	Config   IgnitionConfig `yaml:"config"`
	Proxy    Proxy          `yaml:"proxy"`
	Security Security       `yaml:"security"`
	Timeouts Timeouts       `yaml:"timeouts"`
}

type IgnitionConfig struct {
	Merge   []Resource `yaml:"merge"`
	Replace Resource   `yaml:"replace"`
}

type Security struct {
	TLS TLS `yaml:"tls"`
}

type Storage struct {
	Directories []Directory  `yaml:"directories"`
	Disks       []Disk       `yaml:"disks"`
	Files       []File       `yaml:"files"`
	Filesystems []Filesystem `yaml:"filesystems"`
	Links       []Link       `yaml:"links"`
	Raid        []Raid       `yaml:"raid"`
	Trees       []Tree       `yaml:"trees" butane:"auto_skip"` // Added, not in ignition spec
}

type TLS struct {
	CertificateAuthorities []Resource `yaml:"certificate_authorities"`
}

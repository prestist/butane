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

package v0_3

import (
	v0_2 "github.com/coreos/butane/base/v0_2"
)

// Type aliases for types unchanged from v0_2 that have NO methods
type Device = v0_2.Device
// Disk is redefined below because it contains Partition which was modified in v0_3
type Dropin = v0_2.Dropin
type FilesystemOption = v0_2.FilesystemOption
type Group = v0_2.Group
type HTTPHeader = v0_2.HTTPHeader
type HTTPHeaders = v0_2.HTTPHeaders
type Ignition = v0_2.Ignition
type IgnitionConfig = v0_2.IgnitionConfig
// Link is redefined below because it has methods
type NodeGroup = v0_2.NodeGroup
type NodeUser = v0_2.NodeUser
// Passwd is redefined below because it contains PasswdGroup and PasswdUser which were modified in v0_3
type Passwd struct {
	Groups []PasswdGroup `yaml:"groups"`
	Users  []PasswdUser  `yaml:"users"`
}

// Disk is redefined because it contains Partition which was modified in v0_3
type Disk struct {
	Device     string      `yaml:"device"`
	Partitions []Partition `yaml:"partitions"`
	WipeTable  *bool       `yaml:"wipe_table"`
}

type Proxy = v0_2.Proxy
type Raid = v0_2.Raid
type RaidOption = v0_2.RaidOption
type SSHAuthorizedKey = v0_2.SSHAuthorizedKey
type Security = v0_2.Security
type Systemd = v0_2.Systemd
type TLS = v0_2.TLS
type Timeouts = v0_2.Timeouts
type Unit = v0_2.Unit
type Verification = v0_2.Verification

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

// File must be defined here (not aliased) because it has methods
type File struct {
	Group     NodeGroup  `yaml:"group"`
	Overwrite *bool      `yaml:"overwrite"`
	Path      string     `yaml:"path"`
	User      NodeUser   `yaml:"user"`
	Append    []Resource `yaml:"append"`
	Contents  Resource   `yaml:"contents"`
	Mode      *int       `yaml:"mode"`
}

// Filesystem must be defined here (not aliased) because it has methods
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

// Link must be defined here (not aliased) because it has methods
type Link struct {
	Group     NodeGroup `yaml:"group"`
	Overwrite *bool     `yaml:"overwrite"`
	Path      string    `yaml:"path"`
	User      NodeUser  `yaml:"user"`
	Hard      *bool     `yaml:"hard"`
	Target    string    `yaml:"target"`
}

// Resource must be defined here (not aliased) because it has methods
type Resource struct {
	Compression  *string      `yaml:"compression"`
	HTTPHeaders  HTTPHeaders  `yaml:"http_headers"`
	Source       *string      `yaml:"source"`
	Inline       *string      `yaml:"inline"` // Added, not in ignition spec
	Local        *string      `yaml:"local"`  // Added, not in ignition spec
	Verification Verification `yaml:"verification"`
}

// Tree must be defined here (not aliased) because it has methods
type Tree struct {
	Local string  `yaml:"local"`
	Path  *string `yaml:"path"`
}

// New types in v0_3
type Clevis struct {
	Custom    *Custom `yaml:"custom"`
	Tang      []Tang  `yaml:"tang"`
	Threshold *int    `yaml:"threshold"`
	Tpm2      *bool   `yaml:"tpm2"`
}

type Custom struct {
	Config       string `yaml:"config"`
	NeedsNetwork *bool  `yaml:"needs_network"`
	Pin          string `yaml:"pin"`
}

type Luks struct {
	Clevis     *Clevis      `yaml:"clevis"`
	Device     *string      `yaml:"device"`
	KeyFile    Resource     `yaml:"key_file"`
	Label      *string      `yaml:"label"`
	Name       string       `yaml:"name"`
	Options    []LuksOption `yaml:"options"`
	UUID       *string      `yaml:"uuid"`
	WipeVolume *bool        `yaml:"wipe_volume"`
}

type LuksOption string

type Tang struct {
	Thumbprint *string `yaml:"thumbprint"`
	URL        string  `yaml:"url"`
}

// Modified types in v0_3
type Partition struct {
	GUID               *string `yaml:"guid"`
	Label              *string `yaml:"label"`
	Number             int     `yaml:"number"`
	Resize             *bool   `yaml:"resize"`
	ShouldExist        *bool   `yaml:"should_exist"`
	SizeMiB            *int    `yaml:"size_mib"`
	StartMiB           *int    `yaml:"start_mib"`
	TypeGUID           *string `yaml:"type_guid"`
	WipePartitionEntry *bool   `yaml:"wipe_partition_entry"`
}

type PasswdGroup struct {
	Gid          *int    `yaml:"gid"`
	Name         string  `yaml:"name"`
	PasswordHash *string `yaml:"password_hash"`
	ShouldExist  *bool   `yaml:"should_exist"`
	System       *bool   `yaml:"system"`
}

type PasswdUser struct {
	Gecos             *string            `yaml:"gecos"`
	Groups            []Group            `yaml:"groups"`
	HomeDir           *string            `yaml:"home_dir"`
	Name              string             `yaml:"name"`
	NoCreateHome      *bool              `yaml:"no_create_home"`
	NoLogInit         *bool              `yaml:"no_log_init"`
	NoUserGroup       *bool              `yaml:"no_user_group"`
	PasswordHash      *string            `yaml:"password_hash"`
	PrimaryGroup      *string            `yaml:"primary_group"`
	ShouldExist       *bool              `yaml:"should_exist"`
	SSHAuthorizedKeys []SSHAuthorizedKey `yaml:"ssh_authorized_keys"`
	Shell             *string            `yaml:"shell"`
	System            *bool              `yaml:"system"`
	UID               *int               `yaml:"uid"`
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

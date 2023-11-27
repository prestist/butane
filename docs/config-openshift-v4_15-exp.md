---
# This file is automatically generated from internal/doc and Ignition's
# config/doc.  Do not edit.
title: OpenShift v4.15.0-experimental
parent: Configuration specifications
nav_order: 150
---

# OpenShift Specification v4.15.0-experimental

**Note: This configuration is experimental and has not been stabilized. It is subject to change without warning or announcement.**

The OpenShift configuration is a YAML document conforming to the following specification, with **_italicized_** entries being optional:

<div id="spec-docs"></div>

* **variant** (string): used to differentiate configs for different operating systems. Must be `openshift` for this specification.
* **version** (string): the semantic version of the spec for this document. This document is for version `4.15.0-experimental` and generates Ignition configs with version `3.5.0-experimental`.
* **metadata** (object): metadata about the generated MachineConfig resource. Respected when rendering to a MachineConfig, ignored when rendering directly to an Ignition config.
  * **name** (string): a unique [name](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names) for this MachineConfig resource.
  * **labels** (object): string key/value pairs to apply as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/) to this MachineConfig resource. `machineconfiguration.openshift.io/role` is required.
* **_ignition_** (object): metadata about the configuration itself.
  * **_config_** (object): options related to the configuration.
    * **_merge_** (list of objects): a list of the configs to be merged to the current config.
      * **_source_** (string): the URL of the config. Supported schemes are `http`, `https`, `tftp`, `s3`, `arn`, `gs`, and [`data`](https://tools.ietf.org/html/rfc2397). When using `http`, it is advisable to use the verification option to ensure the contents haven't been modified. Mutually exclusive with `inline` and `local`.
      * **_inline_** (string): the contents of the config. Mutually exclusive with `source` and `local`.
      * **_local_** (string): a local path to the contents of the config, relative to the directory specified by the `--files-dir` command-line argument. Mutually exclusive with `source` and `inline`.
      * **_compression_** (string): the type of compression used on the config (null or gzip). Compression cannot be used with S3.
      * **_http_headers_** (list of objects): a list of HTTP headers to be added to the request. Available for `http` and `https` source schemes only.
        * **name** (string): the header name.
        * **_value_** (string): the header contents.
      * **_verification_** (object): options related to the verification of the config.
        * **_hash_** (string): the hash of the config, in the form `<type>-<value>` where type is either `sha512` or `sha256`. If `compression` is specified, the hash describes the decompressed config.
    * **_replace_** (object): the config that will replace the current.
      * **_source_** (string): the URL of the config. Supported schemes are `http`, `https`, `tftp`, `s3`, `arn`, `gs`, and [`data`](https://tools.ietf.org/html/rfc2397). When using `http`, it is advisable to use the verification option to ensure the contents haven't been modified. Mutually exclusive with `inline` and `local`.
      * **_inline_** (string): the contents of the config. Mutually exclusive with `source` and `local`.
      * **_local_** (string): a local path to the contents of the config, relative to the directory specified by the `--files-dir` command-line argument. Mutually exclusive with `source` and `inline`.
      * **_compression_** (string): the type of compression used on the config (null or gzip). Compression cannot be used with S3.
      * **_http_headers_** (list of objects): a list of HTTP headers to be added to the request. Available for `http` and `https` source schemes only.
        * **name** (string): the header name.
        * **_value_** (string): the header contents.
      * **_verification_** (object): options related to the verification of the config.
        * **_hash_** (string): the hash of the config, in the form `<type>-<value>` where type is either `sha512` or `sha256`. If `compression` is specified, the hash describes the decompressed config.
  * **_timeouts_** (object): options relating to `http` timeouts when fetching files over `http` or `https`.
    * **_http_response_headers_** (integer): the time to wait (in seconds) for the server's response headers (but not the body) after making a request. 0 indicates no timeout. Default is 10 seconds.
    * **_http_total_** (integer): the time limit (in seconds) for the operation (connection, request, and response), including retries. 0 indicates no timeout. Default is 0.
  * **_security_** (object): options relating to network security.
    * **_tls_** (object): options relating to TLS when fetching resources over `https`.
      * **_certificate_authorities_** (list of objects): the list of additional certificate authorities (in addition to the system authorities) to be used for TLS verification when fetching over `https`. All certificate authorities must have a unique `source`, `inline`, or `local`.
        * **_source_** (string): the URL of the certificate bundle (in PEM format). The bundle can contain multiple concatenated certificates. Supported schemes are `http`, `https`, `tftp`, `s3`, `arn`, `gs`, and [`data`](https://tools.ietf.org/html/rfc2397). When using `http`, it is advisable to use the verification option to ensure the contents haven't been modified. Mutually exclusive with `inline` and `local`.
        * **_inline_** (string): the contents of the certificate bundle (in PEM format). The bundle can contain multiple concatenated certificates. Mutually exclusive with `source` and `local`.
        * **_local_** (string): a local path to the contents of the certificate bundle (in PEM format), relative to the directory specified by the `--files-dir` command-line argument. The bundle can contain multiple concatenated certificates. Mutually exclusive with `source` and `inline`.
        * **_compression_** (string): the type of compression used on the certificate bundle (null or gzip). Compression cannot be used with S3.
        * **_http_headers_** (list of objects): a list of HTTP headers to be added to the request. Available for `http` and `https` source schemes only.
          * **name** (string): the header name.
          * **_value_** (string): the header contents.
        * **_verification_** (object): options related to the verification of the certificate bundle.
          * **_hash_** (string): the hash of the certificate bundle, in the form `<type>-<value>` where type is either `sha512` or `sha256`. If `compression` is specified, the hash describes the decompressed certificate bundle.
  * **_proxy_** (object): options relating to setting an `HTTP(S)` proxy when fetching resources.
    * **_http_proxy_** (string): will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by `https_proxy` or `no_proxy`.
    * **_https_proxy_** (string): will be used as the proxy URL for HTTPS requests unless overridden by `no_proxy`.
    * **_no_proxy_** (list of strings): specifies a list of strings to hosts that should be excluded from proxying. Each value is represented by an `IP address prefix (1.2.3.4)`, `an IP address prefix in CIDR notation (1.2.3.4/8)`, `a domain name`, or `a special DNS label (*)`. An IP address prefix and domain name can also include a literal port number `(1.2.3.4:80)`. A domain name matches that name and all subdomains. A domain name with a leading `.` matches subdomains only. For example `foo.com` matches `foo.com` and `bar.foo.com`; `.y.com` matches `x.y.com` but not `y.com`. A single asterisk `(*)` indicates that no proxying should be done.
* **_storage_** (object): describes the desired state of the system's storage devices.
  * **_disks_** (list of objects): the list of disks to be configured and their options. Every entry must have a unique `device`.
    * **device** (string): the absolute path to the device. Devices are typically referenced by the `/dev/disk/by-*` symlinks. The boot disk can be referenced as `/dev/disk/by-id/coreos-boot-disk`.
    * **_wipe_table_** (boolean): whether or not the partition tables shall be wiped. When true, the partition tables are erased before any further manipulation. Otherwise, the existing entries are left intact.
    * **_partitions_** (list of objects): the list of partitions and their configuration for this particular disk. Every partition must have a unique `number`, or if 0 is specified, a unique `label`.
      * **_label_** (string): the PARTLABEL for the partition.
      * **_number_** (integer): the partition number, which dictates its position in the partition table (one-indexed). If zero, use the next available partition slot.
      * **_size_mib_** (integer): the size of the partition (in mebibytes). If zero, the partition will be made as large as possible.
      * **_start_mib_** (integer): the start of the partition (in mebibytes). If zero, the partition will be positioned at the start of the largest block available.
      * **_type_guid_** (string): the GPT [partition type GUID](https://en.wikipedia.org/wiki/GUID_Partition_Table#Partition_type_GUIDs). If omitted, the default will be 0FC63DAF-8483-4772-8E79-3D69D8477DE4 (Linux filesystem data).
      * **_guid_** (string): the GPT unique partition GUID.
      * **_wipe_partition_entry_** (boolean): if true, Ignition will clobber an existing partition if it does not match the config. If false (default), Ignition will fail instead.
      * **_should_exist_** (boolean): whether or not the partition with the specified `number` should exist. If omitted, it defaults to true. If false Ignition will either delete the specified partition or fail, depending on `wipePartitionEntry`. If false `number` must be specified and non-zero and `label`, `start`, `size`, `guid`, and `typeGuid` must all be omitted.
      * **_resize_** (boolean): whether or not the existing partition should be resized. If omitted, it defaults to false. If true, Ignition will resize an existing partition if it matches the config in all respects except the partition size.
  * **_raid_** (list of objects): the list of RAID arrays to be configured. Every RAID array must have a unique `name`.
    * **name** (string): the name to use for the resulting md device.
    * **level** (string): the redundancy level of the array (e.g. linear, raid1, raid5, etc.).
    * **devices** (list of strings): the list of devices (referenced by their absolute path) in the array.
    * **_spares_** (integer): the number of spares (if applicable) in the array.
    * **_options_** (list of strings): any additional options to be passed to mdadm.
  * **_filesystems_** (list of objects): the list of filesystems to be configured. `device` and `format` need to be specified. Every filesystem must have a unique `device`.
    * **device** (string): the absolute path to the device. Devices are typically referenced by the `/dev/disk/by-*` symlinks.
    * **format** (string): the filesystem format (ext4, xfs, vfat, or swap).
    * **_path_** (string): the mount-point of the filesystem while Ignition is running relative to where the root filesystem will be mounted. This is not necessarily the same as where it should be mounted in the real root, but it is encouraged to make it the same.
    * **_wipe_filesystem_** (boolean): whether or not to wipe the device before filesystem creation, see [Ignition's documentation on filesystems](https://coreos.github.io/ignition/operator-notes/#filesystem-reuse-semantics) for more information. Defaults to false.
    * **_label_** (string): the label of the filesystem.
    * **_uuid_** (string): the uuid of the filesystem.
    * **_options_** (list of strings): any additional options to be passed to the format-specific mkfs utility.
    * **_mount_options_** (list of strings): any special options to be passed to the mount command.
    * **_with_mount_unit_** (boolean): whether to additionally generate a generic mount unit for this filesystem or a swap unit for this swap area. If a more specific unit is needed, a custom one can be specified in the `systemd.units` section. The unit will be named with the [escaped](https://www.freedesktop.org/software/systemd/man/systemd-escape.html) version of the `path` or `device`, depending on the unit type. If your filesystem is located on a Tang-backed LUKS device, the unit will automatically require network access if you specify the device as `/dev/mapper/<device-name>` or `/dev/disk/by-id/dm-name-<device-name>`.
  * **_files_** (list of objects): the list of files to be written. Every file, directory and link must have a unique `path`.
    * **path** (string): the absolute path to the file.
    * **_overwrite_** (boolean): whether to delete preexisting nodes at the path. `contents` must be specified if `overwrite` is true. Defaults to false.
    * **_contents_** (object): options related to the contents of the file.
      * **_source_** (string): the URL of the file. Only the [`data`](https://tools.ietf.org/html/rfc2397) scheme is supported. If source is omitted and a regular file already exists at the path, Ignition will do nothing. If source is omitted and no file exists, an empty file will be created. Mutually exclusive with `inline` and `local`.
      * **_inline_** (string): the contents of the file. Mutually exclusive with `source` and `local`.
      * **_local_** (string): a local path to the contents of the file, relative to the directory specified by the `--files-dir` command-line argument. Mutually exclusive with `source` and `inline`.
      * **_compression_** (string): the type of compression used on the file (null or gzip). Compression cannot be used with S3.
      * **_verification_** (object): options related to the verification of the file.
        * **_hash_** (string): the hash of the file, in the form `<type>-<value>` where type is either `sha512` or `sha256`. If `compression` is specified, the hash describes the decompressed file.
    * **_mode_** (integer): the file's permission mode. Setuid/setgid/sticky bits are not supported. If not specified, the permission mode for files defaults to 0644 or the existing file's permissions if `overwrite` is false, `contents` is unspecified, and a file already exists at the path.
    * **_user_** (object): specifies the file's owner.
      * **_id_** (integer): the user ID of the owner.
      * **_name_** (string): the user name of the owner.
    * **_group_** (object): specifies the file's group.
      * **_id_** (integer): the group ID of the group.
      * **_name_** (string): the group name of the group.
  * **_luks_** (list of objects): the list of luks devices to be created. Every device must have a unique `name`.
    * **name** (string): the name of the luks device.
    * **device** (string): the absolute path to the device. Devices are typically referenced by the `/dev/disk/by-*` symlinks.
    * **_key_file_** (object): options related to the contents of the key file.
      * **_source_** (string): the URL of the key file. Supported schemes are `http`, `https`, `tftp`, `s3`, `arn`, `gs`, and [`data`](https://tools.ietf.org/html/rfc2397). When using `http`, it is advisable to use the verification option to ensure the contents haven't been modified. Mutually exclusive with `inline` and `local`.
      * **_inline_** (string): the contents of the key file. Mutually exclusive with `source` and `local`.
      * **_local_** (string): a local path to the contents of the key file, relative to the directory specified by the `--files-dir` command-line argument. Mutually exclusive with `source` and `inline`.
      * **_compression_** (string): the type of compression used on the key file (null or gzip). Compression cannot be used with S3.
      * **_http_headers_** (list of objects): a list of HTTP headers to be added to the request. Available for `http` and `https` source schemes only.
        * **name** (string): the header name.
        * **_value_** (string): the header contents.
      * **_verification_** (object): options related to the verification of the key file.
        * **_hash_** (string): the hash of the key file, in the form `<type>-<value>` where type is either `sha512` or `sha256`. If `compression` is specified, the hash describes the decompressed key file.
    * **_label_** (string): the label of the luks device.
    * **_uuid_** (string): the uuid of the luks device.
    * **_options_** (list of strings): any additional options to be passed to `cryptsetup luksFormat`.
    * **_discard_** (boolean): whether to issue discard commands to the underlying block device when blocks are freed. Enabling this improves performance and device longevity on SSDs and space utilization on thinly provisioned SAN devices, but leaks information about which disk blocks contain data. If omitted, it defaults to false.
    * **_open_options_** (list of strings): any additional options to be passed to `cryptsetup luksOpen`. Supported options will be persistently written to the luks volume.
    * **_wipe_volume_** (boolean): whether or not to wipe the device before volume creation, see [Ignition's documentation on filesystems](https://coreos.github.io/ignition/operator-notes/#filesystem-reuse-semantics) for more information.
    * **_clevis_** (object): describes the clevis configuration for the luks device.
      * **_tang_** (list of objects): describes a tang server. Every server must have a unique `url`.
        * **url** (string): url of the tang server.
        * **thumbprint** (string): thumbprint of a trusted signing key.
        * **_advertisement_** (string): the advertisement JSON. If not specified, the advertisement is fetched from the tang server during provisioning.
      * **_tpm2_** (boolean): whether or not to use a tpm2 device.
      * **_threshold_** (integer): sets the minimum number of pieces required to decrypt the device. Default is 1.
      * **_custom_** (object): overrides the clevis configuration. The `pin` & `config` will be passed directly to `clevis luks bind`. If specified, all other clevis options must be omitted.
        * **pin** (string): the clevis pin.
        * **config** (string): the clevis configuration JSON.
        * **_needs_network_** (boolean): whether or not the device requires networking.
  * **_trees_** (list of objects): a list of local directory trees to be embedded in the config. Symlinks must not be present. Ownership is not preserved. File modes are set to 0755 if the local file is executable or 0644 otherwise. File attributes can be overridden by creating a corresponding entry in the `files` section; such entries must omit `contents`.
    * **local** (string): the base of the local directory tree, relative to the directory specified by the `--files-dir` command-line argument.
    * **_path_** (string): the path of the tree within the target system. Defaults to `/`.
* **_systemd_** (object): describes the desired state of the systemd units.
  * **_units_** (list of objects): the list of systemd units. Every unit must have a unique `name`.
    * **name** (string): the name of the unit. This must be suffixed with a valid unit type (e.g. "thing.service").
    * **_enabled_** (boolean): whether or not the service shall be enabled. When true, the service is enabled. When false, the service is disabled. When omitted, the service is unmodified. In order for this to have any effect, the unit must have an install section.
    * **_mask_** (boolean): whether or not the service shall be masked. When true, the service is masked by symlinking it to `/dev/null`. When false, the service is unmasked by deleting the symlink to `/dev/null` if it exists.
    * **_contents_** (string): the contents of the unit. Mutually exclusive with `contents_local`.
    * **_contents_local_** (string): a local path to the contents of the unit, relative to the directory specified by the `--files-dir` command-line argument. Mutually exclusive with `contents`.
    * **_dropins_** (list of objects): the list of drop-ins for the unit. Every drop-in must have a unique `name`.
      * **name** (string): the name of the drop-in. This must be suffixed with ".conf".
      * **_contents_** (string): the contents of the drop-in. Mutually exclusive with `contents_local`.
      * **_contents_local_** (string): a local path to the contents of the drop-in, relative to the directory specified by the `--files-dir` command-line argument. Mutually exclusive with `contents`.
* **_passwd_** (object): describes the desired additions to the passwd database.
  * **_users_** (list of objects): the list of accounts that shall exist. All users must have a unique `name`.
    * **name** (string): the username for the account. Must be `core`.
    * **_password_hash_** (string): the hashed password for the account.
    * **_ssh_authorized_keys_** (list of strings): a list of SSH keys to be added as an SSH key fragment at `.ssh/authorized_keys.d/ignition` in the user's home directory. All SSH keys must be unique.
    * **_ssh_authorized_keys_local_** (list of strings): a list of local paths to SSH key files, relative to the directory specified by the `--files-dir` command-line argument, to be added as SSH key fragments at `.ssh/authorized_keys.d/ignition` in the user's home directory. All SSH keys must be unique. Each file may contain multiple SSH keys, one per line.
* **_boot_device_** (object): describes the desired boot device configuration. At least one of `luks` or `mirror` must be specified.
  * **_layout_** (string): the disk layout of the target OS image. Supported values are `aarch64`, `ppc64le`, and `x86_64`. Defaults to `x86_64`.
  * **_luks_** (object): describes the clevis configuration for encrypting the root filesystem.
    * **_tang_** (list of objects): describes a tang server. Every server must have a unique `url`.
      * **url** (string): url of the tang server.
      * **thumbprint** (string): thumbprint of a trusted signing key.
      * **_advertisement_** (string): the advertisement JSON. If not specified, the advertisement is fetched from the tang server during provisioning.
    * **_tpm2_** (boolean): whether or not to use a tpm2 device.
    * **_threshold_** (integer): sets the minimum number of pieces required to decrypt the device. Default is 1.
    * **_discard_** (boolean): whether to issue discard commands to the underlying block device when blocks are freed. Enabling this improves performance and device longevity on SSDs and space utilization on thinly provisioned SAN devices, but leaks information about which disk blocks contain data. If omitted, it defaults to false.
  * **_mirror_** (object): describes mirroring of the boot disk for fault tolerance.
    * **_devices_** (list of strings): the list of whole-disk devices (not partitions) to include in the disk array, referenced by their absolute path. At least two devices must be specified.
* **_grub_** (object): describes the desired GRUB bootloader configuration.
  * **_users_** (list of objects): the list of GRUB superusers.
    * **name** (string): the user name.
    * **password_hash** (string): the PBKDF2 password hash, generated with `grub2-mkpasswd-pbkdf2`.
* **_openshift_** (object): describes miscellaneous OpenShift configuration. Respected when rendering to a MachineConfig, ignored when rendering directly to an Ignition config.
  * **_kernel_type_** (string): which kernel to use on the node. Must be `default` or `realtime`.
  * **_kernel_arguments_** (list of strings): arguments to be added to the kernel command line.
  * **_extensions_** (list of strings): RHCOS extensions to be installed on the node.
  * **_fips_** (boolean): whether or not to enable FIPS 140-2 compatibility. If omitted, defaults to false.
# Butane

> **Note:** Butane has been merged into the [Ignition](https://github.com/coreos/ignition) repository.
> All new development, issues, and pull requests should be directed there. The Butane source code now
> lives at [`butane/`](https://github.com/coreos/ignition/tree/main/butane) in the Ignition repo.
>
> **Why?** The separate Butane transpilation step has been a pain point for users adopting Ignition-based
> systems. By merging Butane into Ignition, Ignition can natively accept Butane YAML configs at boot,
> eliminating the need to transpile configs ahead of time. This also removes the circular dependency
> between the two projects and simplifies the development and release process. See
> [fedora-coreos-tracker#2006](https://github.com/coreos/fedora-coreos-tracker/issues/2006) for the
> full discussion.
>
> This repository is maintained for historical reference and for existing Go consumers that import
> `github.com/coreos/butane`. A shim module will be provided to ease migration to the new import
> path `github.com/coreos/ignition/v2/butane/...`.

Butane (formerly the Fedora CoreOS Config Transpiler, FCCT) translates human readable Butane Configs
into machine readable [Ignition](https://github.com/coreos/ignition) Configs. See the [getting
started](docs/getting-started.md) guide for how to use Butane and the [configuration
specifications](docs/specs.md) for everything Butane configs support.

For information on developing Butane, using it as a library, or understanding how the binaries released
in this repository are built, see the [development docs](docs/development.md).

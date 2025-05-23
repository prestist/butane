#!/bin/bash
# Script for generating Fedora releng release signing tickets.  Generated by
# https://github.com/coreos/repo-templates; do not edit downstream.

set -euo pipefail

usage() {
    echo "Usage: $0 test <version-release> <dir>"
    echo "Usage: $0 ticket <version-release>"
    exit 1
}

make_script() {
    sed -e "s/@@VERSION@@/$ver/g" -e "s/@@RELEASE@@/$rel/g" <<'EOF'
#!/bin/bash
set -eux -o pipefail

# Use the Fedora 42 key for the detached signatures
KEYTOSIGNWITH='fedora-42'

VR='@@VERSION@@-@@RELEASE@@.fc42'
RPMKEY='105ef944' # Fedora 42 key

do_sign() {
    # Sign with sigul unless FAKESIGN=1
    if [ ${FAKESIGN:-0} != 1 ]; then
        sigul sign-data -a $KEYTOSIGNWITH "$1" -o "$1.asc"
    else
        echo INVALID > "$1.asc"
    fi
}

# Grab the binaries out of the redistributable rpm
rpm="butane-redistributable-${VR}.noarch.rpm"
koji download-build --key $RPMKEY --rpm $rpm
rpm -qip $rpm | grep -P "^Signature.*${RPMKEY}$" # Verify the output has the key in it
rpm2cpio $rpm | cpio -idv './usr/share/butane/butane-*'

# Rename the binaries
mv usr/share/butane/butane-aarch64-apple-darwin \
    butane-aarch64-apple-darwin
mv usr/share/butane/butane-aarch64-unknown-linux-gnu-static \
    butane-aarch64-unknown-linux-gnu
mv usr/share/butane/butane-ppc64le-unknown-linux-gnu-static \
    butane-ppc64le-unknown-linux-gnu
mv usr/share/butane/butane-s390x-unknown-linux-gnu-static \
    butane-s390x-unknown-linux-gnu
mv usr/share/butane/butane-x86_64-apple-darwin \
    butane-x86_64-apple-darwin
mv usr/share/butane/butane-x86_64-pc-windows-gnu.exe \
    butane-x86_64-pc-windows-gnu.exe
mv usr/share/butane/butane-x86_64-unknown-linux-gnu-static \
    butane-x86_64-unknown-linux-gnu

# Sign them
do_sign butane-aarch64-apple-darwin
do_sign butane-aarch64-unknown-linux-gnu
do_sign butane-ppc64le-unknown-linux-gnu
do_sign butane-s390x-unknown-linux-gnu
do_sign butane-x86_64-apple-darwin
do_sign butane-x86_64-pc-windows-gnu.exe
do_sign butane-x86_64-unknown-linux-gnu

# Fix permissions and clean up
chmod go+r *.asc
rm $rpm; rmdir ./usr/share/butane; rmdir ./usr/share; rmdir ./usr
EOF
}

make_ticket() {
    sed "s/@@VERSION@@/$ver/g" <<'EOF'
TITLE: Create detached signatures for the butane @@VERSION@@ release

Please create detached signatures for the binaries we will upload to GitHub for the `butane` @@VERSION@@ release.  This is a manual process for now, pending the automation discussed in https://pagure.io/robosignatory/issue/53 and https://github.com/coreos/fedora-coreos-tracker/issues/335.

The binaries themselves have been built in koji.  Here is a small script to grab all of the rpms and the files out of the rpms and name them appropriately:

```
EOF
    make_script
    cat <<'EOF'
```

After running this you should end up with a directory with files in it like:

```
$ ls -1
butane-aarch64-apple-darwin
butane-aarch64-apple-darwin.asc
butane-aarch64-unknown-linux-gnu
butane-aarch64-unknown-linux-gnu.asc
butane-ppc64le-unknown-linux-gnu
butane-ppc64le-unknown-linux-gnu.asc
butane-s390x-unknown-linux-gnu
butane-s390x-unknown-linux-gnu.asc
butane-x86_64-apple-darwin
butane-x86_64-apple-darwin.asc
butane-x86_64-pc-windows-gnu.exe
butane-x86_64-pc-windows-gnu.exe.asc
butane-x86_64-unknown-linux-gnu
butane-x86_64-unknown-linux-gnu.asc
```
EOF
}

[ "$#" -lt 2 ] && usage
cmd="$1"
ver="$2"
# Disallow version with preceding "v"
echo "$ver" | grep -q "v" && usage
# Require version-release
echo "$ver" | grep -q "-" || usage
rel="${ver#*-}"
ver="${ver%%-*}"

case "$cmd" in
test)
    [ "$#" != 3 ] && usage
    dir="$3"
    mkdir "$dir"
    make_script > "$dir/script"
    cd "$dir" && FAKESIGN=1 bash script
    ;;
ticket)
    make_ticket
    ;;
*)
    usage
    ;;
esac

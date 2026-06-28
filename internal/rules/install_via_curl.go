package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const installViaCurlPath = "install-via-curl.sh"

const installViaCurlTemplate = `#!/usr/bin/env bash
set -eo pipefail

if [[ ${OS:-} = Windows_NT ]]; then
    echo 'error: please install __NAME__ using Windows Subsystem for Linux'
    exit 1
fi

error() {
    echo "$@" >&2
    exit 1
}

command -v tar >/dev/null || error 'tar is required to install __NAME__'

case $(uname -ms) in
    'Darwin x86_64')
        target=darwin-amd64
    ;;
    'Darwin arm64')
        target=darwin-arm64
    ;;
    'Linux aarch64' | 'Linux arm64')
        target=linux-arm64
    ;;
    'Linux x86_64' | *)
        target=linux-amd64
    ;;
esac

if [[ "$INSTALL_TAG" != "" ]];then
    install_version=$INSTALL_VERSION
    if [[ -z "$install_version" ]];then
        install_version=$INSTALL_TAG
    fi
    install_version=${install_version/#"v"}
    file="__NAME__-v${install_version}-${target}"
    uri="https://github.com/__OWNER__/__REPO__/releases/download/${INSTALL_TAG}/${file}"
else
    latestURL="https://github.com/__OWNER__/__REPO__/releases/latest"
    headers=$(curl "$latestURL" -so /dev/null -D -)
    if [[ "$headers" != *302* ]];then
        error "expect 302 from $latestURL"
    fi

    location=$(echo "$headers"|grep "location: ")
    if [[ -z $location ]];then
        error "expect 302 location from $latestURL"
    fi
    locationURL=${location/#"location: "}
    locationURL=${locationURL/%$'\n'}
    locationURL=${locationURL/%$'\r'}

    versionName=""
    if [[ "$locationURL" = *'/__NAME__-v'* ]];then
        versionName=${locationURL/#*'/__NAME__-v'}
        elif [[ "$locationURL" = *'/tag/v'* ]];then
        versionName=${locationURL/#*'/tag/v'}
    fi

    if [[ -z $versionName ]];then
        error "expect tag format: __NAME__-v1.x.x, actual: $locationURL"
    fi

    file="__NAME__-v${versionName}-${target}"
    uri="$latestURL/download/$file"
fi

install_dir=$HOME/.xgo
bin_dir=$install_dir/bin

if [[ ! -d $bin_dir ]]; then
    mkdir -p "$bin_dir" || error "failed to create install directory \"$bin_dir\""
fi

tmp_dir=$(mktemp -d)
trap 'rm -rf "$tmp_dir"' EXIT

curl --fail --location --progress-bar --output "${tmp_dir}/${file}" "$uri" || error "failed to download __NAME__ from \"$uri\""

chmod +x "${tmp_dir}/${file}"

mv "${tmp_dir}/${file}" "${tmp_dir}/__NAME__"

# detect if we can write to /usr/local/bin without sudo
if touch /usr/local/bin/.write_test 2>/dev/null; then
    rm -f /usr/local/bin/.write_test
    maybe_sudo=
else
    maybe_sudo=sudo
fi

echo "installing __NAME__ to /usr/local/bin"
if [[ -f /usr/local/bin/__NAME__ ]];then
    $maybe_sudo mv /usr/local/bin/{__NAME__,__NAME___backup}
fi
$maybe_sudo install "${tmp_dir}/__NAME__" /usr/local/bin
$maybe_sudo rm -f /usr/local/bin/__NAME___backup || true

echo "Successfully installed, to get started, run:"
echo "  __NAME__ --help"
`

func FixInstallViaCurl(project model.Project, dryRun bool) (model.FixResult, error) {
	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}

	path := filepath.Join(project.Root, installViaCurlPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "install.via.curl",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", installViaCurlPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}

	result := model.FixResult{RuleID: "install.via.curl"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", installViaCurlPath)}
		return result, nil
	}

	content := substituteMeta(installViaCurlTemplate, meta)
	if err := os.WriteFile(path, []byte(content), 0o755); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", installViaCurlPath)}
	return result, nil
}
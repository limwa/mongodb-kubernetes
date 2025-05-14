#!/usr/bin/env bash
set -Eeou pipefail

source scripts/funcs/printing
source scripts/dev/set_env_context.sh

function check_command() {
  command -v "$1" &> /dev/null
}

if ! check_command uname; then
    echo "Command \"uname\" not found, please confirm that you have coreutils installed."
    echo "Or did you forget to run the following command in your terminal?"
    echo
    echo "  echo \"PATH=\"/usr/local/opt/coreutils/libexec/gnubin:\$PATH\"\" >> ~/.bashrc"
    exit 1
fi

function get_os_name() {
  if [ "$(uname)" = "Darwin" ] ; then
    echo "macos"
  elif [ "$(uname)" = "Linux" ] ; then # Ubuntu only
    . /etc/os-release
    echo "${ID}" # ubuntu
  else
    echo "unknown"
  fi
}

function check_python() {
    check_command python3.13 \
        || (check_command python && python --version | grep -q "${PYTHON_VERSION}")
}

# Installing prerequisites
tools="kubectl helm coreutils kind jq shellcheck python@${PYTHON_VERSION}"

if [ "$(get_os_name)" = "macos" ] ; then
  echo "The following tools will be installed using homebrew: ${tools}"
  echo "Note, that you must download 'go' and Docker by yourself"

  # shellcheck disable=SC2086
  brew install ${tools} 2>&1 | prepend "brew install"

elif [ "$(get_os_name)" = "ubuntu" ] ; then # Ubuntu only
  echo "The following tools will be installed: ${tools}"
  echo "Note, that you must download 'go' and Docker by yourself"

  function install_kops() {
    kops_version="$(curl -s https://api.github.com/repos/kubernetes/kops/releases/latest | grep tag_name | cut -d '"' -f 4)"
    curl -Lo kops "https://github.com/kubernetes/kops/releases/download/${kops_version}/kops-linux-amd64"
    echo "hi"
    chmod +x kops
    mv kops "${GOBIN}" || true
  }

  check_command kubectl || sudo snap install kubectl --classic  || true
  check_command kops || install_kops
  check_command helm || sudo snap install helm --classic  || true
  check_command kind || go install sigs.k8s.io/kind
  check_command shellcheck || sudo snap install --channel=edge shellcheck || true
fi

# Verifying prerequisites
echo "Checking if the following tools are installed: ${tools}"
check_command kubectl || (echo "kubectl not found, please install it" && exit 1)
check_command helm || (echo "helm not found, please install it" && exit 1)
check_command kind || (echo "kind not found, please install it" && exit 1)
check_command jq || (echo "jq not found, please install it" && exit 1)
check_command shellcheck || (echo "shellcheck not found, please install it" && exit 1)
check_python || (echo "python@${PYTHON_VERSION} not found, please install it" && exit 1)
echo "All OS prerequisites are installed"

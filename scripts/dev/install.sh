#!/usr/bin/env bash
set -Eeou pipefail

source scripts/funcs/printing
source scripts/dev/set_env_context.sh

source scripts/dev/install_os_prerequisites.sh

echo "Installing Python packages"
scripts/dev/recreate_python_venv.sh 2>&1 | prepend "recreate_python_venv.sh"

title "Tools are installed"

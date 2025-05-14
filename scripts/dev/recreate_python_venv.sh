#!/usr/bin/env bash

# This scripts recreates local python venv in the ${PROJECT_DIR} directory from the current context.

set -Eeou pipefail

source scripts/dev/set_env_context.sh

if [[ -d "${PROJECT_DIR}"/venv ]]; then
  echo "Removing venv..."
  cd "${PROJECT_DIR}"
  rm -rf "venv"
fi

# in our EVG hosts, python versions are always in /opt/python
function get_python_path() {
  possible_commands=(
    "python${PYTHON_VERSION}"
    "python"
    "/opt/python/${PYTHON_VERSION}/bin/python3"
  )

  for command in "${possible_commands[@]}"; do
    if command -v "${command}" &> /dev/null && "${command}" --version | grep -q "3.13"; then
      echo "${command}"
      return
    fi
  done

  return 1
}

python_bin="$(get_python_path)"
echo "Using python from the following path: ${python_bin}"

"${python_bin}" -m venv venv
source venv/bin/activate
pip install --upgrade pip
pip install -r requirements.txt
echo "Python venv was recreated successfully."
echo "Current python path: $(which python)"
python --version

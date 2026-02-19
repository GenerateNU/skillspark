#!/usr/bin/env bash
set -euo pipefail

load_root_env() {
  local script_dir root_dir env_file
  script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
  root_dir="$(cd -- "${script_dir}/.." && pwd)"
  env_file="${root_dir}/.env"

  if [[ -f "${env_file}" ]]; then
    set -a
    # shellcheck disable=SC1090
    . "${env_file}"
    set +a
  else
    echo "warning: ${env_file} not found; continuing without loading root .env" >&2
  fi
}

# If executed directly: load env, then optionally run a command in that context.
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  load_root_env
  if [[ "$#" -gt 0 ]]; then
    exec "$@"
  fi
else
  # If sourced: only load env into the current shell.
  load_root_env
fi

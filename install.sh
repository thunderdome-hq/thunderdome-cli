#!/bin/zsh

setup_go() {
  printf "Checking if go is installed...\n"
  if command -v go &> /dev/null; then
    local version=$( echo `go version | { read _ _ v _; echo ${v#go}; }`)
    printf "Go ${version} is already installed at ${HOME}/${GOPATH}.\n"
    return 0
  fi

  local version="1.20.1"
  local release="go${version}.darwin-arm64.tar.gz"
  local url="https://go.dev/dl/${release}"
  local download="https://dl.google.com/go/${release}"
  printf "Go is not installed, installing go ${version} from ${url} to ${HOME}.\n"
  curl -O ${download}
  tar -C ${HOME} -xzvf "${release}"
  rm "${release}"

  local verify=$( echo `go version | { read _ _ v _; echo ${v#go}; }`)
  printf "Go ${verify} is now installed at ${HOME}/${GOPATH}.\n"
  return 0
}

setup_path() {
  printf "Checking if go bin is in PATH...\n"
  local bin="$(go env GOPATH)/bin"
  if [[ ":${PATH}:" == *":${bin}:"* ]]; then
    printf "Go bin ${bin} is already in PATH.\n"
    return 0
  fi

  printf "Go bin ${bin} is not in PATH, adding it to PATH.\n"
  local shell=$(echo $SHELL)

  local configs=(
    ["/bin/bash"]="$HOME/.bashrc $HOME/.bash_profile"
    ["/bin/zsh"]="$HOME/.zshrc $HOME/.zprofile"
    ["/bin/sh"]="$HOME/.profile"
  )

  if [[ ${configs[$shell]+_} ]]; then
    local command="echo 'export PATH=\$PATH:'\$(go env GOPATH)/bin' >> <file> && source <file>"
    printf "Shell ${shell} is not supported. The go bin must be manually added to PATH using the command\n"
    printf "\t${command}\n"
    printf "with the used shell's configuration file as <file>.\n"
    return 0
  fi

  config_files=${configs[$shell]}
  for file in $config_files; do
    if [! -f "$file" ]; then
      continue
    fi

    printf "Updating path in ${file}.\n"
    echo 'export PATH=$PATH:'${bin}'' >> ${file}
    source ${file}
  done

  printf "Go bin ${bin} is now in PATH.\n"
}

setup_thunderdome() {
  printf "Checking if Thunderdome is installed...\n"
  local bin="$(go env GOPATH)/bin"
  if [ -f "${bin}/thunderdome" ]; then
    printf "Thunderdome is already installed at ${bin}/thunderdome.\n"
    return 0
  fi

  printf "Thunderdome is not installed, installing it to ${bin}/thunderdome.\n"
  go install github.com/thunderdome-hq/thunderdome-cli/thunderdome@latest
  printf "Thunderdome is now installed at ${bin}/thunderdome.\n"
}

printf "Setting up Thunderdome...\n"
setup_go
setup_path
setup_thunderdome
printf "Thunderdome is now ready to be used. Get started by running 'thunderdome --help'\n"
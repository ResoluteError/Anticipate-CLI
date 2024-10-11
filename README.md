# Setup from scratch for your own projects

- `mkdir anticipate`
- `cd anticipate`
- `go mod init anticipate`
- `go get github.com/spf13/cobra`
- `go get github.com/spf13/viper`
- `go get github.com/manifoldco/promptui`
- `goreleaser init`

# Setup necessary for github actions and homebrew deployment

- Update the name_template in [.goreleaser.yaml](./.goreleaser.yaml)
- Create a repository to host the CLI or fork this repository (e.g. `Anticipate-CLI`)
- Create a homebrew tap repository (e.g. `homebrew-[CLI_REPO_NAME]`) (e.g. `homebrew-anticipate-cli`)
- Create a [GitHub Personal Access Token (PAT)](https://github.com/settings/tokens?type=beta)
  - Requires Read & Write permission for the contents of the homebrew repo
  - Requires Read & Write permission for the contents of the main cli repo
- Add the PAT to the repository running this workflow as a secret called `GO_RELEASER_PAT`

# Local Testing

- `go run main.go new 2025-03-17 -d "Birthday"`
- `go run main.go countdown`
- `go run main.go remove`

# Deploying

- Create and push a branch called `release/[VERSION]` (replace version with something like 0.0.1)

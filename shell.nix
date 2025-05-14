{ pkgs ? import <nixpkgs> }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    python313

    fzf
    jq
    awscli2
    kubernetes-helm
    kind
    kubectl
    shellcheck
    telepresence

    # for python dependencies
    openldap.dev
    cyrus_sasl.dev

    (pkgs.writeShellApplication {
      name = "mongodb-rebase";
      runtimeInputs = [ pkgs.git ];

      text = let
        branch_regex = "^nix\\+(.+)$";
      in ''
        if ! git diff --quiet; then
          echo "ERROR: git working directory is dirty"
          exit 1
        fi

        current_branch="$(git rev-parse --abbrev-ref HEAD)"

        if ! echo "$current_branch" | grep -Eq "${branch_regex}"; then
          echo "ERROR: current branch is not a nix-layered branch"
          exit 1
        fi

        upstream_branch="$(echo "$current_branch" | sed -E 's/${branch_regex}/\1/')"

        echo "Syncing changes made in '$current_branch' to '$upstream_branch'"

        git switch --force-create "$upstream_branch" "$current_branch"
        if [[ "$(git rev-parse --abbrev-ref HEAD)" != "$upstream_branch" ]]; then
          echo "ERROR: failed to switch to upstream branch"
          exit 1
        fi

        git rebase --onto master nix
      '';
    })
  ];
}

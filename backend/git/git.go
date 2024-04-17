package git

import (
	"github.com/cble-platform/cble/backend/ent"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func listRefs(url string) ([]*plumbing.Reference, error) {
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	return rem.List(&git.ListOptions{})
}

// Confirms a valid git repo exists at the given URL and is not empty.
//
// URL can be any valid git URL (e.g. https://github.com..., git@github.com:..., etc.)
func RepoExists(url string) bool {
	_, err := listRefs(url)
	return err == nil
}

// Confirms the git repo has a tag/hash with given value.
//
// URL can be any valid git URL (e.g. https://github.com..., git@github.com:..., etc.)
func RepoTagExists(url, tagName string) bool {
	refs, err := listRefs(url)
	if err != nil {
		return false
	}

	// Look at all tag refs to see if tag exists
	for _, ref := range refs {
		if ref.Name() == plumbing.ReferenceName(tagName) {
			return true
		}
	}
	return false
}

// Confirms the Provider has its git config setup properly.
func ValidateProviderGit(entProvider *ent.Provider) bool {
	return RepoTagExists(entProvider.ProviderGitURL, entProvider.ProviderVersion)
}

// Clones the repo associated with a Provider.
func CloneProvider(repoPath string, entProvider *ent.Provider) error {
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:               entProvider.ProviderGitURL,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	return err
}

// Checks out a version of a Provider.
func CheckoutProvider(repoPath string, entProvider *ent.Provider) error {
	// Open the local repo and get worktree
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Update the git repo from the remote
	err = repo.Fetch(&git.FetchOptions{})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}

	// Find the hash for the given tag (or commit hash directly)
	tagHash, err := repo.ResolveRevision(plumbing.Revision(entProvider.ProviderVersion))
	if err != nil {
		return err
	}

	// Checkout the version
	err = w.Checkout(&git.CheckoutOptions{
		Hash: *tagHash,
	})
	if err != nil {
		return err
	}
	return nil
}

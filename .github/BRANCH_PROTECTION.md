# Branch Protection Configuration

This document describes the recommended branch protection rules for this repository.

## Protected Branches

### `internal` branch (Release Branch)
**Critical**: This is the source of truth for all releases. All version tags MUST be created from this branch.

**Recommended Settings:**
- ✅ Require pull request reviews before merging
- ✅ Require status checks to pass before merging
  - Required checks: `lint-and-test`
- ✅ Require branches to be up to date before merging
- ✅ Require conversation resolution before merging
- ✅ Restrict who can push to matching branches (Platform Admins only for direct pushes)
- ✅ Allow force pushes: No
- ✅ Allow deletions: No

### `main` branch (Development Branch)
**Purpose**: Integration branch synced with upstream fork.

**Recommended Settings:**
- ✅ Require pull request reviews before merging (at least 1 approval)
- ✅ Require status checks to pass before merging
  - Required checks: `lint-and-test`
- ✅ Require conversation resolution before merging
- ✅ Allow force pushes: No (except by admins for upstream syncs)
- ✅ Allow deletions: No

## Setting Up Branch Protection

1. Go to repository Settings → Branches
2. Add rule for `internal` branch with settings above
3. Add rule for `main` branch with settings above
4. Save changes

## Why These Rules Matter

- **internal branch protection**: Prevents accidental commits that could break releases
- **Status checks**: Ensures code quality before merging
- **Required reviews**: Catches issues before they reach release branches
- **No force pushes on internal**: Maintains integrity of release history
- **No deletions**: Prevents catastrophic loss of release branches

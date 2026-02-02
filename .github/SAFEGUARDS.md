# Repository Safeguards

This document describes the facilities in place to prevent releases from being tagged from the wrong branch.

## Problem Statement

This repository maintains releases from the `internal` branch rather than `main`. To prevent confusion and mistakes, we've implemented multiple layers of protection.

## Safeguards Implemented

### 1. **Automated Release Script** (`scripts/release.sh`)
**Purpose**: Automates the release process and enforces correct workflow

**Features**:
- ✅ Validates current branch is `internal`
- ✅ Checks version format (vX.Y.Z)
- ✅ Verifies branch is up-to-date with remote
- ✅ Prevents duplicate tags
- ✅ Creates annotated tags
- ✅ Pushes to origin automatically
- ✅ Provides clear error messages

**Usage**:
```bash
./scripts/release.sh v0.0.4
# or
make release VERSION=v0.0.4
```

### 2. **GitHub Actions Validation** (`.github/workflows/validate-tags.yaml`)
**Purpose**: Server-side validation that tags come from `internal`

**Features**:
- ✅ Triggers on all tag pushes
- ✅ Verifies tag was created from `internal` branch
- ✅ Validates version format
- ✅ Provides remediation instructions if validation fails
- ✅ Blocks CI/CD if tag is from wrong branch

**Result**: Even if someone bypasses local checks, GitHub will catch it.

### 3. **Git Pre-Push Hook** (`.git/hooks/pre-push`)
**Purpose**: Local client-side validation before pushing tags

**Features**:
- ✅ Runs automatically before `git push`
- ✅ Checks if pushing a version tag
- ✅ Validates current branch is `internal`
- ✅ Prevents accidental tag pushes from wrong branch
- ✅ Provides clear instructions to fix

**Setup**:
```bash
make setup-hooks
```

### 4. **Branch Protection** (`.github/BRANCH_PROTECTION.md`)
**Purpose**: Document recommended GitHub branch protection rules

**Recommendations**:
- Require PR reviews on `internal`
- Require status checks to pass
- Restrict direct pushes to admins only
- Prevent force pushes and branch deletion
- Require conversation resolution

**Setup**: Manual configuration in GitHub repository settings.

### 5. **Makefile Targets**
**Purpose**: Make correct workflow easy and discoverable

**Targets**:
```bash
make setup-hooks  # Install git hooks
make release VERSION=v0.0.4  # Create a release
```

### 6. **Pull Request Template** (`.github/pull_request_template.md`)
**Purpose**: Remind developers about release process in every PR

**Features**:
- Checklist for regular PRs
- Special checklist for release PRs
- Instructions for using release script
- Warning about branch requirements

### 7. **Comprehensive Documentation**
- [RELEASING.md](RELEASING.md) - Complete release guide
- [README.md](README.md) - Quick reference in main docs
- [.github/BRANCH_PROTECTION.md](.github/BRANCH_PROTECTION.md) - Protection rules

## Defense in Depth

These facilities work together to create multiple layers of protection:

```
Developer Action → Pre-push Hook → GitHub Actions → Branch Protection → Code Review
     (Local)          (Local)         (Remote)          (Remote)        (Human)
```

If one fails, the others catch mistakes.

## Developer Workflow

### For New Contributors
```bash
# First time setup
git clone <repo>
cd go-athenahealth
make setup-hooks
```

### For Regular Development
```bash
# Create feature branch from main
git checkout main
git pull origin main
git checkout -b feature/my-feature

# ... make changes ...

# Create PR to main
git push origin feature/my-feature
```

### For Releases
```bash
# Ensure internal is synced with main
git checkout internal
git pull origin internal

# Merge changes from main (via PR ideally)
# Then tag the release
make release VERSION=v0.0.4
```

## Monitoring

You can verify these safeguards are working by:

1. **Check hook installation**: `ls -la .git/hooks/pre-push`
2. **Test hook locally**: Try to push a tag from main (should fail)
3. **View GitHub Actions**: Check workflow runs on tag pushes

## Maintenance

### When Onboarding New Developers
- Have them run `make setup-hooks` immediately
- Walk through release process
- Show them [RELEASING.md](RELEASING.md)

### When Changing Release Process
- Update all related documentation
- Notify team of changes
- Update this document

### Periodic Checks
- Verify GitHub Actions are enabled
- Confirm branch protection is configured
- Test release script with dry run functionality

## What If Something Goes Wrong?

### Tag Created from Wrong Branch
```bash
# Delete the tag locally and remotely
git tag -d v0.0.4
git push origin :refs/tags/v0.0.4

# Create it from the correct branch
git checkout internal
git tag -a v0.0.4 -m "Release v0.0.4"
git push origin v0.0.4
```

### Hook Not Working
```bash
# Reinstall hooks
make setup-hooks

# Verify installation
ls -la .git/hooks/pre-push
cat .git/hooks/pre-push  # Check contents
```

### GitHub Action Failing
- Check `.github/workflows/validate-tags.yaml`
- View workflow logs in GitHub Actions tab
- Ensure action has proper permissions

## Future Improvements

Potential additions:
- [ ] Automated sync workflow (main → internal)
- [ ] Release notes template generation
- [ ] Semantic version validation
- [ ] Changelog automation
- [ ] Tag signing enforcement
- [ ] Deployment automation after release

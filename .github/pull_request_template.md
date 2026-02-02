## Release Checklist

**Important**: All releases must be tagged from the `internal` branch.

Before creating your pull request, ensure:

- [ ] Changes have been tested locally
- [ ] All tests pass (`make test`)
- [ ] Code has been formatted (`make format`)
- [ ] Dependencies are up to date (`make tidy`)
- [ ] Documentation has been updated if needed

### For Release PRs (merging to `internal`):

- [ ] Version number follows semantic versioning (vX.Y.Z)
- [ ] CHANGELOG or release notes have been prepared
- [ ] Breaking changes are documented
- [ ] Migration guide provided (if applicable)

### Release Process:

After your PR is merged to `internal`:

```bash
# The automated way (recommended):
git checkout internal
git pull origin internal
make release VERSION=v0.0.4

# This will:
# ✓ Validate you're on the internal branch
# ✓ Check version format
# ✓ Create and push the tag
# ✓ Trigger automated workflows
```

**Note**: The release script and GitHub Actions will prevent tags from being created from the wrong branch.

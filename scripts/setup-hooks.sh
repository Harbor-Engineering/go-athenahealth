#!/bin/bash
# Setup git hooks for the repository

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
HOOKS_DIR="$REPO_ROOT/.git/hooks"

echo "Setting up git hooks..."

# Make pre-push hook executable
if [ -f "$HOOKS_DIR/pre-push" ]; then
    chmod +x "$HOOKS_DIR/pre-push"
    echo "✓ pre-push hook enabled"
else
    echo "⚠ pre-push hook not found at $HOOKS_DIR/pre-push"
fi

echo "✓ Git hooks setup complete"

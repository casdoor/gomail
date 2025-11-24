# Semantic Release Setup

This repository uses [semantic-release](https://github.com/semantic-release/semantic-release) for automated version management and package publishing.

## How it Works

When a pull request is merged to the `master` branch:

1. **Commit Analysis**: The semantic-release workflow analyzes commit messages using [Conventional Commits](https://www.conventionalcommits.org/) format
2. **Version Determination**: Based on the commit types, it determines the next version number:
   - `feat:` → Minor version bump (e.g., 2.1.0 → 2.2.0)
   - `fix:`, `perf:`, `docs:`, etc. → Patch version bump (e.g., 2.1.0 → 2.1.1)
   - `BREAKING CHANGE:` → Major version bump (e.g., 2.1.0 → 3.0.0)
   - `chore:` → No release
3. **Changelog Generation**: Automatically generates and updates `CHANGELOG.md`
4. **Git Tag Creation**: Creates a new git tag with the version (e.g., `v2.2.0`)
5. **GitHub Release**: Creates a GitHub release with the changelog

## Commit Message Format

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>
```

### Common Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Test changes
- `build`: Build system changes
- `ci`: CI configuration changes
- `chore`: Maintenance tasks (no release)

### Examples:

```bash
feat: add HTML email template support
fix: correct attachment MIME type detection
docs: update installation instructions
chore: update dependencies
```

### Breaking Changes:

For breaking changes, include `BREAKING CHANGE:` in the commit footer:

```bash
feat: redesign email API

BREAKING CHANGE: The Message.Send() method signature has changed.
Users must now pass a Dialer instance.
```

## Configuration Files

- `.releaserc.json`: Semantic release configuration
- `.github/workflows/release.yml`: GitHub Actions workflow for releases

## Manual Release

Releases are fully automated. Manual releases are not needed. Simply merge PRs with properly formatted commit messages.

## Skipping CI

To skip the CI workflow on release commits, the workflow automatically adds `[skip ci]` to release commit messages.

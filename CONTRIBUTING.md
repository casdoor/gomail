Thank you for contributing to Gomail! Here are a few guidelines:

## Bugs

If you think you found a bug, create an issue and supply the minimum amount
of code triggering the bug so it can be reproduced.


## Fixing a bug

If you want to fix a bug, you can send a pull request. It should contains a
new test or update an existing one to cover that bug.


## New feature proposal

If you think Gomail lacks a feature, you can open an issue or send a pull
request. I want to keep Gomail code and API as simple as possible so please
describe your needs so we can discuss whether this feature should be added to
Gomail or not.


## Commit Message Convention

This project uses [Conventional Commits](https://www.conventionalcommits.org/) for automated versioning and changelog generation.

Please format your commit messages as follows:

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types:
- `feat:` - A new feature (triggers a minor version bump)
- `fix:` - A bug fix (triggers a patch version bump)
- `perf:` - Performance improvements (triggers a patch version bump)
- `refactor:` - Code refactoring without changing functionality (triggers a patch version bump)
- `revert:` - Reverts a previous commit (triggers a patch version bump)
- `docs:` - Documentation changes (no release)
- `style:` - Code style changes (no release)
- `test:` - Adding or updating tests (no release)
- `build:` - Changes to build system or dependencies (no release)
- `ci:` - Changes to CI configuration (no release)
- `chore:` - Other changes that don't modify src or test files (no release)

### Breaking Changes:
Add `BREAKING CHANGE:` in the commit footer to trigger a major version bump.

### Examples:
```
feat: add support for custom headers
fix: resolve attachment encoding issue
docs: update README with new examples
```

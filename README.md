# DevOps - Rewriting WhoKnows_Variations
![Build](https://img.shields.io/github/actions/workflow/status/ASCDevOps/ASC_whoknows/GoCI.yml?branch=master&label=Build&style=for-the-badge)
![Go Version](https://img.shields.io/github/go-mod/go-version/ASCDevOps/ASC_whoknows?filename=go_app/backend/go.mod&style=for-the-badge)
![Release](https://img.shields.io/github/v/release/ASCDevOps/ASC_whoknows_variations?style=for-the-badge)
[![Go CI](https://github.com/ASCDevOps/ASC_whoknows/actions/workflows/GoCI.yml/badge.svg)](https://github.com/ASCDevOps/ASC_whoknows/actions/workflows/GoCI.yml)

[![DeepSource](https://app.deepsource.com/gh/ASCDevOps/ASC_whoknows.svg/?label=resolved+issues&show_trend=true&token=jVwdZ0qK59RhjbIma0eefs0R)](https://app.deepsource.com/gh/ASCDevOps/ASC_whoknows/)

---

## Project Structure (MonoRepo)
* 3 major folders:
1. Documentation
2. python_app
3. go_app

### Documentation
* Folders in documentation has numbers, those are week numbers.
* Each folder is named after the course structure, which follows each week.
* MD is used as format
* Folder has Challenges.md and Choices.md for each week

### Development / Setup
This project uses pre-commit hooks to enforce basic code hygiene and Go formatting before each commit.
Lightweight checks run locally, while full linting and testing are handled in CI.
* Run the following commands from the repository root.

* Windows (PowerShell)
Python must be installed

pip install pre-commit
pre-commit install

* macOS
Using Homebrew:

brew install pre-commit
pre-commit install

To update hooks:
pre-commit autoupdate

After installation, the hooks will automatically run on every git commit.
---

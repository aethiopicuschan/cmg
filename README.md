# cmg

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/cmg.svg)](https://pkg.go.dev/github.com/aethiopicuschan/cmg)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/cmg)](https://goreportcard.com/report/github.com/aethiopicuschan/cmg)
[![CI](https://github.com/aethiopicuschan/cmg/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/cmg/actions/workflows/ci.yaml)

`cmg` is a commit message generator based on git diff using an LLM.

## Installation

```sh
go install github.com/aethiopicuschan/cmg@latest
cmg
```

## Usage

```sh
❯ cmg -h
Commit message generator based on git diff using an LLM

Usage:
  cmg [flags]

Flags:
  -d, --details           include detailed commit body (multi-line commit message)
  -h, --help              help for cmg
  -i, --ignore-unstaged   ignore unstaged changes in the git diff
  -v, --version           version for cmg

❯ cmg
{Generated commit message with one line}

❯ cmg --details
{Generated commit message with detailed description}
```

## Configuration

When you run `cmg` for the first time, it creates a configuration file at `~/.config/cmg/config.json`.

## Supported LLMs

- [x] OpenAI
- [ ] Others

Other providers are planned and partially implemented.

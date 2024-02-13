# nethwv-cli

![nethwv-cli logo](https://github.com/necllmancer/nethwv-cli/assets/96694331/6f51a6d9-357e-467d-af08-0ab402833983)

## Overview
Nether Weave CLI (`nethwv`) is a versatile command-line interface tool designed for aggregating files from GitHub repositories or local directories into a consolidated PDF document. It supports fetching files from specific branches, tags, or directories within GitHub repositories, as well as compiling files from a local directory. This tool is ideal for creating documentation, archives, or project reviews in an accessible PDF format.

## Features
- Clone files from GitHub repositories, specifying branches, tags, or directories.
- Compile files from local directories into a PDF.
- Streamline documentation and review processes by consolidating multiple files into a single document.

## Installation
To install `nethwv`, ensure you have Go installed on your machine and follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the root directory of the project.
3. Run `make install` to install the CLI tool.

## Usage
The `nethwv` CLI tool supports two main commands: `git` for GitHub repositories and `local` for local directories.

### Fetching and Compiling from GitHub Repositories

```shell
nethwv git [options] <user/repo> <output.pdf>
```

Options for `git` command:
- `-b, --branch <branch-name>`: Specify a branch to clone.
- `-t, --tag <tag-name>`: Specify a tag to clone.
- `-d, --directory <directory-path>`: Fetch files from a specific directory in the repository.

### Compiling from Local Directories

```shell
nethwv local -p <path> <output.pdf>
```

Options for `local` command:
- `-p, --path <local-path>`: Specify the path to the local directory.

### Examples
To compile files from the master branch of the `psf/requests` GitHub repository into `output.pdf`:

```shell
nethwv git -b master psf/requests output.pdf
```

To compile files from a specific directory `src` within the `psf/requests` repository:

```shell
nethwv git -d src psf/requests output.pdf
```

To compile files from a local directory `/path/to/files` into `output.pdf`:

```shell
nethwv local -p /path/to/files output.pdf
```

## Building from Source
To build `nethwv` from source:

1. Clone the repository.
2. Navigate to the root directory of the project.
3. Run `make build` to compile the binary. The binary will be located in the `./build` directory.

This updated README provides comprehensive instructions for both the `git` and `local` commands, including usage examples for each, to better guide users in utilizing the `nethwv-cli` tool effectively.
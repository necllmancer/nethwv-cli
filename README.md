# nethwv-cli

![image](https://github.com/necllmancer/nethwv-cli/assets/96694331/6f51a6d9-357e-467d-af08-0ab402833983)

## Overview
Nether Weave CLI is an enhanced command-line interface tool that retrieves files from a specified GitHub repository, branch, or tag and compiles them into a single PDF document. It's ideal for documenting projects, creating archives, or consolidating repository contents for easier review. The tool efficiently handles files across various directories in the repository and allows specifying a particular directory.

## Introduction
`nethwv` is designed to fetch files from a GitHub repository, with options to specify a particular branch, tag, or directory, and generate a PDF containing these files. This tool streamlines the process of aggregating and converting repository content into a more accessible format, providing more control over the content being compiled.

## Installation
To install `nethwv`, follow these steps:

1. Ensure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Run `make install` from the root directory of the project.

## Usage
`nethwv` now supports additional command-line options to specify a branch, tag, or directory. Use the following syntax:

```shell
nethwv [options] <repo-path> <output-pdf-file>
```

Options:
- `-b, --branch <branch-name>`: Specify a branch to clone.
- `-t, --tag <tag-name>`: Specify a tag to clone.
- `-d, --directory <directory-path>`: Fetch files from a specific directory in the repository.

- `<repo-path>`: The path of the GitHub repository in the format `user/repo`.
- `<output-pdf-file>`: The name of the output PDF file.

For example, to fetch files from the master branch of the `psf/requests` repository and generate a PDF named `output.pdf`:

```shell
nethwv -b master psf/requests output.pdf
```

Or, to fetch files from a specific directory `src` in a repository:

```shell
nethwv -d src psf/requests output.pdf
```

## Downloading the README

To download the README file:

1. Visit the GitHub repository page at `https://github.com/necllmancer/nethwv-cli`.
2. Navigate to the `README.md` file.
3. Click on the file to view it.
4. Click the "Raw" button to view the raw file.
5. Right-click and choose "Save as" to download the file.

## Building from Source

To build nethwv from source:
1. Clone the repository.
2. Navigate to the root directory of the project.
3. Run `make build` to build the binary. The binary will be located in the `./build` directory.

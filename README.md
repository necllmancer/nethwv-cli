# nethwv-cli

![image](https://github.com/necllmancer/nethwv-cli/assets/96694331/8b003a9a-8f44-4bc8-87ce-30a50b0b73bb)

## Overview
Nether Weave CLI is a command-line interface tool that retrieves all files from a GitHub repository and compiles them into a single PDF document. Ideal for documenting projects, creating archives, or consolidating repository contents for easier review, it efficiently handles files across various directories in the repository.

## Introduction
`nethwv` is a command-line interface tool designed to fetch files from a specified GitHub repository and generate a PDF containing these files. It streamlines the process of aggregating and converting repository content into a more accessible format.

## Installation
To install `nethwv`, follow these steps:

1. Ensure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Run `make install` from the root directory of the project.

## Usage
To use `nethwv`, run the following command:

```shell
nethwv <repo-path> <output-pdf-file>
```

- <repo-path>: The path of the GitHub repository in the format user/repo.
- <output-pdf-file>: The name of the output PDF file.

For example:
```shell
nethwv psf/requests output.pdf
```

This will fetch files from the psf/requests GitHub repository and generate a PDF named output.pdf.

## Building from Source

To build nethwv from source:
1. Clone the repository.
2. Navigate to the root directory of the project.
3. Run make build to build the binary. The binary will be located in the ./build directory.

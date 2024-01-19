# Get Wiki Data Project

## Overview

The `get-wiki-data` utility functions are designed to download Wikipedia's data dumps.
It focuses on parsing Wikipedia dump status files and downloading the data.
The project also handles downloading of page view data.

## Features

- **Dump Jobs Processing**: Parses dump status files from Wikipedia to identify
  available data dumps.
- **Page View Data Handling**: Parses the page view data, the complete versions.
- **Data Downloading**: Automates the downloading of identified data dumps.

## Components

- `JobStatus`: Struct that represents the status of Wikimedia dump jobs,
  including details like file URLs and job completion status.
- `PageViewDump`: Struct designed to manage page view data, organizing
  information about available page view files.

## Running the Project

To use this project:

1. Ensure Go is installed on your system.
2. Clone this repository and navigate to the `get-wiki-data` project directory.
3. Use the provided structs and functions in your Go programs to download dumps.

## Project Goals

Created to download data wikipedia datasets for other, larger projects.

## Contributions

Suggestions, improvements, and contributions are welcome to enhance the project's
functionality and efficiency.

# Scala Notes Graph Project

## Overview

This project, developed in Scala, is focused on creating a graph structure from
a collection of markdown notes. The primary functionality includes parsing files
to extract internal and external links and constructing a graph that represents
these connections.

## Features

- Parses markdown files to identify links.
- Constructs a graph where each node represents a note, and edges represent links
  between notes.
- Outputs basic representations of the nodes and edges, providing a visual
  insight into the connections among notes.

## Project Structure

- `FileManager`: Handles file operations, reading markdown files from a
  specified directory.
- `LinkFinder`: Extracts internal and external links from the files.
- `Graph`: Constructs the graph based on the links found in the notes.
- `Node` and `Edge`: Represent the nodes and edges of the graph.

## Running the Project

To run this project:

1. Ensure Scala and sbt are installed on your system.
2. Clone this repository and navigate to the project directory.
3. Update the `path` variable in `run` method to point to your markdown notes directory.
4. Execute `sbt run` in the terminal to start the process.

## Final Notes

This project is complete and served as a valuable learning experience in refreshing some Scala basics and exploring graph theory.

# ID Generators

## What is it?

A Go-based project that comprises a set of classes for generating unique IDs.
This project focuses on two distinct algorithms for ID generation:

- [Twitter's Snowflake][1]: A system for generating unique, time-sequential IDs.
- [Flickr's Ticket Server][2]: An approach for creating unique, sequential
  integer IDs.

## Why This Project?

This project was done for fun and as an educational exercise. It allowed me to
delve into unique ID generation methods beyond the conventional UUIDs, exploring
algorithms that yield IDs which are semi-sortable by time and maintain uniqueness.

## How it Works

### Twitter's Snowflake Algorithm

- Utilizes timestamp, machine ID, and sequence number to generate 64-bit unique IDs.
- Configurable parameters for different parts of the ID.

### Flickr's Ticket Server

- Generates sequential integer IDs using a simple, efficient method.
- Implements a channel-based mechanism for ID generation in Go.

## Getting Started

### Prerequisites

- Ensure Go is installed on your system.

### Installation and Usage

- Clone the repository.
- Navigate to the `id-generators` directory.
- Use the `SnowflakeGenerator` or `TicketGenerator` as required in your Go application.

## Project Structure

- **generators package**: Contains the core implementations of both ID
  generation systems.
- **SnowflakeGenerator**: Implements the Snowflake algorithm.
- **TicketGenerator**: Implements the Ticket server algorithm.
- **Generator interface**: Defines the basic structure for the ID generators.

## License

[1]: https://en.wikipedia.org/wiki/Snowflake_ID
[2]: https://code.flickr.net/2010/02/08/ticket-servers-distributed-unique-primary-keys-on-the-cheap/

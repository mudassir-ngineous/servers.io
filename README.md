# Servers.io

A collection of servers implemented in different languages and technologies: Go, Rust, and WebAssembly, more will be added in near future.

## Project Structure

- `/stubs/infra` - Contains infrastructure setup files (docker-compose, pre-setup scripts)
- `/golang` - Contains the Go web server implementation
- `/rust` - Contains the Rust web server implementation
- `/web-assembly` - Contains the WebAssembly server implementation

## Development Setup

### Prerequisites

- git
- docker
- docker-compose

### Setup Instructions

1. Run the pre-setup script `/stubs/infra/pre-setup.sh` to automatically clone the repository and create the shared data directory
2. Refer `stubs/infra/docker-compose.yml`'s inline doc to use it and setup the docker environment

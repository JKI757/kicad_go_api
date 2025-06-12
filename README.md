# KiCad API Go Protobuf Generator

This repository is a build tool that generates Go bindings for the KiCad API protobuf definitions. Te original protobuf definitions come from the main [KiCad project](https://github.com/KiCad/kicad-source-mirror). This tool is designed to make it easier to use KiCad's API from Go applications by providing properly configured Go protobuf bindings.

## Overview

The KiCad project defines its API using Protocol Buffers, but the protobuf files are primarily configured for C++ usage. This project:

1. Takes the protobuf definitions from KiCad
2. Adds necessary Go-specific configurations (like `go_package` options)
3. Generates and organizes the Go code for easy consumption by Go applications

If you're looking to interact with KiCad from Go code, this provides the necessary bindings to do so.

## Project Structure

- `/proto` - Contains the protobuf definition files
  - `/board` - Board-related protobuf definitions
  - `/common` - Common types and utilities
  - `/schematic` - Schematic-related protobuf definitions
- `/go/api` - Generated Go code from protobuf definitions
- `/scripts` - Build and maintenance scripts

## Getting Started

### Prerequisites

- KiCad source code (containing the protobuf definitions)
- CMake (3.10 or higher)
- Protocol Buffers compiler (`protoc`)
- Go 1.16 or higher
- `protoc-gen-go` plugin (automatically installed by the build script)

### Setup

1. Clone the KiCad repository or copy its protobuf definitions into the `/proto` directory
2. Ensure the protobuf files maintain their original directory structure (e.g., `/proto/board`, `/proto/common`, etc.)

### Building

To generate the Go code from the protobuf definitions, run:

```bash
./scripts/generate_proto_go.py
```

This script will:
1. Install the required protoc-gen-go plugin if needed
2. Update go_package options in proto files
3. Generate Go code using CMake
4. Copy the generated files to `/go/api`

### Usage

To use these bindings in your Go project, first install the package:

```bash
go get github.com/JKI757/kicad_go_api
```

Then import the packages you need in your Go code:

Example import paths:
```go
import (
    "github.com/JKI757/kicad_go_api/go/api/board"
    "github.com/JKI757/kicad_go_api/go/api/common"
    "github.com/JKI757/kicad_go_api/go/api/schematic"
)
```

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

When adding new protobuf definitions:
1. Add the .proto file to the appropriate directory under `/proto`
2. Run `./scripts/generate_proto_go.py` to update the generated code
3. Commit both the .proto file and the generated Go code

#!/usr/bin/env python3

import os
import re
import sys
import shutil
import argparse
import subprocess
from pathlib import Path

def find_proto_files(base_dir):
    """Recursively find all .proto files in the given directory."""
    return list(Path(base_dir).rglob("*.proto"))

def get_go_package_path(proto_file, base_dir):
    """Generate the go_package path based on the proto file's location."""
    rel_path = os.path.relpath(os.path.dirname(proto_file), base_dir)
    return f"github.com/kicad/proto/{rel_path}"

def update_proto_file(proto_file, base_dir):
    """Update a single proto file to include or update the go_package option."""
    with open(proto_file, 'r') as f:
        content = f.read()

    # Check if package declaration exists
    package_match = re.search(r'^package\s+([^;]+);', content, re.MULTILINE)
    if not package_match:
        print(f"Warning: No package declaration found in {proto_file}")
        return False

    # Generate go_package path
    go_package = get_go_package_path(proto_file, base_dir)

    # Check if go_package option already exists
    go_package_match = re.search(r'^option\s+go_package\s*=\s*"([^"]+)";', content, re.MULTILINE)
    
    if go_package_match:
        # Update existing go_package
        if go_package_match.group(1) != go_package:
            content = re.sub(
                r'option\s+go_package\s*=\s*"[^"]+";',
                f'option go_package = "{go_package}";',
                content
            )
            print(f"Updated go_package in {proto_file}")
    else:
        # Add new go_package option after the package declaration
        package_end = package_match.end()
        content = (
            content[:package_end] + 
            f'\noption go_package = "{go_package}";' +
            content[package_end:]
        )
        print(f"Added go_package to {proto_file}")

    with open(proto_file, 'w') as f:
        f.write(content)

    return True

def run_command(cmd, cwd=None):
    """Run a command and return its output and success status."""
    try:
        result = subprocess.run(
            cmd,
            cwd=cwd,
            check=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True
        )
        return True, result.stdout
    except subprocess.CalledProcessError as e:
        return False, f"Error: {e.stderr}"

def ensure_protoc_go_plugin():
    """Ensure the protoc-gen-go plugin is installed."""
    cmd = ["go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest"]
    success, output = run_command(cmd)
    if not success:
        print("Failed to install protoc-gen-go:")
        print(output)
        return False
    return True

def copy_generated_files(project_dir):
    """Copy generated Go files from build directory to top-level go/api directory."""
    build_api_dir = os.path.join(project_dir, "build/go/api")
    dest_api_dir = os.path.join(project_dir, "go/api")

    if not os.path.exists(build_api_dir):
        print("No generated files found in build directory")
        return False

    # Remove existing files if any
    if os.path.exists(dest_api_dir):
        shutil.rmtree(dest_api_dir)

    # Create destination directory
    os.makedirs(os.path.dirname(dest_api_dir), exist_ok=True)

    # Copy files
    print("Copying generated files to go/api directory...")
    shutil.copytree(build_api_dir, dest_api_dir)
    return True

def generate_protos(project_dir):
    """Generate Go code from proto files using CMake."""
    build_dir = os.path.join(project_dir, "build")
    
    # Create build directory if it doesn't exist
    os.makedirs(build_dir, exist_ok=True)
    
    # Run cmake
    print("Running cmake...")
    success, output = run_command(["cmake", ".."], build_dir)
    if not success:
        print("CMake configuration failed:")
        print(output)
        return False

    # Run make
    print("\nRunning make...")
    success, output = run_command(["make"], build_dir)
    if not success:
        print("Make failed:")
        print(output)
        return False

    # Copy generated files to top-level directory
    if not copy_generated_files(project_dir):
        return False

    print("\nProto generation completed successfully")
    return True

def main():
    parser = argparse.ArgumentParser(
        description='Generate Go code from proto files, including package updates and compilation'
    )
    parser.add_argument('--project-dir', 
                      default=os.path.dirname(os.path.dirname(os.path.abspath(__file__))),
                      help='Project root directory (default: parent of script directory)')
    parser.add_argument('--only-update-packages', action='store_true',
                      help='Only update go_package options without generating code')
    parser.add_argument('--skip-go-plugin-install', action='store_true',
                      help='Skip installing protoc-gen-go plugin')
    args = parser.parse_args()

    project_dir = os.path.abspath(args.project_dir)
    base_dir = os.path.join(project_dir, 'proto')
    
    if not os.path.isdir(base_dir):
        print(f"Error: {base_dir} is not a directory")
        return 1

    # Step 1: Install protoc-gen-go if needed
    if not args.skip_go_plugin_install:
        print("Ensuring protoc-gen-go is installed...")
        if not ensure_protoc_go_plugin():
            return 1

    # Step 2: Update proto files
    proto_files = find_proto_files(base_dir)
    if not proto_files:
        print(f"No .proto files found in {base_dir}")
        return 1

    print(f"\nFound {len(proto_files)} .proto files")
    success_count = 0
    for proto_file in proto_files:
        if update_proto_file(proto_file, base_dir):
            success_count += 1

    print(f"Successfully processed {success_count} out of {len(proto_files)} files")

    # Step 3: Generate Go code
    if not args.only_update_packages:
        print("\nGenerating Go code from proto files...")
        if not generate_protos(project_dir):
            return 1

    return 0

if __name__ == '__main__':
    exit(main())

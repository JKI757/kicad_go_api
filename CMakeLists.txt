cmake_minimum_required(VERSION 3.10)
project(kicad_api_go)

# Find required protobuf package
find_package(Protobuf REQUIRED)

# Define proto files
set(PROTO_FILES
    common/envelope.proto
    common/types/base_types.proto
    common/types/enums.proto
    common/types/project_settings.proto
    common/commands/base_commands.proto
    common/commands/editor_commands.proto
    common/commands/project_commands.proto
    board/board.proto
    board/board_commands.proto
    board/board_types.proto
    schematic/schematic_types.proto
    schematic/schematic_commands.proto
)

# Set up output directories
set(GO_OUT_DIR ${CMAKE_CURRENT_BINARY_DIR}/go/api)

# Create a list of output files for Go
foreach(PROTO ${PROTO_FILES})
    string(REGEX REPLACE "[.]proto$" ".pb.go" GO_OUT ${PROTO})
    list(APPEND GO_OUT_FILES ${GO_OUT_DIR}/${GO_OUT})
    list(APPEND PROTO_PATHS ${CMAKE_CURRENT_SOURCE_DIR}/proto/${PROTO})
endforeach()

# Generate Go protobuf files
add_custom_command(
    OUTPUT ${GO_OUT_FILES}
    COMMAND ${CMAKE_COMMAND} -E make_directory ${GO_OUT_DIR}
    COMMAND ${Protobuf_PROTOC_EXECUTABLE}
        --go_out=${GO_OUT_DIR}
        --go_opt=paths=source_relative
        --proto_path=${CMAKE_CURRENT_SOURCE_DIR}/proto
        --experimental_allow_proto3_optional
        ${PROTO_FILES}
    DEPENDS ${PROTO_PATHS}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/proto
    COMMENT "Generating Go protobuf files"
)

# Add a custom target that depends on the generated files
add_custom_target(go_proto ALL DEPENDS ${GO_OUT_FILES})

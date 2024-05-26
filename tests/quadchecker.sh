#!/bin/bash

# List of Go files to build
GO_FILES=("quadA.go" "quadB.go" "quadC.go" "quadD.go" "quadE.go")

# Loop through each Go file and build it
for file in "${GO_FILES[@]}"
do
    echo "Building $file..."
    go build "$file"
    if [ $? -eq 0 ]; then
        echo "$file built successfully."
        # Grant execute permissions
        chmod +x "${file%.*}" # Remove the file extension
    else
        echo "Error building $file."
        exit 1
    fi
done

echo "All files built successfully and executable permissions granted."

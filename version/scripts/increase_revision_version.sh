#!/bin/bash

constant_file="version/version.go"

# Get the current version
revision_value=$(sed -n "s/const RevisionVersion = //p" "${constant_file}")

# Increase by one
new_value=$((revision_value + 1))

# Replace the old constants.go to a new one with new value.
sed "s/const RevisionVersion = $revision_value/const RevisionVersion = $new_value/g" "$constant_file" > tempfile && mv tempfile "$constant_file"
echo "Revision version has been increased to $new_value"

major_value=$(sed -n "s/const MajorVersion = //p" "${constant_file}")
minor_value=$(sed -n "s/const MinorVersion = //p" "${constant_file}")
revision_value=$(sed -n "s/const RevisionVersion = //p" "${constant_file}")
echo "The version has been updated to $major_value.$minor_value.$revision_value"

# Set constants.go file to be a staged file
git add version/version.go
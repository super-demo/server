#!/bin/bash

constant_file="version/version.go"

# Get the current version
major_value=$(sed -n "s/const MajorVersion = //p" "${constant_file}")

# Increase by one
new_value=$((major_value + 1))

# Replace the old constants.go to a new one with new value.
sed "s/const MajorVersion = $major_value/const MajorVersion = $new_value/g" "$constant_file" > tempfile && mv tempfile "$constant_file"
echo "Major version has been increased to $new_value"

# Reset minor number to 0 if it needed
minor_value=$(sed -n "s/const MinorVersion = //p" "${constant_file}")
if [ "$minor_value" -gt 0 ]; then
    sed "s/const MinorVersion = $minor_value/const MinorVersion = 0/g" "$constant_file" > tempfile && mv tempfile "$constant_file"
    echo "Minor version has been reset to 0"
fi

# Reset revision number to 0 if it needed
revision_value=$(sed -n "s/const RevisionVersion = //p" "${constant_file}")
if [ "$revision_value" -gt 0 ]; then
    sed "s/const RevisionVersion = $revision_value/const RevisionVersion = 0/g" "$constant_file" > tempfile && mv tempfile "$constant_file"
    echo "Revision version has been reset to 0"
fi

major_value=$(sed -n "s/const MajorVersion = //p" "${constant_file}")
minor_value=$(sed -n "s/const MinorVersion = //p" "${constant_file}")
revision_value=$(sed -n "s/const RevisionVersion = //p" "${constant_file}")
echo "The version has been updated to $major_value.$minor_value.$revision_value"

# Set constants.go file to be a staged file
git add version/version.go
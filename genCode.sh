# Copyright (c) 2014 Aubrey Barnard.  This is free software.  See
# LICENSE for details.

# Generates the code for this package given template files

# Check command line arguments
if [[ $# -ne 1 ]]; then
    echo "$0: Error: Incorrect command line arguments." >&2
    echo "$0: Usage: <all-types-file>" >&2
    exit 2
fi

# Get the file names from the command line
allTypesFile="$1"

# Golang comparable numeric types
intTypes="uint8 uint16 uint32 uint64 uint int8 int16 int32 int64 int"
fltTypes="float32 float64"
numTypes="$intTypes $fltTypes"

# Output the documentation and package declaration
cat <<EOF
// Copyright (c) 2014 Aubrey Barnard.  This is free software.  See
// LICENSE for details.

package ghostlice

EOF
# Output the code for all types
for numType in $numTypes; do
    # Make an uppercase version of the type for use in names
    upNumType=${numType^}
    # Replace all occurrences of "AType" with the upper- or lower-case
    # version of the type name as appropriate.  The distinction is
    # whether the type name follows what could be an identifier.  Remove
    # the header lines from the input so as not to repeat them in the
    # output.
    tail -n +9 "$allTypesFile" | sed -e "s/\([^[:alnum:]_]\)AType/\1$numType/g" -e "s/AType/$upNumType/g"
done
# Output the code specific to int types (none yet) (intTypes.go)
# Output the code specific to float types (none yet) (fltTypes.go)

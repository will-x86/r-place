#!/bin/sh
targets=("c" "clojure" "crystal" "csharp" "go" "http" "java" "javascript" "kotlin" "node" "objc" "ocaml" "php" "powershell" "python" "r" "ruby" "rust" "shell" "swift")

if [ ! -d "examples" ]; then
    mkdir examples
fi

for target in "${targets[@]}"
do
    if [ ! -d "examples/$target" ]; then
        mkdir "examples/$target"
    fi
    npx httpsnippet har.json --target "$target" --output "./examples/$target"
done

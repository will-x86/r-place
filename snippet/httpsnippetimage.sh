#!/bin/sh
targets=("c" "clojure" "crystal" "csharp" "go" "http" "java" "javascript" "kotlin" "node" "objc" "ocaml" "php" "powershell" "python" "r" "ruby" "rust" "shell" "swift")

if [ ! -d "image-examples" ]; then
    mkdir image-examples
fi

for target in "${targets[@]}"
do
    if [ ! -d "image-examples/$target" ]; then
        mkdir "image-examples/$target"
    fi
    httpsnippet har.json --target "$target" --output "./image-examples/$target"
done

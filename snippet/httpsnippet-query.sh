#!/bin/sh
targets=("c" "clojure" "crystal" "csharp" "go" "http" "java" "javascript" "kotlin" "node" "objc" "ocaml" "php" "powershell" "python" "r" "ruby" "rust" "shell" "swift")

if [ ! -d "query-examples" ]; then
    mkdir query-examples
fi

for target in "${targets[@]}"
do
    if [ ! -d "query-examples/$target" ]; then
        mkdir "query-examples/$target"
    fi
    npx httpsnippet har-get.json --target "$target" --output "./query-examples/$target"
done

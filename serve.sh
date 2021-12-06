#!/bin/bash
set -e

git ls-files | entr -r -s 'npx tsc --noEmit && npm run build && go run main.go'

# Consinstent formatting with go fmt

1. Print reformatted file to stdout:

        gofmt main.go

2. Test that all files in package conform to gofmt's standards (useful for CI pipeline):

        test $(gofmt -l . | wc -l) -eq 0

3. Rewrite all files:

        gofmt -w -l .

#!/bin/sh

./tailwindcss -i ./static/css/input.css -o ./static/css/output.css
templ generate
go run main.go

# F5XC Public IP Ranges exporter
This tool takes public F5XC IP ranges from public F5 Xc documentation and generates HCL variables output for later on
usage in TF templates.

## Usage

`go run main.go` will create output file `variables.tf` with public IP prefix per continent.

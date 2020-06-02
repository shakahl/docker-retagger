# docker-retagger

This tool allows you to pull, re-tag, and push the re-tagged docker images on either a mass scale, or as a one-off operation.

## Usage

```bash
Usage of retagger:
  -image string
        the original image to pull
  -infile string
        use an input file rather than one-off flags
  -new-origin string
        the origin to set the image to
```

| flag         | type   | required                                               | help                                                                                  | example                  |
|--------------|--------|--------------------------------------------------------|---------------------------------------------------------------------------------------|--------------------------|
| --image      | string | only if --infile is not used                           | The original image to pull                                                            | `golang:latest`          |
| --new-origin | string | only if --infile is not used                           | The new repository you are pushing to                                                 | `my.awesome.docker.repo` |
| --infile     | string | no, but takes preference over --image and --new-origin | A file with the original image and the new origin seperated by a space (one per line) | `infile.txt`             |

## Installation

*Option 1:*

1) Clone this repo
2) run `make && make install`

This will install a binary called `retagger` into your `/usr/local/bin` directory, so if you need to, make sure your path is updated with this directory:

```bash
cat << EOF >> ~/.bashrc && source ~/.bashrc
export PATH=$PATH:/usr/local/bin
EOF
```

*Option 2:*

Just grab the binary from the releases section here :boom:

*Windows Folks*

You will need to grab the zip file from the releases and put the binary wherever you feel the need :smile:

### Disclaimer

I am well aware this is a poorly constructed CLI :smile:, but it worked in a pinch. I'll try to fix it later.
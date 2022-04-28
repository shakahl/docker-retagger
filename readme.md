# docker-retagger

This tool allows you to pull, re-tag, and push the re-tagged docker images on either a mass scale, or as a one-off operation.

## Usage

| flag          | type   | required                              | help                                                    | example                  |
|---------------|--------|---------------------------------------|---------------------------------------------------------|--------------------------|
| --image       | string | only if --file is not used            | The original image to pull                              | `golang:latest`          |
| --newregistry | string | only if --file is used                | The new repository you are pushing to                   | `my.awesome.docker.repo` |
| --file        | string | no, but takes preference over --image | A file with the original images (one per line)          | `images.list.template`   |
| --skiplogin   | bool   | no                                    | if you want continue without login in original registry | `true / false`           |

## Installation

**Option 1:**

1) Clone this repo
2) run `make && make install`

This will install a binary called `retagger` into your `/usr/local/bin` directory, so if you need to, make sure your path is updated with this directory:

```bash
cat << EOF >> ~/.bashrc && source ~/.bashrc
export PATH=$PATH:/usr/local/bin
EOF
```

**Option 2:**

Just grab the binary from the releases section here :boom:

**Windows Folks:**

You will need to grab the zip file from the releases and put the binary wherever you feel the need :smile:

### Disclaimer

I am well aware this is a poorly constructed CLI :smile:, but it worked in a pinch. I'll try to fix it later.

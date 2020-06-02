package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Image struct {
	Registry string
	User string
	Name string
	Tag string
	Hash string
}

func (I *Image) Marshal() (string, error) {
	if I.User != "" {
		if I.Tag != "" {
			return fmt.Sprintf("%s/%s/%s:%s", I.Registry, I.User, I.Name, I.Tag), nil
		} else if I.Hash != "" {
			return fmt.Sprintf("%s/%s/%s@%s", I.Registry, I.User, I.Name, I.Hash), nil
		} else {
			return "", fmt.Errorf("no Tag or Hash value found")
		}
	} else {
		if I.Tag != "" {
			return fmt.Sprintf("%s/%s:%s", I.Registry, I.Name, I.Tag), nil
		} else if I.Hash != "" {
			return fmt.Sprintf("%s/%s@%s", I.Registry, I.Name, I.Hash), nil
		} else {
			return "", fmt.Errorf("no Tag or Hash value found")
		}
	}
}

func parseImage(image string) (*Image, error) {
	img := &Image{}
	r := regexp.MustCompile("^(((?P<Registry>[a-zA-Z0-9-_]+?(\\.[a-zA-Z0-9]+?)+?\\.[a-zA-Z]{2,})/)?((?P<Name>[a-zA-Z0-9-_]+?)|(?P<UserName>[a-zA-Z0-9-_]+?)/(?P<ImageName>[a-zA-Z-_]+?))((:(?P<Tag>[a-zA-Z0-9-_\\.]+?))|(@(?P<Hash>sha256:[a-z0-9]{64}))))$")
	if r.MatchString(image) {
		matches := r.FindStringSubmatch(image)
		img.Registry = strings.TrimSuffix(matches[2], "/")
		if matches[4] != "" {
			img.Name = strings.TrimPrefix(matches[5], "/")
		} else {
			img.User = strings.TrimPrefix(matches[6], "/")
			img.Name = strings.TrimPrefix(matches[7], "/")
		}
		if matches[9] != "" {
			img.Tag = strings.TrimPrefix(matches[10], ":")
		}
		if matches[11] != "" {
			img.Hash = matches[12]
		}
		return img, nil
	} else {
		return nil, fmt.Errorf("%s does not appear to be a valid container image: %+v", image, img)
	}
}
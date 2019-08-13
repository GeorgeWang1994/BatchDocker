package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

// batch remove image
func (d *Docker) BatchRemoveImage(imageIDs []string) []error {
	var errList []error
	for _, imageID := range imageIDs {
		err := d.RemoveImage(imageID)
		if err != nil {
			errList = append(errList, err)
		}
	}
	return errList
}

// batch get image's info
func (d*Docker) BatchInspectImage(imageIDs []string) []types.ImageInspect {
	var images []types.ImageInspect
	for _, imageID := range imageIDs {
		image, _, err := d.ImageInspectWithRaw(context.TODO(), imageID)
		if err != nil {
			images = append(images, image)
		}
	}
	return images
}

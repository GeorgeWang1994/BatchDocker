package docker

import (
	"bufio"
	"context"
	"github.com/docker/docker/api/types"
)

// image list
func (d *Docker) ListImage(opt types.ImageListOptions) ([]types.ImageSummary, error) {
	return d.ImageList(context.TODO(), opt)
}

// pull image
func (d *Docker) PullImage(name string) error {
	resp, err := d.ImagePull(context.TODO(), name, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	// block until read eof
	scanner := bufio.NewScanner(resp)
	for scanner.Scan() {}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return nil
}

// remove image
func (d *Docker) RemoveImage(imageID string) error {
	_, err := d.ImageRemove(context.TODO(), imageID, types.ImageRemoveOptions{})
	return err
}

// get info of image
func (d*Docker) InspectImage(imageID string) (types.ImageInspect, error) {
	res, _, err := d.ImageInspectWithRaw(context.TODO(), imageID)
	return res, err
}

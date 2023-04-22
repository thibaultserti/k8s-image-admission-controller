package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"k8s-image-admission-controller/pkg/k8simageadmissioncontroller"
)

var imageName = "alpine:3.17.3"

func TestGetManifest(t *testing.T) {
	manifest, err := k8simageadmissioncontroller.GetManifest(imageName)
	if err != nil {
		t.Fatalf("Cannot get manifest: %v", err)
	}
	got, err := json.Marshal(manifest)
	if err != nil {
		t.Errorf("Cannot marshall manifest in got, %v", err)
	}
	want, err := json.Marshal(map[string]interface{}{
		"config": map[string]interface{}{
			"digest":    "sha256:9ed4aefc74f6792b5a804d1d146fe4b4a2299147b0f50eaf2b08435d7b38c27e",
			"mediaType": "application/vnd.docker.container.image.v1+json",
			"size":      1470,
		},
		"layers": []map[string]interface{}{
			{
				"digest":    "sha256:f56be85fc22e46face30e2c3de3f7fe7c15f8fd7c4e5add29d7f64b87abdaa09",
				"mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
				"size":      3374563,
			},
		},
		"mediaType":     "application/vnd.docker.distribution.manifest.v2+json",
		"schemaVersion": 2,
	})
	if err != nil {
		t.Errorf("Cannot marshall manifest in want, %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("got %s want %s", string(got), string(want))
	}
}

func TestGetImageSize(t *testing.T) {
	got, err := k8simageadmissioncontroller.GetImageSize(imageName)
	if err != nil {
		t.Fatalf("Cannot get size: %v", err)
	}
	want := int64(3374563)
	if got != want {
		t.Errorf("got %d want %d given", got, want)
	}
}

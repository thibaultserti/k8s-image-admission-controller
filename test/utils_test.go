package test

import (
	"testing"

	"k8s-image-admission-controller/pkg/k8simageadmissioncontroller"
)

var imageName = "alpine:3.17.3"

func TestPullDockerImage(t *testing.T) {
	err := k8simageadmissioncontroller.PullDockerImage(imageName)
	if err != nil {
		t.Fatalf("Pull error on image : %v", err)
	}

}

func TestGetImageSize(t *testing.T) {

	expectedSize := int64(7_458_929) // Mettre la taille de l'image "alpine" attendue ici

	k8simageadmissioncontroller.PullDockerImage(imageName)

	// Appeler la fonction pour récupérer la taille de l'image
	size, err := k8simageadmissioncontroller.GetImageSize(imageName)
	if err != nil {
		t.Fatalf("Error when getting image size : %v", err)
	}

	// Vérifier que la taille récupérée correspond à la taille attendue
	if size != expectedSize {
		t.Errorf("Image size %s is incorrect. Wait : %d, Got : %d", imageName, expectedSize, size)
	}
}

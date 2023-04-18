package k8simageadmissioncontroller

import (
	"context"
	"fmt"

	logging "github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullDockerImage(imageName string) error {
	ctx := context.Background()

	// Créer un client Docker
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logging.Error("Échec de la création du client Docker")
	}

	// Préparer les options de pull
	options := types.ImagePullOptions{}
	resp, err := cli.ImagePull(ctx, imageName, options)
	if err != nil {
		logging.Error("échec du tirage de l'image Docker")
	}
	defer resp.Close()

	// Lire la sortie du tirage de l'image Docker
	output := make([]byte, 4096)
	for {
		n, err := resp.Read(output)
		if err != nil {
			break
		}
		fmt.Print(string(output[:n]))
	}

	return nil
}

func GetImageSize(imageName string) (int64, error) {
	// Créer un client Docker
	cli, err := client.NewClientWithOpts(client.WithVersion("1.41"))
	if err != nil {
		return 0, err
	}

	// Récupérer les informations de l'image
	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return 0, err
	}

	// Récupérer la taille de l'image à partir des informations de l'image
	size := int64(imageInspect.Size)

	return size, nil
}

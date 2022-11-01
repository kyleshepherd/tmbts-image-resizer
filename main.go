package imageresizer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
)

var (
	storageClient *storage.Client
)

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func init() {
	var err error

	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
}

func ResizeImages(ctx context.Context, e GCSEvent) error {
	inputBlob := storageClient.Bucket(e.Bucket).Object(e.Name)
	r, err := inputBlob.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("NewReader: %v", err)
	}

	outputBlob := storageClient.Bucket(os.Getenv("RESIZED_BUCKET_NAME")).Object(e.Name)
	w := outputBlob.NewWriter(ctx)
	defer w.Close()

	cmd := exec.Command("convert", "-", "-resize", "1000", "-")
	cmd.Stdin = r
	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cmd.Run: %v", err)
	}

	log.Printf("Resized image uploaded to gs://%s/%s\n", outputBlob.BucketName(), outputBlob.ObjectName())

	o := storageClient.Bucket(e.Bucket).Object(e.Name)

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", e.Name, err)
	}

	log.Printf("Blob %v deleted\n", e.Name)

	return nil
}

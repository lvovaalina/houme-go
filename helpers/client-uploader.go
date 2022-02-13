package helpers

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

const (
	projectID  = "eternal-concept-340014" // FILL IN WITH YOURS
	bucketName = "h-test-button"          // FILL IN WITH YOURS
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func NewClientUploader(googleFolder string) *ClientUploader {
	pwd, _ := os.Getwd()
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", pwd+"/configs/google-config.json") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: googleFolder + "/",
	}

}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Fatalln("Could not upload file: ", err)
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalln("Could not close writer:", err)
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (c *ClientUploader) CreateFile(fileName string, name string, email string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	bkt := c.cl.Bucket(c.bucketName)
	obj := bkt.Object(c.uploadPath + fileName + ".txt")
	// Write something to obj.
	// w implements io.Writer.
	w := obj.NewWriter(ctx)
	str := "Filename: " + fileName + " Name: " + name + " Email: " + email
	// Write some text to obj. This will either create the object or overwrite whatever is there already.
	if _, err := fmt.Fprintf(w, str); err != nil {
		// TODO: Handle error.
		log.Fatalln("Could not close writer:", err)
		return fmt.Errorf("Writer.Close: %v", err)
	}
	// Close, just like writing a file.
	if err := w.Close(); err != nil {
		// TODO: Handle error.
	}
	return nil
}

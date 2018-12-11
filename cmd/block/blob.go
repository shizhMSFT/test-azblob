package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// PrepareBlob prepares a blob.
func PrepareBlob() (*storage.Blob, error) {
	cli, err := storage.NewClientFromConnectionString(os.Args[1])
	if err != nil {
		return nil, err
	}

	blobCli := cli.GetBlobService()
	container := blobCli.GetContainerReference("testblock")
	if _, err := container.CreateIfNotExists(nil); err != nil {
		return nil, err
	}
	blob := container.GetBlobReference("hello.yaml")
	blob.Delete(nil)

	return blob, nil
}

// PrintBlob prints blob in text
func PrintBlob(blob *storage.Blob) error {
	content, err := blob.Get(nil)
	if err != nil {
		return err
	}
	defer content.Close()

	text, err := ioutil.ReadAll(content)
	if err != nil {
		return err
	}

	fmt.Println(string(text))
	return nil
}

// NewBlockList creates a new block list
func NewBlockList(IDs ...string) []storage.Block {
	var blocks []storage.Block
	for _, ID := range IDs {
		blocks = append(blocks, storage.Block{
			ID:     ID,
			Status: storage.BlockStatusLatest},
		)
	}
	return blocks
}

// PutBlockList put block list with IDs
func PutBlockList(blob *storage.Blob, IDs ...string) error {
	return blob.PutBlockList(NewBlockList(IDs...), nil)
}

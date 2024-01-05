package main

import (
	"context"
	"log"
	"os"

	"github.com/mholt/archiver/v4"
)

func main() {
	zf, err := os.Create("output.zip")
	if err != nil {
		log.Fatal("error creating zip file", err)
	}

	filesForZip, err := archiver.FilesFromDisk(&archiver.FromDiskOptions{
		FollowSymlinks: false,
	}, map[string]string{
		"toArchive": "toArchive",
	})
	if err != nil {
		log.Fatal("error walking files to create archive", err)

	}

	format := archiver.CompressedArchive{
		Archival: archiver.Zip{
			SelectiveCompression: true,
		},
	}

	// create the archive
	err = format.Archive(context.Background(), zf, filesForZip)
	if err != nil {
		log.Fatal("error creating archive", err)
	}
}

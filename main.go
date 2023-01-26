package main

import (
	"fmt"
	"log"
	"path/filepath"

	"gosync/internals/util"
	"gosync/internals/util/checksum"
)

func main() {
	manifestData, err := util.ReadJSONManifest("manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Starting %s v%s...\n", manifestData.Name, manifestData.Version)

	fmt.Printf("Reading files in directory: %s\n", manifestData.Directories.SourceDirectory)
	sourceFiles, err := checksum.GetFilesInDirectory(manifestData.Directories.SourceDirectory, false)
	if err != nil {
		log.Fatal(err)
	}

	// Evaluate hashes of all files in the source directory.
	sourceHashes := make(map[string]string)
	hashesRunner, err := checksum.GetHashesAlgorithm(manifestData.Settings.HashAlgorithm)
	if err != nil {
		log.Fatal(err)
	}
	hashesRunner(sourceHashes, sourceFiles)

	// Evaluate if the file exists in the destination directories, if not copy it from the source directory.
	// If the file exists verify the hash signature matches, if not then replace the file from the source directory.
	hashRunner, err := checksum.GetHashAlgorithm(manifestData.Settings.HashAlgorithm)
	if err != nil {
		log.Fatal(err)
	}
	updateFiles(hashRunner, sourceHashes, manifestData.Directories.DestinationDirectories)
}

func updateFiles(hashRunner func(sourceFile string) string, sourceHashes map[string]string, targetDirectories []string) {
	for _, dir := range targetDirectories {
		for path, hash := range sourceHashes {
			if util.DirectoryContainsFile(dir, filepath.Base(path)) {
				destFilePath := filepath.Join(dir, filepath.Base(path))
				verifyHash := hashRunner(destFilePath)
				if verifyHash != hash {
					if util.CopyFile(path, destFilePath) {
						fmt.Printf("File Change Detected: [%s, %v -> %v]\n", path, verifyHash, hash)
					}
				}
			} else {
				destFilePath := filepath.Join(dir, filepath.Base(path))
				if util.CopyFile(path, destFilePath) {
					fmt.Printf("File Copied: [%s -> %s]\n", path, destFilePath)
				}
			}
		}

	}
}

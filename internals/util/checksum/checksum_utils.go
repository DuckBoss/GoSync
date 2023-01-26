package checksum

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"gosync/internals/errors/checksumerrors"
	"gosync/internals/util"
	"hash/crc32"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesInDirectory(path string, names_only bool) ([]string, error) {
	if path == "" {
		return nil, &checksumerrors.ChecksumError{Err: errors.New("Source directory file path is not specified.")}
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, &checksumerrors.ChecksumError{Err: err}
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, &checksumerrors.ChecksumError{Err: err}
	}

	var sourceFiles = make([]string, 0, len(files))
	for _, file := range files {
		var filePath string
		if names_only {
			filePath = file.Name()
		} else {
			filePath = filepath.Join(absPath, file.Name())
		}
		sourceFiles = append(sourceFiles, filePath)
	}

	return sourceFiles, nil
}

func getCRC32FromFile(path string) (uint32, error) {
	if path == "" {
		return 0, &checksumerrors.ChecksumError{Err: errors.New("Checksum file path is not specified.")}
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return 0, &checksumerrors.ChecksumError{Err: err}
	}

	checksum, err := getCRC32(fileBytes)
	if err != nil {
		return 0, &checksumerrors.ChecksumError{Err: err}
	}

	return checksum, nil
}

func getCRC32(data []byte) (uint32, error) {
	if len(data) == 0 {
		return 0, &checksumerrors.ChecksumError{Err: errors.New("No data provided for CRC32 checksum computation.")}
	}
	return crc32.ChecksumIEEE(data), nil
}

func getSHA256FromFile(path string) (string, error) {
	if path == "" {
		return "", &checksumerrors.ChecksumError{Err: errors.New("Checksum file path is not specified.")}
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", &checksumerrors.ChecksumError{Err: err}
	}

	checksum, err := getSHA256(fileBytes)
	if err != nil {
		return "", &checksumerrors.ChecksumError{Err: err}
	}

	return checksum, nil
}

func getSHA256(data []byte) (string, error) {
	if len(data) == 0 {
		return "", &checksumerrors.ChecksumError{Err: errors.New("No data provided for SHA256 checksum computation.")}
	}
	bytes := sha256.Sum256(data)
	return fmt.Sprintf("%x", bytes), nil
}

func getFileHashesCRC32(sourceHashes map[string]string, sourceFiles []string) {
	for _, file := range sourceFiles {
		sourceHashes[file] = getFileHashCRC32(file)
		fmt.Printf("File checked: [%s, %v]\n", file, sourceHashes[file])
	}
}

func getFileHashCRC32(sourceFile string) string {
	crc32, err := getCRC32FromFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprint(crc32)
}

func getFileHashesSHA256(sourceHashes map[string]string, sourceFiles []string) {
	for _, file := range sourceFiles {
		sourceHashes[file] = getFileHashSHA256(file)
		fmt.Printf("File checked: [%s, %v]\n", file, sourceHashes[file])
	}
}

func getFileHashSHA256(sourceFile string) string {
	sha256, err := getSHA256FromFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	return sha256
}

func GetHashesAlgorithm(algorithm string) (func(sourceHashes map[string]string, sourceFiles []string), error) {
	algorithm = strings.ToLower(algorithm)
	if len(os.Args) > 1 {
		algorithm = util.GetSysArgs().HashAlgorithm
	}
	switch algorithm {
	case "crc32":
		return getFileHashesCRC32, nil
	case "sha256":
		return getFileHashesSHA256, nil
	default:
		return nil, errors.New("The selected hashing algorithm is not supported.")
	}
}

func GetHashAlgorithm(algorithm string) (func(sourceFile string) string, error) {
	algorithm = strings.ToLower(algorithm)
	switch algorithm {
	case "crc32":
		return getFileHashCRC32, nil
	case "sha256":
		return getFileHashSHA256, nil
	default:
		return nil, errors.New("The selected hashing algorithm is not supported.")
	}
}

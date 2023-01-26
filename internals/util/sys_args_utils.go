package util

import (
	"flag"
)

type SysArgs struct {
	HashAlgorithm string
	ScanInterval  int
}

func GetSysArgs() SysArgs {

	var sysArgsData SysArgs

	flag.StringVar(&sysArgsData.HashAlgorithm, "hash", "crc32", "The hashing algorithm to use.")
	flag.IntVar(&sysArgsData.ScanInterval, "interval", 5, "How frequently to scan the source directory for changes (in seconds).")
	flag.Parse()

	return sysArgsData
}

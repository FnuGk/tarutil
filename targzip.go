package tarutil

import (
	"archive/tar"
	"compress/gzip"
	"io"

	"github.com/pkg/errors"
)

func GZipToTar(inp io.Reader) (*tar.Reader, error) {
	gzf, err := gzip.NewReader(inp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tarReader := tar.NewReader(gzf)
	return tarReader, nil
}

var err error

// FileFunc is the function that should be called on files
type FileFunc func(name string, r io.Reader) error

// DoEachFile calls f for each file in the tarReader
func DoEachFile(tarReader *tar.Reader, f FileFunc) error {
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return errors.WithStack(err)
		}

		switch header.Typeflag {
		case tar.TypeDir: // = directory
		case tar.TypeReg: // = regular file
			r := io.Reader(tarReader)
			err = f(header.Name, r)
			if err != nil {
				return errors.WithStack(err)
			}
		default:
			return errors.Errorf("Unable to determine type %c in file %s", header.Typeflag, header.Name)
		}
	}
	return nil
}

package nborm

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type tempDir string

func NewTempDir(dir, prefix string) (*tempDir, error) {
	path, err := ioutil.TempDir(dir, prefix)
	if err != nil {
		return nil, err
	}
	tempDir := tempDir(path)
	return &tempDir, nil
}

func (d *tempDir) Remove() error {
	return os.RemoveAll(string(*d))
}

func (d *tempDir) CreateTempFile(pattern string) (*file, error) {
	f, err := ioutil.TempFile(string(*d), pattern)
	if err != nil {
		return nil, err
	}
	return fileForWriteFromFile(f), nil
}

func (d *tempDir) CopyFile(filepath string) (*file, error) {
	srcFile, err := newFileForRead(filepath)
	if err != nil {
		return nil, err
	}
	dstFile, err := d.CreateTempFile(fmt.Sprintf("%s_*.go", strings.Trim(filepath, ".go")))
	if err != nil {
		return nil, err
	}
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()
	reader, _ := srcFile.reader()
	writer, _ := dstFile.writer()
	if _, err := io.Copy(writer, reader); err != nil {
		return nil, err
	}
	newFile, err := dstFile.toWrite()
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

func (d *tempDir) concurrentCopy(filepath string, ctx context.Context, errChan chan<- error, fileChan chan<- *file) {
	srcFile, err := newFileForRead(filepath)
	if err != nil {
		errChan <- err
		return
	}
	dstFile, err := d.CreateTempFile(fmt.Sprintf("%s_*.go", strings.Trim(filepath, ".go")))
	if err != nil {
		errChan <- err
		return
	}
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()
	reader, _ := srcFile.reader()
	writer, _ := dstFile.writer()
	bufferSize := reader.Size()
	buffer := make([]byte, bufferSize)
	for {
		select {
		case <-ctx.Done():
			errChan <- errors.New("cancel")
			return
		default:
			n, err := reader.Read(buffer)
			if err != nil {
				if err != io.EOF {
					errChan <- err
					return
				}
				_, err := writer.Write(buffer[:n])
				if err != nil {
					errChan <- err
					return
				}
				newFile, err := dstFile.toRead()
				if err != nil {
					errChan <- err
					return
				}
				fileChan <- newFile
				errChan <- nil
				return
			}
			_, err = writer.Write(buffer[:n])
			if err != nil {
				errChan <- err
				return
			}
		}
	}
}

func (d *tempDir) CopyFiles(filepaths ...string) ([]*file, error) {
	fileChan := make(chan *file, len(filepaths))
	errChan := make(chan error, len(filepaths))
	ctx, cancel := context.WithCancel(context.Background())
	for _, filepath := range filepaths {
		go func(path string) {
			d.concurrentCopy(path, ctx, errChan, fileChan)
		}(filepath)
	}
	var numDone int
	fileList := make([]*file, len(filepaths))
	errStrList := make([]string, 0, len(filepaths))
	var errOccur bool
	for {
		e := <-errChan
		if e != nil {
			errOccur = true
			cancel()
			errStrList = append(errStrList, e.Error())
		} else {
			fileList[numDone] = <-fileChan
		}
		numDone++
		if numDone == len(filepaths) {
			break
		}
	}
	if errOccur {
		return nil, fmt.Errorf("%s", strings.Join(errStrList, ",\n"))
	}
	return fileList, nil
}

type fileType int

const (
	fileForRead fileType = iota
	fileForWrite
	closedFile
)

type file struct {
	*bufio.Reader
	*bufio.Writer
	file *os.File
	typ  fileType
}

func newFileForRead(path string) (*file, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &file{
		bufio.NewReader(f),
		nil,
		f,
		fileForRead,
	}, nil
}

func newFileForWrite(path string) (*file, error) {
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &file{
		nil,
		bufio.NewWriter(f),
		f,
		fileForWrite,
	}, nil
}

func fileForWriteFromFile(f *os.File) *file {
	return &file{
		nil,
		bufio.NewWriter(f),
		f,
		fileForWrite,
	}
}

func fileForReadFromFile(f *os.File) *file {
	return &file{
		bufio.NewReader(f),
		nil,
		f,
		fileForRead,
	}
}

func (f *file) reader() (*bufio.Reader, error) {
	switch f.typ {
	case fileForRead:
		return f.Reader, nil
	case fileForWrite:
		return nil, fmt.Errorf("nborm.file.reader() error: %s is a writing file", f.file.Name())
	default:
		return nil, fmt.Errorf("nborm.file.reader() error: %s has been closed", f.file.Name())
	}
}

func (f *file) writer() (*bufio.Writer, error) {
	switch f.typ {
	case fileForRead:
		return nil, fmt.Errorf("nborm.file.reader() error: %s is a reading file", f.file.Name())
	case fileForWrite:
		return f.Writer, nil
	default:
		return nil, fmt.Errorf("nborm.file.reader() error: %s has been closed", f.file.Name())
	}
}

func (f *file) toRead() (*file, error) {
	switch f.typ {
	case fileForRead:
		return f, nil
	case fileForWrite:
		filename := f.file.Name()
		if err := f.Close(); err != nil {
			return nil, err
		}
		nf, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		return fileForReadFromFile(nf), nil
	default:
		return nil, fmt.Errorf("nborm.file.toRead() error: %s has been closed", f.file.Name())
	}
}

func (f *file) toWrite() (*file, error) {
	switch f.typ {
	case fileForRead:
		filename := f.file.Name()
		if err := f.Close(); err != nil {
			return nil, err
		}
		nf, err := os.OpenFile(filename, os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		return fileForWriteFromFile(nf), nil
	case fileForWrite:
		return f, nil
	default:
		return nil, fmt.Errorf("nborm.file.toWrite() error: %s has been closed", f.file.Name())
	}
}

func (f *file) Close() error {
	switch f.typ {
	case fileForRead:
		f.typ = closedFile
		return f.file.Close()
	case fileForWrite:
		if err := f.Writer.Flush(); err != nil {
			return err
		}
		f.typ = closedFile
		return f.file.Close()
	default:
		return nil
	}
}

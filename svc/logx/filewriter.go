package logx

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ensure we always implement io.WriteCloser
var _ io.WriteCloser = (*FileWriter)(nil)

type FileWriter struct {
	// Filename is the file to write logs to.
	Filename    string
	FilePattern string
	size        int64
	file        *os.File
	mu          sync.Mutex

	millCh    chan bool
	startMill sync.Once
}

var (
	// currentTime exists so it can be mocked out by tests.
	currentTime = time.Now

	// os_Stat exists so it can be mocked out by tests.
	os_Stat = os.Stat
)

func NewFileWriter(filename string, filepattern string) io.Writer {
	return &FileWriter{
		Filename:    filename,
		FilePattern: filepattern,
	}
}

// Write implements io.Writer
func (l *FileWriter) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if err = l.openExistingOrNew(len(p)); err != nil {
		return 0, err
	}

	n, err = l.file.Write(p)
	l.size += int64(n)

	return n, err
}

// Close implements io.Closer, and closes the current logfile.
func (l *FileWriter) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.close()
}

// close closes the file if it is open.
func (l *FileWriter) close() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil
	return err
}

// genFilename generates the name of the logfile from the current time.
func (l *FileWriter) filename() string {
	dir := filepath.Dir(l.Filename)
	filename := filepath.Base(l.Filename)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := currentTime()

	timestamp := t.Format(l.FilePattern)
	return filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ext))

}

// openExistingOrNew opens the logfile if it exists and if the current write
// would not put it over MaxSize.  If there is no such file or the write would
// put it over the MaxSize, a new file is created.
func (l *FileWriter) openExistingOrNew(writeLen int) error {

	filename := l.filename()
	info, err := os_Stat(filename)
	if os.IsNotExist(err) {
		return l.openNew()
	}
	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// if we fail to open the old log file for some reason, just ignore
		// it and open a new log file.
		return l.openNew()
	}
	l.file = file
	l.size = info.Size()
	return nil
}

// dir returns the directory for the current filename.
func (l *FileWriter) dir() string {
	return filepath.Dir(l.filename())
}

func (l *FileWriter) openNew() error {
	err := os.MkdirAll(l.dir(), 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)
	}

	name := l.filename()
	mode := os.FileMode(0644)

	// we use truncate here because this should only get called when we've moved
	// the file ourselves. if someone else creates the file in the meantime,
	// just wipe out the contents.
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	}
	l.file = f
	l.size = 0
	return nil
}

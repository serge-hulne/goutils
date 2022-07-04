package files

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Open(fname string) (error, *os.File) {
	file, err := os.Open(fname)
	check(err)
	return err, file
}

func Close(file *os.File) {
	file.Close()
}

func ReadLines(f *os.File) chan string {
	ch_lines := make(chan string)
	go func() {
		defer close(ch_lines)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch_lines <- scanner.Text()
		}
	}()
	return ch_lines
}

func Read(fname string) (error, string) {
	dat, err := os.ReadFile(fname)
	check(err)
	return err, string(dat)
}

func WriteString(f *os.File, data string) (error, int) {
	n, err := f.WriteString(data)
	check(err)
	f.Sync()
	return err, n
}

func WriteBytes(f *os.File, data []byte) (error, int) {
	n, err := f.Write(data)
	check(err)
	f.Sync()
	return err, n
}

func WriteStringsBuffered(f *os.File, data string) (error, int) {
	w := bufio.NewWriter(f)
	n, err := w.WriteString(data)
	check(err)
	w.Flush()
	return err, n
}

package fileio

include (
	"os"
	"io"
)

func OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	return file
}

func ReadFile(file *os.File) string {
	out := make([]byte, 1024) //Not sure what good start size should be
	data := make([]byte, 100) 
	for {
		_, readErr := file.Read(data)
		if readErr == io.EOF{
			break
		}
		out = append(out, data...)
	}
	return string(out)
}
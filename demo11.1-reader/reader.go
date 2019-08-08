type Reader interface {
    Read(p []byte) (n int, err error)
}

func main() {
    var r io.Reader
    r = os.Stdin
    r = bufio.NewReader(r)
    r = new(bytes.Buffer)
    f, _ := os.Open("test.txt")
    r = bufio.NewReader(f)
}

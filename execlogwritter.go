package mlib


import (
        "bufio"
        "io"
        "sync"
)


type ExecLogWriter struct {
        stdout     io.Reader
	stderr     io.Reader
        dst      * bufio.Writer
        copyJobs sync.WaitGroup
        closed   chan struct{}
}


func NewLogWriter(stdout io.Reader, stderr io.Reader, dst *bufio.Writer) *ExecLogWriter {
        return &ExecLogWriter{
		stdout: stdout,
		stderr: stderr,
		dst:    dst,
		closed: make(chan struct{}),
        }
}


func (c *ExecLogWriter) Run() {
	c.copyJobs.Add(1)
	go c.copySrc(c.stdout)

	c.copyJobs.Add(1)
	go c.copySrc(c.stderr)
}


func (c *ExecLogWriter) copySrc(src io.Reader) {
        defer c.copyJobs.Done()
        reader := bufio.NewReader(src)

        for {
                select {
                case <-c.closed:
                        return
                default:
                        line, err := reader.ReadBytes('\n')

                        if err == nil || len(line) > 0 {
				_, logErr := c.dst.Write(line)
                                if logErr != nil {
                                        Error("Failed to log msg %q  %s", line, logErr)
                                }
                        }

                        if err != nil {
                                if err != io.EOF {
                                        Error("Error scanning log stream: %s", err)
                                }
                                return
                        }
                }
        }
}

func (c *ExecLogWriter) Wait() {
        c.copyJobs.Wait()
}

func (c *ExecLogWriter) Close() {
        select {
        case <-c.closed:
        default:
                close(c.closed)
        }
}

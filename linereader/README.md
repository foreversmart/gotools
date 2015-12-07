## LineReader
---

#### this tool design for reading lines from a io.Reader

* Usage: call NewLineReader to create LineReader, call Reading method manually to read from a reader. Line method will return the line content. this method will block until you get a complete line.

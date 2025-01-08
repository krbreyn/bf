A Brainf*ck interpreter, written in Go.

`compressor.go` and `decompressor.go` implement a dialect of B.F. that allows for more compact expression. There are no measures to prevent it from affecting the text in comments.

It works pretty simply: a compressed input such as `+4>3.<2` is decompressed as `++++>>>.<<`, and vice versa.

package syncpool

import (
	"bufio"
	"io"
	"sync"
)

var (
	DefaultReadBufferSize  = 1024
	DefaultWriteBufferSize = 1024
)

type SyncPool struct {
	ReaderPool    *sync.Pool
	WriterPool    *sync.Pool
	ReaderBufSize int
	WriterBufSize int
}

func NewSyncPool() *SyncPool {
	pool := &SyncPool{}
	pool.ReaderPool = new(sync.Pool)
	pool.WriterPool = new(sync.Pool)
	pool.ReaderBufSize = DefaultReadBufferSize
	pool.WriterBufSize = DefaultWriteBufferSize
	return pool
}

func (self *SyncPool) GetReader(reader io.Reader) *bufio.Reader {
	if object := self.ReaderPool.Get(); object != nil {
		br := object.(*bufio.Reader)
		br.Reset(reader)
		return br
	}
	return bufio.NewReaderSize(reader, self.ReaderBufSize)
}

func (self *SyncPool) PutReader(reader *bufio.Reader) {
	reader.Reset(nil)
	self.ReaderPool.Put(reader)
}

func (self *SyncPool) GetWriter(writer io.Writer) *bufio.Writer {
	if object := self.WriterPool.Get(); object != nil {
		bw := object.(*bufio.Writer)
		bw.Reset(writer)
		return bw
	}
	return bufio.NewWriterSize(writer, self.WriterBufSize)
}

func (self *SyncPool) PutWriter(writer *bufio.Writer) {
	writer.Reset(nil)
	self.WriterPool.Put(writer)
}

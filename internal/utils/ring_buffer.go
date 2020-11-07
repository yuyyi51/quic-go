package utils

type RingBuffer struct {
	buf      []byte
	offset   uint64
	length   int
	capacity int
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		buf:      make([]byte, capacity, capacity),
		capacity: capacity,
	}
}

func (buffer *RingBuffer) remainSize() int {
	return buffer.capacity - buffer.length
}

func (buffer *RingBuffer) getSmallOffset() int {
	return int(buffer.offset % uint64(buffer.capacity))
}

func (buffer *RingBuffer) getLargeOffset() int {
	return int((buffer.offset + uint64(buffer.length)) % uint64(buffer.capacity))
}

func (buffer *RingBuffer) rightWriteSize() int {
	// continuous size in buffer large bound's right
	// may be cut off by small bound or buf's right bound
	if buffer.getSmallOffset() > buffer.getLargeOffset() {
		return buffer.getSmallOffset() - buffer.getLargeOffset()
	}
	if buffer.getSmallOffset() == buffer.getLargeOffset() {
		// full or empty
		if buffer.length != 0 {
			// full
			return 0
		}
		// if it's empty, just go to statement below
	}
	return buffer.capacity - buffer.getLargeOffset()
}

func (buffer *RingBuffer) rightReadSize() int {
	// continuous size in buffer small bounds' right
	// may be cut off by large bound or buf's right bound
	if buffer.getLargeOffset() > buffer.getSmallOffset() {
		return buffer.getLargeOffset() - buffer.getSmallOffset()
	}
	if buffer.getLargeOffset() == buffer.getSmallOffset() {
		// full or empty
		if buffer.length == 0 {
			// empty
			return 0
		}
		// if it's full, just go to statement below
	}
	return buffer.capacity - buffer.getSmallOffset()
}

func (buffer *RingBuffer) inputData(data []byte) {
	// we promise len(data) is smaller than buffer.remainSize()
	if buffer.remainSize() < len(data) {
		panic("must promise len(data) is smaller than buffer.remainSize() when called inputData")
	}
	dataCopied := buffer.rightWriteSize()
	if len(data) < dataCopied {
		dataCopied = len(data)
	}
	copy(buffer.buf[buffer.getLargeOffset():], data[:dataCopied])
	buffer.length += dataCopied
	if dataCopied == len(data) {
		return
	}
	// more data, we need start from left bound of buf
	data = data[dataCopied:]
	dataCopied = buffer.rightWriteSize()
	if len(data) < dataCopied {
		dataCopied = len(data)
	}
	copy(buffer.buf[buffer.getLargeOffset():], data[:dataCopied])
	buffer.length += dataCopied
	if dataCopied <= len(data) {
		panic("must promise len(data) is smaller than buffer.remainSize() when called inputData")
	}
}

func (buffer *RingBuffer) Input(data []byte) (int, bool /* completed */) {
	DefaultLogger.Debugf("Input called len %d", buffer.length)
	defer func(buf *RingBuffer) {
		DefaultLogger.Debugf("Input called done len %d", buf.length)
	}(buffer)
	inputSize := len(data)
	if len(data) > buffer.remainSize() {
		inputSize = buffer.remainSize()
	}
	buffer.inputData(data[:inputSize])
	return inputSize, inputSize >= len(data)
}

func (buffer *RingBuffer) Pop(maxSize int) []byte {
	DefaultLogger.Debugf("Pop called len %d", buffer.length)
	defer func(buf *RingBuffer) {
		defer DefaultLogger.Debugf("Pop called done len %d", buffer.length)
	}(buffer)
	if maxSize <= 0 {
		return nil
	}
	if maxSize > buffer.length {
		maxSize = buffer.length
	}
	ret := make([]byte, maxSize, maxSize)
	dataRead := buffer.rightReadSize()
	if dataRead > maxSize {
		dataRead = maxSize
	}
	copy(ret[:dataRead], buffer.buf[buffer.getSmallOffset():buffer.getSmallOffset()+dataRead])
	buffer.offset += uint64(dataRead)
	buffer.length -= dataRead
	if dataRead == maxSize {
		return ret
	}
	dataRemain := maxSize - dataRead
	copy(ret[dataRead:maxSize], buffer.buf[buffer.getSmallOffset():buffer.getSmallOffset()+dataRemain])
	buffer.offset += uint64(dataRemain)
	buffer.length -= dataRemain
	return ret
}

func (buffer *RingBuffer) Length() int {
	return buffer.length
}

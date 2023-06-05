package biz

import logic "logic/api/v1"

// UploadFileReader 上传文件ReaderWriter
type UploadFileReader struct {
	Stream logic.Logic_UploadFileServer
	Buffer [1 << 22]byte // 4MB
	start  int
	end    int
}

func NewUploadFileReader(stream logic.Logic_UploadFileServer) *UploadFileReader {
	return &UploadFileReader{Stream: stream, Buffer: [1 << 22]byte{}, start: 0, end: 0}
}

// Read 读取数据
func (u *UploadFileReader) Read(p []byte) (n int, err error) {
	if len(p) > len(u.Buffer) {
		req, err := u.Stream.Recv()
		if err != nil {
			return 0, err
		}
		data := req.GetChunkData()
		n = copy(p, data)
		return n, nil
	} else {
		if u.start == u.end {
			req, err := u.Stream.Recv()
			if err != nil {
				return 0, err
			}
			data := req.GetChunkData()
			n = copy(u.Buffer[:], data)
			u.start = 0
			u.end = n
		}
		n = copy(p, u.Buffer[u.start:u.end])
		u.start += n
		return n, nil
	}
}

package common

import v1 "user/api/v1/user"

type UploadAvatarReaderWriter struct {
	Stream v1.User_UploadAvatarServer
	Buffer [1 << 20]byte
	start  int
	end    int
}

func (u *UploadAvatarReaderWriter) Read(p []byte) (n int, err error) {
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

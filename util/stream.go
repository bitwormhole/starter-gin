package util

import (
	"errors"
	"io"
)

// PumpStream 把数据从Reader读出，并写入到Writer
func PumpStream(i io.Reader, o io.Writer, buffer []byte) (int64, error) {

	if i == nil {
		return 0, errors.New("input==nil")
	}
	if o == nil {
		return 0, errors.New("output==nil")
	}
	if buffer == nil {
		buffer = make([]byte, 1024)
	} else if len(buffer) < 1 {
		buffer = make([]byte, 1024)
	}

	var err error = nil
	count := int64(0)

	for {
		cb1, err1 := i.Read(buffer)
		if err1 != nil {
			err = err1
		}
		if cb1 > 0 {
			cb2, err2 := o.Write(buffer[0:cb1])
			if err2 != nil {
				err = err2
			}
			if cb2 != cb1 {
				return count, errors.New("cb1!=cb2")
			}
			count += int64(cb1)
		} else if cb1 == 0 {
			break
		}
		if err != nil {
			return count, err
		}
	}

	return count, nil
}

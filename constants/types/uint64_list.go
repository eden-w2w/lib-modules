package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Uint64List []uint64

func (u *Uint64List) UnmarshalJSON(bytes []byte) error {
	var list = make([]string, 0)
	err := json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}

	*u = make([]uint64, 0)
	for _, v := range list {
		vi, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		*u = append(*u, vi)
	}
	return nil
}

func (u Uint64List) MarshalJSON() ([]byte, error) {
	var list = make([]string, len(u))
	for _, v := range u {
		list = append(list, fmt.Sprintf("%d", v))
	}
	return json.Marshal(list)
}

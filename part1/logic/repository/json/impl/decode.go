package impl

import (
	"encoding/json"
)

func (r *repository) Decode(decoder *json.Decoder, v interface{}) error {
	return decoder.Decode(v)
}

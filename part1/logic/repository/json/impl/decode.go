package impl

import (
	"encoding/json"
)

func (r *repository) Decode(decoder *json.Decoder, v any) error {
	return decoder.Decode(v)
}

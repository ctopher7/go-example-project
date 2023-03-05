package json

import "encoding/json"

type Repository interface {
	Decode(decoder *json.Decoder, v any) error
}

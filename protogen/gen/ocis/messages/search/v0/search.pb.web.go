// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: v0.proto

package v0

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
)

// ResourceIDJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ResourceID. This struct is safe to replace or modify but
// should not be done so concurrently.
var ResourceIDJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ResourceID) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ResourceIDJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ResourceID)(nil)

// ResourceIDJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ResourceID. This struct is safe to replace or modify but
// should not be done so concurrently.
var ResourceIDJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ResourceID) UnmarshalJSON(b []byte) error {
	return ResourceIDJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ResourceID)(nil)

// ReferenceJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Reference. This struct is safe to replace or modify but
// should not be done so concurrently.
var ReferenceJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Reference) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ReferenceJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Reference)(nil)

// ReferenceJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Reference. This struct is safe to replace or modify but
// should not be done so concurrently.
var ReferenceJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Reference) UnmarshalJSON(b []byte) error {
	return ReferenceJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Reference)(nil)

// EntityJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Entity. This struct is safe to replace or modify but
// should not be done so concurrently.
var EntityJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Entity) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := EntityJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Entity)(nil)

// EntityJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Entity. This struct is safe to replace or modify but
// should not be done so concurrently.
var EntityJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Entity) UnmarshalJSON(b []byte) error {
	return EntityJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Entity)(nil)

// MatchJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Match. This struct is safe to replace or modify but
// should not be done so concurrently.
var MatchJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Match) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := MatchJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Match)(nil)

// MatchJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Match. This struct is safe to replace or modify but
// should not be done so concurrently.
var MatchJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Match) UnmarshalJSON(b []byte) error {
	return MatchJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Match)(nil)

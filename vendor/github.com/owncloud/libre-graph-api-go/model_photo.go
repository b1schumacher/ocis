/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// checks if the Photo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Photo{}

// Photo The photo resource provides photo and camera properties, for example, EXIF metadata, on a driveItem.
type Photo struct {
	// Camera manufacturer. Read-only.
	CameraMake *string `json:"cameraMake,omitempty"`
	// Camera model. Read-only.
	CameraModel *string `json:"cameraModel,omitempty"`
	// The denominator for the exposure time fraction from the camera. Read-only.
	ExposureDenominator *float64 `json:"exposureDenominator,omitempty"`
	// The numerator for the exposure time fraction from the camera. Read-only.
	ExposureNumerator *float64 `json:"exposureNumerator,omitempty"`
	// The F-stop value from the camera. Read-only.
	FNumber *float64 `json:"fNumber,omitempty"`
	// The focal length from the camera. Read-only.
	FocalLength *float64 `json:"focalLength,omitempty"`
	// The ISO value from the camera. Read-only.
	Iso *int32 `json:"iso,omitempty"`
	// The orientation value from the camera. Read-only.
	Orientation *int32 `json:"orientation,omitempty"`
	// Represents the date and time the photo was taken. Read-only.
	TakenDateTime *time.Time `json:"takenDateTime,omitempty"`
}

// NewPhoto instantiates a new Photo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoto() *Photo {
	this := Photo{}
	return &this
}

// NewPhotoWithDefaults instantiates a new Photo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhotoWithDefaults() *Photo {
	this := Photo{}
	return &this
}

// GetCameraMake returns the CameraMake field value if set, zero value otherwise.
func (o *Photo) GetCameraMake() string {
	if o == nil || IsNil(o.CameraMake) {
		var ret string
		return ret
	}
	return *o.CameraMake
}

// GetCameraMakeOk returns a tuple with the CameraMake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetCameraMakeOk() (*string, bool) {
	if o == nil || IsNil(o.CameraMake) {
		return nil, false
	}
	return o.CameraMake, true
}

// HasCameraMake returns a boolean if a field has been set.
func (o *Photo) HasCameraMake() bool {
	if o != nil && !IsNil(o.CameraMake) {
		return true
	}

	return false
}

// SetCameraMake gets a reference to the given string and assigns it to the CameraMake field.
func (o *Photo) SetCameraMake(v string) {
	o.CameraMake = &v
}

// GetCameraModel returns the CameraModel field value if set, zero value otherwise.
func (o *Photo) GetCameraModel() string {
	if o == nil || IsNil(o.CameraModel) {
		var ret string
		return ret
	}
	return *o.CameraModel
}

// GetCameraModelOk returns a tuple with the CameraModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetCameraModelOk() (*string, bool) {
	if o == nil || IsNil(o.CameraModel) {
		return nil, false
	}
	return o.CameraModel, true
}

// HasCameraModel returns a boolean if a field has been set.
func (o *Photo) HasCameraModel() bool {
	if o != nil && !IsNil(o.CameraModel) {
		return true
	}

	return false
}

// SetCameraModel gets a reference to the given string and assigns it to the CameraModel field.
func (o *Photo) SetCameraModel(v string) {
	o.CameraModel = &v
}

// GetExposureDenominator returns the ExposureDenominator field value if set, zero value otherwise.
func (o *Photo) GetExposureDenominator() float64 {
	if o == nil || IsNil(o.ExposureDenominator) {
		var ret float64
		return ret
	}
	return *o.ExposureDenominator
}

// GetExposureDenominatorOk returns a tuple with the ExposureDenominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetExposureDenominatorOk() (*float64, bool) {
	if o == nil || IsNil(o.ExposureDenominator) {
		return nil, false
	}
	return o.ExposureDenominator, true
}

// HasExposureDenominator returns a boolean if a field has been set.
func (o *Photo) HasExposureDenominator() bool {
	if o != nil && !IsNil(o.ExposureDenominator) {
		return true
	}

	return false
}

// SetExposureDenominator gets a reference to the given float64 and assigns it to the ExposureDenominator field.
func (o *Photo) SetExposureDenominator(v float64) {
	o.ExposureDenominator = &v
}

// GetExposureNumerator returns the ExposureNumerator field value if set, zero value otherwise.
func (o *Photo) GetExposureNumerator() float64 {
	if o == nil || IsNil(o.ExposureNumerator) {
		var ret float64
		return ret
	}
	return *o.ExposureNumerator
}

// GetExposureNumeratorOk returns a tuple with the ExposureNumerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetExposureNumeratorOk() (*float64, bool) {
	if o == nil || IsNil(o.ExposureNumerator) {
		return nil, false
	}
	return o.ExposureNumerator, true
}

// HasExposureNumerator returns a boolean if a field has been set.
func (o *Photo) HasExposureNumerator() bool {
	if o != nil && !IsNil(o.ExposureNumerator) {
		return true
	}

	return false
}

// SetExposureNumerator gets a reference to the given float64 and assigns it to the ExposureNumerator field.
func (o *Photo) SetExposureNumerator(v float64) {
	o.ExposureNumerator = &v
}

// GetFNumber returns the FNumber field value if set, zero value otherwise.
func (o *Photo) GetFNumber() float64 {
	if o == nil || IsNil(o.FNumber) {
		var ret float64
		return ret
	}
	return *o.FNumber
}

// GetFNumberOk returns a tuple with the FNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetFNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.FNumber) {
		return nil, false
	}
	return o.FNumber, true
}

// HasFNumber returns a boolean if a field has been set.
func (o *Photo) HasFNumber() bool {
	if o != nil && !IsNil(o.FNumber) {
		return true
	}

	return false
}

// SetFNumber gets a reference to the given float64 and assigns it to the FNumber field.
func (o *Photo) SetFNumber(v float64) {
	o.FNumber = &v
}

// GetFocalLength returns the FocalLength field value if set, zero value otherwise.
func (o *Photo) GetFocalLength() float64 {
	if o == nil || IsNil(o.FocalLength) {
		var ret float64
		return ret
	}
	return *o.FocalLength
}

// GetFocalLengthOk returns a tuple with the FocalLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetFocalLengthOk() (*float64, bool) {
	if o == nil || IsNil(o.FocalLength) {
		return nil, false
	}
	return o.FocalLength, true
}

// HasFocalLength returns a boolean if a field has been set.
func (o *Photo) HasFocalLength() bool {
	if o != nil && !IsNil(o.FocalLength) {
		return true
	}

	return false
}

// SetFocalLength gets a reference to the given float64 and assigns it to the FocalLength field.
func (o *Photo) SetFocalLength(v float64) {
	o.FocalLength = &v
}

// GetIso returns the Iso field value if set, zero value otherwise.
func (o *Photo) GetIso() int32 {
	if o == nil || IsNil(o.Iso) {
		var ret int32
		return ret
	}
	return *o.Iso
}

// GetIsoOk returns a tuple with the Iso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetIsoOk() (*int32, bool) {
	if o == nil || IsNil(o.Iso) {
		return nil, false
	}
	return o.Iso, true
}

// HasIso returns a boolean if a field has been set.
func (o *Photo) HasIso() bool {
	if o != nil && !IsNil(o.Iso) {
		return true
	}

	return false
}

// SetIso gets a reference to the given int32 and assigns it to the Iso field.
func (o *Photo) SetIso(v int32) {
	o.Iso = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *Photo) GetOrientation() int32 {
	if o == nil || IsNil(o.Orientation) {
		var ret int32
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetOrientationOk() (*int32, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *Photo) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given int32 and assigns it to the Orientation field.
func (o *Photo) SetOrientation(v int32) {
	o.Orientation = &v
}

// GetTakenDateTime returns the TakenDateTime field value if set, zero value otherwise.
func (o *Photo) GetTakenDateTime() time.Time {
	if o == nil || IsNil(o.TakenDateTime) {
		var ret time.Time
		return ret
	}
	return *o.TakenDateTime
}

// GetTakenDateTimeOk returns a tuple with the TakenDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Photo) GetTakenDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TakenDateTime) {
		return nil, false
	}
	return o.TakenDateTime, true
}

// HasTakenDateTime returns a boolean if a field has been set.
func (o *Photo) HasTakenDateTime() bool {
	if o != nil && !IsNil(o.TakenDateTime) {
		return true
	}

	return false
}

// SetTakenDateTime gets a reference to the given time.Time and assigns it to the TakenDateTime field.
func (o *Photo) SetTakenDateTime(v time.Time) {
	o.TakenDateTime = &v
}

func (o Photo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Photo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CameraMake) {
		toSerialize["cameraMake"] = o.CameraMake
	}
	if !IsNil(o.CameraModel) {
		toSerialize["cameraModel"] = o.CameraModel
	}
	if !IsNil(o.ExposureDenominator) {
		toSerialize["exposureDenominator"] = o.ExposureDenominator
	}
	if !IsNil(o.ExposureNumerator) {
		toSerialize["exposureNumerator"] = o.ExposureNumerator
	}
	if !IsNil(o.FNumber) {
		toSerialize["fNumber"] = o.FNumber
	}
	if !IsNil(o.FocalLength) {
		toSerialize["focalLength"] = o.FocalLength
	}
	if !IsNil(o.Iso) {
		toSerialize["iso"] = o.Iso
	}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.TakenDateTime) {
		toSerialize["takenDateTime"] = o.TakenDateTime
	}
	return toSerialize, nil
}

type NullablePhoto struct {
	value *Photo
	isSet bool
}

func (v NullablePhoto) Get() *Photo {
	return v.value
}

func (v *NullablePhoto) Set(val *Photo) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoto) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoto(val *Photo) *NullablePhoto {
	return &NullablePhoto{value: val, isSet: true}
}

func (v NullablePhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
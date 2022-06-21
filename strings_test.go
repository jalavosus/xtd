package xtd_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"

	"github.com/jalavosus/xtd"
)

type FromStringConstraint interface {
	constraints.Signed | constraints.Unsigned | constraints.Float
}

type FromStringTestCase[T FromStringConstraint] struct {
	name    string
	arg     string
	want    T
	wantErr bool
}

func TestFloatFromString(t *testing.T) {
	float32Tests := []FromStringTestCase[float32]{
		{
			name:    "valid",
			arg:     "42.069",
			want:    42.069,
			wantErr: false,
		},
		{
			name:    "max float32",
			arg:     fmt.Sprintf("%f", math.MaxFloat32),
			want:    math.MaxFloat32,
			wantErr: false,
		},
		{
			name:    "max float64 (fails)",
			arg:     fmt.Sprintf("%f", math.MaxFloat64),
			want:    0.,
			wantErr: true,
		},
	}

	float64Tests := []FromStringTestCase[float64]{
		{
			name:    "valid",
			arg:     "42.069",
			want:    42.069,
			wantErr: false,
		},
		{
			name:    "max float32",
			arg:     fmt.Sprintf("%f", math.MaxFloat32),
			want:    math.MaxFloat32,
			wantErr: false,
		},
		{
			name:    "max float64 (fails)",
			arg:     fmt.Sprintf("%f", math.MaxFloat64),
			want:    math.MaxFloat64,
			wantErr: false,
		},
	}

	t.Run("float32", testFloatFromString(float32Tests))
	t.Run("float64", testFloatFromString(float64Tests))
}

func testFloatFromString[T constraints.Float](tests []FromStringTestCase[T]) func(*testing.T) {
	return func(t *testing.T) {
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got, err := xtd.FloatFromString[T](tc.arg)

				if tc.wantErr {
					assert.Error(t, err)
					return
				}

				assert.NoError(t, err)

				assert.Equal(t, tc.want, got)
			})
		}
	}
}

func TestIntFromString(t *testing.T) {
	intTestCases := []FromStringTestCase[int]{
		{
			name:    "valid",
			arg:     "5",
			want:    5,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     "-5",
			want:    -5,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", math.MaxInt),
			want:    math.MaxInt,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	int8TestCases := []FromStringTestCase[int8]{
		{
			name:    "valid",
			arg:     "127",
			want:    127,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     "-128",
			want:    -128,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", math.MaxInt),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	int16TestCases := []FromStringTestCase[int16]{
		{
			name:    "valid",
			arg:     "-32768",
			want:    -32768,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", math.MaxInt16),
			want:    math.MaxInt16,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", math.MaxInt32),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	int32TestCases := []FromStringTestCase[int32]{
		{
			name:    "valid",
			arg:     "-2147483648",
			want:    -2147483648,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", math.MaxInt32),
			want:    math.MaxInt32,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", math.MaxInt64),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	int64TestCases := []FromStringTestCase[int64]{
		{
			name:    "valid",
			arg:     "-9223372036854775808",
			want:    -9223372036854775808,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", math.MaxInt64),
			want:    math.MaxInt64,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", uint64(math.MaxUint64)),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	t.Run("int", testIntFromString(intTestCases))
	t.Run("int8", testIntFromString(int8TestCases))
	t.Run("int16", testIntFromString(int16TestCases))
	t.Run("int32", testIntFromString(int32TestCases))
	t.Run("int64", testIntFromString(int64TestCases))
}

func testIntFromString[T constraints.Signed](tests []FromStringTestCase[T]) func(*testing.T) {
	return func(t *testing.T) {
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got, err := xtd.IntFromString[T](tc.arg)

				if tc.wantErr {
					assert.Error(t, err)
					return
				}

				assert.NoError(t, err)

				assert.Equal(t, tc.want, got)
			})
		}
	}
}

func TestUintFromString(t *testing.T) {
	uintTestCases := []FromStringTestCase[uint]{
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", uint(math.MaxUint)),
			want:    math.MaxUint,
			wantErr: false,
		},
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", uint64(math.MaxUint64)),
			want:    math.MaxUint64,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     "-9223372036854775808",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	uint8TestCases := []FromStringTestCase[uint8]{
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", uint8(math.MaxUint8)),
			want:    math.MaxUint8,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     "-9223372036854775808",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", uint16(math.MaxUint16)),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	uint16TestCases := []FromStringTestCase[uint16]{
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", uint16(math.MaxUint16)),
			want:    math.MaxUint16,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     "-9223372036854775808",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", uint32(math.MaxUint32)),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	uint32TestCases := []FromStringTestCase[uint32]{
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", uint32(math.MaxUint32)),
			want:    math.MaxUint32,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     "-9223372036854775808",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     fmt.Sprintf("%d", uint64(math.MaxUint64)),
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	uint64TestCases := []FromStringTestCase[uint64]{
		{
			name:    "valid",
			arg:     fmt.Sprintf("%d", uint64(math.MaxUint64)),
			want:    math.MaxUint64,
			wantErr: false,
		},
		{
			name:    "invalid",
			arg:     "-9223372036854775808",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid",
			arg:     "420.69",
			want:    0,
			wantErr: true,
		},
	}

	t.Run("uint", testUintFromString(uintTestCases))
	t.Run("uint8", testUintFromString(uint8TestCases))
	t.Run("uint16", testUintFromString(uint16TestCases))
	t.Run("uint32", testUintFromString(uint32TestCases))
	t.Run("uint64", testUintFromString(uint64TestCases))
}

func testUintFromString[T constraints.Unsigned](tests []FromStringTestCase[T]) func(*testing.T) {
	return func(t *testing.T) {
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got, err := xtd.UintFromString[T](tc.arg)

				if tc.wantErr {
					assert.Error(t, err)
					return
				}

				assert.NoError(t, err)

				assert.Equal(t, tc.want, got)
			})
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "hello",
			arg:  "hello",
			want: "olleh",
		},
		{
			name: "puppies",
			arg:  "puppies",
			want: "seippup",
		},
		{
			name: "033334776GGre",
			arg:  "033334776GGre",
			want: "erGG677433330",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, xtd.ReverseString(tt.arg))
		})
	}
}

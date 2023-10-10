package types

import (
	"testing"

	"github.com/massalabs/station-massa-wallet/pkg/types/object"
	"github.com/stretchr/testify/assert"
)

func TestAddress_UnmarshalText(t *testing.T) {
	tests := []struct {
		name         string
		text         []byte
		expectedType object.Kind
		expectedData []byte
		expectedErr  error
	}{
		{
			name:         "Valid user address account bonjour",
			text:         []byte("AU1uSeKuYRQkC8fHWoohBmq8WKYxGTXuzaTcGomAoLXTGdq8kEsR"),
			expectedType: object.UserAddress,
			expectedData: []byte{0x77, 0x13, 0x86, 0x8f, 0xe5, 0x5a, 0xd1, 0xdb, 0x9c, 0x8, 0x30, 0x7c, 0x61, 0x5e, 0xdf, 0xc0, 0xc8, 0x3b, 0x5b, 0xd9, 0x88, 0xec, 0x2e, 0x3c, 0xe9, 0xe4, 0x1c, 0xf1, 0xf9, 0x4d, 0xc5, 0xd1},
			expectedErr:  nil,
		},
		{
			name:         "Valid user address",
			text:         []byte("AU12NuzyNHR5xujwATTi4CFCeLeQ6wdBcQVxohg1WNgoUArKKZYZk"),
			expectedType: object.UserAddress,
			expectedData: []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
			expectedErr:  nil,
		},
		{
			name:         "Valid smart contract",
			text:         []byte("AS1pzpVk79cubdM7Zk9pNdg1JR4qKooVa6usyquiCEQhPeRWD9k7"),
			expectedType: object.SmartContractAddress,
			expectedData: []byte{0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				Object: &object.Object{},
			}
			if err := a.UnmarshalText(tt.text); err != tt.expectedErr {
				t.Errorf("Address.UnmarshalText() error = %v, wantErr %v", err, tt.expectedErr)
			}
			assert.Equal(t, tt.expectedType, a.Kind, "Should be %s, got %s", tt.expectedType, a.Kind)
			assert.Equal(t, tt.expectedData, a.Data)
		})
	}
}

func TestAddress_MashalText(t *testing.T) {
	tests := []struct {
		name         string
		address      Address
		expectedData []byte
		expectedErr  error
	}{
		{
			name: "Valid user address",
			address: Address{
				Object: &object.Object{
					Kind:    object.UserAddress,
					Version: 0x00,
					Data:    []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
				},
			},
			expectedData: []byte("AU12NuzyNHR5xujwATTi4CFCeLeQ6wdBcQVxohg1WNgoUArKKZYZk"),
			expectedErr:  nil,
		},
		{
			name: "Valid smart contract address",
			address: Address{
				Object: &object.Object{
					Kind:    object.SmartContractAddress,
					Version: 0x00,
					Data:    []byte{0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
				},
			},
			expectedData: []byte("AS1pzpVk79cubdM7Zk9pNdg1JR4qKooVa6usyquiCEQhPeRWD9k7"),
			expectedErr:  nil,
		},
		{
			name: "Invalid version",
			address: Address{
				Object: &object.Object{
					Kind:    object.SmartContractAddress,
					Version: 0x01,
					Data:    []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
				},
			},
			expectedData: []byte(""),
			expectedErr:  object.ErrUnsupportedVersion,
		},
		{
			name: "Invalid Type",
			address: Address{
				Object: &object.Object{
					Kind:    object.EncryptedPrivateKey,
					Version: 0x00,
					Data:    []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
				},
			},
			expectedData: []byte(""),
			expectedErr:  object.ErrInvalidType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.address.MarshalText()
			assert.Equal(t, string(tt.expectedData), string(got))
			assert.Equal(t, err, tt.expectedErr)
		})
	}
}

func TestAddress_MarshalBinary(t *testing.T) {
	tests := []struct {
		name         string
		address      Address
		expectedData []byte
		expectedErr  error
	}{
		{
			name: "Valid user address",
			address: Address{
				Object: &object.Object{
					Kind:    object.UserAddress,
					Version: 0x00,
					Data:    []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
				},
			},
			expectedData: []byte{0x00, 0x00, 0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
			expectedErr:  nil,
		},
		{
			name: "Valid smart contract address",
			address: Address{
				Object: &object.Object{
					Kind:    object.SmartContractAddress,
					Version: 0x00,
					Data:    []byte{0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
				},
			},
			expectedData: []byte{0x01, 0x00, 0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedErr:  nil,
		},
		{
			name: "Invalide version",
			address: Address{
				Object: &object.Object{
					Kind:    object.SmartContractAddress,
					Version: 0x01,
					Data:    []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
				},
			},
			expectedData: nil,
			expectedErr:  object.ErrUnsupportedVersion,
		},
		{
			name: "Invalide Type",
			address: Address{
				Object: &object.Object{
					Kind:    object.EncryptedPrivateKey,
					Version: 0x00,
					Data:    []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
				},
			},
			expectedData: nil,
			expectedErr:  object.ErrInvalidType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.address.MarshalBinary()
			assert.Equal(t, tt.expectedData, got)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestAddress_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name         string
		data         []byte
		expectedType object.Kind
		expectedData []byte
		expectedErr  error
	}{
		{
			name:         "Valid user address",
			data:         []byte{0x00, 0x00, 0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
			expectedType: object.UserAddress,
			expectedData: []byte{0xb5, 0x74, 0x3a, 0xf0, 0x51, 0x96, 0xf1, 0xa3, 0xf7, 0x4a, 0xe9, 0xb4, 0xd9, 0xb9, 0x9c, 0x78, 0xfe, 0x15, 0x97, 0x38, 0xe4, 0xcc, 0x58, 0x49, 0x43, 0xd5, 0xdc, 0xd0, 0x43, 0x5a, 0x9c, 0xad},
			expectedErr:  nil,
		},
		{
			name:         "Valid smart contract",
			data:         []byte{0x01, 0x00, 0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedType: object.SmartContractAddress,
			expectedData: []byte{0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedErr:  nil,
		},
		{
			name:         "Invalid type",
			data:         []byte{0x02, 0x00, 0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedType: object.Unknown,
			expectedData: nil,
			expectedErr:  object.ErrInvalidType,
		},
		{
			name:         "Unsupported version",
			data:         []byte{0x01, 0x01, 0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedType: object.SmartContractAddress,
			expectedData: []byte{0x6c, 0xfb, 0x97, 0x95, 0x87, 0xd6, 0x2f, 0xa8, 0x2a, 0x35, 0xb9, 0x3a, 0x1e, 0x3d, 0x8e, 0xe3, 0xab, 0x7f, 0x63, 0x17, 0x61, 0xce, 0x5f, 0x8d, 0xc2, 0xce, 0x10, 0xec, 0x58, 0xb1, 0x64, 0x8f},
			expectedErr:  object.ErrUnsupportedVersion,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				Object: &object.Object{},
			}
			if err := a.UnmarshalBinary(tt.data); err != tt.expectedErr {
				t.Errorf("Address.UnmarshalBinary() error = %v, wantErr %v", err, tt.expectedErr)
			}
			assert.Equal(t, tt.expectedType, a.Kind, "Should be %s, got %s", tt.expectedType, a.Kind)
			assert.Equal(t, tt.expectedData, a.Data)
		})
	}
}

func TestNewAddressFromPublicKey(t *testing.T) {
	samplePublicKey := &PublicKey{
		Object: &object.Object{
			Kind:    object.PublicKey,
			Version: 0x00,
			Data:    []byte{45, 150, 188, 218, 203, 190, 65, 56, 44, 162, 62, 82, 227, 210, 25, 108, 186, 101, 231, 161, 172, 210, 9, 223, 201, 92, 107, 50, 182, 161, 138, 147},
		},
	}

	expectedKind := object.UserAddress
	expectedVersion := AddressLastVersion

	// Call NewAddressFromPublicKey with the sample PublicKey
	address := NewAddressFromPublicKey(samplePublicKey)

	t.Run("Validate Kind And Version", func(t *testing.T) {
		assert.Equal(t, expectedKind, address.Kind)
		assert.Equal(t, byte(expectedVersion), address.Version)
	})

	t.Run("Validate Data", func(t *testing.T) {
		// assert with the text representation:
		text, err := address.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, string(text), "AU1uSeKuYRQkC8fHWoohBmq8WKYxGTXuzaTcGomAoLXTGdq8kEsR")

		// assert with the binary data:
		data := []byte{0x77, 0x13, 0x86, 0x8f, 0xe5, 0x5a, 0xd1, 0xdb, 0x9c, 0x8, 0x30, 0x7c, 0x61, 0x5e, 0xdf, 0xc0, 0xc8, 0x3b, 0x5b, 0xd9, 0x88, 0xec, 0x2e, 0x3c, 0xe9, 0xe4, 0x1c, 0xf1, 0xf9, 0x4d, 0xc5, 0xd1}
		assert.Equal(t, data, address.Data)
	})
}

package random

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	service "github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type kmsSource struct {
	service.KMSAPI
}

func NewKMS(api service.KMSAPI) *kmsSource {
	return &kmsSource{api}
}

func (k *kmsSource) Fill(buf []byte) error {

	buf2, err := k.Generate(uint32(len(buf)))

	if err != nil {
		return err
	}

	copy(buf2, buf)

	return nil

}

func (k *kmsSource) Generate(bytes uint32) ([]byte, error) {

	req := &kms.GenerateRandomInput{
		NumberOfBytes: aws.Int64(int64(bytes)),
	}

	res, err := k.GenerateRandom(req)

	if err != nil {
		return nil, err
	}

	return res.Plaintext, nil

}

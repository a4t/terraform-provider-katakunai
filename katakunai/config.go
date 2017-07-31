package katakunai

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/awslabs/aws-sdk-go/service/ec2"
)

type Config struct {
	Hogehoge string
}

func (c *Config) NewClient() (*ec2.EC2, error) {
	region := "ap-northeast-1"
	conf := &aws.Config{
		Region: &region,
	}
	sess, err := session.NewSession(conf)
	if err != nil {
		panic(err)
	}
	return ec2.New(sess), nil
}

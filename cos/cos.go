package cos

import (
	"cld/settings"
	"time"

	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

var c *sts.Client
var opt *sts.CredentialOptions

func Init(cfg *settings.CosConfig) error {
	var resource []string

	for _, v := range cfg.Resource.AllowKey {
		resource = append(resource, "qcs::cos:"+cfg.Resource.Region+":uid/"+cfg.Resource.Appid+":"+cfg.Resource.Bucket+v)
	}

	policy := &sts.CredentialPolicy{
		Statement: []sts.CredentialPolicyStatement{
			{
				Action:   cfg.Action,
				Effect:   "allow",
				Resource: resource,
			},
		},
	}
	opt = &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          cfg.Resource.Region,
		Policy:          policy,
	}

	c = sts.NewClient(cfg.TmpSecret.Id, cfg.TmpSecret.Key, nil)

	_, err := c.GetCredential(opt)
	if err != nil {
		return err
	}
	return nil
}

func GetKey() (*sts.CredentialResult, error) {
	res, err := c.GetCredential(opt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

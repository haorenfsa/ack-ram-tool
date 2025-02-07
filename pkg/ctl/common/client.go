package common

import (
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	"log"
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/alibabacloudsdkgo/helper/env"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/aliyuncli"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	"github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

func NewClient(regionId, credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName string) (*openapi.Client, error) {
	crd, err := getCredential(credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName, version.BinName())
	if err != nil {
		return nil, err
	}
	return openapi.NewClient(&client.Config{
		RegionId:   tea.String(regionId),
		Credential: crd,
		UserAgent:  tea.String(version.UserAgent()),
	})
}

func getCredential(credentialFilePath, aliyuncliConfigFilePath, aliyuncliProfileName, sessionName string) (credentials.Credential, error) {
	if credentialFilePath == "" && aliyuncliConfigFilePath == "" {
		if sessionName != "" {
			_ = os.Setenv(env.EnvRoleSessionName, sessionName)
		}
		if cred, err := env.NewCredential(); err == nil && cred != nil {
			log.Println("use credentials from environment variables")
			return cred, err
		}
	}
	if aliyuncliConfigFilePath == "" {
		aliyuncliConfigFilePath, _ = utils.ExpandPath("~/.aliyun/config.json")
	}

	acli, err := aliyuncli.NewCredentialHelper(aliyuncliConfigFilePath, aliyuncliProfileName)
	if err == nil && acli != nil {
		log.Printf("use credentials from aliyun cli (%s)", aliyuncliConfigFilePath)
		return acli.GetCredentials()
	}

	if credentialFilePath != "" {
		if _, err := os.Stat(credentialFilePath); err == nil {
			_ = os.Setenv(credentials.ENVCredentialFile, credentialFilePath)
		}
	} else {
		path, err := utils.ExpandPath(credentials.PATHCredentialFile)
		if err == nil {
			credentialFilePath = path
		}
	}
	log.Printf("use default credentials from %s", credentialFilePath)
	return credentials.NewCredential(nil)
}

func GetClientOrDie() *openapi.Client {
	c, err := NewClient(
		ctl.GlobalOption.Region,
		ctl.GlobalOption.CredentialFilePath,
		ctl.GlobalOption.AliyuncliConfigFilePath,
		ctl.GlobalOption.AliyuncliProfileName,
	)
	if err != nil {
		ExitByError(fmt.Sprintf("init client failed: %+v", err))
	}
	return c
}

package handler

import (
	"fmt"
	"github.com/gyepisam/secretsmanagerenv/pkg/aws"
	"os"
	"os/exec"
	"strings"
)

func RunCommandWithSecret(secrets []string, region string, args []string, upcase bool, prefix string) error {
	var env []string

    for _, secret := range secrets {
      if data, err := aws.GetSecretData(secret, region); err != nil {
        return err
      } else {
        moreEnv := mapToEnv(data, upcase, prefix)
        env = append(env, moreEnv...)
      }
    }

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func mapToEnv(m map[string]interface{}, upcase bool, prefix string) []string {
	var ret []string
	var k string
	for key, value := range m {
		k = key
		if upcase {
			k = strings.ToUpper(k)
		}
		if len(prefix) > 0 {
			k = prefix + k
		}
		keyval := fmt.Sprintf("%s=%v", k, value)
		ret = append(ret, keyval)
	}
	return ret
}

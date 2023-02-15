package tfcmd

func getImpersonationToken(executor Executor, accountToImpersonate string) (string, error) {
	cmdArgs := []string{"auth", "print-access-token", "--impersonate-service-account=" + accountToImpersonate}
	result, err := executor.Execute("gcloud", cmdArgs...)

	return string(result), err
}

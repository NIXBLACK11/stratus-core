package ping

func CheckStatus(url string, checks []string) []string{
	triggerd := make([]string, 0)
	
	for _, checkFor := range checks {
		switch checkFor {
		case "WebsiteUnreachable":
			if isWebsiteUnreachable(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "SSLCertificateExpiring":
			if isSSLCertificateExpiring(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "ServerResponseTimeHigh":
			if isServerResponseTimeHigh(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "BrokenLinks":
			if hasBrokenLinks(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "InternalServerError":
			if isInternalServerError(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "UnauthorizedAccess":
			if hasUnauthorizedAccess(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "DNSResolutionIssue":
			if isDNSResolutionIssue(url) {
				triggerd = append(triggerd, checkFor)
			}
		case "ServiceUnavailable":
			if isServiceUnavailable(url) {
				triggerd = append(triggerd, checkFor)
			}
		}
	}

	return triggerd
}
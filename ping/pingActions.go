package ping

import (
    "crypto/tls"
    "net/http"
    "time"
    "strings"
)

func isWebsiteUnreachable(url string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(url)
    if err != nil || resp.StatusCode != http.StatusOK {
        return true
    }
    return false
}

func isSSLCertificateExpiring(url string) bool {
    conn, err := tls.Dial("tcp", url+":443", nil)
    if err != nil {
        return true
    }
    expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
    daysUntilExpiry := expiry.Sub(time.Now()).Hours() / 24
    return daysUntilExpiry < 30 // Alert if SSL certificate expires in less than 30 days
}

func isServerResponseTimeHigh(url string) bool {
    start := time.Now()
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    _, err := client.Get(url)
    if err != nil {
        return true
    }
    duration := time.Since(start)
    return duration > 2*time.Second // Alert if response time exceeds 2 seconds
}

func hasBrokenLinks(url string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(url)
    if err != nil || resp.StatusCode == http.StatusNotFound {
        return true
    }
    return false
}

func isInternalServerError(url string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(url)
    if err != nil || resp.StatusCode == http.StatusInternalServerError {
        return true
    }
    return false
}

func hasUnauthorizedAccess(url string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(url)
    if err != nil || resp.StatusCode == http.StatusUnauthorized {
        return true
    }
    return false
}

func isDNSResolutionIssue(url string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    _, err := client.Get(url)
    if strings.Contains(err.Error(), "no such host") {
        return true
    }
    return false
}

func isServiceUnavailable(url string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(url)
    if err != nil || resp.StatusCode == http.StatusServiceUnavailable {
        return true
    }
    return false
}

/*
Not implemented yet

Check for server downtime (same as unreachable in this simple context)
func isServerDown(url string) bool {
    return isWebsiteUnreachable(url)
}

Check for high traffic spike (simplified)
func hasHighTraffic(url string) bool {
    // This would typically involve analyzing logs, here just a placeholder
    return false
}

Check for database connection error (simplified)
func isDatabaseConnectionError(url string) bool {
    // This would involve checking a database, here just a placeholder
    return false
}

func hasContentChanged(url string) bool {
    // This would involve tracking content, here just a placeholder
    return false
}

Check for suspicious activity (simplified)
func hasSuspiciousActivity(url string) bool {
    // This would involve security monitoring, here just a placeholder
    return false
}

Check for slow page load times
func isPageLoadTimeSlow(url string) bool {
    return isServerResponseTimeHigh(url)
}
*/
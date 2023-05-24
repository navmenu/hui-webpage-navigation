package utils_kratos

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/pkg/errors"
)

func GetCookie(ctx context.Context) (string, error) {
	tp, ok := transport.FromServerContext(ctx)
	if !ok {
		return "", errors.Errorf("no transporter from context")
	}
	cookie := tp.RequestHeader().Get("cookie")
	if cookie == "" {
		return "", errors.Errorf("no cookie")
	}
	return cookie, nil
}

func GetIp(ctx context.Context) (string, error) {
	tp, ok := transport.FromServerContext(ctx)
	if !ok {
		return "", errors.Errorf("no transporter from context")
	}
	ht, ok := tp.(*http.Transport)
	if !ok {
		return "", errors.Errorf("transporter is not http transport type")
	}
	ip := GetHttpRequestIp(ht.Request())
	if ip == "" {
		return "", errors.Errorf("no ip")
	}
	return ip, nil
}

func GetHttpRequestIp(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}

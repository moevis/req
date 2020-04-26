package req

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// create a default client
func newClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Jar:       jar,
		Transport: transport,
		Timeout:   2 * time.Minute,
	}
}

// Client return the default underlying http client
func (r *Req) Client() *http.Client {
	if r.client == nil {
		r.client = newClient()
	}
	return r.client
}

// Client return the default underlying http client
func Client() *http.Client {
	return std.Client()
}

// SetClient sets the underlying http.Client.
func (r *Req) SetClient(client *http.Client) {
	r.client = client // use default if client == nil
}

// SetClient sets the default http.Client for requests.
func SetClient(client *http.Client) {
	std.SetClient(client)
}

// SetFlags control display format of *Resp
func (r *Req) SetFlags(flags int) {
	r.flag = flags
}

// SetFlags control display format of *Resp
func SetFlags(flags int) {
	std.SetFlags(flags)
}

// Flags return output format for the *Resp
func (r *Req) Flags() int {
	return r.flag
}

// Flags return output format for the *Resp
func Flags() int {
	return std.Flags()
}

func (r *Req) getTransport() *http.Transport {
	trans, _ := r.Client().Transport.(*http.Transport)
	return trans
}

// EnableInsecureTLS allows insecure https
func (r *Req) EnableInsecureTLS(enable bool) {
	trans := r.getTransport()
	if trans == nil {
		return
	}
	if trans.TLSClientConfig == nil {
		trans.TLSClientConfig = &tls.Config{}
	}
	trans.TLSClientConfig.InsecureSkipVerify = enable
}

func EnableInsecureTLS(enable bool) {
	std.EnableInsecureTLS(enable)
}

// EnableCookieenable or disable cookie manager
func (r *Req) EnableCookie(enable bool) {
	if enable {
		jar, _ := cookiejar.New(nil)
		r.Client().Jar = jar
	} else {
		r.Client().Jar = nil
	}
}

// EnableCookieenable or disable cookie manager
func EnableCookie(enable bool) {
	std.EnableCookie(enable)
}

// SetTimeout sets the timeout for every request
func (r *Req) SetTimeout(d time.Duration) {
	r.Client().Timeout = d
}

// SetJSONEncoder sets the custmized encorder for json
func SetJSONEncoder(enc jsoniter.API) {
	std.jsonEncorder = enc
}

// SetJSONDecoder sets the custmized decorder for json
func SetJSONDecoder(enc jsoniter.API) {
	std.jsonDecorder = enc
}

// SetTimeout sets the timeout for every request
func SetTimeout(d time.Duration) {
	std.SetTimeout(d)
}

// SetProxyUrl set the simple proxy with fixed proxy url
func (r *Req) SetProxyUrl(rawurl string) error {
	trans := r.getTransport()
	if trans == nil {
		return errors.New("req: no transport")
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	trans.Proxy = http.ProxyURL(u)
	return nil
}

// SetProxyUrl set the simple proxy with fixed proxy url
func SetProxyUrl(rawurl string) error {
	return std.SetProxyUrl(rawurl)
}

// SetProxy sets the proxy for every request
func (r *Req) SetProxy(proxy func(*http.Request) (*url.URL, error)) error {
	trans := r.getTransport()
	if trans == nil {
		return errors.New("req: no transport")
	}
	trans.Proxy = proxy
	return nil
}

// SetProxy sets the proxy for every request
func SetProxy(proxy func(*http.Request) (*url.URL, error)) error {
	return std.SetProxy(proxy)
}

type xmlEncOpts struct {
	prefix string
	indent string
}

func (r *Req) getXMLEncOpts() *xmlEncOpts {
	if r.xmlEncOpts == nil {
		r.xmlEncOpts = &xmlEncOpts{}
	}
	return r.xmlEncOpts
}

// SetXMLIndent sets the encoder to generate XML in which each element
// begins on a new indented line that starts with prefix and is followed by
// one or more copies of indent according to the nesting depth.
func (r *Req) SetXMLIndent(prefix, indent string) {
	opts := r.getXMLEncOpts()
	opts.prefix = prefix
	opts.indent = indent
}

// SetXMLIndent sets the encoder to generate XML in which each element
// begins on a new indented line that starts with prefix and is followed by
// one or more copies of indent according to the nesting depth.
func SetXMLIndent(prefix, indent string) {
	std.SetXMLIndent(prefix, indent)
}

// SetProgressInterval sets the progress reporting interval of both
// UploadProgress and DownloadProgress handler
func (r *Req) SetProgressInterval(interval time.Duration) {
	r.progressInterval = interval
}

// SetProgressInterval sets the progress reporting interval of both
// UploadProgress and DownloadProgress handler for the default client
func SetProgressInterval(interval time.Duration) {
	std.SetProgressInterval(interval)
}

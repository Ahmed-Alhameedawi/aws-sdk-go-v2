// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_XmlTimestamps_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *XmlTimestampsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Tests how normal timestamps are serialized
		"XmlTimestamps": {
			Params: &XmlTimestampsInput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <normal>2014-04-29T18:30:38Z</normal>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"XmlTimestampsWithDateTimeFormat": {
			Params: &XmlTimestampsInput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <dateTime>2014-04-29T18:30:38Z</dateTime>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of date-time on the target shape works like
		// normal timestamps
		"XmlTimestampsWithDateTimeOnTargetFormat": {
			Params: &XmlTimestampsInput{
				DateTimeOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <dateTimeOnTarget>2014-04-29T18:30:38Z</dateTimeOnTarget>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"XmlTimestampsWithEpochSecondsFormat": {
			Params: &XmlTimestampsInput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <epochSeconds>1398796238</epochSeconds>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of epoch-seconds on the target shape works
		"XmlTimestampsWithEpochSecondsOnTargetFormat": {
			Params: &XmlTimestampsInput{
				EpochSecondsOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <epochSecondsOnTarget>1398796238</epochSecondsOnTarget>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of http-date works
		"XmlTimestampsWithHttpDateFormat": {
			Params: &XmlTimestampsInput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <httpDate>Tue, 29 Apr 2014 18:30:38 GMT</httpDate>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of http-date on the target shape works
		"XmlTimestampsWithHttpDateOnTargetFormat": {
			Params: &XmlTimestampsInput{
				HttpDateOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <httpDateOnTarget>Tue, 29 Apr 2014 18:30:38 GMT</httpDateOnTarget>
			</XmlTimestampsInputOutput>
			`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actualReq := &http.Request{}
			serverURL := "http://localhost:8888/"
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               protocoltesthttp.NewClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.XmlTimestamps(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyprivateprotocol.AddCaptureRequestMiddleware(stack, actualReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_XmlTimestamps_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlTimestampsOutput
	}{
		// Tests how normal timestamps are serialized
		"XmlTimestamps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <normal>2014-04-29T18:30:38Z</normal>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"XmlTimestampsWithDateTimeFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <dateTime>2014-04-29T18:30:38Z</dateTime>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time on the target shape works like
		// normal timestamps
		"XmlTimestampsWithDateTimeOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <dateTimeOnTarget>2014-04-29T18:30:38Z</dateTimeOnTarget>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				DateTimeOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"XmlTimestampsWithEpochSecondsFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <epochSeconds>1398796238</epochSeconds>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds on the target shape works
		"XmlTimestampsWithEpochSecondsOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <epochSecondsOnTarget>1398796238</epochSecondsOnTarget>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				EpochSecondsOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date works
		"XmlTimestampsWithHttpDateFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <httpDate>Tue, 29 Apr 2014 18:30:38 GMT</httpDate>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date on the target shape works
		"XmlTimestampsWithHttpDateOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <httpDateOnTarget>Tue, 29 Apr 2014 18:30:38 GMT</httpDateOnTarget>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				HttpDateOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params XmlTimestampsInput
			result, err := client.XmlTimestamps(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}

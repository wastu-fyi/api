package resp

import (
	"net/http"
	"strings"

	gh "github.com/goravel/framework/contracts/http"
)

type Body struct {
	StatusCode int         `json:"statusCode"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

var defaultCodes = map[int]string{
	http.StatusOK:                  "OK",
	http.StatusCreated:             "CREATED",
	http.StatusAccepted:            "ACCEPTED",
	http.StatusBadRequest:          "BAD_REQUEST",
	http.StatusUnauthorized:        "UNAUTHORIZED",
	http.StatusForbidden:           "FORBIDDEN",
	http.StatusNotFound:            "NOT_FOUND",
	http.StatusMethodNotAllowed:    "METHOD_NOT_ALLOWED",
	http.StatusNotAcceptable:       "NOT_ACCEPTABLE",
	http.StatusTooManyRequests:     "TOO_MANY_REQUESTS",
	http.StatusInternalServerError: "INTERNAL_SERVER_ERROR",
	http.StatusNotImplemented:      "NOT_IMPLEMENTED",
	http.StatusServiceUnavailable:  "SERVICE_UNAVAILABLE",
}

type Option func(*Body, *int)

func WithCode(code string) Option {
	return func(b *Body, _ *int) {
		if code != "" {
			b.Code = strings.ToUpper(code)
		}
	}
}

func WithMessage(msg string) Option {
	return func(b *Body, _ *int) {
		if msg != "" {
			b.Message = msg
		}
	}
}

func WithStatus(status int) Option {
	return func(_ *Body, s *int) {
		*s = status
	}
}

func Write(ctx gh.Context, status int, message string, data interface{}, opts ...Option) gh.Response {
	code := defaultCodes[status]

	if code == "" {
		code = "UNKNOWN"
	}

	body := Body{
		StatusCode: status,
		Code:       code,
		Message:    message,
		Data:       data,
	}

	for _, opt := range opts {
		opt(&body, &status)
	}

	body.StatusCode = status
	return ctx.Response().Json(status, body)
}

func OK(ctx gh.Context, data interface{}, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusOK, msg, data, opts...)
}

func Created(ctx gh.Context, data interface{}, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusCreated, msg, data, opts...)
}

func Accepted(ctx gh.Context, data interface{}, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusAccepted, msg, data, opts...)
}

func BadRequest(ctx gh.Context, msg string, data interface{}, opts ...Option) gh.Response {
	return Write(ctx, http.StatusBadRequest, msg, data, opts...)
}

func Unauthorized(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusUnauthorized, msg, nil, opts...)
}

func Forbidden(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusForbidden, msg, nil, opts...)
}

func NotFound(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusNotFound, msg, nil, opts...)
}

func MethodNotAllowed(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusMethodNotAllowed, msg, nil, opts...)
}

func NotAcceptable(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusNotAcceptable, msg, nil, opts...)
}

func TooManyRequests(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusTooManyRequests, msg, nil, opts...)
}

func InternalServerError(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusInternalServerError, msg, nil, opts...)
}

func NotImplemented(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusNotImplemented, msg, nil, opts...)
}

func ServiceUnavailable(ctx gh.Context, msg string, opts ...Option) gh.Response {
	return Write(ctx, http.StatusServiceUnavailable, msg, nil, opts...)
}

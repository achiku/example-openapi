// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeAddPetRequest(
	req *Pet,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateUserRequest(
	req *User,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateUsersWithArrayInputRequest(
	req []User,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		e.ArrStart()
		for _, elem := range req {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCreateUsersWithListInputRequest(
	req []User,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		e.ArrStart()
		for _, elem := range req {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePlaceOrderRequest(
	req *Order,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdatePetRequest(
	req *Pet,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUpdatePetWithFormRequest(
	req OptUpdatePetWithFormReq,
	r *http.Request,
) error {
	const contentType = "application/x-www-form-urlencoded"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "name" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Name.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "status" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.Status.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	encoded := q.Values().Encode()
	ht.SetBody(r, strings.NewReader(encoded), contentType)
	return nil
}

func encodeUpdateUserRequest(
	req *User,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := new(jx.Encoder)
	{
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeUploadFileRequest(
	req OptUploadFileReq,
	r *http.Request,
) error {
	const contentType = "multipart/form-data"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	request := req.Value

	q := uri.NewFormEncoder(map[string]string{})
	{
		// Encode "additionalMetadata" form field.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "additionalMetadata",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}
		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := request.AdditionalMetadata.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "encode query")
		}
	}
	body, boundary := ht.CreateMultipartBody(func(w *multipart.Writer) error {
		if val, ok := request.File.Get(); ok {
			if err := val.WriteMultipart("file", w); err != nil {
				return errors.Wrap(err, "write \"file\"")
			}
		}
		if err := q.WriteMultipart(w); err != nil {
			return errors.Wrap(err, "write multipart")
		}
		return nil
	})
	ht.SetCloserBody(r, body, mime.FormatMediaType(contentType, map[string]string{"boundary": boundary}))
	return nil
}
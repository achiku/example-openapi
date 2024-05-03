// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeAddPetResponse(response AddPetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *AddPetMethodNotAllowed:
		w.WriteHeader(405)
		span.SetStatus(codes.Error, http.StatusText(405))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreateUserResponse(response *CreateUserDef, w http.ResponseWriter, span trace.Span) error {
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodeCreateUsersWithArrayInputResponse(response *CreateUsersWithArrayInputDef, w http.ResponseWriter, span trace.Span) error {
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodeCreateUsersWithListInputResponse(response *CreateUsersWithListInputDef, w http.ResponseWriter, span trace.Span) error {
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodeDeleteOrderResponse(response DeleteOrderRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteOrderBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *DeleteOrderNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeletePetResponse(response *DeletePetBadRequest, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(400)
	span.SetStatus(codes.Error, http.StatusText(400))

	return nil
}

func encodeDeleteUserResponse(response DeleteUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *DeleteUserNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeFindPetsByStatusResponse(response FindPetsByStatusRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FindPetsByStatusOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *FindPetsByStatusBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeFindPetsByTagsResponse(response FindPetsByTagsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FindPetsByTagsOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *FindPetsByTagsBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetInventoryResponse(response GetInventoryOK, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

func encodeGetOrderByIdResponse(response GetOrderByIdRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Order:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetOrderByIdBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *GetOrderByIdNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetPetByIdResponse(response GetPetByIdRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetPetByIdBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *GetPetByIdNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetUserByNameResponse(response GetUserByNameRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *User:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetUserByNameBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *GetUserByNameNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeLoginUserResponse(response LoginUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *LoginUserOKHeaders:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Set-Cookie" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Set-Cookie",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.SetCookie.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Set-Cookie header")
				}
			}
			// Encode "X-Expires-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Expires-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.XExpiresAfter.Get(); ok {
						return e.EncodeValue(conv.DateTimeToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode X-Expires-After header")
				}
			}
			// Encode "X-Rate-Limit" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Rate-Limit",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.XRateLimit.Get(); ok {
						return e.EncodeValue(conv.Int32ToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode X-Rate-Limit header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		e.Str(response.Response)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *LoginUserOKApplicationXMLHeaders:
		w.Header().Set("Content-Type", "application/xml")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Set-Cookie" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Set-Cookie",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.SetCookie.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Set-Cookie header")
				}
			}
			// Encode "X-Expires-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Expires-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.XExpiresAfter.Get(); ok {
						return e.EncodeValue(conv.DateTimeToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode X-Expires-After header")
				}
			}
			// Encode "X-Rate-Limit" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "X-Rate-Limit",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.XRateLimit.Get(); ok {
						return e.EncodeValue(conv.Int32ToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode X-Rate-Limit header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response.Response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *LoginUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeLogoutUserResponse(response *LogoutUserDef, w http.ResponseWriter, span trace.Span) error {
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil
}

func encodePlaceOrderResponse(response PlaceOrderRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Order:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PlaceOrderBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdatePetResponse(response UpdatePetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UpdatePetBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *UpdatePetNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *UpdatePetMethodNotAllowed:
		w.WriteHeader(405)
		span.SetStatus(codes.Error, http.StatusText(405))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdatePetWithFormResponse(response *UpdatePetWithFormMethodNotAllowed, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(405)
	span.SetStatus(codes.Error, http.StatusText(405))

	return nil
}

func encodeUpdateUserResponse(response UpdateUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UpdateUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *UpdateUserNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUploadFileResponse(response *ApiResponse, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := new(jx.Encoder)
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	return nil
}

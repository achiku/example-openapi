// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteOrderParams is parameters of deleteOrder operation.
type DeleteOrderParams struct {
	// ID of the order that needs to be deleted.
	OrderId string
}

func unpackDeleteOrderParams(packed middleware.Parameters) (params DeleteOrderParams) {
	{
		key := middleware.ParameterKey{
			Name: "orderId",
			In:   "path",
		}
		params.OrderId = packed[key].(string)
	}
	return params
}

func decodeDeleteOrderParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteOrderParams, _ error) {
	// Decode path: orderId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "orderId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.OrderId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "orderId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeletePetParams is parameters of deletePet operation.
type DeletePetParams struct {
	APIKey OptString
	// Pet id to delete.
	PetId int64
}

func unpackDeletePetParams(packed middleware.Parameters) (params DeletePetParams) {
	{
		key := middleware.ParameterKey{
			Name: "api_key",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.APIKey = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "petId",
			In:   "path",
		}
		params.PetId = packed[key].(int64)
	}
	return params
}

func decodeDeletePetParams(args [1]string, argsEscaped bool, r *http.Request) (params DeletePetParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: api_key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "api_key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotAPIKeyVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAPIKeyVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.APIKey.SetTo(paramsDotAPIKeyVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "api_key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: petId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "petId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteUserParams is parameters of deleteUser operation.
type DeleteUserParams struct {
	// The name that needs to be deleted.
	Username string
}

func unpackDeleteUserParams(packed middleware.Parameters) (params DeleteUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "username",
			In:   "path",
		}
		params.Username = packed[key].(string)
	}
	return params
}

func decodeDeleteUserParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteUserParams, _ error) {
	// Decode path: username.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "username",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Username = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "username",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// FindPetsByStatusParams is parameters of findPetsByStatus operation.
type FindPetsByStatusParams struct {
	// Status values that need to be considered for filter.
	//
	// Deprecated: schema marks this parameter as deprecated.
	Status []FindPetsByStatusStatusItem
}

func unpackFindPetsByStatusParams(packed middleware.Parameters) (params FindPetsByStatusParams) {
	{
		key := middleware.ParameterKey{
			Name: "status",
			In:   "query",
		}
		params.Status = packed[key].([]FindPetsByStatusStatusItem)
	}
	return params
}

func decodeFindPetsByStatusParams(args [0]string, argsEscaped bool, r *http.Request) (params FindPetsByStatusParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: status.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: false,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotStatusVal FindPetsByStatusStatusItem
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotStatusVal = FindPetsByStatusStatusItem(c)
						return nil
					}(); err != nil {
						return err
					}
					params.Status = append(params.Status, paramsDotStatusVal)
					return nil
				})
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Status == nil {
					return errors.New("nil is invalid value")
				}
				var failures []validate.FieldError
				for i, elem := range params.Status {
					if err := func() error {
						if err := elem.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "status",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// FindPetsByTagsParams is parameters of findPetsByTags operation.
type FindPetsByTagsParams struct {
	// Tags to filter by.
	Tags []string
}

func unpackFindPetsByTagsParams(packed middleware.Parameters) (params FindPetsByTagsParams) {
	{
		key := middleware.ParameterKey{
			Name: "tags",
			In:   "query",
		}
		params.Tags = packed[key].([]string)
	}
	return params
}

func decodeFindPetsByTagsParams(args [0]string, argsEscaped bool, r *http.Request) (params FindPetsByTagsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: tags.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "tags",
			Style:   uri.QueryStyleForm,
			Explode: false,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotTagsVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotTagsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.Tags = append(params.Tags, paramsDotTagsVal)
					return nil
				})
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Tags == nil {
					return errors.New("nil is invalid value")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "tags",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetOrderByIdParams is parameters of getOrderById operation.
type GetOrderByIdParams struct {
	// ID of pet that needs to be fetched.
	OrderId int64
}

func unpackGetOrderByIdParams(packed middleware.Parameters) (params GetOrderByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "orderId",
			In:   "path",
		}
		params.OrderId = packed[key].(int64)
	}
	return params
}

func decodeGetOrderByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetOrderByIdParams, _ error) {
	// Decode path: orderId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "orderId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.OrderId = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        true,
					Max:           5,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.OrderId)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "orderId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetPetByIdParams is parameters of getPetById operation.
type GetPetByIdParams struct {
	// ID of pet to return.
	PetId int64
}

func unpackGetPetByIdParams(packed middleware.Parameters) (params GetPetByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "petId",
			In:   "path",
		}
		params.PetId = packed[key].(int64)
	}
	return params
}

func decodeGetPetByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetPetByIdParams, _ error) {
	// Decode path: petId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "petId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetUserByNameParams is parameters of getUserByName operation.
type GetUserByNameParams struct {
	// The name that needs to be fetched. Use user1 for testing.
	Username string
}

func unpackGetUserByNameParams(packed middleware.Parameters) (params GetUserByNameParams) {
	{
		key := middleware.ParameterKey{
			Name: "username",
			In:   "path",
		}
		params.Username = packed[key].(string)
	}
	return params
}

func decodeGetUserByNameParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUserByNameParams, _ error) {
	// Decode path: username.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "username",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Username = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "username",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// LoginUserParams is parameters of loginUser operation.
type LoginUserParams struct {
	// The user name for login.
	Username string
	// The password for login in clear text.
	Password string
}

func unpackLoginUserParams(packed middleware.Parameters) (params LoginUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "username",
			In:   "query",
		}
		params.Username = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "password",
			In:   "query",
		}
		params.Password = packed[key].(string)
	}
	return params
}

func decodeLoginUserParams(args [0]string, argsEscaped bool, r *http.Request) (params LoginUserParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: username.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "username",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Username = c
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        regexMap["^[a-zA-Z0-9]+[a-zA-Z0-9\\.\\-_]*[a-zA-Z0-9]+$"],
				}).Validate(string(params.Username)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "username",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: password.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "password",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Password = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "password",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// UpdatePetWithFormParams is parameters of updatePetWithForm operation.
type UpdatePetWithFormParams struct {
	// ID of pet that needs to be updated.
	PetId int64
}

func unpackUpdatePetWithFormParams(packed middleware.Parameters) (params UpdatePetWithFormParams) {
	{
		key := middleware.ParameterKey{
			Name: "petId",
			In:   "path",
		}
		params.PetId = packed[key].(int64)
	}
	return params
}

func decodeUpdatePetWithFormParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdatePetWithFormParams, _ error) {
	// Decode path: petId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "petId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateUserParams is parameters of updateUser operation.
type UpdateUserParams struct {
	// Name that need to be deleted.
	Username string
}

func unpackUpdateUserParams(packed middleware.Parameters) (params UpdateUserParams) {
	{
		key := middleware.ParameterKey{
			Name: "username",
			In:   "path",
		}
		params.Username = packed[key].(string)
	}
	return params
}

func decodeUpdateUserParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateUserParams, _ error) {
	// Decode path: username.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "username",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Username = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "username",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UploadFileParams is parameters of uploadFile operation.
type UploadFileParams struct {
	// ID of pet to update.
	PetId int64
}

func unpackUploadFileParams(packed middleware.Parameters) (params UploadFileParams) {
	{
		key := middleware.ParameterKey{
			Name: "petId",
			In:   "path",
		}
		params.PetId = packed[key].(int64)
	}
	return params
}

func decodeUploadFileParams(args [1]string, argsEscaped bool, r *http.Request) (params UploadFileParams, _ error) {
	// Decode path: petId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "petId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
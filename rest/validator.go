package rest

import (
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	if err := ValidateStruct(body); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(err, ","),
		}
	}

	return nil

}

func ParseAndValidatePartially(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	if err := ValidateStructPartially(body); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(err, ","),
		}
	}

	return nil
}

func ValidateStruct(model interface{}) []string {
	var errors []string
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, "`%v` doesn't satisfy the `%v` constraint", err.Field(), err.Tag())
		}
	}
	return errors
}

func ValidateStructPartially(model interface{}) []string {
	var errors []string
	validate := validator.New()
	validate.SetTagName("partial_validate")
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, "`%v` doesn't satisfy the `%v` constraint", err.Field(), err.Tag())
		}
	}
	return errors
}

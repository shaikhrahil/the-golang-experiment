package rest

import (
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

// parseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func parseBody(ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) error {
	if err := parseBody(ctx, body); err != nil {
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

func ParseAndValidatePartially(ctx *fiber.Ctx, body interface{}) error {
	if err := parseBody(ctx, body); err != nil {
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

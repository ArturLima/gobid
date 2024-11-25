package user

import (
	"context"

	"github.com/ArturLima/gobid/internal/validator"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "email must be a valid")
	eval.CheckField(validator.NotBlank(req.Password), "password", "this field cannot be blank")

	return eval
}

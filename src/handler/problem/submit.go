package problem

import (
	"encoding/base64"

	"leita/src/entity"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// SubmitProblem godoc
//
//	@Summary		Submit a problem solution
//	@Description	Submit code for a specific problem to be judged.
//	@Accept			json
//	@Produce		json
//	@Tags			Problem
//	@Param			problemId	path		string						true	"Problem ID"
//	@Param			requestBody	body		entity.SubmitProblemRequest	true	"Solution code and metadata"
//	@Success		200			{object}	entity.SubmitProblemResponse
//	@Failure		400			{object}	entity.SubmitProblemResponse
//	@Failure		500			{object}	entity.SubmitProblemResponse
//	@Router			/problem/submit/{problemId} [post]
func (handler *Handler) SubmitProblem() fiber.Handler {
	return func(c fiber.Ctx) error {
		var req entity.SubmitProblemRequest
		if err := c.Bind().Body(&req); err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(entity.SubmitProblemResponse{
				Error: err.Error(),
			})
		}

		code, err := base64.StdEncoding.DecodeString(req.Code)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(entity.SubmitProblemResponse{
				Error: err.Error(),
			})
		}

		dto := entity.SubmitProblemDTO{
			ProblemId: c.Params("problemId"),
			SubmitId:  req.SubmitId,
			Language:  req.Language,
			Code:      code,
			Limit:     req.Limit,
		}

		result, usedTime, usedMemory, err := handler.service.SubmitProblem(dto)
		if result == entity.JudgeUnknown {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(entity.SubmitProblemResponse{
				Result: entity.JudgeUnknown.String(),
				Error:  err.Error(),
			})
		}

		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		return c.Status(fiber.StatusOK).JSON(entity.SubmitProblemResponse{
			Result:     result.String(),
			Error:      errStr,
			UsedTime:   usedTime,
			UsedMemory: usedMemory,
		})
	}
}

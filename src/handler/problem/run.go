package problem

import (
	"encoding/base64"

	"leita/src/entity"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// RunProblem godoc
//
//	@Summary		Run code with test cases
//	@Description	Execute code against provided test cases for testing purposes.
//	@Accept			json
//	@Produce		json
//	@Tags			Problem
//	@Param			problemId	path		string					true	"Problem ID"
//	@Param			requestBody	body		entity.RunProblemRequest	true	"Code and test cases"
//	@Success		200			{object}	[]entity.RunProblemResponse
//	@Failure		400			{object}	[]entity.RunProblemResponse
//	@Failure		500			{object}	[]entity.RunProblemResponse
//	@Router			/problem/run/{problemId} [post]
func (handler *Handler) RunProblem() fiber.Handler {
	return func(c fiber.Ctx) error {
		var req entity.RunProblemRequest
		if err := c.Bind().Body(&req); err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON([]entity.RunProblemResponse{
				{Error: err.Error()},
			})
		}

		code, err := base64.StdEncoding.DecodeString(req.Code)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON([]entity.RunProblemResponse{
				{Error: err.Error()},
			})
		}

		dto := entity.RunProblemDTO{
			ProblemId: c.Params("problemId"),
			Language:  req.Language,
			Code:      code,
			Limit:     req.Limit,
			TestCases: req.TestCases,
		}

		results := handler.service.RunProblem(dto)

		responses := make([]entity.RunProblemResponse, 0, len(results))
		for _, result := range results {
			errStr := ""
			if result.Error != nil {
				errStr = result.Error.Error()
			}
			responses = append(responses, entity.RunProblemResponse{
				Result: result.Result.String(),
				Error:  errStr,
				Output: result.Output,
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses)
	}
}

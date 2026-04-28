package problem

import (
	"math"

	. "leita/src/entities"
	. "leita/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// RunProblem godoc
//
//	@Accept		json
//	@Produce	json
//	@Tags		Problem
//	@Param		problemId	path		string				true	"problemId"
//	@Param		requestBody	body		RunProblemRequest	true	"requestBody"
//	@Success	200			{object}	[]RunProblemResponse
//	@Failure	500			{object}	[]RunProblemResponse
//	@Router		/problem/run/{problemId} [post]
func (handler *ProblemHandler) RunProblem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RunProblemRequest
		if err := c.BodyParser(&req); err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON([]RunProblemResponse{
				{
					Error: err.Error(),
				},
			})
		}

		problemId := c.Params("problemId")
		language := req.Language
		code := DecodeBase64([]byte(req.Code))
		testCases := req.TestCases
		submitId := RandomInt(int(math.Pow10(11)), int(math.Pow10(12)-1))
		command := Commands[language]
		buildCmd := ReplaceCommand(command.BuildCmd, "run", submitId)
		runCmd := ReplaceCommand(command.RunCmd, "run", submitId)
		deleteCmd := ReplaceCommand(command.DeleteCmd, "run", submitId)

		runProblemDTO := RunProblemDTO{
			ProblemId: problemId,
			SubmitId:  submitId,
			Language:  language,
			Code:      code,
			Limit:     req.Limit,
			TestCases: testCases,
			BuildCmd:  buildCmd,
			RunCmd:    runCmd,
			DeleteCmd: deleteCmd,
		}

		results := handler.service.RunProblem(runProblemDTO)

		responses := make([]RunProblemResponse, 0, len(results))
		for _, result := range results {
			responses = append(responses, RunProblemResponse{
				Result: result.Result.String(),
				Error: ErrStrIfNotNil(result.Error),
				Output: result.Output,
			})
		}

		return c.Status(fiber.StatusOK).JSON(responses)
	}
}

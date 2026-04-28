package problem

import (
	. "leita/src/entities"
	. "leita/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// SubmitProblem godoc
//
//	@Accept		json
//	@Produce	json
//	@Tags		Problem
//	@Param		problemId	path		string					true	"problemId"
//	@Param		requestBody	body		SubmitProblemRequest	true	"requestBody"
//	@Success	200			{object}	SubmitProblemResponse
//	@Failure	500			{object}	SubmitProblemResponse
//	@Router		/problem/submit/{problemId} [post]
func (handler *ProblemHandler) SubmitProblem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req SubmitProblemRequest
		if err := c.BodyParser(&req); err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(SubmitProblemResponse{
				Error: err.Error(),
			})
		}

		problemId := c.Params("problemId")
		submitId := req.SubmitId
		language := req.Language
		code := DecodeBase64([]byte(req.Code))
		command := Commands[language]
		buildCmd := ReplaceCommand(command.BuildCmd, "submit", submitId)
		runCmd := ReplaceCommand(command.RunCmd, "submit", submitId)
		deleteCmd := ReplaceCommand(command.DeleteCmd, "submit", submitId)

		submitProblemDTO := SubmitProblemDTO{
			ProblemId: problemId,
			SubmitId:  submitId,
			Language:  language,
			Code:      code,
			Limit:     req.Limit,
			BuildCmd:  buildCmd,
			RunCmd:    runCmd,
			DeleteCmd: deleteCmd,
		}

		result, usedTime, usedMemory, err := handler.service.SubmitProblem(submitProblemDTO)
		if result == JudgeUnknown {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(SubmitProblemResponse{
				Result: JudgeUnknown.String(),
				Error:  err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(SubmitProblemResponse{
			Result:     result.String(),
			Error: ErrStrIfNotNil(err),
			UsedTime:   usedTime,
			UsedMemory: usedMemory,
		})
	}
}

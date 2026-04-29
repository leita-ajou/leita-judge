package executor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"leita/src/entity"

	"github.com/gofiber/fiber/v3/log"
)

type Executor interface {
	Build(buildCmd []string) (entity.JudgeResultEnum, error)
	Run(runCmd []string, input []byte, timeLimit int) (entity.JudgeResultEnum, []byte, int64, int64, error)
	Delete(deleteCmd []string) error
}

type OsExecutor struct{}

func NewOsExecutor() *OsExecutor {
	return &OsExecutor{}
}

func (e *OsExecutor) Build(buildCmd []string) (entity.JudgeResultEnum, error) {
	log.Info("--------------------------------")
	log.Info("소스 코드 빌드 중...")

	if len(buildCmd) == 0 {
		log.Info("빌드 생략")
		return entity.JudgeCorrect, nil
	}

	var stderr bytes.Buffer
	cmd := exec.Command(buildCmd[0], buildCmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		compileError := fmt.Errorf("\n%w\n%s", err, stderr.String())
		log.Error(compileError)
		return entity.JudgeCompileError, compileError
	}

	log.Info("소스 코드 빌드 완료!")
	return entity.JudgeCorrect, nil
}

func (e *OsExecutor) Run(runCmd []string, input []byte, timeLimit int) (entity.JudgeResultEnum, []byte, int64, int64, error) {
	log.Info("프로그램 실행 중...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeLimit)*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, runCmd[0], runCmd[1:]...)
	cmd.Stdin = bytes.NewReader(input)

	var outputBuffer bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		log.Error(err)
		return entity.JudgeRuntimeError, nil, 0, 0, err
	}

	startTime := time.Now()
	err := cmd.Wait()
	usedTime := time.Since(startTime).Milliseconds()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Error(ctx.Err().Error())
		return entity.JudgeTimeOut, nil, 0, 0, ctx.Err()
	}

	if err != nil {
		runtimeError := fmt.Errorf("\n%w\n%s", err, stderr.String())
		log.Error(runtimeError)
		return entity.JudgeRuntimeError, nil, 0, 0, runtimeError
	}

	output := outputBuffer.Bytes()
	output = bytes.TrimRight(output, "\n\r\t ")

	return entity.JudgeCorrect, output, usedTime, 0, nil
}

func (e *OsExecutor) Delete(deleteCmd []string) error {
	log.Info("--------------------------------")
	log.Info("생성된 실행 파일 삭제 중...")

	if len(deleteCmd) == 0 {
		log.Info("삭제 생략")
		return nil
	}

	var stderr bytes.Buffer
	cmd := exec.Command(deleteCmd[0], deleteCmd[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		deleteError := fmt.Errorf("\n%w\n%s", err, stderr.String())
		log.Error(deleteError)
		return deleteError
	}

	log.Info("실행 파일 삭제 완료!")
	return nil
}

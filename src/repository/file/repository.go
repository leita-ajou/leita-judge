package file

import (
	"os"
	"path/filepath"
	"strconv"

	"leita/src/entity"
	"leita/src/language"

	"github.com/gofiber/fiber/v3/log"
)

type Repository interface {
	SaveSourceCode(submitId int, code []byte, lang, judgeType string) error
	SaveTestCases(submitId int, inputs, outputs [][]byte, judgeType string) error
	ReadInput(submitId int, index int, judgeType string) ([]byte, error)
	ReadOutput(submitId int, index int, judgeType string) ([]byte, error)
	GetTestCaseCount(submitId int, judgeType string) (int, error)
}

type LocalRepository struct{}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{}
}

func (r *LocalRepository) SaveSourceCode(submitId int, code []byte, lang, judgeType string) error {
	log.Info("--------------------------------")
	log.Info("소스 코드 저장 중...")

	dir := filepath.Join(judgeType, strconv.Itoa(submitId))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Error(err)
		return err
	}

	path := filepath.Join(dir, language.FileName+"."+language.FileExtension(lang))
	if err := os.WriteFile(path, code, 0644); err != nil {
		log.Error(err)
		return err
	}

	log.Info("소스 코드 저장 완료!")
	return nil
}

func (r *LocalRepository) SaveTestCases(submitId int, inputs, outputs [][]byte, judgeType string) error {
	log.Info("--------------------------------")
	log.Info("테스트 케이스 저장 중...")

	inDir := filepath.Join(judgeType, strconv.Itoa(submitId), "in")
	if err := os.MkdirAll(inDir, os.ModePerm); err != nil {
		log.Error(err)
		return err
	}

	outDir := filepath.Join(judgeType, strconv.Itoa(submitId), "out")
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Error(err)
		return err
	}

	for i := range inputs {
		inPath := filepath.Join(inDir, strconv.Itoa(i)+".in")
		if err := os.WriteFile(inPath, inputs[i], 0644); err != nil {
			log.Error(err)
			return err
		}

		outPath := filepath.Join(outDir, strconv.Itoa(i)+".out")
		if err := os.WriteFile(outPath, outputs[i], 0644); err != nil {
			log.Error(err)
			return err
		}
	}

	log.Info(len(inputs), "개 테스트 케이스 저장 완료!")
	return nil
}

func (r *LocalRepository) ReadInput(submitId int, index int, judgeType string) ([]byte, error) {
	path := filepath.Join(judgeType, strconv.Itoa(submitId), "in", strconv.Itoa(index)+".in")
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, nil
}

func (r *LocalRepository) ReadOutput(submitId int, index int, judgeType string) ([]byte, error) {
	path := filepath.Join(judgeType, strconv.Itoa(submitId), "out", strconv.Itoa(index)+".out")
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, nil
}

func (r *LocalRepository) GetTestCaseCount(submitId int, judgeType string) (int, error) {
	path := filepath.Join(judgeType, strconv.Itoa(submitId), "in")
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return len(entries), nil
}

// TestCasesFromRequest는 entity.TestCase 슬라이스를 입력/출력 바이트 슬라이스로 변환한다.
func TestCasesFromRequest(testCases []entity.TestCase) (inputs [][]byte, outputs [][]byte) {
	inputs = make([][]byte, len(testCases))
	outputs = make([][]byte, len(testCases))
	for i, tc := range testCases {
		inputs[i] = []byte(tc.Input)
		outputs[i] = []byte(tc.Output)
	}
	return
}

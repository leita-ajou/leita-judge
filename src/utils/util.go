package utils

import (
	"encoding/base64"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func DecodeBase64(data []byte) []byte {
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(decodedData, data)
	if err != nil {
		log.Error(err)
		return []byte{}
	}
	decodedData = decodedData[:n]

	return decodedData
}

func EncodeBase64(data []byte) []byte {
	encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encodedData, data)

	return encodedData
}

func FileExtension(language string) string {
	switch language {
	case "C":
		return "c"
	case "CPP":
		return "cpp"
	case "GO":
		return "go"
	case "JAVA":
		return "java"
	case "JAVASCRIPT":
		return "js"
	case "KOTLIN":
		return "kt"
	case "PYTHON":
		return "py"
	case "SWIFT":
		return "swift"
	default:
		return "error"
	}
}

func GetTestCaseNum(path string) (int, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return len(entries), nil
}

func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func MakeDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func ReplaceCommand(args []string, judgeType string, submitID int) []string {
	replacer := strings.NewReplacer(
		"{JUDGE_TYPE}", judgeType,
		"{SUBMIT_ID}", strconv.Itoa(submitID),
	)
	replaced := make([]string, len(args))
	for i, arg := range args {
		replaced[i] = replacer.Replace(arg)
	}
	return replaced
}

func ErrStrIfNotNil(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

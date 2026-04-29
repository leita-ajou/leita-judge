package language

import (
	"strconv"
	"strings"
)

const FileName = "Main"

type Command struct {
	BuildCmd  []string
	RunCmd    []string
	DeleteCmd []string
}

var Commands = map[string]Command{
	"C": {
		BuildCmd:  []string{"gcc", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".c", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName, "-O2", "-Wall", "-lm", "-static", "-std=gnu99"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
	},
	"CPP": {
		BuildCmd:  []string{"g++", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".cpp", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName, "-O2", "-Wall", "-lm", "-static", "-std=gnu++17"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
	},
	"JAVA": {
		BuildCmd:  []string{"javac", "-J-Xms1024m", "-J-Xmx1920m", "-J-Xss512m", "-encoding", "UTF-8", "-d", "bin", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".java"},
		RunCmd:    []string{"java", "-Xms1024m", "-Xmx1920m", "-Xss512m", "-Dfile.encoding=UTF-8", "-XX:+UseSerialGC", "-cp", "bin", FileName},
		DeleteCmd: []string{"rm", "-r", "{JUDGE_TYPE}/{SUBMIT_ID}/bin"},
	},
	"PYTHON": {
		BuildCmd:  []string{},
		RunCmd:    []string{"python3", "-W", "ignore", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".py"},
		DeleteCmd: []string{},
	},
	"JAVASCRIPT": {
		BuildCmd:  []string{},
		RunCmd:    []string{"node", "--stack-size=65536", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".js"},
		DeleteCmd: []string{},
	},
	"GO": {
		BuildCmd:  []string{"go", "build", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName, "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".go"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
	},
	"KOTLIN": {
		BuildCmd:  []string{"kotlinc", "-J-Xms1024m", "-J-Xmx1920m", "-J-Xss512m", "-include-runtime", "-d", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".jar", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".kt"},
		RunCmd:    []string{"java", "-Xms1024m", "-Xmx1920m", "-Xss512m", "-Dfile.encoding=UTF-8", "-XX:+UseSerialGC", "-jar", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".jar"},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".jar"},
	},
	"SWIFT": {
		BuildCmd:  []string{"swiftc", "-O", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName, "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName + ".swift"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FileName},
	},
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

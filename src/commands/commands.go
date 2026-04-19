package commands

const FILE_NAME = "Main"

type Command struct {
	BuildCmd  []string
	RunCmd    []string
	DeleteCmd []string
}

var Commands = map[string]Command{
	"C": {
		BuildCmd:  []string{"gcc", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".c", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME, "-O2", "-Wall", "-lm", "-static", "-std=gnu99"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
	},
	"CPP": {
		BuildCmd:  []string{"g++", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".cpp", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME, "-O2", "-Wall", "-lm", "-static", "-std=gnu++17"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
	},
	"JAVA": {
		BuildCmd:  []string{"javac", "-J-Xms1024m", "-J-Xmx1920m", "-J-Xss512m", "-encoding", "UTF-8", "-d", "bin", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".java"},
		RunCmd:    []string{"java", "-Xms1024m", "-Xmx1920m", "-Xss512m", "-Dfile.encoding=UTF-8", "-XX:+UseSerialGC", "-cp", "bin", FILE_NAME},
		DeleteCmd: []string{"rm", "-r", "{JUDGE_TYPE}/{SUBMIT_ID}/bin"},
	},
	"PYTHON": {
		BuildCmd:  []string{},
		RunCmd:    []string{"python3", "-W", "ignore", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".py"},
		DeleteCmd: []string{},
	},
	"JAVASCRIPT": {
		BuildCmd:  []string{},
		RunCmd:    []string{"node", "--stack-size=65536", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".js"},
		DeleteCmd: []string{},
	},
	"GO": {
		BuildCmd:  []string{"go", "build", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME, "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".go"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
	},
	"KOTLIN": {
		BuildCmd:  []string{"kotlinc", "-J-Xms1024m", "-J-Xmx1920m", "-J-Xss512m", "-include-runtime", "-d", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".jar", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".kt"},
		RunCmd:    []string{"java", "-Xms1024m", "-Xmx1920m", "-Xss512m", "-Dfile.encoding=UTF-8", "-XX:+UseSerialGC", "-jar", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".jar"},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".jar"},
	},
	"SWIFT": {
		BuildCmd:  []string{"swiftc", "-O", "-o", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME, "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME + ".swift"},
		RunCmd:    []string{"{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
		DeleteCmd: []string{"rm", "{JUDGE_TYPE}/{SUBMIT_ID}/" + FILE_NAME},
	},
}

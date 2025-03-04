package models

type Resigter struct {
	AppName   string   `json:"appName"`
	Functions []string `json:"functions"`
}

type Call struct {
	Caller       string                 `json:"caller"`
	TargetApp    string                 `json:"targetApp"`
	FunctionName string                 `json:"functionName"`
	Payload      map[string]interface{} `json:"payload"`
}

package main

type Answer struct {
	Error         bool   `json:"error"`
	ErrorReason   string `json:"error_reason"`
	Command       string `json:"command"`
	CommandResult string `json:"command_result"`
}

type Request struct {
	Commands []string `json:"commands"`
}

type Response struct {
	Answers []Answer `json:"answers"`
}

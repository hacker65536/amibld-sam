package main

type BuildEvent struct {
	BuildStatus           string                `json:"build-status"`
	AdditionalInformation AdditionalInformation `json:"additional-information"`
	BuildID               string                `json:"build-id"`
	ProjectName           string                `json:"project-name"`
}

type AdditionalInformation struct {
	// Initiator   string
	Environment    Environment `json:"environment"`
	BuildStartTime string      `json:"build-start-time"`
	Phases         []Phase     `json:"phases"`
	Logs           Logs        `json:"logs"`
}

type Phase struct {
	StartTime         string `json:"start-time"`
	EndTime           string `json:"end-time"`
	DurationInSeconds int    `json:"duration-in-seconds"`
	PhaseType         string `json:phase-type"`
	PhaseStatus       string `json:"phase-status"`
}

type Environment struct {
	EnvironmentVariables []EnvironmentVariables `json:"environment-variables"`
	// Image                string
}

type EnvironmentVariables struct {
	Name  string
	Type  string
	Value string
}

type Logs struct {
	GroupName  string `json:"group-name"`
	DeepLink   string `json:"deep-link"`
	StreamName string `json:"stream-name"`
}

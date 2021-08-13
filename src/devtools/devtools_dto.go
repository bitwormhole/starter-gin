package devtools

type RequestRecordDTO struct {
	TimeBegin int64
	TimeEnd   int64
	Method    string
	Path      string
}

type RequestSumDTO struct {
	TimeBegin int64
	TimeEnd   int64
	Count     int64
	Method    string
	Path      string
}

type ModuleInfoDTO struct {
	Index    string
	Name     string
	Revision string
	Version  string
}

type ComponentInfoDTO struct {
	Type  string
	ID    string
	Class []string
	Scope string
}

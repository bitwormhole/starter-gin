package devtools

////////////////////////////////////////////////////////////////////////////////

// RequestEndpoint
type RequestEndpoint struct {
	Method string
	Path   string

	Count        int64
	FirstRequest *RequestRecordDTO
	LastRequest  *RequestRecordDTO
}

// Add 添加记录
func (inst *RequestEndpoint) Add(rec *RequestRecordDTO) int64 {
	if inst.FirstRequest == nil {
		inst.FirstRequest = rec
	}
	inst.LastRequest = rec
	inst.Count++
	return inst.Count
}

// GetResult 取结果
func (inst *RequestEndpoint) GetResult() *RequestSumDTO {
	sum := &RequestSumDTO{}
	sum.Count = inst.Count
	sum.Method = inst.Method
	sum.Path = inst.Path
	sum.TimeBegin = inst.getTime(inst.FirstRequest, false)
	sum.TimeEnd = inst.getTime(inst.LastRequest, true)
	return sum
}

func (inst *RequestEndpoint) getTime(rec *RequestRecordDTO, useEnding bool) int64 {
	if rec == nil {
		return 0
	}
	if !useEnding {
		return rec.TimeBegin
	}
	return rec.TimeEnd
}

////////////////////////////////////////////////////////////////////////////////

// RequestAccumulator 请求累加器
type RequestAccumulator struct {
	endpoints map[string]*RequestEndpoint
}

func (inst *RequestAccumulator) keyFor(rec *RequestRecordDTO) string {
	method := rec.Method
	path := rec.Path
	return path + "#" + method
}

func (inst *RequestAccumulator) reset() {
	all := inst.getEndpoints()
	for _, item := range all {
		item.Count = 0
		item.FirstRequest = nil
		item.LastRequest = nil
	}
}

func (inst *RequestAccumulator) getEndpoints() map[string]*RequestEndpoint {
	table := inst.endpoints
	if table == nil {
		table = make(map[string]*RequestEndpoint)
		inst.endpoints = table
	}
	return table
}

func (inst *RequestAccumulator) getEndpoint(rec *RequestRecordDTO, create bool) *RequestEndpoint {
	key := inst.keyFor(rec)
	table := inst.getEndpoints()
	ep := table[key]
	if ep == nil && create {
		ep = &RequestEndpoint{}
		ep.Method = rec.Method
		ep.Path = rec.Path
		table[key] = ep
	}
	return ep
}

func (inst *RequestAccumulator) Add(rec *RequestRecordDTO) *RequestEndpoint {
	if rec == nil {
		return nil
	}
	ep := inst.getEndpoint(rec, true)
	ep.Add(rec)
	return ep
}

////////////////////////////////////////////////////////////////////////////////

// Code generated by "stringer -type=RecordType"; DO NOT EDIT.

package perffile

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RecordTypeMmap-1]
	_ = x[RecordTypeLost-2]
	_ = x[RecordTypeComm-3]
	_ = x[RecordTypeExit-4]
	_ = x[RecordTypeThrottle-5]
	_ = x[RecordTypeUnthrottle-6]
	_ = x[RecordTypeFork-7]
	_ = x[RecordTypeRead-8]
	_ = x[RecordTypeSample-9]
	_ = x[recordTypeMmap2-10]
	_ = x[RecordTypeAux-11]
	_ = x[RecordTypeItraceStart-12]
	_ = x[RecordTypeLostSamples-13]
	_ = x[RecordTypeSwitch-14]
	_ = x[RecordTypeSwitchCPUWide-15]
	_ = x[RecordTypeNamespaces-16]
	_ = x[RecordTypeKsymbol-17]
	_ = x[RecordTypeBPFEvent-18]
	_ = x[RecordTypeCGroup-19]
	_ = x[RecordTypeTextPoke-20]
	_ = x[recordTypeUserStart-64]
	_ = x[recordTypeAttr-64]
	_ = x[recordTypeEventType-65]
	_ = x[recordTypeTracingData-66]
	_ = x[recordTypeBuildID-67]
	_ = x[recordTypeFinishedRound-68]
	_ = x[recordTypeIDIndex-69]
	_ = x[RecordTypeAuxtraceInfo-70]
	_ = x[RecordTypeAuxtrace-71]
	_ = x[RecordTypeAuxtraceError-72]
	_ = x[recordTypeThreadMap-73]
	_ = x[recordTypeCPUMap-74]
	_ = x[recordTypeStatConfig-75]
	_ = x[recordTypeStat-76]
	_ = x[recordTypeStatRound-77]
	_ = x[recordTypeEventUpdate-78]
	_ = x[recordTypeTimeConv-79]
	_ = x[recordTypeHeaderFeature-80]
}

const (
	_RecordType_name_0 = "RecordTypeMmapRecordTypeLostRecordTypeCommRecordTypeExitRecordTypeThrottleRecordTypeUnthrottleRecordTypeForkRecordTypeReadRecordTypeSamplerecordTypeMmap2RecordTypeAuxRecordTypeItraceStartRecordTypeLostSamplesRecordTypeSwitchRecordTypeSwitchCPUWideRecordTypeNamespacesRecordTypeKsymbolRecordTypeBPFEventRecordTypeCGroupRecordTypeTextPoke"
	_RecordType_name_1 = "recordTypeUserStartrecordTypeEventTyperecordTypeTracingDatarecordTypeBuildIDrecordTypeFinishedRoundrecordTypeIDIndexRecordTypeAuxtraceInfoRecordTypeAuxtraceRecordTypeAuxtraceErrorrecordTypeThreadMaprecordTypeCPUMaprecordTypeStatConfigrecordTypeStatrecordTypeStatRoundrecordTypeEventUpdaterecordTypeTimeConvrecordTypeHeaderFeature"
)

var (
	_RecordType_index_0 = [...]uint16{0, 14, 28, 42, 56, 74, 94, 108, 122, 138, 153, 166, 187, 208, 224, 247, 267, 284, 302, 318, 336}
	_RecordType_index_1 = [...]uint16{0, 19, 38, 59, 76, 99, 116, 138, 156, 179, 198, 214, 234, 248, 267, 288, 306, 329}
)

func (i RecordType) String() string {
	switch {
	case 1 <= i && i <= 20:
		i -= 1
		return _RecordType_name_0[_RecordType_index_0[i]:_RecordType_index_0[i+1]]
	case 64 <= i && i <= 80:
		i -= 64
		return _RecordType_name_1[_RecordType_index_1[i]:_RecordType_index_1[i+1]]
	default:
		return "RecordType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

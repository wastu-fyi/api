package enums

type StudyInfo struct {
	Level string
	Name  string
}

var studies = map[int]StudyInfo{
	113: {Level: "S1", Name: "Teknik Tekstil"},
	115: {Level: "S1", Name: "Teknik Industri"},
	123: {Level: "D3", Name: "Teknik Mesin"},
	125: {Level: "S1", Name: "Teknik Mesin"},
	133: {Level: "D3", Name: "Manajemen Industri"},
	135: {Level: "S1", Name: "Teknik Informatika"},
}

func StudyInfoByCode(code int) (StudyInfo, bool) {
	info, ok := studies[code]
	return info, ok
}

func StudyName(code int) string {
	if info, ok := studies[code]; ok {
		return info.Name
	}
	return ""
}

func StudyLevel(code int) string {
	if info, ok := studies[code]; ok {
		return info.Level
	}
	return ""
}

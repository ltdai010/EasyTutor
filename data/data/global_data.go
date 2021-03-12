package data

type Subject string

const (
	MATH = "math"
	PHYSIC = "physic"
	CHEMISTRY = "chemistry"
	LITERATURE = "literature"
	HISTORY = "history"
	GEOGRAPHY = "geography"
	BIOLOGY = "biology"
	OTHER = "other"
	ENGLISH = "english"
	CHINESE = "chinese"
	GERMAN = "german"
	FRENCH = "french"
	JAPANESE = "japanese"
	APPRENTICE = "apprentice"
)

func NewSubject(sSub string) Subject {
	switch sSub {
	case MATH:
		return MATH
	case PHYSIC:
		return PHYSIC
	case CHEMISTRY:
		return CHEMISTRY
	case LITERATURE:
		return LITERATURE
	case HISTORY:
		return HISTORY
	case GEOGRAPHY:
		return GEOGRAPHY
	case BIOLOGY:
		return BIOLOGY
	case OTHER:
		return OTHER
	case ENGLISH:
		return ENGLISH
	case CHINESE:
		return CHINESE
	case GERMAN:
		return GERMAN
	case FRENCH:
		return FRENCH
	case JAPANESE:
		return JAPANESE
	case APPRENTICE:
		return APPRENTICE
	default:
		return ""
	}
}


type Method string

const (
	Online = "online"
	Offline = "offline"
)

func NewMethod(sMethod string) Method {
	switch sMethod {
	case Offline:
		return Offline
	case Online:
		return Online
	default:
		return ""
	}
}

type Gender string

const (
	Male = "male"
	Female = "female"
)

func NewGender(sGender string) Gender {
	switch sGender {
	case Male:
		return Male
	case Female:
		return Female
	default:
		return ""
	}
}

type Schedule struct {
	MonMorning   bool `json:"monday_morning"`
	MonAfternoon bool `json:"monday_afternoon"`
	MonNight     bool `json:"monday_night"`
	TueMorning   bool `json:"tue_morning"`
	TueAfternoon bool `json:"tue_afternoon"`
	TueNight     bool `json:"tue_night"`
	WedMorning   bool `json:"wed_morning"`
	WedAfternoon bool `json:"wed_afternoon"`
	WedNight     bool `json:"wed_night"`
	ThuMorning   bool `json:"thu_morning"`
	ThuAfternoon bool `json:"thu_afternoon"`
	ThuNight     bool `json:"thu_night"`
	FriMorning   bool `json:"fri_morning"`
	FriAfternoon bool `json:"fri_afternoon"`
	FriNight     bool `json:"fri_night"`
	SatMorning   bool `json:"sat_morning"`
	SatAfternoon bool `json:"sat_afternoon"`
	SatNight     bool `json:"sat_night"`
	SunMorning   bool `json:"sun_morning"`
	SunAfternoon bool `json:"sun_afternoon"`
	SunNight     bool `json:"sun_night"`
}

type Graduation string

const (
	UnderGraduated = "under_graduated"
	TeacherUnderGraduated = "teacher_under_graduated"
	UpperGraduated = "upper_graduated"
	TeacherUpperGraduated = "teacher_upper_graduated"
)

func NewGraduation(sGraduation string) Graduation {
	switch sGraduation {
	case UnderGraduated:
		return UnderGraduated
	case TeacherUnderGraduated:
		return TeacherUnderGraduated
	case UpperGraduated:
		return UpperGraduated
	case TeacherUpperGraduated:
		return TeacherUpperGraduated
	default:
		return ""
	}
}

type Status string

const (
	Open = "open"
	Close = "close"
)
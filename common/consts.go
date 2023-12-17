package common

const (
	GraphStatusInit   = int64(0)
	GraphStatusOK     = int64(1)
	GraphStatusUpdate = int64(2)
)

const (
	TypeString     = 0
	TypeDouble     = 1
	TypeInt        = 2
	TypeStringList = 3
	TypeDoubleList = 4
	TypeLong       = 5
)

func Type2String(t int64) string {
	switch t {
	case TypeInt:
		return "int"
	case TypeString:
		return "string"
	case TypeDouble:
		return "double"
	default:
		return "string"
	}
}

const (
	BaseTag    = "base"
	KeyDeg     = "deg"
	KeyID      = "id"
	KeySrc     = "source"
	KeyTgt     = "target"
	DefaultDeg = "0"
)

const (
	DirectionNormal  = ""
	DirectionReverse = "REVERSELY"
	DirectionBoth    = "BIDIRECT"

	MaxGoDistance      = 5
	MaxMatchCandidates = 5
)

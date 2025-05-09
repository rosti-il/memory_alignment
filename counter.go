package counter

type Counter interface {
	Inc(int)
	Get() int32
}

package driver

type PathInfo struct {
	Key string      `json:"key"`
	Sub []*PathInfo `json:"children"`
}

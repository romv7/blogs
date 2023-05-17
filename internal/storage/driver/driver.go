package driver

// PathInfo is heavily used by the driver's method Describe (See StorageDriver interface)
// Describe must be a recursive function (possibly an iterative as well) that
// returns an array of this type []*PathInfo. Although poorly designed,
// I find it useful in the long run.
type PathInfo struct {
	Key string      `json:"key"`
	Sub []*PathInfo `json:"children"`
}

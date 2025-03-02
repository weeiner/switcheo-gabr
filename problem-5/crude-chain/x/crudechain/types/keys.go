package types

const (
	// ModuleName defines the module name
	ModuleName = "crudechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_crudechain"

	// ResourceKey is used to uniquely identify resources within the system.
	ResourceKey = "Resource/value/"

	// ResourceCountKey is used to track the ID of the latest resource.
	ResourceCountKey = "Resource/count/"
)

var (
	ParamsKey = []byte("p_crudechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

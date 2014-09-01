package storage

type VersionStorageFactory struct {
	StorageType    string
	DefaultVersion string
}

func (f VersionStorageFactory) GetVersionStorage() VersionStorage {
	switch f.StorageType {
	case "local":
		return NewVersionStorageLocal(f.DefaultVersion)
	}
	return &VersionStorageFile{}
}

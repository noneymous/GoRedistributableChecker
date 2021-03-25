package redistributable

import (
	"golang.org/x/sys/windows/registry"
	"strings"
)

const (
	VC2005x86 = iota
	VC2005x64
	VC2008x86
	VC2008x64
	VC2010x86
	VC2010x64
	VC2012x86
	VC2012x64
	VC2013x86
	VC2013x64
	VC2015x86
	VC2015x64
	VC2017x86
	VC2017x64
	VC2015to2019x86
	VC2015to2019x64
)

func IsInstalled(redistributableVersion int) bool {

	// Execute check
	switch redistributableVersion {
	case VC2015to2019x86:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\DevDiv\\VC\\Servicing\\14.0\\RuntimeMinimum", "Version")
		if strings.HasPrefix(val, "14") {
			return true
		}
		break

	case VC2015to2019x64:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\DevDiv\\VC\\Servicing\\14.0\\RuntimeMinimum", "Version")
		if strings.HasPrefix(val, "14") {
			return true
		}
		break

	case VC2017x86:
		val := readString(registry.CLASSES_ROOT, "Installer\\Dependencies\\,,x86,14.0,bundle", "Version")
		if strings.HasPrefix(val, "14") {
			return true
		}
		val = readString(registry.CLASSES_ROOT, "Installer\\Dependencies\\VC,redist.x86,x86,14.16,bundle", "Version") //changed in 14.16.x
		if strings.HasPrefix(val, "14") {
			return true
		}
		break

	case VC2017x64:
		val := readString(registry.CLASSES_ROOT, "Installer\\Dependencies\\,,amd64,14.0,bundle", "Version")
		if strings.HasPrefix(val, "14") {
			return true
		}
		val = readString(registry.CLASSES_ROOT, "Installer\\Dependencies\\VC,redist.x64,amd64,14.16,bundle", "Version") //changed in 14.16.x
		if strings.HasPrefix(val, "14") {
			return true
		}
		break

	case VC2015x86:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Dependencies\\{e2803110-78b3-4664-a479-3611a381656a}", "Version")
		if strings.HasPrefix(val, "14") {
			return true
		}
		break

	case VC2015x64:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Dependencies\\{d992c12e-cab2-426f-bde3-fb8c53950b0d}", "Version")
		if strings.HasPrefix(val, "14") {
			return true
		}
		break

	case VC2013x86:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Dependencies\\{f65db027-aff3-4070-886a-0d87064aabb1}", "Version")
		if strings.HasPrefix(val, "12") {
			return true
		}
		break

	case VC2013x64:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Dependencies\\{050d4fc8-5d48-4b8f-8972-47c82c46020f}", "Version")
		if strings.HasPrefix(val, "12") {
			return true
		}
		break

	case VC2012x86:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Dependencies\\{33d1fd90-4274-48a1-9bc1-97e33d9c2d6f}", "Version")
		if strings.HasPrefix(val, "11") {
			return true
		}
		break

	case VC2012x64:
		val := readString(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Dependencies\\{ca67548a-5ebe-413a-b50c-4b9ceb6d66c6}", "Version")
		if strings.HasPrefix(val, "11") {
			return true
		}
		break

	case VC2010x86:
		val := readInt(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Products\\1D5E3C0FEDA1E123187686FED06E995A", "Version")
		if val > 1 {
			return true
		}
		break

	case VC2010x64:
		val := readInt(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Products\\1926E8D15D0BCE53481466615F760A7F", "Version")
		if val > 1 {
			return true
		}
		break

	case VC2008x86:
		val := readInt(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Products\\6E815EB96CCE9A53884E7857C57002F0", "Version")
		if val > 1 {
			return true
		}
		break

	case VC2008x64:
		val := readInt(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Products\\67D6ECF5CD5FBA732B8B22BAC8DE1B4D", "Version")
		if val > 1 {
			return true
		}
		break

	case VC2005x86:
		val := readInt(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Products\\c1c4f01781cc94c4c8fb1542c0981a2a", "Version")
		if val > 1 {
			return true
		}
		break

	case VC2005x64:
		val := readInt(registry.LOCAL_MACHINE, "SOFTWARE\\Classes\\Installer\\Products\\1af2a8da7e60d0b429d7e6453b3d0182", "Version")
		if val > 1 {
			return true
		}
		break

	default:
		return false
	}

	// Return false if no suitable registry keys were found
	return false
}

func readString(root registry.Key, path, key string) string {

	// Attach to registry
	k, errOpen := registry.OpenKey(root, path, registry.QUERY_VALUE)
	if errOpen != nil {
		return ""
	}

	// Make sure registry key gets closed on exit
	defer func() { _ = k.Close() }()

	// Read value
	val, _, errGet := k.GetStringValue(key)
	if errGet != nil {
		return ""
	}

	// Return read value
	return val
}

func readInt(root registry.Key, path, key string) uint64 {

	// Attach to registry
	k, errOpen := registry.OpenKey(root, path, registry.QUERY_VALUE)
	if errOpen != nil {
		return 0
	}

	// Make sure registry key gets closed on exit
	defer func() { _ = k.Close() }()

	// Read value
	val, _, errGet := k.GetIntegerValue(key)
	if errGet != nil {
		return 0
	}

	// Return read value
	return val
}

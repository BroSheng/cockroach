// Code generated by "stringer"; DO NOT EDIT.

package descpb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BaseFormatVersion-1]
	_ = x[FamilyFormatVersion-2]
	_ = x[InterleavedFormatVersion-3]
}

const _FormatVersion_name = "BaseFormatVersionFamilyFormatVersionInterleavedFormatVersion"

var _FormatVersion_index = [...]uint8{0, 17, 36, 60}

func (i FormatVersion) String() string {
	i -= 1
	if i >= FormatVersion(len(_FormatVersion_index)-1) {
		return "FormatVersion(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FormatVersion_name[_FormatVersion_index[i]:_FormatVersion_index[i+1]]
}

// Code generated by "stringer -type ScmId"; DO NOT EDIT.

package eddn

import "strconv"

const _ScmId_name = "SblackmarketScommoditySjournalSoutfittingSshipyard"

var _ScmId_index = [...]uint8{0, 12, 22, 30, 41, 50}

func (i ScmId) String() string {
	if i < 0 || i >= ScmId(len(_ScmId_index)-1) {
		return "ScmId(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ScmId_name[_ScmId_index[i]:_ScmId_index[i+1]]
}

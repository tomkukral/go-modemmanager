// Code generated by "stringer -type=MMModemAccessTechnology"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModemAccessTechnologyUnknown-0]
	_ = x[MmModemAccessTechnologyPots-1]
	_ = x[MmModemAccessTechnologyGsm-2]
	_ = x[MmModemAccessTechnologyGsmCompact-4]
	_ = x[MmModemAccessTechnologyGprs-8]
	_ = x[MmModemAccessTechnologyEdge-16]
	_ = x[MmModemAccessTechnologyUmts-32]
	_ = x[MmModemAccessTechnologyHsdpa-64]
	_ = x[MmModemAccessTechnologyHsupa-128]
	_ = x[MmModemAccessTechnologyHspa-256]
	_ = x[MmModemAccessTechnologyHspaPlus-512]
	_ = x[MmModemAccessTechnology1xrtt-1024]
	_ = x[MmModemAccessTechnologyEvdo0-2048]
	_ = x[MmModemAccessTechnologyEvdoa-4096]
	_ = x[MmModemAccessTechnologyEvdob-8192]
	_ = x[MmModemAccessTechnologyLte-16384]
	_ = x[MmModemAccessTechnologyAny-4294967295]
}

const _MMModemAccessTechnology_name = "MmModemAccessTechnologyUnknownMmModemAccessTechnologyPotsMmModemAccessTechnologyGsmMmModemAccessTechnologyGsmCompactMmModemAccessTechnologyGprsMmModemAccessTechnologyEdgeMmModemAccessTechnologyUmtsMmModemAccessTechnologyHsdpaMmModemAccessTechnologyHsupaMmModemAccessTechnologyHspaMmModemAccessTechnologyHspaPlusMmModemAccessTechnology1xrttMmModemAccessTechnologyEvdo0MmModemAccessTechnologyEvdoaMmModemAccessTechnologyEvdobMmModemAccessTechnologyLteMmModemAccessTechnologyAny"

var _MMModemAccessTechnology_map = map[MMModemAccessTechnology]string{
	0:          _MMModemAccessTechnology_name[0:30],
	1:          _MMModemAccessTechnology_name[30:57],
	2:          _MMModemAccessTechnology_name[57:83],
	4:          _MMModemAccessTechnology_name[83:116],
	8:          _MMModemAccessTechnology_name[116:143],
	16:         _MMModemAccessTechnology_name[143:170],
	32:         _MMModemAccessTechnology_name[170:197],
	64:         _MMModemAccessTechnology_name[197:225],
	128:        _MMModemAccessTechnology_name[225:253],
	256:        _MMModemAccessTechnology_name[253:280],
	512:        _MMModemAccessTechnology_name[280:311],
	1024:       _MMModemAccessTechnology_name[311:339],
	2048:       _MMModemAccessTechnology_name[339:367],
	4096:       _MMModemAccessTechnology_name[367:395],
	8192:       _MMModemAccessTechnology_name[395:423],
	16384:      _MMModemAccessTechnology_name[423:449],
	4294967295: _MMModemAccessTechnology_name[449:475],
}

func (i MMModemAccessTechnology) String() string {
	if str, ok := _MMModemAccessTechnology_map[i]; ok {
		return str
	}
	return "MMModemAccessTechnology(" + strconv.FormatInt(int64(i), 10) + ")"
}
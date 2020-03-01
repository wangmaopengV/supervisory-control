package agreement

import (
	"bytes"
	"encoding/binary"
	"supervisory-control/config"
	pbSC "supervisory-control/proto"
	"supervisory-control/sql"
)

func BytesToInt(b []byte) int64 {

	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	_ = binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return x
}

func AnalysisAgreement(req []byte) []byte {

	res := []byte{req[0], req[1], req[2], req[3], req[4], req[5], 0x00, req[10], req[11]}
	uid := BytesToInt(req[2:6])
	time := BytesToInt(req[6:10])

	data := &sql.DeviceItem{
		UID:  uid,
		Type: int(pbSC.SCDeviceType_SC_Type_LOCK),
	}
	if err := config.GlobalConfig.DBClient.FindDeviceTable(data); err != nil {

		if err.Error() != "empty" {
			data.Status = 0x01
			data.CreateTime = time
			if err := config.GlobalConfig.DBClient.InsertDeviceTable(data); err != nil {
				return res
			}
		} else {
			return res
		}
	}

	_ = config.GlobalConfig.DBClient.InsertDeviceMonitorTable(&sql.DeviceMonitorItem{
		DeviceUID: data.UID,
		LoginTime: time,
	})

	switch data.Status {
	case int(pbSC.SCDeviceStatus_SC_STATUS_ENABLE):
		res[6] = 0x01
	case int(pbSC.SCDeviceStatus_SC_STATUS_DISABLE):
		res[6] = 0x02
	case int(pbSC.SCDeviceStatus_SC_STATUS_RESTART):
		res[6] = 0x03
	default:
	}

	return res
}

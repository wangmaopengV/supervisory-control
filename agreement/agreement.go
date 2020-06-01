package agreement

import (
	"encoding/binary"
	"supervisory-control/config"
	pbSC "supervisory-control/proto"
	"supervisory-control/sql"
	log "github.com/sirupsen/logrus"
	"sync"
)

type TaskAgreement struct {
	req  []byte
	res  []byte
	sync sync.WaitGroup
}

var taskChan chan *TaskAgreement

func RunTask(meta config.SCAgreementMeta) {

	run := func() {
		for true {
			select {

			case task := <-taskChan:
				task.res = analysisAgreement(task.req)
				task.sync.Done()

			default:
			}
		}
	}

	taskChan = make(chan *TaskAgreement, meta.TaskSize)
	for i := 0; i < int(meta.TaskThread); i++ {
		go run()
	}
}

func DealTask(req []byte) []byte {

	task := &TaskAgreement{
		req: req,
	}
	task.sync.Add(1)
	taskChan <- task
	task.sync.Wait()

	return task.res
}

func analysisAgreement(req []byte) []byte {

	log.Debug(req)

	res := []byte{req[0], req[1], req[2], req[3], req[4], req[5], 0x00, req[10], req[11]}
	uid := int64(binary.BigEndian.Uint32(req[2:6]))
	time := int64(binary.BigEndian.Uint32(req[6:10]))

	log.Debug(uid, time)

	data := &sql.DeviceItem{
		UID:  uid,
		Type: int(pbSC.SCDeviceType_SC_Type_LOCK),
	}
	if err := config.GlobalConfig.DBClient.FindDeviceTable(data); err != nil {

		if err.Error() == "empty" {
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

	log.Debug(res)
	return res
}

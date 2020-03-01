package sql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type DeviceMonitorItem struct {
	DeviceUID int64
	LoginTime int64
}

func (db *DBClient) CreateDeviceMonitorTable() {

	sql := "CREATE TABLE IF NOT EXISTS device_monitor(" +
		"uid BIGINT NOT NULL AUTO_INCREMENT, " +
		"device_uid BIGINT NOT NULL, " +
		"login_time BIGINT NOT NULL, " +
		"PRIMARY KEY (uid));"

	stmt, err := db.dbClient.Prepare(sql)
	if err != nil {
		log.Fatal("CreateDeviceMonitorTable ", err, sql)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Warn(err)
		}
	}()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("CreateDeviceMonitorTable ", err, sql)
	}
}

func (db *DBClient) InsertDeviceMonitorTable(data *DeviceMonitorItem) error {

	stmt, err := db.dbClient.Prepare("insert device_monitor set device_uid=?,login_time=?")
	if err != nil {
		log.Error("InsertDeviceMonitorTable ", err)
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Warn(err)
		}
	}()

	_, err = stmt.Exec(data.DeviceUID, data.LoginTime)
	if err != nil {
		log.Error("InsertDeviceMonitorTable ", err)
		return err
	}
	return nil
}

func (db *DBClient) ScreenDeviceMonitorTable(data *DeviceMonitorItem) ([]DeviceMonitorItem, error) {

	sql := fmt.Sprintf("SELECT login_time FROM device_monitor where device_uid = %d", data.DeviceUID)
	rows, err := db.dbClient.Query(sql)
	if err != nil {
		log.Error("ScreenDeviceMonitorTable ", err, sql)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Warn(err)
		}
	}()

	var dataS []DeviceMonitorItem
	var temp DeviceMonitorItem
	for rows.Next() {
		if err := rows.Scan(&temp.LoginTime); err != nil {
			return nil, err
		}
		dataS = append(dataS, temp)
	}

	return dataS, nil
}

package sql

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type DeviceItem struct {
	UID        int64
	Status     int
	Type       int
	Meta       string
	CreateTime int64
}

func (db *DBClient) CreateDeviceTable() {

	sql := "CREATE TABLE IF NOT EXISTS device(" +
		"uid BIGINT NOT NULL, " +
		"status INT NOT NULL, " +
		"type INT NOT NULL, " +
		"meta VARCHAR(4096), " +
		"create_time BIGINT, " +
		"PRIMARY KEY (uid));"

	stmt, err := db.dbClient.Prepare(sql)
	if err != nil {
		log.Fatal("CreateDeviceTable ", err, sql)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Warn(err)
		}
	}()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("CreateDeviceTable ", err, sql)
	}
}

func (db *DBClient) InsertDeviceTable(data *DeviceItem) error {

	stmt, err := db.dbClient.Prepare("insert device set uid=?,status=?,type=?,meta=?,create_time=?")
	if err != nil {
		log.Error("InsertDeviceTable ", err)
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Warn(err)
		}
	}()

	_, err = stmt.Exec(data.UID, data.Status, data.Type, data.Meta, data.CreateTime)
	if err != nil {
		log.Error("InsertDeviceTable ", err)
		return err
	}
	return nil
}

func (db *DBClient) UpdateDeviceTable(data *DeviceItem) error {

	stmt, err := db.dbClient.Prepare("update device set status=? where uid=?")
	if err != nil {
		log.Error("UpdateDeviceTable ", err)
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Warn(err)
		}
	}()

	_, err = stmt.Exec(data.Status, data.UID)
	if err != nil {
		log.Error("UpdateDeviceTable ", err)
		return err
	}
	return nil
}

func (db *DBClient) FindDeviceTable(data *DeviceItem) error {

	sql := fmt.Sprintf("SELECT status FROM device where uid = %d", data.UID)
	rows, err := db.dbClient.Query(sql)
	if err != nil {
		log.Error("FindDeviceTable ", err, sql)
		return err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Warn(err)
		}
	}()

	if rows.Next() {

		if err := rows.Scan(&data.Status); err != nil {
			return err
		}
		return nil
	}
	return errors.New("empty")
}

func (db *DBClient) CensusDeviceTable(start int64, end int64) ([]DeviceItem, error) {

	sql := fmt.Sprintf("SELECT uid,status,type,meta,create_time FROM device where create_time > %d and create_time < %d", start, end)
	rows, err := db.dbClient.Query(sql)
	if err != nil {
		log.Error("CensusDeviceTable ", err, sql)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Warn(err)
		}
	}()

	var dataS []DeviceItem
	var temp DeviceItem
	for rows.Next() {
		if err := rows.Scan(&temp.UID, &temp.Status, &temp.Type, &temp.Meta, &temp.CreateTime); err != nil {
			return nil, err
		}
		dataS = append(dataS, temp)
	}

	return dataS, nil
}

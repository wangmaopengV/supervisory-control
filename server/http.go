package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"supervisory-control/config"
	pbSC "supervisory-control/proto"
	"supervisory-control/sql"
)

type SupervisoryControlService struct {
}

func (sc *SupervisoryControlService) QueryDevice(ctx context.Context, request *pbSC.QueryDeviceRequest) (*pbSC.QueryDeviceResponse, error) {

	device := &sql.DeviceItem{UID: request.UID}
	if err := config.GlobalConfig.DBClient.FindDeviceTable(device); err != nil {
		return &pbSC.QueryDeviceResponse{
			Result: &pbSC.SCResult{
				Msg:  err.Error(),
				Code: pbSC.SCErrorCode_SC_QUERY_DB_FAIL,
			},
		}, nil
	}

	res := &pbSC.QueryDeviceResponse{
		Result: &pbSC.SCResult{
			Msg:  "SUCCESS",
			Code: pbSC.SCErrorCode_SC_SUCCESS,
		},
		Device: &pbSC.SCDeviceMeta{
			UID:        device.UID,
			Status:     pbSC.SCDeviceStatus(device.Status),
			Type:       pbSC.SCDeviceType(device.Type),
			Meta:       device.Meta,
			CreateTime: &timestamp.Timestamp{Seconds: device.CreateTime},
		},
	}

	return res, nil
}

func (sc *SupervisoryControlService) ControlDevice(ctx context.Context, request *pbSC.ControlDeviceRequest) (*pbSC.ControlDeviceResponse, error) {

	device := &sql.DeviceItem{UID: request.UID}
	if err := config.GlobalConfig.DBClient.FindDeviceTable(device); err != nil {
		return &pbSC.ControlDeviceResponse{
			Result: &pbSC.SCResult{
				Msg:  err.Error(),
				Code: pbSC.SCErrorCode_SC_QUERY_DB_FAIL,
			},
		}, nil
	}

	device.Status = int(request.Status)
	if err := config.GlobalConfig.DBClient.UpdateDeviceTable(device); err != nil {
		return &pbSC.ControlDeviceResponse{
			Result: &pbSC.SCResult{
				Msg:  err.Error(),
				Code: pbSC.SCErrorCode_SC_INSERT_DB_FAIL,
			},
		}, nil
	}

	return &pbSC.ControlDeviceResponse{
		Result: &pbSC.SCResult{
			Msg:  "SUCCESS",
			Code: pbSC.SCErrorCode_SC_SUCCESS,
		},
	}, nil
}

func (sc *SupervisoryControlService) CensusDevice(ctx context.Context, request *pbSC.CensusDeviceRequest) (*pbSC.CensusDeviceResponse, error) {

	devices, err := config.GlobalConfig.DBClient.CensusDeviceTable(int64(request.StartTime), int64(request.EndTime))
	if err != nil {
		return &pbSC.CensusDeviceResponse{
			Result: &pbSC.SCResult{
				Msg:  err.Error(),
				Code: pbSC.SCErrorCode_SC_QUERY_DB_FAIL,
			},
		}, nil
	}

	data := make([]*pbSC.SCDeviceMeta, len(devices))
	for key, value := range devices {

		data[key] = &pbSC.SCDeviceMeta{
			UID:        value.UID,
			Status:     pbSC.SCDeviceStatus(value.Status),
			Type:       pbSC.SCDeviceType(value.Type),
			Meta:       value.Meta,
			CreateTime: &timestamp.Timestamp{Seconds: value.CreateTime},
		}
	}

	return &pbSC.CensusDeviceResponse{
		Result: &pbSC.SCResult{
			Msg:  "SUCCESS",
			Code: pbSC.SCErrorCode_SC_SUCCESS,
		},
		TotalDevice: int32(len(devices)),
		Devices:     data,
	}, nil
}

func (sc *SupervisoryControlService) QueryDeviceMonitor(ctx context.Context, request *pbSC.QueryDeviceMonitorRequest) (*pbSC.QueryDeviceMonitorResponse, error) {

	device := &sql.DeviceItem{UID: int64(request.UID)}
	if err := config.GlobalConfig.DBClient.FindDeviceTable(device); err != nil {
		return &pbSC.QueryDeviceMonitorResponse{
			Result: &pbSC.SCResult{
				Msg:  err.Error(),
				Code: pbSC.SCErrorCode_SC_QUERY_DB_FAIL,
			},
		}, nil
	}

	times, err := config.GlobalConfig.DBClient.ScreenDeviceMonitorTable(&sql.DeviceMonitorItem{DeviceUID: int64(request.UID)})
	if err != nil {
		return &pbSC.QueryDeviceMonitorResponse{
			Result: &pbSC.SCResult{
				Msg:  err.Error(),
				Code: pbSC.SCErrorCode_SC_QUERY_DB_FAIL,
			},
		}, nil
	}

	data := make([]*timestamp.Timestamp, len(times))
	for key, value := range times {
		data[key] = &timestamp.Timestamp{Seconds: value.LoginTime}
	}

	return &pbSC.QueryDeviceMonitorResponse{
		Result: &pbSC.SCResult{
			Msg:  "SUCCESS",
			Code: pbSC.SCErrorCode_SC_SUCCESS,
		},
		Device: &pbSC.SCDeviceMeta{
			UID:        device.UID,
			Status:     pbSC.SCDeviceStatus(device.Status),
			Type:       pbSC.SCDeviceType(device.Type),
			Meta:       device.Meta,
			CreateTime: &timestamp.Timestamp{Seconds: device.CreateTime},
		},
		Login: data,
	}, nil
}

func newSupervisoryControlService() *SupervisoryControlService {

	return new(SupervisoryControlService)
}

func RunServer(gRpcPort int32, httpPort int32) {

	gRpcServer := grpc.NewServer()
	pbSC.RegisterSControllerServiceServer(gRpcServer, newSupervisoryControlService())
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", gRpcPort))

	go func() {

		log.Info("gRpcServer port ", gRpcPort)
		err := gRpcServer.Serve(lis)
		if err != nil {
			log.Fatal(" gRpcServer failed ", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	err := pbSC.RegisterSControllerServiceHandlerFromEndpoint(
		ctx, mux, fmt.Sprintf(":%d", gRpcPort),
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("httpServer port ", httpPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
	if err != nil {
		log.Fatal(" listenAndServe failed ", err)
	}
}

package INIT

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"relay-chain/module/db"
	"relay-chain/module/plugin/chainmaker"
	"relay-chain/module/plugin/fabric"
	"relay-chain/module/plugin/flato"
	"relay-chain/module/server"
	"relay-chain/module/server/config"
	"relay-chain/module/service"
	"relay-chain/module/types/pb"
	"relay-chain/module/verifier"
	"syscall"
	"time"
)

const (
	ChainMaker = "chainmaker"
	Fabric     = "fabric"
	Flato      = "flato"
)

var GConfig *config.Config

func Init(confPath string) {
	var err error
	//初始化配置
	GConfig, err = config.InitConfig(confPath)
	/*GConfig = {
		RPCAdd: "localhost:123X",
		ChainType: "chainmaker",
		SdkConfPath: "./configs/chainmaker/chainX/sdk_config_org1_client1.yml"
		ChainId: "chainX",
		CrossTransactionContractAddress: "",
		DbFile: "./store/chainX_db",
		Mode: "listen",
	}*/
	if err != nil {
		log.Fatalln(err)
	}

	server.DefaultServer.DB = db.GetDB(GConfig.DbFile)
	log.Printf("[success start db] at [%s]\n", GConfig.DbFile)

	server.DefaultServer.Mode = GConfig.Mode
	server.DefaultServer.Verifier = &verifier.Verifier{}

	switch GConfig.ChainType {
	case ChainMaker:
		cmPlugin := chainmaker.ChainMaker{
			ChainId: GConfig.ChainId,
		}

		err = cmPlugin.InitClient(GConfig.SdkConfPath)
		if err != nil {
			log.Fatalln(err)
		}

		server.DefaultServer.Plugin = &cmPlugin
	case Fabric:
		fabPlugin := fabric.Fabric{
			ChainId: GConfig.ChainId,
		}
		err = fabPlugin.InitClient(GConfig.SdkConfPath)
		if err != nil {
			log.Fatalln(err)
		}

		server.DefaultServer.Plugin = &fabPlugin
	case Flato:
		flatoPlugin := flato.Flato{
			ChainId: GConfig.ChainId,
		}

		flato.CrossTransactionContractAddress = GConfig.CrossTransactionContractAddress

		err = flatoPlugin.InitClient(GConfig.SdkConfPath)
		if err != nil {
			log.Fatalln(err)
		}

		server.DefaultServer.Plugin = &flatoPlugin
	default:
		log.Fatalf("unsupport chain type [%s]", GConfig.ChainType)
	}
}

func Start(UnLog bool) {
	if GConfig.Mode == "listen" {
		go server.DefaultServer.Plugin.SubscribeCrossEvent()
		log.Printf("[success start chain listen] at [%s (%s)] for cross tx\n", GConfig.ChainId, GConfig.ChainType)
	} else if GConfig.Mode == "consensus" {

	} else {
		log.Fatalf("do not support mode [%s]", GConfig.Mode)
	}

	go startGrpc()

	time.Sleep(time.Millisecond * 10)
	log.Printf("========== Everything is ok! Working Mode: [%s] ==========", GConfig.Mode)

	if UnLog {
		log.Println("[log is disable]")
		log.SetOutput(io.Discard)
	}

	errorC := make(chan error, 1)
	go handleExitSignal(errorC)
	// will block until Interrupt
	e := <-errorC
	if e != nil {
		return
	}
	if err := server.DefaultServer.DB.Close(); err != nil {
		log.Print("fail close db ")
		return
	}
	log.Println("========== All is stopped! Done ==========")
}

func handleExitSignal(exitC chan<- error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	defer signal.Stop(signalChan)

	for range signalChan {
		exitC <- nil
	}
}

func startGrpc() {
	lis, err := net.Listen("tcp", GConfig.RPCAdd)
	if err != nil {
		log.Fatalf("failed to start grpc listen: %v", err)
	}
	s := grpc.NewServer()
	//注册实现服务端具体服务的结构体Service（位于service包下）
	pb.RegisterServiceServer(s, &service.Service{})
	log.Printf("[grpc success start listening] at [%v]", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc listen: %v", err)
	}
}

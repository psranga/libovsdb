package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/ovn-org/libovsdb/cache"
	"github.com/ovn-org/libovsdb/client"
	"github.com/ovn-org/libovsdb/example/vswitchd"
	"github.com/ovn-org/libovsdb/model"
	"github.com/ovn-org/libovsdb/ovsdb"
)

// Silly game that detects creation of Bridge named "stop" and exits
// Just a demonstration of how an app can use libovsdb library to configure and manage OVS
const (
	bridgeTable = "Bridge"
	ovsTable    = "Open_vSwitch"
)

var rootUUID string
var connection = flag.String("ovsdb", "unix:/var/run/openvswitch/db.sock", "OVSDB connection string")
var nwdevice = flag.String("device", "eth1", "local interface to add to ovs (must exist)")
var mac = flag.String("mac", "01:02:03:04:05:06", "mac address to set in ovs")
var lsp = flag.String("lsp", "lsp00", "logical switch port to set in ovs")

// func play(ovs client.Client) {
// 	go processInput(ovs)
// 	for model := range update {
// 		bridge := model.(*vswitchd.Bridge)
// 		if bridge.Name == "stop" {
// 			fmt.Printf("Bridge stop detected: %+v\n", *bridge)
// 			ovs.Disconnect()
// 			quit <- true
// 		} else {
// 			fmt.Printf("Current list of bridges:\n")
// 			var bridges []vswitchd.Bridge
// 			if err := ovs.List(context.Background(), &bridges); err != nil {
// 				log.Fatal(err)
// 			}
// 			for _, b := range bridges {
// 				fmt.Printf("UUID: %s  Name: %s\n", b.UUID, b.Name)
// 			}
// 		}
// 	}
// }

func createBridge(ovs client.Client, bridgeName string) {
	bridge := vswitchd.Bridge{
		UUID: uuid.NewString(),
		Name: bridgeName,
	}
	insertOp, err := ovs.Create(&bridge)
	if err != nil {
		log.Fatal(err)
	}

	ovsRow := vswitchd.OpenvSwitch{
		UUID: rootUUID,
	}
	mutateOps, err := ovs.Where(&ovsRow).Mutate(&ovsRow, model.Mutation{
		Field:   &ovsRow.Bridges,
		Mutator: "insert",
		Value:   []string{bridge.UUID},
	})
	if err != nil {
		log.Fatal(err)
	}

	operations := append(insertOp, mutateOps...)
	reply, err := ovs.Transact(context.TODO(), operations...)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ovsdb.CheckOperationResults(reply, operations); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bridge Addition Successful : ", reply[0].UUID.GoUUID)
}

// sudo ovs-vsctl -- --id=@if0 create Interface name=eth1 -- --id=@port0 create Port name=eth1 interfaces=@if0 -- add Bridge obr0 ports @port0
func main() {
	flag.Parse()

	clientDBModel, err := vswitchd.FullDatabaseModel()
	if err != nil {
		log.Fatal("Unable to create DB model ", err)
	}
	log.Println("Got client db model via FullDatabaseModel.")

	ovs, err := client.NewOVSDBClient(clientDBModel, client.WithEndpoint(*connection))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created db client object for connection to: %s\n", *connection)

	err = ovs.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Success connecting to db. connection: %s\n", *connection)
	defer ovs.Disconnect()

	ovs.Cache().AddEventHandler(&cache.EventHandlerFuncs{
		AddFunc: func(table string, model model.Model) {
			log.Printf("AddFunc: Received update in table: %s", table)
			log.Printf("AddFunc: Received update: %v", model)
		},
	})
	_, err = ovs.MonitorAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success MonitorAll.")

	// Get root UUID
	for uuid := range ovs.Cache().Table("Open_vSwitch").Rows() {
		rootUUID = uuid
		log.Printf("Got root UUID: %s\n", rootUUID)
	}
	log.Printf("Final root UUID is: %s\n", rootUUID)

	if len(*nwdevice) <= 0 {
		log.Fatal("Need non-empty --device argument.")
	}
	log.Printf("Creating interface named: %s", *nwdevice)

	log.Println("Exiting.")
}

// func main() {
// 	flag.Parse()
// 	quit = make(chan bool)
// 	update = make(chan model.Model)

// 	clientDBModel, err := model.NewClientDBModel("Open_vSwitch",
// 		map[string]model.Model{bridgeTable: &vswitchd.Bridge{}, ovsTable: &vswitchd.OpenvSwitch{}})
// 	if err != nil {
// 		log.Fatal("Unable to create DB model ", err)
// 	}

// 	ovs, err := client.NewOVSDBClient(clientDBModel, client.WithEndpoint(*connection))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = ovs.Connect(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ovs.Disconnect()

// 	ovs.Cache().AddEventHandler(&cache.EventHandlerFuncs{
// 		AddFunc: func(table string, model model.Model) {
// 			if table == bridgeTable {
// 				update <- model
// 			}
// 		},
// 	})
// 	_, err = ovs.Monitor(
// 		context.TODO(),
// 		ovs.NewMonitor(
// 			client.WithTable(&vswitchd.OpenvSwitch{}),
// 			client.WithTable(&vswitchd.Bridge{}),
// 		),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Get root UUID
// 	for uuid := range ovs.Cache().Table("Open_vSwitch").Rows() {
// 		rootUUID = uuid
// 	}

// 	fmt.Println(`Silly game of stopping this app when a Bridge with name "stop" is monitored !`)
// 	go play(ovs)
// 	<-quit
// }

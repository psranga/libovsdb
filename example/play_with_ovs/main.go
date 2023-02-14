package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/ovn-org/libovsdb/cache"
	"github.com/ovn-org/libovsdb/client"
	"github.com/ovn-org/libovsdb/example/vswitchd"
	"github.com/ovn-org/libovsdb/model"
	"github.com/ovn-org/libovsdb/ovsdb"
)

// Silly game that detects creation of Bridge named "stop" and exits
// Just a demonstration of how an app can use libovsdb library to configure and manage OVS
// const (
// 	bridgeTable = "Bridge"
// 	ovsTable    = "Open_vSwitch"
// )

var rootUUID string
var connection = flag.String("ovsdb", "unix:/var/run/openvswitch/db.sock", "OVSDB connection string")
var nwdevice = flag.String("device", "eth1", "local interface to add to ovs (must exist)")
var bridgeName = flag.String("bridge", "br-int", "logical bridge to create the port (must exist)")

// var mac = flag.String("mac", "01:02:03:04:05:06", "mac address to set in ovs")
// var lsp = flag.String("lsp", "lsp00", "logical switch port to set in ovs")

var namedUUIDCounter uint32

func NamedUUID() string {
	return fmt.Sprintf("xu%010d", atomic.AddUint32(&namedUUIDCounter, 1))
}

// GenerateMac generates mac address.
func GenerateMac() string {
	prefix := "00:00:00"
	b := make([]byte, 3)
	rand.Read(b)
	mac := fmt.Sprintf("%s:%02X:%02X:%02X", prefix, b[0], b[1], b[2])
	return mac
}

func fmtErrorf(format string, a ...any) error {
	// log.Printf("Returning Error: ")
	log.Printf(format, a...)
	// log.Printf("\n")
	return fmt.Errorf(format, a...)
}

func GetBridge(ovs client.Client, bridgeName string, ignoreNotFound bool) (*vswitchd.Bridge, error) {
	ctx := context.Background()

	log.Printf("GetBridge: looking for bridge: %s, ignoreNotFound: %t", bridgeName, ignoreNotFound)
	br := &vswitchd.Bridge{}
	api := ovs.WhereAll(br, model.Condition{
		Field:    &br.Name,
		Function: ovsdb.ConditionEqual,
		Value:    bridgeName,
	})

	var results []*vswitchd.Bridge
	err := api.List(ctx, &results)
	if err != nil {
		return nil, fmtErrorf("Error when listing bridges: %v\n", err)
	}

	log.Printf("List returned: %d rows looking for %s.\n", len(results), bridgeName)

	if len(results) <= 0 {
		if ignoreNotFound {
			log.Printf("Zero-sized List for %s, returning nil without error\n", bridgeName)
			return nil, nil
		}
		return nil, fmtErrorf("zero-sized list: %d for bridge %s", len(results), bridgeName)
	}

	return results[0], nil
}

func BridgeExists(ovs client.Client, name string) (bool, error) {
	br, err := GetBridge(ovs, name, true)
	return br != nil, err
}

// sudo ovs-vsctl -- --id=@if0 create Interface name=eth1 -- --id=@port0 create Port name=eth1 interfaces=@if0 -- add Bridge obr0 ports @port0
func createPortOnBridge(ovs client.Client, bridgeName, interfaceName string) error {
	log.Printf("createInterface(): creating interface named: %s on bridge: %s", *nwdevice, bridgeName)

	br, err := GetBridge(ovs, bridgeName, false)
	if br == nil || err != nil {
		if err != nil {
			return fmtErrorf("Bridge %s does not exist. Err: %v\n", bridgeName, err)
		}
	}
	operations := []ovsdb.Operation{}

	// newInterface := vswitchd.Interface{
	// 	UUID: NamedUUID(),
	// 	Name: interfaceName,
	// }
	// interfaceOps, err := ovs.Create(&newInterface)
	// if err != nil {
	// 	return fmtErrorf("Error creating interface: %v", err)
	// }
	// fmt.Printf("Created insertOps: # %d\n", len(interfaceOps))
	// operations = append(operations, interfaceOps...)

	// mac := GenerateMac()
	// newPort := vswitchd.Port{
	// 	UUID: NamedUUID(),
	// 	// Name:       "p" + interfaceName,
	// 	MAC: &mac,
	// 	// Interfaces: []string{newInterface.UUID},
	// }
	// portOps, err := ovs.Create(&newPort)
	// if err != nil {
	// 	return fmtErrorf("Error creating port: %v", err)
	// }
	// fmt.Printf("Created PortOps: # %d\n", len(portOps))
	// operations = append(operations, portOps...)

	newInterface := vswitchd.Interface{
		UUID: NamedUUID(),
		Name: interfaceName,
	}
	mac := GenerateMac()
	newPort := vswitchd.Port{
		UUID:       NamedUUID(),
		Name:       interfaceName,
		MAC:        &mac,
		Interfaces: []string{newInterface.UUID},
	}
	// for i := 0; i < 4096; i++ {
	// 	newPort.CVLANs[i] = i
	// 	newPort.Trunks[i] = i
	// }
	createOps, err := ovs.Create(&newInterface, &newPort)
	if err != nil {
		return fmtErrorf("Error creating interface and port: %v", err)
	}
	fmt.Printf("Created createOps: # %d\n", len(createOps))
	operations = append(operations, createOps...)

	br2 := &vswitchd.Bridge{
		UUID: br.UUID,
	}
	bridgeMutateOps, err := ovs.
		Where(br2).
		Mutate(br2,
			model.Mutation{
				Field:   &br2.Ports,
				Mutator: "insert",
				Value:   []string{newPort.UUID},
			})
	if err != nil {
		return fmtErrorf("Error creating bridge mutation ops: %v", err)
	}
	fmt.Printf("Created bridgeMutateOps: # %d\n", len(bridgeMutateOps))
	operations = append(operations, bridgeMutateOps...)

	// operations := append(insertOp, mutateOps...)
	fmt.Printf("Transacting operations: # %d\n", len(operations))

	reply, err := ovs.Transact(context.TODO(), operations...)
	if err != nil {
		return fmtErrorf("Error in Transact. err: %v", err)
	}
	if _, err := ovsdb.CheckOperationResults(reply, operations); err != nil {
		return fmtErrorf("Error from CheckOperationResults. err: %v", err)
	}
	fmt.Println("Interface Addition Successful : ", reply[0].UUID.GoUUID)

	return nil
}

// func createInterface(ovs client.Client, interfaceName string) {
// 	log.Printf("createInterface(): creating interface named: %s", *nwdevice)

// 	newInterface := vswitchd.Interface{
// 		UUID: "atif0",
// 		Name: interfaceName,
// 	}
// 	insertOps, err := ovs.Create(&newInterface)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// ovsRow := vswitchd.OpenvSwitch{
// 	// 	UUID: rootUUID,
// 	// }
// 	// mutateOps, err := ovs.Where(&ovsRow).Mutate(&ovsRow, model.Mutation{
// 	// 	Field:   &ovsRow.Bridges,
// 	// 	Mutator: "insert",
// 	// 	Value:   []string{bridge.UUID},
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// operations := append(insertOp, mutateOps...)
// 	operations := insertOps
// 	reply, err := ovs.Transact(context.TODO(), operations...)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if _, err := ovsdb.CheckOperationResults(reply, operations); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Interface Addition Successful : ", reply[0].UUID.GoUUID)
// }

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
			log.Printf("AddFunc: Received add in table: %s\n", table)
			log.Printf("AddFunc: Received add: %v\n", model)
		},
		UpdateFunc: func(table string, old, new model.Model) {
			log.Printf("UpdateFunc: Received update in table: %s\n", table)
			log.Printf("UpdateFunc: Received update from: %v\n", new)
			log.Printf("UpdateFunc: Received update to: %v\n", new)
		},
		DeleteFunc: func(table string, model model.Model) {
			log.Printf("DeleteFunc: Received delete in table: %s\n", table)
			log.Printf("DeleteFunc: Received delete from: %v\n", model)
		},
	})
	_, err = ovs.MonitorAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success MonitorAll.")

	log.Println("Sleeping 2 seconds.")
	time.Sleep(2 * time.Second)
	log.Println("Done sleeping 2 seconds.")

	// Get root UUID
	for uuid := range ovs.Cache().Table("Open_vSwitch").Rows() {
		rootUUID = uuid
		log.Printf("Got root UUID: %s\n", rootUUID)
	}
	log.Printf("Final root UUID is: %s\n", rootUUID)

	if len(*bridgeName) <= 0 {
		log.Fatal("Need non-empty --bridge argument.")
	}
	log.Printf("About to modify bridge named: %s", *bridgeName)

	if len(*nwdevice) <= 0 {
		log.Fatal("Need non-empty --device argument.")
	}
	log.Printf("About to create interface named: %s", *nwdevice)

	createPortOnBridge(ovs, *bridgeName, *nwdevice)

	log.Println("Sleeping 4 seconds.")
	time.Sleep(4 * time.Second)
	log.Println("Done sleeping 4 seconds.")

	log.Println("Exiting.")
}

ovsdb-tool create ovs.ovsdb /usr/share/openvswitch/vswitch.ovsschema
./run/ovs/bin/ovs-vsctl --db=tcp:127.0.0.1:6641 show

# MapperSDK Example
Use a sample device to guide you to use MapperSDK

## Function introduction
Generate random numbers and report them to the cloud, and control 
the range of generating numbers according to the value returned by the cloud

## How to run this code
```shell
cd cmd
```
```shell
go run main.go --v 4 // Debug level log
```

## MQTT Information
### MQTT Security Configuration
Developers need to configure their own authentication keys. Mapper-sdk enables server-side security authentication by default, but does not enable client-side authentication. If necessary, we will add it in the next version.
Developers need to provide the root certificate path, client key path and client certificate path in the `config.yaml` file.
### Send
The program can collect information from device.Then send to mqtt broker according to the ```CollectCycle``` in configmap  
If you set the log level to 4, you can see the information in the terminal.
### Set
The program can subscribe to messages sent from the cloud and perform corresponding tasks. You can use these topics  
1. ```$hw/events/device/``` + integer-generator-01+```/twin/update/delta```  to set VirtualDevice's limit
2. ```"$hw/events/device/```+integer-generator-01+```/state/update"``` to get VirtualDevice's state
3. ```"$hw/events/node/```+ integer-generator-03+```/membership/updated"``` to check the JSON file and add or delete the device according to the payload content
## RestFul API Usage Example
### HTTPS Security Configuration
Developers need to configure the authentication key by themselves. By default, mapper-sdk enables https two-way security authentication for restful interfaces. The authentication of the client to the server needs to be implemented by the developer.
Developers need to provide the root certificate path, server key path and server certificate path in the config.yaml file.  
This example provides the security key of the client side.You can find them in [res/https-client-key](./res/https-client-key).Configure the key to your HTTPS client to access the following restful API
### <font color=green>**GET**</font>   GetFloatData
```https://127.0.0.1:1215/api/v1/device/id/integer-generator-01/random-float```

### <font color=orange>**POST**</font> AddDevice

```https://127.0.0.1:1215/api/v1/callback/device```
#### Body
```json
{
    "deviceInstances": 
        {
            "id": "integer-generator-05",
            "name": "random",
            "protocol": "virtual-protocol",
            "model": "random-01",
            "twins": [
                {
                    "propertyName": "random-int",
                    "desired": {
                        "value": "30",
                        "metadata": {
                            "timestamp": "1550049403598",
                            "type": "int"
                        }
                    },
                    "reported": {
                        "value": "30",
                        "metadata": {
                            "timestamp": "1550049403598",
                            "type": "int"
                        }
                    }
                },
                {
                    "propertyName": "random-float",
                    "desired": {
                        "value": "30",
                        "metadata": {
                            "timestamp": "1550049403598",
                            "type": "float"
                        }
                    },
                    "reported": {
                        "value": "30",
                        "metadata": {
                            "timestamp": "1550049403598",
                            "type": "float"
                        }
                    }
                }
            ],
            "propertyVisitors": [
                {
                    "name": "random-int",
                    "propertyName": "random-int",
                    "modelName": "random-01",
                    "protocol": "random",
                    "CollectCycle": 2000000000,
                    "visitorConfig": {
                        "dataType": "int"
                    }
                },
                {
                    "name": "random-float",
                    "propertyName": "random-float",
                    "modelName": "random-01",
                    "protocol": "random",
                    "CollectCycle": 3000000000,
                    "visitorConfig": {
                        "dataType": "float"
                    }
                }
            ]
        },
	"deviceModels": [{
		"name": "random-01",
		"properties": [{
			"name": "random-int",
			"dataType": "int",
			"description": "Random number of virtual device production",
			"accessMode": "ReadWrite",
			"defaultValue": 0,
			"minimum": 0,
			"maximum": 1000
		},{
			"name": "random-float",
			"dataType": "float",
			"description": "Random number of virtual device production",
			"accessMode": "ReadOnly",
			"defaultValue": 0,
			"minimum": 0,
			"maximum": 1000
		}]
	},{
		"name": "random-02",
		"properties": [{
			"name": "random-int",
			"dataType": "int",
			"description": "Random number of virtual device production",
			"accessMode": "ReadWrite",
			"defaultValue": 0,
			"minimum": 0,
			"maximum": 100
		},{
			"name": "random-float",
			"dataType": "float",
			"description": "Random number of virtual device production",
			"accessMode": "ReadOnly",
			"defaultValue": 0,
			"minimum": 0,
			"maximum": 100
		}]
	}],
	"protocols": [{
		"name": "virtual-protocol",
		"protocol": "random",
		"protocolConfig": {
			"deviceId": 1
		},
		"protocolCommonConfig": {
		}
	},{
		"name": "virtual-protocol-test",
		"protocol": "random",
		"protocolConfig": {
			"deviceId": 2
		},
		"protocolCommonConfig": {
		}
	}]
    
}
```

### <font color=#60D6F4>**PUT**</font> WriteData
```https://127.0.0.1:1215/api/v1/device/id/integer-generator-01?random-int=1```
#### QueryParams:random-int

### <font color=#FF5555>**DEL**</font>  RemoveDevice
```https://127.0.0.1:1215/api/v1/callback/device/id/integer-generator-02```


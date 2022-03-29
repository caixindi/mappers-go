![Gitee Latest Dev Tag](https://img.shields.io/badge/latest--dev-v0.0.1-orange) ![Gitee go.mod Go version](https://img.shields.io/badge/Go-v1.17-brightgreen) [![Gitee License](https://camo.githubusercontent.com/3e671e69d5fad7978893d028dcdeb3af16edb20b61f23cd276f738a76f33f3cf/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6963656e73652f6b756265656467652f6b756265656467652e7376673f7374796c653d666c61742d737175617265)](https://gitee.com/ascend/mappers-go-sample/blob/mapper-go-sdk/LICENSE)
# MapperSDK

## OverView
This repository is a set of Go packages that can be used to build Go-based mapper for use within the KubeEdge framework.

## Usage
Developers need to provide [configmap](./example/virtualDevice/res/virtual_device_test.json) and [configuration file](./example/virtualDevice/res/config.yaml),and then

Developers can make their own mapper by implementing the [ProtocolDriver](pkg/models/protocoldriver.go) interface for their desired IoT protocol, and the [Bootstrap](pkg/service/main.go) function to start the Device Service. 

Please see the [VirtualDevice](example/virtualDevice/) mapper provided in this repository as an example.

## Command Line Options

      --config-file string          Config file name (default "..\\res\\config.yaml")
      --mqtt-address string         MQTT broker address
      --mqtt-certification string   certification file path
      --mqtt-password string        password
      --mqtt-privatekey string      private key file path
      --mqtt-username string        username
      --v string                    log level (default "1")


## Supported MQTT topics
	TopicTwinUpdateDelta = "$hw/events/device/%s/twin/update/delta"
	TopicTwinUpdate      = "$hw/events/device/%s/twin/update"
	TopicStateUpdate     = "$hw/events/device/%s/state/update"
	TopicDataUpdate      = "$ke/events/device/%s/data/update"
	TopicDeviceUpdate    = "$hw/events/node/#"
### in addition
If you want to accept large packets over HTTPS instead of mqtt, you can set ```CollectCycle``` to ```-1``` in configmap.  
Then the twin that ```CollectCycle``` be sett to ```-1``` will not be actively reported to mqtt broker
## Supported RESTful API
The URLs listed below are given in the form of local IP. You can use these services from any network accessible to mapper   

Port ```1215``` is enabled by default.      

 

```deviceInstances-ID```  
according to your own CRD definition  

```propertyName```  
according to your own CRD definition  

If you have any questions,you can see examples in the [example](example/virtualDevice/README.md)  

The functions and urls are as follows. 
1. Detect whether the RESTful service starts normally  
   Method=<font color=green>**GET**</font>   
   https://127.0.0.1:1215/api/v1/ping

2. Get device's property  
Method=<font color=green>**GET**</font>  
https://127.0.0.1:1215/api/v1/device/id/deviceInstances-ID/propertyName

3. Set device's config  
Method=<font color=#60D6F4>**PUT**</font> 
https://127.0.0.1:1215/api/v1/device/id/deviceInstances-ID?propertyName=Value
4. Add a deviceInstance  
Method=<font color=orange>**POST**</font>  
https://127.0.0.1:1215/api/v1/callback/device  
You must provide a JSON body that conforms to the CRD definition
5. Delete a deviceInstance  
   Method=<font color=#FF5555>**DEL**</font>
   https://127.0.0.1:1215/api/v1/callback/device/id/deviceInstances-ID

## File Check
If you want to verify whether the configuration file you provided is 
parsed correctly. Put the file in the appropriate package
(```config.yaml``` should be placed config package, ```configmap.json``` should be placed configmap package),
and you can directly view the output results through ```make test```
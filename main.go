package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/goburrow/modbus"
	"math"
	"time"
)

//type a map.s.int32

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/
const version = "0.00.1"

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

type jan_type_of_data int

var jan_float jan_type_of_data = 1
var jan_short jan_type_of_data = 2

type Modbusparam struct {
	Num        int
	Id         uint16
	Name       string
	TypeOfData jan_type_of_data
}

type Modbusparams []Modbusparam

var paramName0 = Modbusparams{
	{0, 1000, "Voltage0", jan_float},
	{1, 1002, "Voltage1", jan_float},
	{2, 1004, "Voltage2", jan_float},
	{3, 1006, "Voltage3", jan_float},
	{4, 1008, "Voltage4", jan_float},
	{5, 1010, "Voltage5", jan_float},
	{6, 1272, "Mains_frequency", jan_float},
	{7, 1394, "Vol_of_fund_wave0", jan_float},
	{8, 1396, "Vol_of_fund_wave1", jan_float},
	{9, 1398, "Vol_of_fund_wave2", jan_float},
	{10, 1400, "Vol_of_fund_wave3", jan_float},
	{11, 1402, "Vol_of_fund_wave4", jan_float},
	{12, 1404, "Vol_of_fund_wave5", jan_float},
	{13, 1406, "Pha_ang_of_volt0", jan_float},
	{14, 1408, "Pha_ang_of_volt1", jan_float},
	{15, 1410, "Pha_ang_of_volt2", jan_float},
	{16, 1412, "Pha_ang_of_volt3", jan_float},
	{17, 1414, "Pha_ang_of_volt4", jan_float},
	{18, 1416, "Pha_ang_of_volt5", jan_float},
	{19, 1012, "Current0", jan_float},
	{20, 1014, "Current1", jan_float},
	{21, 1016, "Current2", jan_float},
	{22, 1052, "Real_power0", jan_float},
	{23, 1054, "Real_power1", jan_float},
	{24, 1056, "Real_power2", jan_float},
	{25, 1092, "Reactive_power0", jan_float},
	{26, 1094, "Reactive_power1", jan_float},
	{27, 1096, "Reactive_power2", jan_float},
	{28, 1132, "Apparent_power0", jan_float},
	{29, 1134, "Apparent_power1", jan_float},
	{30, 1136, "Apparent_power2", jan_float},
	{31, 1172, "Power_factor0", jan_float},
	{32, 1174, "Power_factor1", jan_float},
	{33, 1176, "Power_factor2", jan_float},
	{34, 1212, "Energy0", jan_float},
	{35, 1214, "Energy1", jan_float},
	{36, 1216, "Energy2", jan_float},
	{37, 1252, "Reset_energy0", jan_short},
	{38, 1253, "Reset_energy1", jan_short},
	{39, 1254, "Reset_energy2", jan_short},
	{40, 1274, "Curr_of_fund_wave0", jan_float},
	{41, 1276, "Curr_of_fund_wave1", jan_float},
	{42, 1278, "Curr_of_fund_wave2", jan_float},
	{43, 1314, "cos_Phi0", jan_float},
	{44, 1316, "cos_Phi1", jan_float},
	{45, 1318, "cos_Phi2", jan_float},
	{46, 1354, "Pha_ang_of_curr0", jan_float},
	{47, 1356, "Pha_ang_of_curr1", jan_float},
	{48, 1358, "Pha_ang_of_curr2", jan_float},
}

var paramName1 = Modbusparams{
	{0, 1000, "Voltage0", jan_float},
	{1, 1002, "Voltage1", jan_float},
	{2, 1004, "Voltage2", jan_float},
	{3, 1006, "Voltage3", jan_float},
	{4, 1008, "Voltage4", jan_float},
	{5, 1010, "Voltage5", jan_float},
	{6, 1272, "Mains_frequency", jan_float},
	{7, 1394, "Vol_of_fund_wave0", jan_float},
	{8, 1396, "Vol_of_fund_wave1", jan_float},
	{9, 1398, "Vol_of_fund_wave2", jan_float},
	{10, 1400, "Vol_of_fund_wave3", jan_float},
	{11, 1402, "Vol_of_fund_wave4", jan_float},
	{12, 1404, "Vol_of_fund_wave5", jan_float},
	{13, 1406, "Pha_ang_of_volt0", jan_float},
	{14, 1408, "Pha_ang_of_volt1", jan_float},
	{15, 1410, "Pha_ang_of_volt2", jan_float},
	{16, 1412, "Pha_ang_of_volt3", jan_float},
	{17, 1414, "Pha_ang_of_volt4", jan_float},
	{18, 1416, "Pha_ang_of_volt5", jan_float},
	{19, 1018, "Current3", jan_float},
	{20, 1020, "Current4", jan_float},
	{21, 1022, "Current5", jan_float},
	{22, 1058, "Real_power3", jan_float},
	{23, 1060, "Real_power4", jan_float},
	{24, 1062, "Real_power5", jan_float},
	{25, 1098, "Reactive_power3", jan_float},
	{26, 1100, "Reactive_power4", jan_float},
	{27, 1102, "Reactive_power5", jan_float},
	{28, 1138, "Apparent_power3", jan_float},
	{29, 1140, "Apparent_power4", jan_float},
	{30, 1142, "Apparent_power5", jan_float},
	{31, 1178, "Power_factor3", jan_float},
	{32, 1180, "Power_factor4", jan_float},
	{33, 1182, "Power_factor5", jan_float},
	{34, 1218, "Energy3", jan_float},
	{35, 1220, "Energy4", jan_float},
	{36, 1222, "Energy5", jan_float},
	{37, 1255, "Reset_energy3", jan_short},
	{38, 1256, "Reset_energy4", jan_short},
	{39, 1257, "Reset_energy5", jan_short},
	{40, 1280, "Curr_of_fund_wave3", jan_float},
	{41, 1282, "Curr_of_fund_wave4", jan_float},
	{42, 1284, "Curr_of_fund_wave5", jan_float},
	{43, 1320, "cos_Phi3", jan_float},
	{44, 1322, "cos_Phi4", jan_float},
	{45, 1324, "cos_Phi5", jan_float},
	{46, 1360, "Pha_ang_of_curr3", jan_float},
	{47, 1362, "Pha_ang_of_curr4", jan_float},
	{48, 1364, "Pha_ang_of_curr5", jan_float},
}

var paramName2 = Modbusparams{
	{0, 1000, "Voltage0", jan_float},
	{1, 1002, "Voltage1", jan_float},
	{2, 1004, "Voltage2", jan_float},
	{3, 1006, "Voltage3", jan_float},
	{4, 1008, "Voltage4", jan_float},
	{5, 1010, "Voltage5", jan_float},
	{6, 1272, "Mains_frequency", jan_float},
	{7, 1394, "Vol_of_fund_wave0", jan_float},
	{8, 1396, "Vol_of_fund_wave1", jan_float},
	{9, 1398, "Vol_of_fund_wave2", jan_float},
	{10, 1400, "Vol_of_fund_wave3", jan_float},
	{11, 1402, "Vol_of_fund_wave4", jan_float},
	{12, 1404, "Vol_of_fund_wave5", jan_float},
	{13, 1406, "Pha_ang_of_volt0", jan_float},
	{14, 1408, "Pha_ang_of_volt1", jan_float},
	{15, 1410, "Pha_ang_of_volt2", jan_float},
	{16, 1412, "Pha_ang_of_volt3", jan_float},
	{17, 1414, "Pha_ang_of_volt4", jan_float},
	{18, 1416, "Pha_ang_of_volt5", jan_float},
	{19, 1024, "Current6", jan_float},
	{20, 1026, "Current7", jan_float},
	{21, 1028, "Current8", jan_float},
	{22, 1064, "Real_power6", jan_float},
	{23, 1066, "Real_power7", jan_float},
	{24, 1068, "Real_power8", jan_float},
	{25, 1104, "Reactive_power6", jan_float},
	{26, 1106, "Reactive_power7", jan_float},
	{27, 1108, "Reactive_power8", jan_float},
	{28, 1144, "Apparent_power6", jan_float},
	{29, 1146, "Apparent_power7", jan_float},
	{30, 1148, "Apparent_power8", jan_float},
	{31, 1184, "Power_factor6", jan_float},
	{32, 1186, "Power_factor7", jan_float},
	{33, 1188, "Power_factor8", jan_float},
	{34, 1224, "Energy6", jan_float},
	{35, 1226, "Energy7", jan_float},
	{36, 1228, "Energy8", jan_float},
	{37, 1258, "Reset_energy6", jan_short},
	{38, 1259, "Reset_energy7", jan_short},
	{39, 1260, "Reset_energy8", jan_short},
	{40, 1286, "Curr_of_fund_wave6", jan_float},
	{41, 1288, "Curr_of_fund_wave7", jan_float},
	{42, 1290, "Curr_of_fund_wave8", jan_float},
	{43, 1326, "cos_Phi6", jan_float},
	{44, 1328, "cos_Phi7", jan_float},
	{45, 1330, "cos_Phi8", jan_float},
	{46, 1366, "Pha_ang_of_curr6", jan_float},
	{47, 1368, "Pha_ang_of_curr7", jan_float},
	{48, 1370, "Pha_ang_of_curr8", jan_float},
}

var paramName3 = Modbusparams{
	{0, 1000, "Voltage0", jan_float},
	{1, 1002, "Voltage1", jan_float},
	{2, 1004, "Voltage2", jan_float},
	{3, 1006, "Voltage3", jan_float},
	{4, 1008, "Voltage4", jan_float},
	{5, 1010, "Voltage5", jan_float},
	{6, 1272, "Mains_frequency", jan_float},
	{7, 1394, "Vol_of_fund_wave0", jan_float},
	{8, 1396, "Vol_of_fund_wave1", jan_float},
	{9, 1398, "Vol_of_fund_wave2", jan_float},
	{10, 1400, "Vol_of_fund_wave3", jan_float},
	{11, 1402, "Vol_of_fund_wave4", jan_float},
	{12, 1404, "Vol_of_fund_wave5", jan_float},
	{13, 1406, "Pha_ang_of_volt0", jan_float},
	{14, 1408, "Pha_ang_of_volt1", jan_float},
	{15, 1410, "Pha_ang_of_volt2", jan_float},
	{16, 1412, "Pha_ang_of_volt3", jan_float},
	{17, 1414, "Pha_ang_of_volt4", jan_float},
	{18, 1416, "Pha_ang_of_volt5", jan_float},
	{19, 1030, "Current9", jan_float},
	{20, 1032, "Current10", jan_float},
	{21, 1034, "Current11", jan_float},
	{22, 1070, "Real_power9", jan_float},
	{23, 1072, "Real_power10", jan_float},
	{24, 1074, "Real_power11", jan_float},
	{25, 1110, "Reactive_power9", jan_float},
	{26, 1112, "Reactive_power10", jan_float},
	{27, 1114, "Reactive_power11", jan_float},
	{28, 1150, "Apparent_power9", jan_float},
	{29, 1152, "Apparent_power10", jan_float},
	{30, 1154, "Apparent_power11", jan_float},
	{31, 1190, "Power_factor9", jan_float},
	{32, 1192, "Power_factor10", jan_float},
	{33, 1194, "Power_factor11", jan_float},
	{34, 1230, "Energy9", jan_float},
	{35, 1232, "Energy10", jan_float},
	{36, 1234, "Energy11", jan_float},
	{37, 1261, "Reset_energy9", jan_short},
	{38, 1262, "Reset_energy10", jan_short},
	{39, 1263, "Reset_energy11", jan_short},
	{40, 1292, "Curr_of_fund_wave9", jan_float},
	{41, 1294, "Curr_of_fund_wave10", jan_float},
	{42, 1296, "Curr_of_fund_wave11", jan_float},
	{43, 1332, "cos_Phi9", jan_float},
	{44, 1334, "cos_Phi10", jan_float},
	{45, 1336, "cos_Phi11", jan_float},
	{46, 1372, "Pha_ang_of_curr9", jan_float},
	{47, 1374, "Pha_ang_of_curr10", jan_float},
	{48, 1376, "Pha_ang_of_curr11", jan_float},
}

var paramName4 = Modbusparams{
	{0, 1000, "Voltage0", jan_float},
	{1, 1002, "Voltage1", jan_float},
	{2, 1004, "Voltage2", jan_float},
	{3, 1006, "Voltage3", jan_float},
	{4, 1008, "Voltage4", jan_float},
	{5, 1010, "Voltage5", jan_float},
	{6, 1272, "Mains_frequency", jan_float},
	{7, 1394, "Vol_of_fund_wave0", jan_float},
	{8, 1396, "Vol_of_fund_wave1", jan_float},
	{9, 1398, "Vol_of_fund_wave2", jan_float},
	{10, 1400, "Vol_of_fund_wave3", jan_float},
	{11, 1402, "Vol_of_fund_wave4", jan_float},
	{12, 1404, "Vol_of_fund_wave5", jan_float},
	{13, 1406, "Pha_ang_of_volt0", jan_float},
	{14, 1408, "Pha_ang_of_volt1", jan_float},
	{15, 1410, "Pha_ang_of_volt2", jan_float},
	{16, 1412, "Pha_ang_of_volt3", jan_float},
	{17, 1414, "Pha_ang_of_volt4", jan_float},
	{18, 1416, "Pha_ang_of_volt5", jan_float},
	{19, 1036, "Current12", jan_float},
	{20, 1038, "Current13", jan_float},
	{21, 1040, "Current14", jan_float},
	{22, 1076, "Real_power12", jan_float},
	{23, 1078, "Real_power13", jan_float},
	{24, 1080, "Real_power14", jan_float},
	{25, 1116, "Reactive_power12", jan_float},
	{26, 1118, "Reactive_power13", jan_float},
	{27, 1120, "Reactive_power14", jan_float},
	{28, 1156, "Apparent_power12", jan_float},
	{29, 1158, "Apparent_power13", jan_float},
	{30, 1160, "Apparent_power14", jan_float},
	{31, 1196, "Power_factor12", jan_float},
	{32, 1198, "Power_factor13", jan_float},
	{33, 1200, "Power_factor14", jan_float},
	{34, 1236, "Energy12", jan_float},
	{35, 1238, "Energy13", jan_float},
	{36, 1240, "Energy14", jan_float},
	{37, 1264, "Reset_energy12", jan_short},
	{38, 1265, "Reset_energy13", jan_short},
	{39, 1266, "Reset_energy14", jan_short},
	{40, 1298, "Curr_of_fund_wave12", jan_float},
	{41, 1300, "Curr_of_fund_wave13", jan_float},
	{42, 1302, "Curr_of_fund_wave14", jan_float},
	{43, 1338, "cos_Phi12", jan_float},
	{44, 1340, "cos_Phi13", jan_float},
	{45, 1342, "cos_Phi14", jan_float},
	{46, 1378, "Pha_ang_of_curr12", jan_float},
	{47, 1380, "Pha_ang_of_curr13", jan_float},
	{48, 1382, "Pha_ang_of_curr14", jan_float},
}

var paramName5 = Modbusparams{
	{0, 1000, "Voltage0", jan_float},
	{1, 1002, "Voltage1", jan_float},
	{2, 1004, "Voltage2", jan_float},
	{3, 1006, "Voltage3", jan_float},
	{4, 1008, "Voltage4", jan_float},
	{5, 1010, "Voltage5", jan_float},
	{6, 1272, "Mains_frequency", jan_float},
	{7, 1394, "Vol_of_fund_wave0", jan_float},
	{8, 1396, "Vol_of_fund_wave1", jan_float},
	{9, 1398, "Vol_of_fund_wave2", jan_float},
	{10, 1400, "Vol_of_fund_wave3", jan_float},
	{11, 1402, "Vol_of_fund_wave4", jan_float},
	{12, 1404, "Vol_of_fund_wave5", jan_float},
	{13, 1406, "Pha_ang_of_volt0", jan_float},
	{14, 1408, "Pha_ang_of_volt1", jan_float},
	{15, 1410, "Pha_ang_of_volt2", jan_float},
	{16, 1412, "Pha_ang_of_volt3", jan_float},
	{17, 1414, "Pha_ang_of_volt4", jan_float},
	{18, 1416, "Pha_ang_of_volt5", jan_float},
	{19, 1042, "Current15", jan_float},
	{20, 1044, "Current16", jan_float},
	{21, 1046, "Current17", jan_float},
	{22, 1048, "Current18", jan_float},
	{23, 1050, "Current19", jan_float},
	{24, 1082, "Real_power15", jan_float},
	{25, 1084, "Real_power16", jan_float},
	{26, 1086, "Real_power17", jan_float},
	{27, 1088, "Real_power18", jan_float},
	{28, 1090, "Real_power19", jan_float},
	{29, 1122, "Reactive_power15", jan_float},
	{30, 1124, "Reactive_power16", jan_float},
	{31, 1126, "Reactive_power17", jan_float},
	{32, 1128, "Reactive_power18", jan_float},
	{33, 1130, "Reactive_power19", jan_float},
	{34, 1162, "Apparent_power15", jan_float},
	{35, 1164, "Apparent_power16", jan_float},
	{36, 1166, "Apparent_power17", jan_float},
	{37, 1168, "Apparent_power18", jan_float},
	{38, 1170, "Apparent_power19", jan_float},
	{39, 1202, "Power_factor15", jan_float},
	{40, 1204, "Power_factor16", jan_float},
	{41, 1206, "Power_factor17", jan_float},
	{42, 1208, "Power_factor18", jan_float},
	{43, 1210, "Power_factor19", jan_float},
	{44, 1242, "Energy15", jan_float},
	{45, 1244, "Energy16", jan_float},
	{46, 1246, "Energy17", jan_float},
	{47, 1248, "Energy18", jan_float},
	{48, 1250, "Energy19", jan_float},
	{49, 1267, "Reset_energy15", jan_short},
	{50, 1268, "Reset_energy16", jan_short},
	{51, 1269, "Reset_energy17", jan_short},
	{52, 1270, "Reset_energy18", jan_short},
	{53, 1271, "Reset_energy19", jan_short},
	{54, 1304, "Curr_of_fund_wave15", jan_float},
	{55, 1306, "Curr_of_fund_wave16", jan_float},
	{56, 1308, "Curr_of_fund_wave17", jan_float},
	{57, 1310, "Curr_of_fund_wave18", jan_float},
	{58, 1312, "Curr_of_fund_wave19", jan_float},
	{59, 1344, "cos_Phi15", jan_float},
	{60, 1346, "cos_Phi16", jan_float},
	{61, 1348, "cos_Phi17", jan_float},
	{62, 1350, "cos_Phi18", jan_float},
	{63, 1352, "cos_Phi19", jan_float},
	{64, 1384, "Pha_ang_of_curr15", jan_float},
	{65, 1386, "Pha_ang_of_curr16", jan_float},
	{66, 1388, "Pha_ang_of_curr17", jan_float},
	{67, 1390, "Pha_ang_of_curr18", jan_float},
	{68, 1392, "Pha_ang_of_curr19", jan_float},
}

func main() {

	addressIP := flag.String("ip", "127.0.0.1", "a string")
	tcpPort := flag.String("port", "502", "a string")
	slaveID := flag.Int("id", 1, "an int")
	typeOfdata := flag.Uint("type", 0, "an uint")
	flag.Parse()
	serverParam := fmt.Sprint(*addressIP, ":", *tcpPort)
	s := byte(*slaveID)
	var tOfdata = *typeOfdata

	switch tOfdata {
	case 0:
		var dataA, dataB = creatData(paramName0, serverParam, s)
		PrintJSON(dataA, dataB, paramName0)
	case 1:
		var dataA, dataB = creatData(paramName1, serverParam, s)
		PrintJSON(dataA, dataB, paramName1)
	case 2:
		var dataA, dataB = creatData(paramName2, serverParam, s)
		PrintJSON(dataA, dataB, paramName2)
	case 3:
		var dataA, dataB = creatData(paramName3, serverParam, s)
		PrintJSON(dataA, dataB, paramName3)
	case 4:
		var dataA, dataB = creatData(paramName4, serverParam, s)
		PrintJSON(dataA, dataB, paramName4)
	case 5:
		var dataA, dataB = creatData(paramName5, serverParam, s)
		PrintJSON(dataA, dataB, paramName5)
	}
}

func PrintJSON(dataF []float32, dataS []uint16, param Modbusparams) {
	/*fmt.Println(dataF)
	fmt.Println(len(dataF))
	fmt.Println(dataS)
	fmt.Println(len(dataS))*/
	for l := 0; l < len(dataF)+3; l++ {
		if l == 0 {
			fmt.Printf("{\"%d_%s\":", param[l].Id, param[l].Name)
			fmt.Print(dataF[l])
		} else if l > 0 && l < 37 {
			fmt.Printf(",\"%d_%s\":", param[l].Id, param[l].Name)
			fmt.Print(dataF[l])
		} else if l >= 37 && l < 40 {
			fmt.Printf(",\"%d_%s\":", param[l].Id, param[l].Name)
			fmt.Print(dataS[l-37])
		} else if l >= 40 {
			fmt.Printf(",\"%d_%s\":", param[l].Id, param[l].Name)
			fmt.Print(dataF[l-3])
		}
	}
	fmt.Printf(",\"version\":\"%s\"}", version)
}

func ModbusQuery(address uint16, quantity uint16, serverParams string, slaveID byte) []byte {
	handler := modbus.NewTCPClientHandler(serverParams)
	handler.SlaveId = slaveID
	handler.Timeout = 2 * time.Second

	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	results, err := client.ReadHoldingRegisters(address, quantity)
	if err != nil {
		fmt.Printf("{\"status\":\"error\", \"error\":\"%s\"}", err)
		//fmt.Printf("%s\n", err)
	}
	//fmt.Printf("%v", results)
	return results
}

func creatData(params Modbusparams, serverParam string, s byte) (dataFloat []float32, dataShort []uint16) {
	//var dataFloat []float32
	//var dataShort []uint16

	tmpdata := ModbusQuery(params[0].Id, 6*2, serverParam, s)
	i := 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}

	tmpdata = ModbusQuery(params[6].Id, 1*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[7].Id, 12*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[19].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[22].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[25].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[28].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[31].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	tmpdata = ModbusQuery(params[34].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}

	tmpdata = ModbusQuery(params[37].Id, 3, serverParam, s) //Janitsa short
	i = 0
	for i < len(tmpdata) {
		a := binary.BigEndian.Uint16(tmpdata[i : i+2])
		dataShort = append(dataShort, a)
		i += 2
	}

	tmpdata = ModbusQuery(params[40].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}

	tmpdata = ModbusQuery(params[43].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}

	tmpdata = ModbusQuery(params[46].Id, 3*2, serverParam, s)
	i = 0
	for i < len(tmpdata) {
		a := Float32frombytes(tmpdata[i : i+4])
		if math.IsNaN(float64(a)) {
			dataFloat = append(dataFloat, 0)
		} else {
			dataFloat = append(dataFloat, a)
		}
		i += 4
	}
	return dataFloat, dataShort
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

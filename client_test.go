package cgminer_client

import (
	"testing"
)

func TestProcessChipStat(t *testing.T) {
	jsonStr := `{"STATUS":[{"STATUS":"S","When":1517894216,"Code":9,"Msg":"4 ASC(s)","Description":"cgminer 4.10.0"}],"DEVS":[{"ASC":0,"Name":"ttyS1","ID":1,"1_accept":154,"2_accept":132,"3_accept":138,"4_accept":129,"5_accept":134,"6_accept":124,"7_accept":53,"8_accept":2,"9_accept":151,"10_accept":147,"11_accept":133,"12_accept":2,"13_accept":108,"14_accept":131,"15_accept":126,"16_accept":143,"17_accept":124,"18_accept":125,"19_accept":151,"20_accept":123,"21_accept":124,"22_accept":111,"23_accept":128,"24_accept":140,"25_accept":174,"26_accept":139,"27_accept":53,"28_accept":26,"29_accept":41,"30_accept":52,"31_accept":121,"32_accept":157,"33_accept":123,"34_accept":69,"35_accept":1,"36_accept":0},{"ASC":1,"Name":"ttyS2","ID":2,"1_accept":146,"2_accept":149,"3_accept":124,"4_accept":158,"5_accept":164,"6_accept":148,"7_accept":152,"8_accept":13,"9_accept":151,"10_accept":156,"11_accept":137,"12_accept":133,"13_accept":143,"14_accept":149,"15_accept":144,"16_accept":138,"17_accept":125,"18_accept":143,"19_accept":178,"20_accept":160,"21_accept":154,"22_accept":111,"23_accept":118,"24_accept":153,"25_accept":167,"26_accept":144,"27_accept":56,"28_accept":45,"29_accept":98,"30_accept":93,"31_accept":112,"32_accept":114,"33_accept":141,"34_accept":0,"35_accept":0,"36_accept":0},{"ASC":2,"Name":"ttyS3","ID":3,"1_accept":146,"2_accept":174,"3_accept":157,"4_accept":176,"5_accept":147,"6_accept":187,"7_accept":165,"8_accept":150,"9_accept":161,"10_accept":122,"11_accept":157,"12_accept":144,"13_accept":156,"14_accept":177,"15_accept":175,"16_accept":100,"17_accept":152,"18_accept":186,"19_accept":126,"20_accept":150,"21_accept":157,"22_accept":137,"23_accept":143,"24_accept":150,"25_accept":174,"26_accept":130,"27_accept":161,"28_accept":134,"29_accept":80,"30_accept":119,"31_accept":92,"32_accept":161,"33_accept":90,"34_accept":152,"35_accept":83,"36_accept":0},{"ASC":3,"Name":"ttyS4","ID":4,"1_accept":163,"2_accept":151,"3_accept":161,"4_accept":44,"5_accept":168,"6_accept":82,"7_accept":166,"8_accept":8,"9_accept":155,"10_accept":151,"11_accept":135,"12_accept":35,"13_accept":146,"14_accept":132,"15_accept":125,"16_accept":152,"17_accept":156,"18_accept":151,"19_accept":142,"20_accept":150,"21_accept":163,"22_accept":112,"23_accept":142,"24_accept":164,"25_accept":143,"26_accept":144,"27_accept":105,"28_accept":49,"29_accept":46,"30_accept":116,"31_accept":150,"32_accept":130,"33_accept":151,"34_accept":5,"35_accept":0,"36_accept":4}],"id":1}`
	actualResult, err := processChipStat(jsonStr)
	if err != nil {
		t.Fatal(err)
	}

	csLen := len(actualResult.ChipStats)
	if csLen != 4 {
		t.Fatalf("Length is not 4, it is %s", csLen)
	}
}

func TestProcessDevs(t *testing.T) {
	jsonStr := `{"STATUS":[{"STATUS":"S","When":1517895412,"Code":9,"Msg":"4 ASC(s)","Description":"cgminer 4.10.0"}],"DEVS":[{"ASC":0,"Name":"ttyS1","ID":1,"Enabled":"Y","Status":"Alive","temperature":32.00,"MHS av":55.66,"MHS 5s":67.57,"MHS 1m":83.52,"MHS 5m":67.57,"MHS 15m":65.00,"Accepted":3810,"Rejected":21,"Hardware Errors":1179,"Utility":1.68,"Last Share Pool":0,"Last Share Time":1517895408,"Total MH":7584488.0000,"Diff1 Work":805,"Difficulty Accepted":2212.00000000,"Difficulty Rejected":14.00000000,"Last Share Difficulty":0.50000000,"Last Valid Work":1517895408,"Device Hardware%":59.4254,"Device Rejected%":1.7391,"Device Elapsed":136276},{"ASC":1,"Name":"ttyS2","ID":2,"Enabled":"Y","Status":"Alive","temperature":32.00,"MHS av":61.79,"MHS 5s":82.24,"MHS 1m":122.96,"MHS 5m":82.24,"MHS 15m":78.80,"Accepted":4345,"Rejected":14,"Hardware Errors":1068,"Utility":1.91,"Last Share Pool":0,"Last Share Time":1517895408,"Total MH":8421135.0000,"Diff1 Work":949,"Difficulty Accepted":2546.62500000,"Difficulty Rejected":8.00000000,"Last Share Difficulty":0.50000000,"Last Valid Work":1517895408,"Device Hardware%":52.9499,"Device Rejected%":0.8430,"Device Elapsed":136276},{"ASC":2,"Name":"ttyS3","ID":3,"Enabled":"Y","Status":"Alive","temperature":32.00,"MHS av":75.71,"MHS 5s":91.10,"MHS 1m":105.56,"MHS 5m":91.10,"MHS 15m":91.26,"Accepted":5099,"Rejected":22,"Hardware Errors":562,"Utility":2.24,"Last Share Pool":0,"Last Share Time":1517895387,"Total MH":10317967.0000,"Diff1 Work":1113,"Difficulty Accepted":2976.87500000,"Difficulty Rejected":12.12500000,"Last Share Difficulty":0.50000000,"Last Valid Work":1517895387,"Device Hardware%":33.5522,"Device Rejected%":1.0894,"Device Elapsed":136276},{"ASC":3,"Name":"ttyS4","ID":4,"Enabled":"Y","Status":"Alive","temperature":32.00,"MHS av":62.12,"MHS 5s":90.15,"MHS 1m":109.07,"MHS 5m":90.15,"MHS 15m":91.30,"Accepted":4238,"Rejected":8,"Hardware Errors":721,"Utility":1.87,"Last Share Pool":0,"Last Share Time":1517895395,"Total MH":8465538.0000,"Diff1 Work":848,"Difficulty Accepted":2435.75000000,"Difficulty Rejected":4.75000000,"Last Share Difficulty":0.50000000,"Last Valid Work":1517895395,"Device Hardware%":45.9528,"Device Rejected%":0.5601,"Device Elapsed":136276}],"id":1}`
	actualResult, err := processDevs(jsonStr)
	if err != nil {
		t.Fatal(err)
	}

	csLen := len(actualResult.Devs)
	if csLen != 4 {
		t.Fatalf("Length is not 4, it is %s", csLen)
	}

	// Make sure the temperature vs. Temperature works ok
	for _, dev := range actualResult.Devs {
		tmp := dev.Temperature
		if tmp == 0 {
			t.Fatalf("Temp should not be zero: %s", tmp)
		}
	}
}

func TestProcessSummary(t *testing.T) {
	jsonStr := `{"STATUS":[{"STATUS":"S","When":1517965137,"Code":11,"Msg":"Summary","Description":"cgminer 4.10.0"}],"SUMMARY":[{"Elapsed":48417,"MHS av":335.86,"MHS 5s":357.15,"MHS 1m":323.17,"MHS 5m":357.15,"MHS 15m":339.55,"temper1":31.00,"temper2":31.00,"temper3":31.00,"temper4":31.00,"Found Blocks":0,"Getworks":3236,"Accepted":8215,"Rejected":75,"Hardware Errors":1654,"Utility":10.18,"Discarded":48716,"Stale":0,"Get Failures":0,"Local Work":77113,"Remote Failures":0,"Network Blocks":366,"Total MH":16261155.0000,"Work Utility":1.39,"Difficulty Accepted":4538.37500000,"Difficulty Rejected":42.87500000,"Difficulty Stale":0.00000000,"Best Share":8322,"Device Hardware%":59.6036,"Device Rejected%":3.8247,"Pool Rejected%":0.9359,"Pool Stale%":0.0000,"Last getwork":1517965137}],"id":1}`
	actualResult, err := processSummary(jsonStr)
	if err != nil {
		t.Fatal(err)
	}

	elapsed := actualResult.Summary[0].Elapsed
	if elapsed != 48417 {
		t.Fatalf("Elapsed is not 48417, it is %s", elapsed)
	}
}

//func TestQuit(t *testing.T) {
//	client := New("10.10.0.11", 4028, 5)
//	err := client.Quit()
//	if err != nil {
//		t.Fatal(err)
//	}
//}

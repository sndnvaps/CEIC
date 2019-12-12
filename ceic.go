package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
{"shuju":
[
	{"id":"38461","CATA_ID":"CC20191212100824.00","SAVE_TIME":"2019-12-12 10:12:35","O_TIME":"2019-12-12 10:08:23","EPI_LAT":"33.06","EPI_LON":"121.37","EPI_DEPTH":10,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"3.5","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u6c5f\u82cf\u76d0\u57ce\u5e02\u4e1c\u53f0\u5e02\u6d77\u57df","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-12 10:17:41","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CC20191212100824"},
	{"id":"38460","CATA_ID":"CD20191212084919.00","SAVE_TIME":"2019-12-12 08:55:03","O_TIME":"2019-12-12 08:49:18","EPI_LAT":"36.94","EPI_LON":"78.78","EPI_DEPTH":10,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"4.0","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u65b0\u7586\u548c\u7530\u5730\u533a\u76ae\u5c71\u53bf","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-12 09:00:09","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CD20191212084919"},{"id":"38459","CATA_ID":"CD20191211173056.00","SAVE_TIME":"2019-12-11 17:33:38","O_TIME":"2019-12-11 17:30:56","EPI_LAT":"42.55","EPI_LON":"83.04","EPI_DEPTH":10,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"3.6","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u65b0\u7586\u963f\u514b\u82cf\u5730\u533a\u5e93\u8f66\u53bf","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-11 17:38:38","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CD20191211173056"},{"id":"38458","CATA_ID":"CD20191211091356.00","SAVE_TIME":"2019-12-11 09:18:48","O_TIME":"2019-12-11 09:13:56","EPI_LAT":"39.97","EPI_LON":"77.99","EPI_DEPTH":16,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"3.8","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u65b0\u7586\u5580\u4ec0\u5730\u533a\u5df4\u695a\u53bf","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-11 09:23:44","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CD20191211091356"},{"id":"38457","CATA_ID":"CC20191211055829.00","SAVE_TIME":"2019-12-11 06:22:14","O_TIME":"2019-12-11 05:58:28","EPI_LAT":"35.35","EPI_LON":"26.40","EPI_DEPTH":60,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"5.3","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u5e0c\u814a\u514b\u91cc\u7279\u5c9b","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-11 06:27:10","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CC20191211055829"},{"id":"38456","CATA_ID":"CD20191211041208.00","SAVE_TIME":"2019-12-11 04:20:27","O_TIME":"2019-12-11 04:12:07","EPI_LAT":"42.97","EPI_LON":"84.41","EPI_DEPTH":17,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"3.3","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u65b0\u7586\u5df4\u97f3\u90ed\u695e\u5dde\u548c\u9759\u53bf","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-11 04:25:22","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CD20191211041208"},{"id":"38455","CATA_ID":"CC20191211010460.00","SAVE_TIME":"2019-12-11 01:27:25","O_TIME":"2019-12-11 01:04:59","EPI_LAT":"30.71","EPI_LON":"141.80","EPI_DEPTH":10,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"5.9","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u65e5\u672c\u672c\u5dde\u4e1c\u5357\u6d77\u57df","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-11 01:32:19","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CC20191211010460"},{"id":"38454","CATA_ID":"CD20191211000252.00","SAVE_TIME":"2019-12-11 00:08:08","O_TIME":"2019-12-11 00:02:52","EPI_LAT":"24.89","EPI_LON":"99.17","EPI_DEPTH":8,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"3.4","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u4e91\u5357\u4fdd\u5c71\u5e02\u65bd\u7538\u53bf","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-11 00:13:01","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CD20191211000252"},{"id":"38453","CATA_ID":"CD20191210222800.00","SAVE_TIME":"2019-12-10 22:33:51","O_TIME":"2019-12-10 22:28:00","EPI_LAT":"40.30","EPI_LON":"116.37","EPI_DEPTH":18,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"2.0","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u5317\u4eac\u660c\u5e73\u533a","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-10 22:38:44","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CD20191210222800"},{"id":"38452","CATA_ID":"CC20191210164041.00","SAVE_TIME":"2019-12-10 17:11:02","O_TIME":"2019-12-10 16:40:40","EPI_LAT":"-20.82","EPI_LON":"168.65","EPI_DEPTH":30,"AUTO_FLAG":"M","EQ_TYPE":"M","O_TIME_FRA":"0","M":"5.6","M_MS":"0","M_MS7":"0","M_ML":"0","M_MB":"0","M_MB2":"0","SUM_STN":"0","LOC_STN":"0","LOCATION_C":"\u6d1b\u4e9a\u8482\u7fa4\u5c9b\u9644\u8fd1\u6d77\u57df","LOCATION_S":"","CATA_TYPE":"","SYNC_TIME":"2019-12-10 17:15:52","IS_DEL":"","EQ_CATA_TYPE":"","NEW_DID":"CC20191210164041"}
	],
	"jieguo":"\u6700\u8fd148\u5c0f\u65f6\u5730\u9707\u4fe1\u606f",
	"page":"",
"num":1}
*/

type Data struct {
	Id         string  `json:"id"`
	CataId     string  `json:"CATA_ID"`
	SaveTime   string  `json:"SAVE_TIME"`
	OTime      string  `json:"O_TIME"`
	EpiLat     string  `json:"EPI_LAT"`
	EpiLon     string  `json:"EPI_LON"`
	EpiDepth   float32 `json:"EPI_DEPTH"`
	AutoFlag   string  `json:"AUTO_FLAG"`
	OTimeFra   string  `json:"O_TIME_FRA"`
	EqType     string  `json:"EQ_TYPE"`
	M          string  `json:"M"`
	MMS        string  `json:"M_MS"`
	MMS7       string  `json:"M_MS7"`
	MML        string  `json:"M_ML"`
	MMB        string  `json:"M_MB"`
	SumStn     string  `json:"SUM_STN"`
	LocStn     string  `json:"LOC_STN"`
	LocationC  string  `json:"LOCATION_C"`
	LocationS  string  `json:"LOCATION_S,omitempty"`
	CataType   string  `json:"CATA_TYPE,omitempty"`
	SyncTime   string  `json:"SYNC_TIME"`
	IsDel      string  `json:"IS_DEL"`
	EqCataType string  `json:"EQ_CATA_TYPE,omitempty"`
	NewDid     string  `json:"NEW_DID"`
}

type RespBody struct {
	Data   []Data `json:"shuju"`
	Result string `json:"jieguo"`
	Page   string `json:"page,omitempty"`
	Num    int    `json:"num"`
}

func LoadURL(link string) (RespBody, error) {
	var data RespBody

	resp, err := http.Get(link)
	if err != nil {
		return RespBody{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		//替换掉body中的开始和结尾的括号
		tmp := string(body)
		tmp = strings.Replace(tmp, "(", "", -1)
		tmp = strings.Replace(tmp, ")", "", -1)
		err = json.Unmarshal([]byte(tmp), &data)
		if err != nil {
			return data, err
		}
	}

	return data, err
}

func main() {
	data, err := LoadURL("http://www.ceic.ac.cn/ajax/speedsearch?num=2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data.Result)
	CeicDatas := data.Data
	for index := 0; index < len(CeicDatas); index++ {
		tmp := fmt.Sprintf("%s在%s发生里氏%s地震", CeicDatas[index].LocationC, CeicDatas[index].SyncTime, CeicDatas[index].M)
		fmt.Println(tmp)
	}

}

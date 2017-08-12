package config

import (
	"time"
)

const (

	//-------LOGGING-------//
	FILEv     bool = false
	CHANNELv  bool = false
	USERv     bool = false
	COMMENTSv bool = false
	INDEXv    bool = false
	GROUPSv   bool = false
	VIDEOv    bool = false
	TRACEv    bool = true
	STATIONv  bool = false
	REGISTERv bool = true
	BROWSEv   bool = true

	//------ENVIRONMENT--------//

	LOCAL            bool = false //all applications running on a 'nix based machine
	UPLOAD_DISBURSED bool = true //upload is on Beta
	FULL_DISBURSED   bool = false //webserver running locally, everything else running remotely
	ON_A             bool = false //production/test
	ON_B             bool = false //production/test //labs
	ON_BETA          bool = false //beta/server
	ON_LAB2          bool = false
	ON_WWW           bool = false

	//------CONSTANTS-------//

	BETA      string = "beta.cut2it.com"
	LABS      string = "10.0.0.60"
	WWW       string = "www.cut2it.com"
	LAB2      string = "10.0.0.76"
	LOCALHOST string = "localhost"
	MO        string = "localhost" // DO NOT CHANGE
	UP        string = "localhost" //Uploader is running on the BETA server currently )2015-22-8
	LIVEIP    string = "beta.cut2it.com"

	MOPORT         string = ":9393"
	UPPORT         string = ":9001"
	MAILPORT       string = ":9002"
	WEBPORT        string = ":9094"
	LIVEPORT       string = ":9010"
	BIWAY_RPC_PORT string = ":9301"
	MU_RPC_PORT    string = ":9292"

	DOC_CDN   = "https://cdn.cut2it.com/"
	MEDIA_CDN = "https://mmcdn.cut2it.com/"
	LOCAL_CDN = "http://localhost:9010/"

	MaxSleep time.Duration = 300
)

//----INTERNAL IDENTIFIER----//
var (
	MOIP, //Medulla-Oblongata source
	UPIP, //Uploader source
	WEBHOSTURL, //Web Host URL example: http://localhost:9090
	MAILIP, //Mailer Source
	LIVEURL string //Live Streaming Server URL
	ALLOW_ORIGIN_STR string
)

package config

//TODO- add command line flagging for setting the environment
func init() {

	switch {
	case LOCAL:
		MOIP = LOCALHOST + MOPORT
		UPIP = LOCALHOST + UPPORT
		MAILIP = LOCALHOST + MAILPORT
		WEBHOSTURL = "https://" + LOCALHOST + WEBPORT
		LIVEURL = "https://" + LOCALHOST + LIVEPORT
		ALLOW_ORIGIN_STR = "http://localhost:9090"
	case FULL_DISBURSED:
		MOIP = BETA + MOPORT
		UPIP = BETA + UPPORT
		MAILIP = BETA + MAILPORT
		WEBHOSTURL = "https://" + LOCALHOST + WEBPORT
		LIVEURL = "https://" + BETA + LIVEPORT
		ALLOW_ORIGIN_STR = "https://beta.cut2it.com"
	case UPLOAD_DISBURSED:
		MOIP = LOCALHOST + MOPORT
		UPIP = BETA + UPPORT
		MAILIP = LOCALHOST + MAILPORT
		WEBHOSTURL = "https://" + LOCALHOST + WEBPORT
		LIVEURL = "https://" + LOCALHOST + LIVEPORT
		ALLOW_ORIGIN_STR = "https://beta.cut2it.com"
	case ON_A:
		MOIP = MO + MOPORT
		UPIP = UP + UPPORT
		WEBHOSTURL = "https:cut2it.com//"
		LIVEURL = "https://" + LIVEIP + LIVEPORT
		ALLOW_ORIGIN_STR = "https://cut2it.com"
	case ON_WWW:
		MOIP = "10.0.0.60" + MOPORT
		UPIP = "10.0.0.60" + UPPORT
		MAILIP = "10.0.0.60" + MAILPORT
		WEBHOSTURL = "https://www.cut2it.com"
		LIVEURL = "https://" + LIVEIP + LIVEPORT
		ALLOW_ORIGIN_STR = "https://www.cut2it.com"
	case ON_B: //CURRENT CONFIGURATION; THIS IS TEMPORARY 25 8 2015 ::Tim
		MOIP = "10.0.0.60" + MOPORT
		UPIP = "10.0.0.60" + UPPORT
		MAILIP = "10.0.0.60" + MAILPORT
		WEBHOSTURL = "https://labs.cut2it.com"
		LIVEURL = "https://" + LIVEIP + LIVEPORT
		ALLOW_ORIGIN_STR = "https://labs.cut2it.com"
	case ON_LAB2:
		MOIP = "10.0.0.76" + MOPORT
		UPIP = "10.0.0.76" + UPPORT
		MAILIP = "10.0.0.76" + MAILPORT
		WEBHOSTURL = "https://labs2.cut2it.com"
		LIVEURL = "https://" + LIVEIP + LIVEPORT
		ALLOW_ORIGIN_STR = "https://labs2.cut2it.com"
	case ON_BETA:
		MOIP = LOCALHOST + MOPORT
		UPIP = LOCALHOST + UPPORT
		MAILIP = LOCALHOST + MAILPORT
		WEBHOSTURL = "https://" + BETA //No Port
		LIVEURL = "https://" + BETA + LIVEPORT
		ALLOW_ORIGIN_STR = "https://beta.cut2it.com"
	}

	EVILCORP = make(map[string]string)
	EVILCORP[linux] = "/mnt/EvilCorp/"
	EVILCORP[windows] = "C:/mnt/EvilCorp/"

	MEDIACLUSTER = make(map[string]string)
	MEDIACLUSTER[linux] = "/mnt/mediacluster/"
	MEDIACLUSTER[windows] = "C:/mnt/mediacluster/"

	Architecture = &setUP{
		SELF_OS: windows, //CHANGE THIS VARIABLE DEPENING ON LOCAL SYSTEM
	}

	if ON_BETA == true || ON_LAB2 == true || ON_A == true || ON_B == true || ON_WWW == true {
		Architecture.SELF_OS = linux
	}

	EVILCORP_DIR = EVILCORP[Architecture.SELF_OS]
	MEDIACLUSTER_DIR = MEDIACLUSTER[Architecture.SELF_OS]
}

//ADD TO TRANSCODING UPLOADER
//get ping from server
//add to list
//new jobs go to new server
//old server threshold rops
//pings that its droped
//direct new traffic to old server
//stop directing traffic to new servre
//drop new server from serverlist

//this will be needed to take advantage of the scaling
//architecture that is being implemented 23 8 2015

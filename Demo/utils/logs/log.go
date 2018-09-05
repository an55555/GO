package logs

import (
	// "errors"
	"fmt"
	// "io"

	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

//nbabxcbmvzozcadd
func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="/logs/roll.log" maxsize="100000" maxrolls="5"/>
        <filter levels="critical">
            <file path="logs/roll.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="207230373@qq.com" sendername="ShortUrl API" hostname="smtp.qq.com" hostport="587" username="207230373@qq.com" password="odcjjdfarcetbjdd">
                <recipient address="156081289@qq.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
        <format id="critical" format="%File %FullPath %Func %Msg%n" />
        <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library logs output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library logs.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

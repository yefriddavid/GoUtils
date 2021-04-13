package GoUtils


import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)


func CreateAwsSession(awsProfile string, awsRegion string) *session.Session {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(awsRegion),
		},
		Profile: awsProfile,
	})
	fmt.Println(err)

	return sess

}


/*func loadSetting() (config Config, err error) {

	v := viper.New()
	if *configFile == "" {
		v.AddConfigPath("./")
	} else {
		dir, file := path.Split(*configFile)
		ext := path.Ext(file)
		var absoluteFileName string
		if ext == "" {
			absoluteFileName = file
		} else {
			absoluteFileName = strings.TrimRight(file, ext)
		}
		v.AddConfigPath(dir)
		v.SetConfigName(absoluteFileName)

	}
	err = v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	v.Unmarshal(&config)
	return
}*/


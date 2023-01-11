package configuration

import "github.com/sirupsen/logrus"

type Config struct {
    Forum struct {
        Url      string
        UserName string
        UserPass string
        Login    string
    }
    Clients struct {
        Urls []string
    }
    Log struct {
        LogLevel  string
        LogFormat string
    }
    Logger *logrus.Logger
}

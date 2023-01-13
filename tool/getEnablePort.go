package tool

import "net"

func getPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

/** 获取一个可用端口 */
func GetEnablePort() int {
	port, err := getPort()
	if err != nil {
		panic(err)
	}
	return port
}

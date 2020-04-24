package graceful


import (
	"net/http"
	"net"
	"flag"
	"log"
	"os"
	"syscall"
	"context"
	"time"
	"os/signal"
	"errors"
	"os/exec"
)
var (
	SERVER   *http.Server
	LISTENER net.Listener = nil
	GRACEFUL = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)


func StartServer(addr string, handler http.Handler)  {
	var err error
	flag.Parse()
	SERVER = &http.Server{Addr: addr, Handler: handler}

	if *GRACEFUL{
		log.Println("listening on the existing file descriptor 3")
		f := os.NewFile(3, "")
		LISTENER, err = net.FileListener(f)
	}else{
		log.Println("listening on a new file descriptor")
		LISTENER, err = net.Listen("tcp", SERVER.Addr)
	}

	if err != nil{
		log.Fatalf("listener error: %v", err)
	}

	go func() {
		err = SERVER.Serve(LISTENER)
		log.Printf("server.Serve err: %v\n", err)
	}()

	// 监听信号
	handleSignal()

	log.Println("signal end")
}




func handleSignal()  {
	ch := make(chan os.Signal, 1)
	// 监听信号
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal receive: %v\n", sig)
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM: // 终止进程执行
			log.Println("shutdown")
			signal.Stop(ch)
			SERVER.Shutdown(ctx)
			log.Println("graceful shutdown")
			return
		case syscall.SIGUSR2: // 进程热重启
			log.Println("reload")
			err := reload() // 执行热重启函数
			if err != nil {
				log.Fatalf("graceful reload error: %v", err)
			}
			SERVER.Shutdown(ctx)
			log.Println("graceful reload")
			return
		}
	}
}

func reload() error {
	tl, ok := LISTENER.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}
	// 获取 socket 描述符
	f, err := tl.File()
	if err != nil {
		return err
	}
	// 设置传递给子进程的参数（包含 socket 描述符）
	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout         // 标准输出
	cmd.Stderr = os.Stderr         // 错误输出
	cmd.ExtraFiles = []*os.File{f} // 文件描述符
	// 新建并执行子进程
	return cmd.Start()

}
package pprof
// add by stefan 20190606 16:12
import (
	"runtime/pprof"
	"os"
	"path"
	"time"
	"strings"
	"fmt"
	//"log"
	"context"
	"sync"
)

const (
	const_PProfWriteInterval = int32(60*5)
)

type TPProfMgr struct {
	ctx 	context.Context
	wg  	sync.WaitGroup
}

func (self *TPProfMgr) StartPProf(ctx context.Context){
	self.ctx = ctx
	self.wg.Add(1)
	cpu := createCpu()
	mem := createMem()
	go self.loop(cpu, mem)
}

func (self *TPProfMgr) exitPProf(cpu, mem *os.File){
	if cpu != nil {
		pprof.StopCPUProfile()
		cpu.Close()
	}
	if mem != nil {
		pprof.WriteHeapProfile(mem)
		mem.Close()
	}
	self.wg.Wait()
}

func (self *TPProfMgr) loop(cpu, mem *os.File){
	defer self.wg.Done()
	t := time.NewTicker(time.Duration(const_PProfWriteInterval))
	for {
		select {
		case <-self.ctx.Done():
			self.exitPProf(cpu, mem)
		case <-t.C:
			// do nothing...
		}
	}
}

func Newpprof(file string) (retfile string){
	timeformat := time.Now().Format("20120101_120100")
	retfile = timeformat+"_"+file
	execpath, err := os.Executable()
	if err != nil {
		return
	}
	execpath = strings.Replace(execpath, "\\", "/", -1)
	_, sfile := path.Split(execpath)
	arrfile := strings.Split(sfile, ".")
	retfile = fmt.Sprintf("./tmp/%s_%v.prof", arrfile[0], retfile)
	return
}

func createCpu() (file *os.File){
	cpuf := Newpprof("cpu")
	f, err := os.OpenFile(cpuf, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("cpu pprof open fail, err: ", err)
		return
	}
	pprof.StartCPUProfile(f)
	return f
}

func createMem()(file *os.File){
	cpuf := Newpprof("mem")
	f, err := os.OpenFile(cpuf, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("mem pprof open fail, err: ", err)
		return
	}
	return f
}
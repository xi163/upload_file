package uploader

import (
	"time"

	"github.com/xi123/libgo/core/base/cc"
	"github.com/xi123/libgo/core/base/run"
	"github.com/xi123/libgo/core/base/timer"
)

// <summary>
// Args 协程启动参数
// <summary>
type Args struct {
	state    cc.AtomFlag
	stopping cc.Singal
}

func newArgs(proc run.Proc) run.Args {
	s := &Args{
		state:    cc.NewAtomFlag(),
		stopping: cc.NewSingal(),
	}
	return s
}

func (s *Args) SetState(busy bool) {
	if busy {
		s.state.Set()
	} else {
		s.state.Reset()
	}
}

func (s *Args) Busing() bool {
	return s.state.IsSet()
}

func (s *Args) Quit() bool {
	s.stopping.Signal()
	return true
}

func (s *Args) Trigger() <-chan time.Time {
	return nil
}

func (s *Args) TimerCallback() (handler timer.TimerCallback) {
	return
}

func (s *Args) RunAfter(delay int32, args ...any) uint32 {
	return 0
}

func (s *Args) RunAfterWith(delay int32, handler timer.TimerCallback, args ...any) uint32 {
	return 0
}

func (s *Args) RunEvery(delay, interval int32, args ...any) uint32 {
	return 0
}

func (s *Args) RunEveryWith(delay, interval int32, handler timer.TimerCallback, args ...any) uint32 {
	return 0
}

func (s *Args) RemoveTimer(timerID uint32) {
}

func (s *Args) RemoveTimers() {
}

func (s *Args) Duration() time.Duration {
	return 0
}

func (s *Args) Reset(d time.Duration) {

}

func (s *Args) Add(args ...any) {

}

package common

const (
	//任务保存路径
	JOB_SAVE_DIR = "/cron/jobs/"
    //任务杀死通知路径
	JOB_KILL_DIR = "/cron/killer/"
	//保存任务事件
	JOB_EVENT_DELETE = -1
	JOB_EVNET_SAVE = 1
	JOB_ENENT_KILL = 2
	//锁目录
   JOB_LOCK_DIR = "/cron/lock/"
)

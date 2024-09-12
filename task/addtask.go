package task

import (
	"ydsd_gin/tools/news"
)

const (
	reportEvent = "【 *report 】"
	clickEvent  = "【 *click 】"
)

// AddReportTask 加入报告任务.
func AddReportTask(acn string, queryBody []byte) error {
	var err error
	reportType := ""
	taskid := UniqueID(reportType)
	t := NewTask(reportType, queryBody, ID(taskid))
	appname := ""
	// 处理任务
	if _, err = Client().Enqueue(t, TaskOption(reportType)...); err != nil {
		addReportError(appname, err)
		return err
	}
	// logger.Infof(appname+reportEvent+"成功将任务加入队列：%+v -- %+v", taskid, string(info.Payload))
	return nil
}

// AddClickTask 加入点击任务.
func AddClickTask(acn string, queryBody []byte) error {
	var err error
	clickType := ""
	taskid := UniqueID(clickType)
	t := NewTask(clickType, queryBody, ID(taskid))
	appname := ""
	// 处理任务
	if _, err = Client().Enqueue(t, TaskOption(clickType)...); err != nil {
		addClickError(appname, err)
		return err
	}
	// logger.Infof(appname+clickEvent+"成功将任务加入队列：%+v -- %+v", taskid, string(info.Payload))
	return nil
}

// addClickError 记录点击任务错误.
func addClickError(app string, err error) {
	errmsg := app + clickEvent + " 任务加入队列失败: " + err.Error()
	news.ErrRecord(errmsg)
}

// addReportError 记录报告任务错误.
func addReportError(app string, err error) {
	errmsg := app + reportEvent + " 任务加入队列失败:  " + err.Error()
	news.ErrRecord(errmsg)
}

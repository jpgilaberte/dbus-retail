package agent

var(
	WorkerIOInstance *WorkerIO
)

type MsgWorkerIOResponse struct{
	Err             chan error
	ResponseUnit    *Unit
}

type MsgWorkerIORequest struct{
	Res         *MsgWorkerIOResponse
	CurrentUnit *Unit
}

type WorkerIO struct{
	Units      map[string]Unit
	AddUnit    chan *MsgWorkerIORequest
	UpdateUnit chan *MsgWorkerIORequest
	StopUnit   chan *MsgWorkerIORequest
	RemoveUnit chan *MsgWorkerIORequest
	PurgeUnit  chan *MsgWorkerIORequest
	ListUnit   chan *MsgWorkerIORequest
}

func NewWorkerIO() *WorkerIO{
	if WorkerIOInstance == nil{
		WorkerIOInstance = new(WorkerIO)
		WorkerIOInstance.AddUnit = make(chan *MsgWorkerIORequest)
		WorkerIOInstance.UpdateUnit = make(chan *MsgWorkerIORequest)
		WorkerIOInstance.StopUnit = make(chan *MsgWorkerIORequest)
		WorkerIOInstance.RemoveUnit = make(chan *MsgWorkerIORequest)
		WorkerIOInstance.PurgeUnit = make(chan *MsgWorkerIORequest)
		WorkerIOInstance.ListUnit = make(chan *MsgWorkerIORequest)
		go WorkerIOInstance.Run()
	}
	return WorkerIOInstance
}

func (w *WorkerIO) Run(){
	for{
		select {
		case m := <-w.AddUnit:
			m.Res.ResponseUnit = m.CurrentUnit
			m.Res.Err <- addUnit(m)
		case m := <-w.UpdateUnit:
			m.Res.ResponseUnit = m.CurrentUnit
			m.Res.Err <- updateUnit(m)
		case m := <-w.StopUnit:
			m.Res.ResponseUnit = m.CurrentUnit
			m.Res.Err <- stopUnit(m)
		case m := <-w.RemoveUnit:
			m.Res.ResponseUnit = m.CurrentUnit
			m.Res.Err <- removeUnit(m)
		case m := <-w.PurgeUnit:
			m.Res.ResponseUnit = m.CurrentUnit
			m.Res.Err <- purgeUnit(m)
		}
	}
}

func addUnit(mreq *MsgWorkerIORequest)(error){
	fs := []func()(err error){mreq.CurrentUnit.CreateUnitFile, mreq.CurrentUnit.CreateLinkUnitFile, mreq.CurrentUnit.startUnit}
	return pipelineUnitFunctions(mreq, fs, statusStarting, statusRunning, statusError)
}


func updateUnit(mreq *MsgWorkerIORequest)(error){
	// create file v2
	// create link v2
	// start process
	// set status
	/*
	mres.ResponseUnit = mreq.CurrentUnit
	fs := []func()(err error){mreq.CurrentUnit.CreateUnitFile, mreq.CurrentUnit.CreateLinkUnitFile, mreq.CurrentUnit.startUnit}
	mres.Err = pipelineUnitFunctions(mreq, fs, statusStarting, statusRunning, statusError)
	*/
	return nil
}

func stopUnit(mreq *MsgWorkerIORequest)(error){
	//stop u
	// set status
	/*
	mres.ResponseUnit = mreq.CurrentUnit
	fs := []func()(err error){mreq.CurrentUnit.CreateUnitFile, mreq.CurrentUnit.CreateLinkUnitFile, mreq.CurrentUnit.startUnit}
	mres.Err = pipelineUnitFunctions(mreq, fs, statusStarting, statusRunning, statusError)
	*/
	return nil
}

func removeUnit(mreq *MsgWorkerIORequest)(error){
	//stop u vn
	//remove u vn
	/*
	// set status
	mres.ResponseUnit = mreq.CurrentUnit
	fs := []func()(err error){mreq.CurrentUnit.CreateUnitFile, mreq.CurrentUnit.CreateLinkUnitFile, mreq.CurrentUnit.startUnit}
	mres.Err = pipelineUnitFunctions(mreq, fs, statusStarting, statusRunning, statusError)
	*/
	return nil
}

func purgeUnit(mreq *MsgWorkerIORequest)(error){
	//stopAll u
	//removeAll u
	// set status
	/*
	mres.ResponseUnit = mreq.CurrentUnit
	fs := []func()(err error){mreq.CurrentUnit.CreateUnitFile, mreq.CurrentUnit.CreateLinkUnitFile, mreq.CurrentUnit.startUnit}
	mres.Err = pipelineUnitFunctions(mreq, fs, statusStarting, statusRunning, statusError)
	*/
	return nil
}

func pipelineUnitFunctions (m *MsgWorkerIORequest, fs []func()(err error), initialState Status, finishState Status, errorState Status)(err error){
	m.CurrentUnit.Status = initialState
	for _, f := range fs{
		if err = f(); err != nil{
			m.CurrentUnit.Status = errorState
			return
		}
	}
	m.CurrentUnit.Status = finishState
	return
}

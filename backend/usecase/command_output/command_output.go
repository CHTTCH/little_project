package command_output

type CommandOutput struct {
	Id      string
	Result  bool
	Message string
}

func (i *CommandOutput) SetId(id string) {
	i.Id = id
}

func (i *CommandOutput) GetId() string {
	return i.Id
}

func (i *CommandOutput) SetResult(result bool) {
	i.Result = result
}

func (i *CommandOutput) GetResult() bool {
	return i.Result
}

func (i *CommandOutput) SetMessage(message string) {
	i.Message = message
}

func (i *CommandOutput) GetMessage() string {
	return i.Message
}

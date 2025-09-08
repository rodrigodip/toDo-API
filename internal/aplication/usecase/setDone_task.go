package usecase

func (ct *TaskRepository) SetTaskDone(id string) error {
	err := ct.Repository.SetTaskDone(id)
	if err != nil {
		return err
	}
	return nil
}

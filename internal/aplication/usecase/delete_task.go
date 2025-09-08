package usecase

func (ct *TaskRepository) DeleteTask(id string) error {
	err := ct.Repository.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

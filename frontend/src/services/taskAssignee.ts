import AbstractService from './abstractService'
import TaskAssigneeModel from '@/models/taskAssignee'
import type {ITaskAssignee} from '@/modelTypes/ITaskAssignee'

export default class TaskAssigneeService extends AbstractService<ITaskAssignee> {
	constructor() {
		super({
			getAll: '/tasks/{id}/assignees',
			create: '/tasks/{taskId}/assignees',
			delete: '/tasks/{taskId}/assignees/{userId}',
		})
	}

	modelFactory(data: Partial<ITaskAssignee>): TaskAssigneeModel {
		return new TaskAssigneeModel(data)
	}
}
import AbstractService from './abstractService'
import TaskModel from '@/models/task'
import type {ITask} from '@/modelTypes/ITask'
import TaskAssigneeService from './taskAssignee'

export default class TaskService extends AbstractService<ITask> {
	constructor() {
		super({
			getAll: '/projects/{projectId}/tasks',
			get: '/tasks/{id}',
			create: '/projects/{projectId}/tasks',
			update: '/tasks/{id}',
			delete: '/tasks/{id}',
		})
	}

	modelFactory(data: Partial<ITask>): TaskModel {
		return new TaskModel(data)
	}

	beforeUpdate(model: ITask) {
		// Fix repeatAfter
		if (
			typeof model.repeatAfter === 'object' &&
			model.repeatAfter !== null &&
			'amount' in model.repeatAfter
		) {
			model.repeatAfter = model.repeatAfter.amount
		}
	
	
		return {
			id: model.id,
			title: model.title,
			description: model.description,
			done: model.done,
			due_date: model.dueDate,
			start_date: model.startDate,
			end_date: model.endDate,
			priority: model.priority,
			percent_done: model.percentDone,
			hex_color: model.hexColor,
			estimation: model.estimation,
			repeat_after: model.repeatAfter,
			repeat_mode: model.repeatMode,
			project_id: model.projectId,
			is_favorite: model.isFavorite,
			cover_image_attachment_id: model.coverImageAttachmentId,
		} as ITask
	}

	beforeCreate(task: ITask) {
		task.done = false
		return this.beforeUpdate(task)
	}

	async update(model: ITask) {
		const updatedModel = await super.update(model)

		try {
			const assigneeService = new TaskAssigneeService()
			const assignees = await assigneeService.getAll({id: model.id})
			updatedModel.assignees = assignees || []
		} catch (e) {
			// If no assignees or error, use empty array or existing
			updatedModel.assignees = model.assignees || []
		}

		return updatedModel
	}

	async create(task: ITask) {
		const createdTask = await super.create(task)

		try {
			const assigneeService = new TaskAssigneeService()
			const assignees = await assigneeService.getAll({id: createdTask.id})
			createdTask.assignees = assignees || []
		} catch (e) {
			// If no assignees or error, use empty array
			createdTask.assignees = []
		}

		return createdTask
	}
}
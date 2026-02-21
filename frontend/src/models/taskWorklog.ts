import AbstractModel from './abstractModel'
import UserModel from './user'

import type {ITaskWorklog} from '@/modelTypes/ITaskWorklog'
import type {IUser} from '@/modelTypes/IUser'

export default class TaskWorklogModel extends AbstractModel<ITaskWorklog> implements ITaskWorklog {
	id = 0
	taskId = 0
	duration  = 0
	description = ''
	logDate = ''
	user: IUser | null = null

	created: Date = new Date()
	updated: Date = new Date()

	constructor(data: Partial<ITaskWorklog> = {}) {
		super()
		this.assignData(data)

		this.created = new Date(this.created)
		this.updated = new Date(this.updated)
		
		if (this.user !== null) {
			this.user = new UserModel(this.user)
		}
	}
}
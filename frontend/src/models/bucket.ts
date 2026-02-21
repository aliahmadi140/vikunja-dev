import AbstractModel from './abstractModel'
import UserModel from './user'
import TaskModel from '@/models/task'

import type {IBucket} from '@/modelTypes/IBucket'
import type {IUser} from '@/modelTypes/IUser'

export default class BucketModel extends AbstractModel<IBucket> implements IBucket {
	id = 0
	title = ''
	projectId = ''
	limit = 0
	tasks = []
	position = 0
	count = 0

	createdBy: IUser = null
	created: Date = null
	updated: Date = null

	constructor(data: Partial<IBucket>) {
		super()
		this.assignData(data)

		this.tasks = (data.tasks || []).map(t => new TaskModel(t))

		this.createdBy = this.createdBy ? new UserModel(this.createdBy) : null
		this.created = this.created ? new Date(this.created) : null
		this.updated = this.updated ? new Date(this.updated) : null
	}
}
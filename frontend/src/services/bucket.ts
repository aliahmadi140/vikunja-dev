import AbstractService from './abstractService'
import BucketModel from '../models/bucket'
import TaskModel from '@/models/task'
import type { IBucket } from '@/modelTypes/IBucket'

export default class BucketService extends AbstractService<IBucket> {
	constructor() {
		super({
			getAll: '/projects/{projectId}/views/{projectViewId}/buckets',
			create: '/projects/{projectId}/views/{projectViewId}/buckets',
			update: '/projects/{projectId}/views/{projectViewId}/buckets/{id}',
			delete: '/projects/{projectId}/views/{projectViewId}/buckets/{id}',
		})
	}

	modelFactory(data: Partial<IBucket>) {
		return new BucketModel(data)
	}

	beforeUpdate(model) {
		model.tasks = model.tasks?.map(t => new TaskModel(t))
		return model
	}
}
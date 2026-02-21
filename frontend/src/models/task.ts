import {PRIORITIES, type Priority} from '@/constants/priorities'

import type {ITask} from '@/modelTypes/ITask'
import type {ILabel} from '@/modelTypes/ILabel'
import type {IUser} from '@/modelTypes/IUser'
import type {IAttachment} from '@/modelTypes/IAttachment'
import type {IProject} from '@/modelTypes/IProject'
import type {ISubscription} from '@/modelTypes/ISubscription'
import type {IBucket} from '@/modelTypes/IBucket'
import type {ITaskWorklog} from '@/modelTypes/ITaskWorklog'

import type {IRepeatAfter} from '@/types/IRepeatAfter'
import type {IRelationKind} from '@/types/IRelationKind'
import {TASK_REPEAT_MODES, type IRepeatMode} from '@/types/IRepeatMode'

import {parseDateOrNull} from '@/helpers/parseDateOrNull'
import {secondsToPeriod} from '@/helpers/time/period'

import AbstractModel from './abstractModel'
import LabelModel from './label'
import UserModel from './user'
import AttachmentModel from './attachment'
import SubscriptionModel from './subscription'
import type {ITaskReminder} from '@/modelTypes/ITaskReminder'
import TaskReminderModel from '@/models/taskReminder'
import TaskCommentModel from '@/models/taskComment'
import TaskWorklogModel from '@/models/taskWorklog'

export function getHexColor(hexColor: string): string | undefined {
	if (hexColor === '' || hexColor === '#') {
		return undefined
	}

	return hexColor
}

export function parseRepeatAfter(repeatAfterSeconds: number): IRepeatAfter {
	const period = secondsToPeriod(repeatAfterSeconds)

	return {
		type: period.unit,
		amount: period.amount,
	}
}

export function getTaskIdentifier(task: ITask | null | undefined): string {
	if (task === null || typeof task === 'undefined') {
		return ''
	}

	if (task.identifier === '') {
		return `#${task.index}`
	}

	return task.identifier
}

export default class TaskModel extends AbstractModel<ITask> implements ITask {
	id = 0
	title = ''
	description = ''
	done = false
	doneAt: Date | null = null
	priority: Priority = PRIORITIES.UNSET
	labels: ILabel[] = []
	assignees: IUser[] = []

	dueDate: Date | null = null
	startDate: Date | null = null
	endDate: Date | null = null

	estimation: number | null = null

	repeatAfter: number | IRepeatAfter = 0
	repeatFromCurrentDate = false
	repeatMode: IRepeatMode = TASK_REPEAT_MODES.REPEAT_MODE_DEFAULT
	reminders: ITaskReminder[] = []
	parentTaskId: ITask['id'] = 0
	hexColor = ''
	percentDone = 0
	relatedTasks: Partial<Record<IRelationKind, ITask[]>> = {}
	attachments: IAttachment[] = []
	coverImageAttachmentId: IAttachment['id'] = null
	identifier = ''
	index = 0
	isFavorite = false
	subscription: ISubscription = null

	position = 0

	reactions = {}
	comments = []

	worklogs: ITaskWorklog[] = []

	createdBy: IUser = null
	created: Date = null
	updated: Date = null

	projectId: IProject['id'] = 0
	bucketId: IBucket['id'] = 0

	constructor(data: Partial<ITask> = {}) {
		super()
		
		this.assignData(data)
	

		this.done = Boolean(this.done)
this.priority = Number.isFinite(this.priority)
	? this.priority
	: PRIORITIES.UNSET

		this.id = Number(this.id)
		this.title = this.title?.trim()
		this.doneAt = parseDateOrNull(this.doneAt)

	

		this.labels = this.labels
			.map(l => new LabelModel(l))
			.sort((a, b) => a.title.localeCompare(b.title))

		this.assignees = this.assignees.map(a => new UserModel(a))

		this.dueDate = parseDateOrNull(this.dueDate)
		this.startDate = parseDateOrNull(this.startDate)
		this.endDate = parseDateOrNull(this.endDate)

		this.estimation = typeof this.estimation === 'number'
			? this.estimation
			: null

		this.repeatAfter = parseRepeatAfter(this.repeatAfter as number)

		this.reminders = this.reminders.map(r => new TaskReminderModel(r))

		if (this.hexColor !== '' && this.hexColor.substring(0, 1) !== '#') {
			this.hexColor = '#' + this.hexColor
		}

		Object.keys(this.relatedTasks).forEach(relationKind => {
			this.relatedTasks[relationKind] = this.relatedTasks[relationKind].map(
				t => new TaskModel(t),
			)
		})

		this.attachments = this.attachments.map(a => new AttachmentModel(a))

		if (this.identifier === `-${this.index}`) {
			this.identifier = ''
		}

		if (this.subscription !== null && typeof this.subscription !== 'undefined') {
			this.subscription = new SubscriptionModel(this.subscription)
		}

		this.createdBy = this.createdBy
	? new UserModel(this.createdBy)
	: null
		this.created = new Date(this.created)
		this.updated = new Date(this.updated)

		this.projectId = Number(this.projectId)

		this.comments = (data.comments || []).map(c => new TaskCommentModel(c))

		this.reactions = {}
		Object.keys(data.reactions || {}).forEach(reaction => {
			this.reactions[reaction] =
				data.reactions[reaction].map(u => new UserModel(u))
		})

		this.worklogs = (data.worklogs || []).map(
			w => new TaskWorklogModel(w),
		)
	}

	getTextIdentifier() {
		return getTaskIdentifier(this)
	}

	getHexColor() {
		return getHexColor(this.hexColor)
	}
}
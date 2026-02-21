import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import AbstractService from './abstractService'
import TaskWorklogModel from '@/models/taskWorklog'
import type {ITaskWorklog} from '@/modelTypes/ITaskWorklog'

dayjs.extend(utc)

type LogDateLike = string | Date | null | undefined

type WorklogPayload = {
    duration: number
    description?: string | null
    log_date?: string
}

const toUtcIsoString = (value: LogDateLike): string | undefined => {
    if (!value) {
        return undefined
    }

    const parsed = dayjs(value)
    if (!parsed.isValid()) {
        return undefined
    }

    return parsed.utc().format('YYYY-MM-DDTHH:mm:ss[Z]')
}

const buildPayload = (_taskId: number, worklog: ITaskWorklog): WorklogPayload => {
    const logDate = worklog.logDate ?? new Date()

    return {
        duration: worklog.duration,
        description: worklog.description,
        log_date: new Date(logDate).toISOString(),
    }
}

export default class TaskWorklogService extends AbstractService<ITaskWorklog> {
    constructor() {
        super({
            getAll: '/tasks/{taskId}/worklogs',
            create: '/tasks/{taskId}/worklogs',
            update: '/tasks/{taskId}/worklogs/{id}',
            delete: '/tasks/{taskId}/worklogs/{id}',
        })
    }

    modelFactory(data: Partial<ITaskWorklog>): TaskWorklogModel {
        return new TaskWorklogModel(data)
    }

    async getAllByTask(taskId: number): Promise<TaskWorklogModel[]> {
        const response = await this.http({
            url: `/tasks/${taskId}/worklogs`,
            method: 'GET',
        })
        return response.data.map((worklog: ITaskWorklog) => this.modelFactory(worklog))
    }

    async createForTask(taskId: number, worklog: ITaskWorklog): Promise<TaskWorklogModel> {
        const payload = buildPayload(taskId, worklog)
        const response = await this.http({
            url: `/tasks/${taskId}/worklogs`,
            method: 'POST',
            data: payload,
        })
        return this.modelFactory(response.data)
    }

    async updateForTask(taskId: number, worklog: ITaskWorklog): Promise<TaskWorklogModel> {
        const payload = buildPayload(taskId, worklog)
        const response = await this.http({
            url: `/tasks/${taskId}/worklogs/${worklog.id}`,
            method: 'PUT',
            data: payload,
        })
        return this.modelFactory(response.data)
    }

    async deleteForTask(taskId: number, worklogId: number): Promise<void> {
        await this.http({
            url: `/tasks/${taskId}/worklogs/${worklogId}`,
            method: 'DELETE',
        })
    }
}
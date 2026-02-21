import type {IAbstract} from './IAbstract'
import type {IUser} from './IUser'

export interface ITaskWorklog extends IAbstract {
	id: number
	taskId: number
	duration: number
	description: string
	logDate: string
	user?: IUser
	
	created: Date
	updated: Date
}
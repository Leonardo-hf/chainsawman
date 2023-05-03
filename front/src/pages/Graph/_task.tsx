export const TaskTypeMap = {
    0: {
        text: '未完成',
        color: 'grey',
    },
    1: {
        text: '完成',
        color: 'green',
    }
}

export enum TaskType {
    unfinished,
    finished,
}

export function getTaskTypeDesc(status: TaskType) {
    return TaskTypeMap[status]
}
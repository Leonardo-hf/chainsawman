export const getTag = (graph: number, node: number) => graph.toString() + '@' + node.toString()

export const formatDate = (time: number) => {
    const t = new Date(time * 1000);
    const tf = (i: number) => (i < 10 ? '0' : '') + i
    const year = t.getUTCFullYear()
    const thisYear = new Date().getUTCFullYear()
    if (year == thisYear) {
        return (t.getUTCMonth() + 1) + '月' + t.getUTCDay() + '日 ' +
            tf(t.getUTCHours()) + ':' + tf(t.getUTCMinutes())
    }
    return year + '年' + (t.getUTCMonth() + 1) + '月' + t.getUTCDay() + '日 '
}

export const formatNumber = (v: any) => Math.floor(v).toString()

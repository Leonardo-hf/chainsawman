
export const formatDate = (time: number) => {
    const t = new Date(time);
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

export const UUID = () =>
    'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0,
            v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });

export const getRandomColor = function () {
    const text = '00000' + (Math.random() * 0x1000000 << 0).toString(16)
    return '#' + text.substring(text.length - 6);
}
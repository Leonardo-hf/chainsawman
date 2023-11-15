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

function hslToRgb(h: number, s: number, l: number) {
    l /= 100
    const a = s * Math.min(l, 1 - l) / 100;
    const f = (n: number) => {
        const k = (n + h / 30) % 12;
        const color = l - a * Math.max(Math.min(k - 3, 9 - k, 1), -1);
        return Math.round(255 * color).toString(16).padStart(2, '0');   // convert to Hex and prefix "0" if needed
    };
    return `#${f(0)}${f(8)}${f(4)}`
}

export const getNRandomColor = function (n: number) {
    const step = 360 / n
    let i = 0
    const colors = []
    while (i < 360) {
        const h = i
        const s = 90 + Math.random() * 10
        const l = 50 + Math.random() * 10
        i += step
        colors.push(hslToRgb(h, s, l))
    }
    return colors
}

const toTimestamp = t => {
    return parseInt(new Date(t).getTime() / 1000)
}
const now = () => {
    return parseInt(new Date().getTime() / 1000)
}
const relativeTime = t => {
    let timestamp = toTimestamp(t)
    let n = now()
    let diff = n - timestamp
    let minute = 60;
    let hour = minute * 60;
    let day = hour * 24;
    let month = day * 30;

    let monthC = diff / month;
    let dayC = diff / day;
    let hourC = diff / hour;
    let minC = diff / minute;

    let str = "错误"
    if (monthC > 12) {
        str = parseInt(monthC / 12) + " 年前";
    } else if (monthC >= 1) {
        str = parseInt(monthC) + " 月前";
    } else if (dayC >= 1) {
        str = parseInt(dayC) + " 天前";
    } else if (hourC >= 1) {
        str =  parseInt(hourC) + " 小时前";
    } else if (minC >= 1) {
        str = parseInt(minC) + " 分钟前";
    }else {
        str = '刚刚';
    }
    return str
}


// 打开新窗口
function openWin (url) {
    uni.navigateTo({url: url})
}


module.exports = {
    toTimestamp,
    relativeTime,
    openWin,
}

const prevEvents = function () {
    if (p <= 0) return
    p--
    fetchEvents()
}

const nextEvents = function () {
    p++
    fetchEvents()
}

const searchWipe = function () {
    document.querySelector('div.search>input').value = ''
    search()
}

const search = function () {
    const filter = document.querySelector('div.search>input').value
    if (filter.length < 3 && filter.length > 0) {
        return
    }
    p = 0
    fetchEvents()
}

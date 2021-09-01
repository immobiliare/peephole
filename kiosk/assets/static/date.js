const isISODate = d => new Date(d) !== 'Invalid Date' && !isNaN(new Date(d)) && d === new Date(d).toISOString()

const dateFormat = d => `${d.getFullYear()}-${span(d.getMonth() + 1)}-${span(d.getDate())} ${span(d.getHours())}:${span(d.getMinutes())}:${span(d.getSeconds())}`

const span = i => `${i < 10 ? '0' : ''}${i}`
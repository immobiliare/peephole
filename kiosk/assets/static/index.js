const isISODate = date => new Date(date) !== 'Invalid Date' && !isNaN(new Date(date)) && date === new Date(date).toISOString()

const syntaxHighlight = function (json) {
  json = JSON.stringify(json, null, 2)
  json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
  const result = json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+-]?\d+)?)/g, function (match) {
    let cls = ''
    if (/^"/.test(match)) {
      if (/:$/.test(match)) {
        cls = 'key'
        match = match.substring(1, match.length - 2)
      } else {
        if (/"[\d]+"/.test(match)) {
          cls = 'number'
        } else {
          cls = 'string'
        }
        match = match.substring(1, match.length - 1)
        if (isISODate(match)) {
          cls = 'date'
        }
      }
    } else if (/true|false/.test(match)) {
      cls = 'boolean'
    } else if (/null/.test(match)) {
      cls = 'null'
    }
    if (cls === 'key') {
      return '<span class="' + cls + '">' + match + '</span>:'
    } else if (cls === 'string') {
      return '"<span class="' + cls + '">' + match + '</span>"'
    } else {
      return '<span class="' + cls + '">' + match + '</span>'
    }
  })
  return result
}

const dialog = function (jid) {
  if (jid === null) {
    return
  }

  fetch(`/events/${jid}`)
    .then((response) => {
      return response.json()
    }).then((json) => {
      const body = document.getElementsByTagName('body')[0]
      body.innerHTML = `
        <div class="event-dialog">
          <a class="dismiss" onclick="dismiss()">close</a>
          <pre>` + syntaxHighlight(JSON.parse(json.RawData)) + `</pre>
        </div>` + body.innerHTML
    })
}

const dismiss = function () {
  const dialog = document.getElementsByClassName('event-dialog')
  if (dialog.length > 0) {
    dialog[0].remove()
  }
}

const setLoader = function (value) {
  const loader = document.querySelector('.loading')
  if (value == null) {
    loader.classList.add('gone')
  } else {
    loader.classList.remove('gone')
    loader.innerHTML = value
  }
}

const addEvent = function (e) {
  setLoader(null)
  e.Timestamp = new Date(e.Timestamp)
  const events = document.getElementsByTagName('ul')[0]
  const caption = new Date(e.Timestamp).toLocaleString().replace(',', ' at ') +
    ' <i>—</i> <pre>' + e.Minion + '</pre>'
  events.innerHTML = `
    <li>
      <div class="event">
          <span class="caption">${caption}</span>
          <span class="function">${e.Function}</span>
          <a class="show" onclick="dialog('${e.Jid}')">show</a>
      </div>
    </li>` + events.innerHTML

  if (events.children.length > 15) {
    events.children[events.children.length - 1].remove()
  }
}

const fetchEvents = function (filter) {
  let path = '/events'
  if (filter !== null && filter !== '') {
    path += '?q=' + filter
  }

  fetch(path)
    .then((response) => {
      return response.json()
    }).then((json) => {
      if (json.length > 0) {
        for (let i = 0; i < json.length; i++) {
          addEvent(json[i])
        }
      } else {
        setLoader('no result found')
      }
    })
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

  document.querySelector('.event-list').innerHTML = ''
  setLoader('...')
  fetchEvents(filter)
}

let es = null

window.onload = async function () {
  fetchEvents(null)
  es = new EventSource('/stream')
  es.onerror = function () {
    document.querySelector('span.liveness').classList.toggle('dead')
  }
  es.addEventListener('event', function (e) {
    addEvent(JSON.parse(e.data))
  }, false)
}

window.onbeforeunload = async function () {
  if (es != null) {
    es.close()
  }
}

document.onkeydown = async function (e) {
  switch (e.which) {
    case 27:
      dismiss()
      break
    case 191:
      // placeholder to make a call
      // to dialog/searchWipe function
      // (for linting purposes)
      dialog(null)
      searchWipe()
      break
    default:
      break
  }
}

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

const eventCaption = function (e) {
  return `${new Date(e.Timestamp).toLocaleString().replace(',', ' at ')} <i>—</i> <pre>${e.Minion}</pre>`
}

const eventFunction = function (e) {
  const fnParts = e.Function.split(' (')
  if (fnParts.length === 1) {
    return e.Function
  }

  return `${fnParts[0]} <pre>(${fnParts[1]}</pre>`
}

const toggleButton = function (status, selector, callback) {
  const el = document.querySelector(selector)
  if (status) { // enable
    el.classList.remove('disabled')
    el.disabled = false
  } else { // disable
    document.querySelector(selector).classList.add('disabled')
    el.disabled = true
  }
}

const addEvent = function (e) {
  setLoader(null)
  const events = document.getElementsByTagName('ul')[0]
  events.innerHTML = `
    <li>
    <div class="event">
      <span class="caption">${eventCaption(e)}</span>
      <span class="function">${eventFunction(e)}</span>
      <a class="show" onclick="dialog('${e.Jid}')">show</a>
    </div>
    </li>` + events.innerHTML

  if (events.children.length > 7) {
    events.children[events.children.length - 1].remove()
  }
}

const fetchEvents = function () {
  let path = '/events?'
  const q = document.querySelector('div.search>input').value
  if (q !== null && q !== '') {
    path += '&q=' + q
  }
  if (p !== null && p >= 0) {
    path += '&p=' + p
  }

  document.querySelector('.event-list').innerHTML = ''
  document.querySelector('div.pager>span').innerText = p + 1
  setLoader('...')

  fetch(path)
    .then((response) => {
      return response.json()
    }).then((json) => {
      if (json.events.length > 0) {
        for (let i = 0; i < json.events.length; i++) {
          addEvent(json.events[i])
        }
      } else {
        setLoader('no result found')
      }

      if (!json.has_next) {
        toggleButton(false, 'button.right')
      } else {
        toggleButton(true, 'button.right')
      }

      if (p === 0) {
        toggleButton(false, 'button.left')
      } else {
        toggleButton(true, 'button.left')
      }
    })
    .catch(function () {
      setLoader('unable to query')
    })
}

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

let es = null
let p = 0

window.onload = async function () {
  fetchEvents()
  es = new EventSource('/stream')
  es.onerror = function () {
    document.querySelector('span.liveness').classList.toggle('dead')
  }
  es.addEventListener('event', function (e) {
    if (p === 0) {
      addEvent(JSON.parse(e.data))
    }
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
      nextEvents()
      prevEvents()
      break
    default:
      break
  }
}

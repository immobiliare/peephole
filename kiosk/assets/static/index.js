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
  if (value == null) {
    const sheet = window.document.styleSheets[0]
    sheet.insertRule('.loading { display: none }', sheet.cssRules.length)
  } else {
    document.getElementsByClassName('loading')[0].innerHTML = value
  }
}

const addEvent = function (e) {
  setLoader(null)
  document.getElementsByClassName('event-list')[0].innerHTML = `
    <li>
      <div class="event">
          <span class="minion">${e.Minion}</span>
          <span class="function">${e.Function}</span>
          <a class="show" onclick="dialog('${e.Jid}')">show</a>
      </div>
    </li>` + document.getElementsByClassName('event-list')[0].innerHTML
}

let es = null

window.onload = async function () {
  fetch('/events')
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
      // to dialog function
      // (for linting purposes)
      dialog(null)
      break
    default:
      break
  }
}

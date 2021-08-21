var dialog = function (jid) {
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
            <pre>` + JSON.stringify(JSON.parse(json.RawData), null, 2) + `</pre>
        </div> ` + body.innerHTML
    })
}

var dismiss = function () {
  const dialog = document.getElementsByClassName('event-dialog')
  if (dialog.length > 0) {
    dialog[0].remove()
  }
}

var setLoader = function (value) {
  if (value == null) {
    const sheet = window.document.styleSheets[0]
    sheet.insertRule('.loading { display: none }', sheet.cssRules.length)
  } else {
    document.getElementsByClassName('loading')[0].innerHTML = value
  }
}

var addEvent = function (e) {
  setLoader(null)
  document.getElementsByClassName('event-list')[0].innerHTML = `
        <li>
        <div class="event">
            <span class="minion">${e.Minion}</span>
            <span class="function">${e.Function}</span>
            <a class="show" onclick="dialog('${e.Jid}')">show</a>
        </div>
    </li> ` + document.getElementsByClassName('event-list')[0].innerHTML
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
  es.onopen = function () {
    console.log('stream opened')
  }
  es.onerror = function () {
    console.log('stream closed')
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

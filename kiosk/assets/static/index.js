const setLoader = function (value) {
  const loader = document.querySelector('.loading')
  if (value == null) {
    loader.classList.add('gone')
  } else {
    loader.classList.remove('gone')
    loader.innerHTML = value
  }
}

const toggleButton = function (status, selector) {
  const el = document.querySelector(selector)
  if (status) { // enable
    el.classList.remove('disabled')
    el.disabled = false
  } else { // disable
    document.querySelector(selector).classList.add('disabled')
    el.disabled = true
  }
}

const fetchEvents = function () {
  const q = document.querySelector('div.search>input').value
  document.querySelector('.event-list').innerHTML = ''
  setLoader('...')
  apiEvents(q, p)
    .then((json) => {
      document.querySelector('div.pager>span').innerText = p + 1
      if (json.Events.length > 0) {
        for (let i = 0; i < json.Events.length; i++) {
          eventAdd(json.Events[i])
        }
      } else {
        setLoader('no result found')
      }

      if (!json.HasNext) {
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
    .catch(function (err) {
      console.log(err)
      setLoader('unable to query')
    })
}

let es = null
let p = 0

window.onload = async function () {
  fetchEvents()
  es = new EventSource('/stream')
  es.onerror = function () {
    document.querySelector('div.liveness').classList.toggle('dead')
  }
  es.addEventListener('event', function (e) {
    const q = document.querySelector('div.search>input').value
    if (p === 0 && q.length === 0) {
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
      dialogDismiss()
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

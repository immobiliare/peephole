const apiEvent = async function (jid) {
  if (jid === null) return

  return fetch(`/events/${jid}`)
    .then(r => r.json())
}

const apiEvents = async function (query, page) {
  let path = '/events?'
  if (query !== null && query !== '') path += '&q=' + query
  if (page !== null && page >= 0) path += '&p=' + page

  return fetch(path)
    .then(r => r.json())
}


const eventCaption = function (e) {
  return `${dateFormat(new Date(e.Timestamp))} <i class="spacer">—</i> <pre>${e.Minion}</pre>`
}

const eventFunction = function (e) {
  const prefix = `<span class="${e.Success ? 'success">&#10003;' : 'failure">&#10799;'}</span> <i class="spacer">-</i> ${e.Function}`
  if (e.Args !== null && e.Args.length >= 0) {
    return `${prefix} <i class="spacer">-</i> <pre>${e.Args.join(', ').replace(/(_)?orch(estrate)?\./, '')}</pre>`
  }
  return `${prefix}`
}

const eventAdd = function (e) {
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
    toggleButton(true, 'button.right')
  }
}

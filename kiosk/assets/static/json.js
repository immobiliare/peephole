const jsonHighlight = function (json) {
  return JSON.stringify(json, null, 2)
    .replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
    .replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+-]?\d+)?)/g, function (match) {
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
      } else if (/true/.test(match)) {
        cls = 'boolean success'
      } else if (/false/.test(match)) {
        cls = 'boolean failure'
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
}

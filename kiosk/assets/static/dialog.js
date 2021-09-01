const dialog = function (jid) {
  apiEvent(jid)
    .then((json) => {
      dialogShow(jsonHighlight(json))
    })
    .catch(function (err) {
      console.log(err)
      dialogShow("unable to query")
    })
}

const dialogShow = function (content) {
  document.querySelector('.event-dialog>pre').innerHTML = content
  document.querySelector('.event-dialog').classList.add('show')
}

const dialogDismiss = function () {
  document.querySelector('.event-dialog').classList.remove('show')
}

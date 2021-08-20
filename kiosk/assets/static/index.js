var es = null;

function dialog(jid) {
    fetch(`/events/${jid}`)
        .then((response) => {
            return response.json();
        }).then((json) => {
            let body = document.getElementsByTagName('body')[0];
            body.innerHTML = `
            <div class="event-dialog">
                <a class="dismiss" onclick="dismiss()">close</a>
                <pre>` + JSON.stringify(JSON.parse(json.RawData), null, 2) + `</pre>
            </div> ` + body.innerHTML;
        });
}

function dismiss() {
    let dialog = document.getElementsByClassName('event-dialog');
    if (dialog.length > 0) {
        dialog[0].remove();
    }
};

function setLoader(value) {
    if (value == null) {
        var sheet = window.document.styleSheets[0];
        sheet.insertRule('.loading { display: none; }', sheet.cssRules.length);
    } else {
        document.getElementsByClassName('loading')[0].innerHTML = value;
    }
}

function addEvent(e) {
    setLoader(null);
    document.getElementsByClassName('event-list')[0].innerHTML = `
        <li>
        <div class="event">
            <span class="minion">${e.Minion}</span>
            <span class="function">${e.Function}</span>
            <a class="show" onclick="dialog('${e.Jid}')">show</a>
        </div>
    </li> ` + document.getElementsByClassName('event-list')[0].innerHTML;
}

window.onload = async function () {
    fetch('/events')
        .then((response) => {
            return response.json();
        }).then((json) => {
            if (json.length > 0) {
                for (var i = 0; i < json.length; i++) {
                    addEvent(json[i]);
                }
            } else {
                setLoader("no result found");
            }
        });
    es = new EventSource("/stream");
    es.onopen = function () {
        console.log('stream opened');
    }
    es.onerror = function () {
        console.log('stream closed');
    }
    es.addEventListener('event', function (e) {
        addEvent(JSON.parse(e.data))
    }, false);
}

window.onbeforeunload = async function () {
    if (es != null) {
        es.close();
    }
}

document.onkeydown = async function (e) {
    if (e.which == 27) {
        dismiss();
    }
}
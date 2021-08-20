var es = null;

function dialog(jid) {
    console.log('dialog placeholder for jid ' + jid);
}

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
            <a class="show" onclick="dialog(${e.Jid})">show</a>
        </div>
    </li>` + document.getElementsByClassName('event-list')[0].innerHTML;
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
    es.onopen = function () { }
    es.onerror = function () { }
    es.addEventListener('event', function (e) {
        addEvent(JSON.parse(e.data))
    }, false);
}

window.onbeforeunload = async function () {
    if (es != null) {
        es.close();
    }
}

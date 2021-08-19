// (new EventSource("/stream")).addEventListener("event", function (e) {
//     document.body.innerHTML += e.data + "</br>";
// });
function dialog(jid) {
    console.log('dialog for jid ' + jid);
}

function addEvent(json) {
    events = document.getElementsByClassName('event-list')[0];
    for (var i = 0; i < json.length; i++) {
        var event = json[i];
        events.innerHTML += `
        <li>
            <div class="event">
                <span class="minion">` + event.Minion + `</span>
                <span class="function">` + event.Function + `</span>
                <a class="show" onclick="dialog(` + event.Jid + `)">show</a>
            </div>
        </li>`;
    }
}

window.onload = async function () {
    fetch('/events')
        .then((response) => {
            return response.json();
        }).then((json) => {
            loader = document.getElementsByClassName('loading')[0];
            if (json.length > 0) {
                loader.remove();
                addEvent(json);
            } else {
                loader.innerHTML = "no result found";
            }
        });
}

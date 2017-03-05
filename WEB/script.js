/**
 * Created by wery_a on 04/03/2017.
 */

function updateDirectory() {
    var xhr = new XMLHttpRequest();
    xhr.withCredentials = true;

    xhr.addEventListener("readystatechange", function () {
        if (this.readyState === 4) {
            var json = JSON.parse(this.responseText);
            updatesFiles(json.files);
        }
    });

    xhr.open("GET", pathJoin("/api/list", path));
    xhr.send();
}

function updatesFiles(files) {
    var dom = document.getElementById("files");
    dom.innerHTML = "";
    for (var i = 0; i < files.length; ++i) {
        var file = files[i];
        var a = document.createElement("a");
        if (file.is_dir) {
            a.setAttribute("href", pathJoin(path, file.name));
            a.className = "item folder";
        } else {
            a.className = "item file";
        }
        a.setAttribute("data-url", pathJoin(path, file.name));
        var span = document.createElement("span");
        span.className = "filename";
        span.innerText = file.name;
        a.appendChild(span);
        dom.appendChild(a);
    }
    console.log(dom);
}
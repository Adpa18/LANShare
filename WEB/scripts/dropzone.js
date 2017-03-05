/**
 * Created by wery_a on 04/03/2017.
 */


(function () {

    var dropZone = document.getElementById("dropzone");

    var inner = document.getElementById("inner");
    var inputFile = document.getElementById("fileInput");

    function showDropZone() {
        dropZone.className = "dropzone active";
    }

    function hideDropZone() {
        dropZone.className = "dropzone";
    }

    function allowDrag(e) {
        if (true) {  // Test that the item being dragged is a valid one
            e.dataTransfer.dropEffect = "copy";
            e.preventDefault();
        }
    }

    function handleDrop(e) {
        e.preventDefault();
        hideDropZone();

        uploadFiles(e);
    }

    // 1
    window.addEventListener("dragenter", function (e) {
        showDropZone();
    });

    // 2
    dropZone.addEventListener("dragenter", allowDrag);
    dropZone.addEventListener("dragover", allowDrag);

    // 3
    dropZone.addEventListener("dragleave", function (e) {
        hideDropZone();
    });

    // 4
    dropZone.addEventListener("drop", handleDrop);

    inner.addEventListener("click", function (e) {
        inputFile.click();
    });

    inputFile.addEventListener("change", function (e) {
        uploadFiles(e);
    });

    var progressBar = document.getElementById("progressBar");
    var progresses = {};

    function uploadFiles(e) {
        var files = e.target.files || e.dataTransfer.files;

        // process all File objects
        for (var i = 0, f; f = files[i]; i++) {
            uploadFile(f, path);
        }
        inputFile.value = null;
    }

    function uploadFile(file, to) {
        var data = new FormData();
        data.append("uploadFile", file);

        var xhr = new XMLHttpRequest();
        xhr.withCredentials = true;

        {
            var id = Math.random();

            addProgressBar(id, file.name);
            xhr.addEventListener("readystatechange", function () {
                if (this.readyState === 4) {
                    setTimeout(function () {
                        var elem = progresses[id]["root"];
                        elem.style.width = 0;
                        elem.style.left = "100%";
                        elem.style.borderRadius = "10px";
                        progresses[id]["bar"].style.borderRadius = "10px";
                        setTimeout(function () {
                            elem.parentNode.removeChild(elem);
                            delete progresses[id];
                        }, 4000);
                    }, 500);
                    updateDirectory();
                }
            });
            xhr.upload.onprogress = function (e) {
                var percentage = Math.ceil((e.loaded / e.total ) * 100) + "%"
                progresses[id]["percentage"].innerText = percentage;
                progresses[id]["bar"].style.width = percentage;
            };
        }


        xhr.open("POST", pathJoin("/api/upload", to));
        xhr.send(data);
    }

    function addProgressBar(id, filename) {
        var divRoot = document.createElement("div");
        divRoot.className = "progressBar";

        var divContent = document.createElement("div");
        divContent.className = "contentBar";

        var divBar = document.createElement("div");
        divBar.className = "bar";
        divContent.appendChild(divBar);

        var spanFilename = document.createElement("span");
        spanFilename.className = "filename";
        divContent.appendChild(spanFilename);

        divRoot.appendChild(divContent);

        var divPercentage = document.createElement("div");
        divPercentage.className = "percentage";

        spanFilename.innerText = filename;
        divPercentage.innerText = 0 + "%";

        divRoot.appendChild(divPercentage);

        progressBar.appendChild(divRoot);

        progresses[id] = {
            "root": divRoot,
            "bar": divBar,
            "percentage": divPercentage
        };
    }
})();
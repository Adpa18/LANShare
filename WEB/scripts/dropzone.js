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

    function uploadFiles(e) {
        var files = e.target.files || e.dataTransfer.files;

        // process all File objects
        for (var i = 0, f; f = files[i]; i++) {
            uploadFile(f, path);
        }
    }

    function uploadFile(file, to) {
        var data = new FormData();
        data.append("uploadFile", file);

        var xhr = new XMLHttpRequest();
        xhr.withCredentials = true;

        xhr.addEventListener("readystatechange", function () {
            if (this.readyState === 4) {
                console.log(this.responseText);
                updateDirectory();
            }
        });

        xhr.open("POST", pathJoin("/api/upload", to));
        xhr.send(data);
    }
})();
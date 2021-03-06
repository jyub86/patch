///////////////////////////////
///////////////
// init check
const autosaveCheck = document.getElementById('autosave_check')
const ffmpeg = document.getElementById('ffmpeg_path')
const ffprobe = document.getElementById('ffprobe_path')
const oiioTool = document.getElementById('oiioTool_path')
const ocio = document.getElementById('ocio_path')
const preset = document.getElementById('preset')
const presetName = document.getElementById('preset_name')
const presetDelete = document.getElementById('preset_delete')
const splitter = document.getElementById('splitter')
const shotname = document.getElementById('shotname')
const seqsplitter = document.getElementById('seqsplitter')
const seqpad = document.getElementById('seqpad')
const startat = document.getElementById('startat')
const prefix = document.getElementById('prefix')
const suffix = document.getElementById('suffix')
const thumbCheck = document.getElementById('thumb_check')
const thumbColorspaceIn = document.getElementById('thumb_colorspace_in')
const thumbColorspaceOut = document.getElementById('thumb_colorspace_out')
const thumbWidth = document.getElementById('thumb_width')
const thumbHeight = document.getElementById('thumb_height')
const thumbDir = document.getElementById('thumb_dir')
const thumbName = document.getElementById('thumb_name')
const thumbExt = document.getElementById('thumb_ext')
const thumbPath = document.getElementById('thumb_path')
const thumbClass = document.querySelectorAll(".thumb")
const plateCheck = document.getElementById('plate_check')
const plateColorspaceIn = document.getElementById('plate_colorspace_in')
const plateColorspaceOut = document.getElementById('plate_colorspace_out')
const plateWidth = document.getElementById('plate_width')
const plateHeight = document.getElementById('plate_height')
const plateDir = document.getElementById('plate_dir')
const plateName = document.getElementById('plate_name')
const plateExt = document.getElementById('plate_ext')
const platePath = document.getElementById('plate_path')
const plateClass = document.querySelectorAll(".plate")
const videoCheck = document.getElementById('video_check')
const videoColorspaceIn = document.getElementById('video_colorspace_in')
const videoColorspaceOut = document.getElementById('video_colorspace_out')
const videoWidth = document.getElementById('video_width')
const videoHeight = document.getElementById('video_height')
const videoFps = document.getElementById('video_fps')
const videoCodec = document.getElementById('video_codec')
const videoDir = document.getElementById('video_dir')
const videoName = document.getElementById('video_name')
const videoExt = document.getElementById('video_ext')
const videoPath = document.getElementById('video_path')
const proxyDir = document.getElementById('proxy_dir')
const proxyName = document.getElementById('proxy_name')
const proxyExt = document.getElementById('proxy_ext')
const proxyPath = document.getElementById('proxy_path')
const videoClass = document.querySelectorAll(".video")
let osSep = ""
let initDatas = []

// set init data
initSummit()

function initSummit() {
    fetch("/init")
    .then(response => response.json())
    .then(data => init(data))
};

// get init data from user directroy
function init(data) {
    if ("error" in data) {
        message(data["error"], "alert-danger")
        return
    }
    initDatas = data["data"]
    if (!initDatas) {
        document.getElementById("modal_btn").click()
    } else {
        presetSetup()
        setInitData(initDatas[0])
    }
    if (autosaveCheck.checked) {
        // load autoSave
        fetch("/load")
        .then(response => response.json())
        .then(data => addData(data["data"], false))
    }
}

function presetSetup() {
    // remove all childs
    while ( preset.hasChildNodes() ) {
        preset.removeChild( preset.firstChild );
    }
    for (let i = 0; i < initDatas.length; i++) {
        preset.add(new Option(initDatas[i]["presetname"], initDatas[i]["presetname"]))
    }
    preset.add(new Option("new", "new"))
}

// set init data
function setInitData(initData) {
    autosaveCheck.checked = initData["autosavecheck"]
    ffmpeg.value = initData["ffmpeg"]
    ffprobe.value = initData["ffprobe"]
    oiioTool.value = initData["oiiotool"]
    ocio.value = initData["ocio"]
    presetName.value = initData["presetname"]
    splitter.value = initData["splitter"]
    shotname.value = initData["shotname"]
    seqsplitter.value = initData["seqsplitter"]
    seqpad.value = initData["seqpad"]
    startat.value = initData["startat"]
    thumbCheck.checked = initData["thumbcheck"]
    thumbColorspaceIn.value = initData["thumbcolorspacein"]
    thumbColorspaceOut.value = initData["thumbcolorspaceout"]
    thumbWidth.value = initData["thumbwidth"]
    thumbHeight.value = initData["thumbheight"]
    thumbDir.value = initData["thumbdir"]
    thumbName.value = initData["thumbname"]
    for (let opt, j = 0; opt = thumbExt.options[j]; j++) {
        if (opt.value == initData["thumbext"]) {
            thumbExt.selectedIndex = j;
            break;
        }
    }
    thumbPath.value = initData["thumbpath"]
    plateCheck.checked = initData["platecheck"]
    plateColorspaceIn.value = initData["platecolorspacein"]
    plateColorspaceOut.value = initData["platecolorspaceout"]
    plateWidth.value = initData["platewidth"]
    plateHeight.value = initData["plateheight"]
    plateDir.value = initData["platedir"]
    plateName.value = initData["platename"]
    plateExt.value = initData["plateext"]
    platePath.value = initData["platepath"]
    videoCheck.checked = initData["videocheck"]
    videoColorspaceIn.value = initData["videocolorspacein"]
    videoColorspaceOut.value = initData["videocolorspaceout"]
    videoWidth.value = initData["videowidth"]
    videoHeight.value = initData["videoheight"]
    videoFps.value = initData["videofps"]
    for (let opt, j = 0; opt = videoCodec.options[j]; j++) {
        if (opt.value == initData["videocodec"]) {
            thumbExt.selectedIndex = j;
            break;
        }
    }
    videoDir.value = initData["videodir"]
    videoName.value = initData["videoname"]
    for (let opt, j = 0; opt = videoExt.options[j]; j++) {
        if (opt.value == initData["videoext"]) {
            videoExt.selectedIndex = j;
            break;
        }
    }
    videoPath.value = initData["videopath"]
    proxyDir.value = initData["proxydir"]
    proxyName.value = initData["proxyname"]
    let proxyopts = proxyExt.options;
    for (let opt, j = 0; opt = proxyopts[j]; j++) {
        if (opt.value == initData["proxyext"]) {
            proxyExt.selectedIndex = j;
            break;
        }
    }
    proxyPath.value = initData["proxypath"]

    // check thumbnail option
    if (!thumbCheck.checked) {
        for (let i = 0; i <thumbClass.length; i++) {
            thumbClass[i].style.display = 'none';
        }
    }
    // check plate option
    if (!plateCheck.checked) {
        for (let i = 0; i <plateClass.length; i++) {
            plateClass[i].style.display = 'none';
        }
    }
    // check video option
    if (!videoCheck.checked) {
        for (let i = 0; i <videoClass.length; i++) {
            videoClass[i].style.display = 'none';
        }
    }
    presetInputShow()
}

///////////////////////////////
///////////////
// Modal

// open modal window
document.querySelector("#modal_btn").addEventListener("click", function(){
    const modal = new bootstrap.Modal(document.getElementById('preferences_modal'))
    modal.show()
    // get os separator
    fetch("/ossep")
    .then(response => response.json())
    .then(function(data) {
        osSep = data["data"]
    })
    // set prefix, suffix
    prefix.value = shotname.value.split(splitter.value, 2)[0]
    suffix.value = shotname.value.split(splitter.value, 2)[1]
})

// save init data when modal closed.
document.querySelector('#modal_close').addEventListener('click', function (event) {
    if (!presetName.value) {
        alert("No preset name. Can not save it.");
        return
    }
    fetch("/savepreset", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(currentPreferences())
    })
    .then(response => response.json())
    .catch(error => console.error('Error:', error))
    .then(function(data) {
        if ("error" in data) {
            message(data["error"], "alert-danger")
            return
        }
        initSummit()
        message("Preferences saved", "alert-success")
    })
})

function currentPreferences() {
    let pref = {
        "presetname" : presetName.value,
        "autosavecheck" : autosaveCheck.checked,
        "ffmpeg" : ffmpeg.value,
        "ffprobe" : ffprobe.value,
        "oiiotool" : oiioTool.value,
        "ocio" : ocio.value,
        "splitter" : splitter.value,
        "shotname" : shotname.value,
        "seqsplitter" : seqsplitter.value,
        "seqpad" : parseInt(seqpad.value),
        "startat" : parseInt(startat.value),
        "thumbcheck" : thumbCheck.checked,
        "thumbcolorspacein" : thumbColorspaceIn.value,
        "thumbcolorspaceout" : thumbColorspaceOut.value,
        "thumbwidth" : parseInt(thumbWidth.value),
        "thumbheight" : parseInt(thumbHeight.value),
        "thumbdir" : thumbDir.value,
        "thumbname" : thumbName.value,
        "thumbext" : thumbExt.options[thumbExt.selectedIndex].value,
        "thumbpath" : thumbPath.value,
        "platecheck" : plateCheck.checked,
        "platecolorspacein" : plateColorspaceIn.value,
        "platecolorspaceout" : plateColorspaceOut.value,
        "platewidth" : parseInt(plateWidth.value),
        "plateheight" : parseInt(plateHeight.value),
        "platedir" : plateDir.value,
        "platename" : plateName.value,
        "plateExt" : plateExt.options[plateExt.selectedIndex].value,
        "platepath" : platePath.value,
        "videocheck" : videoCheck.checked,
        "videocolorspacein" : videoColorspaceIn.value,
        "videocolorspaceout" : videoColorspaceOut.value,
        "videowidth" : parseInt(videoWidth.value),
        "videoheight" : parseInt(videoHeight.value),
        "videofps" : parseFloat(videoFps.value),
        "videocodec" : videoCodec.options[videoCodec.selectedIndex].value,
        "videodir" : videoDir.value,
        "videoname" : videoName.value,
        "videoext" : videoExt.options[videoExt.selectedIndex].value,
        "videopath" : videoPath.value,
        "proxydir" : proxyDir.value,
        "proxyname" : proxyName.value,
        "proxyext" : proxyExt.options[proxyExt.selectedIndex].value,
        "proxypath" : proxyPath.value,
    }
    return pref
}


document.querySelector('#preferences_modal').addEventListener('hidden.bs.modal', function (event) {
    err = "Preference setup is not completed. It may not operate correctly."
    isMissing = false;
    if (!ffmpeg.value) {
        isMissing = true;
    } else if (!ffprobe.value) {
        isMissing = true;
    } else if (!oiioTool.value) {
        isMissing = true;
    } else if (!ocio.value) {
        isMissing = true;
    } else if (!seqsplitter.value) {
        isMissing = true;
    } else if (!seqpad.value) {
        isMissing = true;
    }
    if (isMissing) {
        message(err, "alert-danger")
    }
})

// preset changed
preset.addEventListener("change", function(){
    let presetValue = preset.options[preset.selectedIndex].value
    for (let i = 0; i < initDatas.length; i++) {
        if (initDatas[i]["presetname"] === presetValue) {
            setInitData(initDatas[i])
            break
        }
    }
    presetInputShow()
})

// preset input element hide and show
function presetInputShow() {
    const presetClass = document.querySelectorAll(".preset_names")
    if (preset.options[preset.selectedIndex].value != "new") {
        for (let i = 0; i <presetClass.length; i++) {
            presetClass[i].hidden = true;
        }
    } else {
        for (let i = 0; i <presetClass.length; i++) {
            presetClass[i].hidden = false;
        }
    }
}
presetDelete.addEventListener("click", function(){
    let thisPreset = preset.options[preset.selectedIndex].value
    let an = confirm(`Delete ${thisPreset} preset?`);
    if (!an) {
        return
    }
    fetch("/deletepreset", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(currentPreferences())
    })
    .then(response => response.json())
    .catch(error => console.error('Error:', error))
    .then(function(data) {
        if ("error" in data) {
            message(data["error"], "alert-danger")
            return
        }
        initSummit()
        message("Preferences Deleted", "alert-success")
    })
})

// splitter change
splitter.addEventListener("keyup", function(){
    prefix.value = shotname.value.split(splitter.value, 2)[0]
    suffix.value = shotname.value.split(splitter.value, 2)[1]
    if (!splitter.value) {
        prefix.value = "";
        suffix.value = "";
    }
    updateThumbPath()
    updatePlatePath()
    updateVideoPath()
    updateProxyPath()
})

// shotname change
shotname.addEventListener("keyup", function(){
    prefix.value = shotname.value.split(splitter.value, 2)[0]
    suffix.value = shotname.value.split(splitter.value, 2)[1]
    updateThumbPath()
    updatePlatePath()
    updateVideoPath()
    updateProxyPath()
})

// seq splitter change
seqsplitter.addEventListener("keyup", function(){
    updateThumbPath()
    updatePlatePath()
    updateVideoPath()
    updateProxyPath()
})

// seq pad change
seqpad.addEventListener("keyup", function(){
    if (!seqpad.value.match(/^[0-9]+$/)) {
        seqpad.classList.add("is-invalid")
    } else {
        seqpad.classList.remove("is-invalid")
        updateThumbPath()
        updatePlatePath()
        updateVideoPath()
        updateProxyPath()
    }
})

// start at change
startat.addEventListener("keyup", function(){
    if (!startat.value.match(/^[0-9]+$/)) {
        startat.classList.add("is-invalid")
    } else {
        startat.classList.remove("is-invalid")
    }
})

// thumbnail
thumbCheck.addEventListener("change", function(){
    if (thumb_check.checked) {
        for (let i = 0; i <thumbClass.length; i++) {
            thumbClass[i].hidden = false;
        }
    } else {
        for (let i = 0; i <thumbClass.length; i++) {
            thumbClass[i].hidden = true;
        }
    }
})

thumbWidth.addEventListener("keyup", function(){
    if (!thumbWidth.value.match(/^[0-9]+$/)) {
        thumbWidth.classList.add("is-invalid")
    } else {
        thumbWidth.classList.remove("is-invalid")
    }
})

thumbHeight.addEventListener("keyup", function(){
    if (!thumbHeight.value.match(/^[0-9]+$/)) {
        thumbHeight.classList.add("is-invalid")
    } else {
        thumbHeight.classList.remove("is-invalid")
    }
})

thumbDir.addEventListener("keyup", function(){
    updateThumbPath()
})

thumbName.addEventListener("keyup", function(){
    updateThumbPath()
})

thumbExt.addEventListener("change", function(){
    updateThumbPath()
})

// plate
plateCheck.addEventListener("change", function(){
    if (plate_check.checked) {
        for (let i = 0; i <plateClass.length; i++) {
            plateClass[i].hidden = false;
        }
    } else {
        for (let i = 0; i <plateClass.length; i++) {
            plateClass[i].hidden = true;
        }
    }
})

plateWidth.addEventListener("keyup", function(){
    if (!plateWidth.value.match(/^[0-9]+$/)) {
        plateWidth.classList.add("is-invalid")
    } else {
        plateWidth.classList.remove("is-invalid")
    }
})

plateHeight.addEventListener("keyup", function(){
    if (!plateHeight.value.match(/^[0-9]+$/)) {
        plateHeight.classList.add("is-invalid")
    } else {
        plateHeight.classList.remove("is-invalid")
    }
})

plateDir.addEventListener("keyup", function(){
    updatePlatePath()
})

plateName.addEventListener("keyup", function(){
    updatePlatePath()
})

plateExt.addEventListener("change", function(){
    updatePlatePath()
})

// video
videoCheck.addEventListener("change", function(){
    if (video_check.checked) {
        for (let i = 0; i <videoClass.length; i++) {
            videoClass[i].hidden = false;
        }
    } else {
        for (let i = 0; i <videoClass.length; i++) {
            videoClass[i].hidden = true;
        }
    }
})

videoWidth.addEventListener("keyup", function(){
    if (!videoWidth.value.match(/^[0-9]+$/)) {
        videoWidth.classList.add("is-invalid")
    } else {
        videoWidth.classList.remove("is-invalid")
    }
})

videoHeight.addEventListener("keyup", function(){
    if (!videoHeight.value.match(/^[0-9]+$/)) {
        videoHeight.classList.add("is-invalid")
    } else {
        videoHeight.classList.remove("is-invalid")
    }
})

videoFps.addEventListener("keyup", function(){
    if (!videoFps.value.match(/^[0-9]+(\.[0-9]+)?$/)) {
        videoFps.classList.add("is-invalid")
    } else {
        videoFps.classList.remove("is-invalid")
    }
})

videoDir.addEventListener("keyup", function(){
    updateVideoPath()
})

videoName.addEventListener("keyup", function(){
    updateVideoPath()
})

videoExt.addEventListener("change", function(){
    updateVideoPath()
    if (videoExt.options[videoExt.selectedIndex].value == ".mp4") {
        for (let i = videoCodec.options.length-1; i >= 0; i--) {
            videoCodec.removeChild(videoCodec.options[i])
        }
        videoCodec.add(new Option('h264', 'h264'))
    } else {
        for (let i = videoCodec.options.length-1; i >= 0; i--) {
            videoCodec.removeChild(videoCodec.options[i])
        }
        videoCodec.add(new Option('h264', 'h264'))
        videoCodec.add(new Option('proresLT', 'proresLT'))
        videoCodec.add(new Option('proresHQ', 'proresHQ'))
        videoCodec.add(new Option('prores4444', 'prores4444'))
    }
})

// proxy
proxyDir.addEventListener("keyup", function(){
    updateProxyPath()
})

proxyName.addEventListener("keyup", function(){
    updateProxyPath()
})

proxyExt.addEventListener("change", function(){
    updateProxyPath()
})

function updateThumbPath() {
    return thumbPath.value = replaceName(thumbDir.value) + osSep + replaceName(thumbName.value) + thumbExt.options[thumbExt.selectedIndex].value
}

function updatePlatePath() {
    return platePath.value = replaceName(plateDir.value) + osSep + replaceName(plateName.value) + seqsplitter.value + "%0" + seqpad.value + "d" + plateExt.options[plateExt.selectedIndex].value
}

function updateVideoPath() {
    return videoPath.value = replaceName(videoDir.value) + osSep + replaceName(videoName.value) + videoExt.options[videoExt.selectedIndex].value
}

function updateProxyPath() {
    return proxyPath.value = replaceName(proxyDir.value) + osSep + replaceName(proxyName.value) + seqsplitter.value + "%0" + seqpad.value + "d" + proxyExt.options[proxyExt.selectedIndex].value
}

function replaceName(name) {
    let rename = name.replace(/<SHOTNAME>/, shotname.value)
    rename = rename.replace(/<PREFIX>/, prefix.value)
    rename = rename.replace(/<SUFFIX>/, suffix.value)
    return rename
}

///////////////////////////////
///////////////
// Table

// table data
let tabledata = [{id:1, path:"", framein:0, frameout:0, framerange:0, pad:0, timecodein:"00:00:00", timecodeout:"00:00:00", width:0, height:0, ext:"", fps:0, codec:"",
shotname:"", trimin:0, trimout:0, trimintc:"", trimouttc:"", colorin:"", colorout:"", rewidth:0, reheight:0, pub:true, log:""}];

// last index number of table
let lastIndex = 0

// add options
let addOption = "";

//// table menu
// cell right click
const cellContextMenu = [
    {
        label:"Select all",
        action:function(e, cell){
            table.selectRow("visible");
        }
    },
    {
        label:"Deselect all",
        action:function(e, cell){
            table.deselectRow();
        }
    },
    {
        label:"Duplicate Row",
        action:function(e, cell){
            let data = cell.getRow().getData()
            data["id"] = lastIndex
            lastIndex = lastIndex + 1;
            table.addRow(data, true, cell.getRow().getIndex());
            autoSave()
        }
    },
    {
        label:"Delete Row",
        action:function(e, cell){
            cell.getRow().delete();
            document.getElementById("select-stats").innerHTML = table.getSelectedRows().length;
            autoSave()
        }
    },
    {
        label:"Delete Selected Row",
        action:function(e, cell){
            table.deleteRow(table.getSelectedRows());
            document.getElementById("select-stats").innerHTML = table.getSelectedRows().length;
            autoSave()
        }
    },
]

// edit cell right click
const editContextMenu = [
    {
        label:"Select all",
        action:function(e, cell){
            table.selectRow("visible");
        }
    },
    {
        label:"Deselect all",
        action:function(e, cell){
            table.deselectRow();
        }
    },
    {
        label:"Fill data from preference",
        action:function(e, cell){
            let thisRow = cell.getRow()
            let data = thisRow.getData()
            data["colorin"] = plateColorspaceIn.value
            data["colorout"] = plateColorspaceOut.value
            data["rewidth"] = parseInt(plateWidth.value)
            data["reheight"] = parseInt(plateHeight.value)
            table.updateRow(thisRow.getIndex(), data);
            let selRows = table.getSelectedRows();
            if (selRows) {
                for(var i=0; i<selRows.length; i++) {
                    data = selRows[i].getData()
                    data["colorin"] = plateColorspaceIn.value
                    data["colorout"] = plateColorspaceOut.value
                    data["rewidth"] = parseInt(plateWidth.value)
                    data["reheight"] = parseInt(plateHeight.value)
                    table.updateRow(selRows[i].getIndex(), data);
                }
            }
            table.redraw(true);
            autoSave();
        }
    },
    {
        label:"Auto fill names",
        action:function(e, cell){
            const userConfirm2 = document.getElementById("user_confirm2")
            userConfirm2.style.zIndex = 999;
            var toast = new bootstrap.Toast(userConfirm2)
            userConfirm2.style.left = (mouseX) + 'px';
            userConfirm2.style.top = (mouseY) + 'px';
            toast.show()
        }
    },
]

let table = new Tabulator("#table", {
    data:tabledata,
    autoColumns:true,
    layout:"fitDataFill",
    addRowPos:"top",
    history:true,
    pagination:"local",
    paginationSize:10,
    paginationSizeSelector:[10, 20, 50, 100],
    selectable:true,
    selectableRollingSelection:false,
    initialSort:[
        {column:"path", dir:"asc"},
    ],
    autoColumnsDefinitions:[
        {title:"ID", field:"id", headerVertical:true, hozAlign:"center", visible:false, editable:false, download:false},
        {title:"FILE PATH", field:"path", headerVertical:true, hozAlign:"left", visible:true, editable:false, download:true,
        contextMenu:cellContextMenu},
        {title:"FRAME IN", field:"framein", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, formatter:addPadding, contextMenu:cellContextMenu},
        {title:"FRAME OUT", field:"frameout", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, formatter:addPadding, contextMenu:cellContextMenu},
        {title:"FRAME RANGE", field:"framerange", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"SEQUENCE PAD", field:"pad", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"TIMECODE IN", field:"timecodein", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"TIMECODE OUT", field:"timecodeout", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"WIDTH", field:"width", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"HEIGHT", field:"height", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"EXTENTION", field:"ext", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"FPS", field:"fps", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"CODEC", field:"codec", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
        {title:"SHOT NAME", field:"shotname", headerVertical:true, hozAlign:"left", editable:true, editor:"input", download:true,
        contextMenu:editContextMenu},
        {title:"TRIM IN(frame)", field:"trimin", headerVertical:true, hozAlign:"center", editable:true, editor:"number", download:true,
        formatter:trimCheck, contextMenu:editContextMenu},
        {title:"TRIM OUT(frame)", field:"trimout", headerVertical:true, hozAlign:"center", editable:true, editor:"number", download:true,
        formatter:trimCheck, contextMenu:editContextMenu},
        {title:"TRIM IN(timecode)", field:"trimintc", headerVertical:true, hozAlign:"center", editable:true, editor:"input", download:true,
        editorParams:{ //masking
            mask:"99:99:99:99",
            maskAutoFill:true,
        }, contextMenu:editContextMenu},
        {title:"TRIM OUT(timecode)", field:"trimouttc", headerVertical:true, hozAlign:"center", editable:true, editor:"input", download:true,
        editorParams:{ // masking
            mask:"99:99:99:99",
            maskAutoFill:true,
        }, contextMenu:editContextMenu},
        {title:"COLORSPACE IN", field:"colorin", headerVertical:true, hozAlign:"center", editable:true, editor:"input", download:true,
        contextMenu:editContextMenu},
        {title:"COLORSPACE OUT", field:"colorout", headerVertical:true, hozAlign:"center", editable:true, editor:"input", download:true,
        contextMenu:editContextMenu},
        {title:"RESIZE WIDTH", field:"rewidth", headerVertical:true, hozAlign:"center", editable:true, editor:"number", download:true,
        formatter:addPadding, contextMenu:editContextMenu},
        {title:"RESIZE HEIGHT", field:"reheight", headerVertical:true, hozAlign:"center", editable:true, editor:"number", download:true,
        formatter:addPadding, contextMenu:editContextMenu},
        {title:"PUBLISH", field:"pub", headerVertical:true, hozAlign:"center", formatter:"tickCross", editor:"tickCross", download:true, contextMenu:cellContextMenu},
        {title:"LOG", field:"log", headerVertical:true, hozAlign:"center", visible:true, editable:false, download:true, contextMenu:cellContextMenu},
    ],
    dataChanged:function(data){
        //data - the updated table data
        document.getElementById("total-stats").innerHTML = table.getDataCount();
    },
    cellEdited:function(cell){
        // multi edit
        let selRows = table.getSelectedRows();
        selRows.push(cell.getRow())
        for(var i=0; i<selRows.length; i++) {
            let data = selRows[i].getData()
            let field = cell.getField()
            let value = cell.getValue()
            data[field] = value
            data["log"] = ""
            // pub check
            if (!isShotname(data) && data["pub"]) {
                data["pub"] = false;
                data["log"] = "shotname does not exist."
            }
            if (isSamename(data) && data["pub"]) {
                data["pub"] = false;
                data["log"] = "same shotname exists."
            }
            // seq code check
            if (field == "trimintc" && !isTimecode(data)) {
                data[field] = "";
            }
            if (field == "trimouttc" && !isTimecode(data)) {
                data[field] = "";
            }
            table.updateRow(selRows[i].getIndex(), data);
        }
        table.redraw(true);
        autoSave();
    },
    cellClick:function(e, cell){
        // prevent row select for editable cells
        if (cell.getColumn().getDefinition()["editable"]) {
            let row = cell.getRow();
            row.toggleSelect();
        }
    },
    rowSelectionChanged:function(data, rows){
        //update selected row counter on selection change
    	document.getElementById("select-stats").innerHTML = data.length;
    },
});

//// formatter functions
// add padding
function addPadding(cell) {
    let data = cell.getRow().getData();
    return fillZero(data["pad"], cell.getValue()) 
}

// trim check
function trimCheck(cell) {
    let value = cell.getValue();
    if (!value) {
        let data = cell.getRow().getData();
        data[cell.getField()] = 0;
        try {table.updateRow(cell.getRow().getIndex(), data);}
        catch {}
        return 0
    }
    return value
}

//// support functions
// check shotname is not empty
function isShotname(data) {
    if (data.shotname == undefined || data.shotname == "") {
        return false;
    }
    return true;
}

// check timecode is not empty
function isTimecode(data) {
    if (data.timecodein == "00:00:00:00" && data.timecodeout == "00:00:00:00") {
        return false;
    }
    return true;
}

// check same shotname
function isSamename(data) {
    if (!data.shotname) {
        return false;
    }
    let sameNames = table.searchData("shotname", "=", data.shotname);
    if (sameNames.length > 1) {
        return true;
    }
    return false;
}

// auto fill
document.getElementById("toast_fill").addEventListener("click", function(){
    let fillName = document.getElementById("toast_name").value
    let fillStart = document.getElementById("toast_start").value
    let fillPad = String(fillStart).length;
    let fillStep = document.getElementById("toast_step").value
    let selRows = table.getSelectedRows();
    if (selRows) {
        let num = parseInt(fillStart)
        for(var i=0; i<selRows.length; i++) {
            let fullName = fillName + fillZero(fillPad, String(num))
            data = selRows[i].getData()
            data["shotname"] = fullName
            table.updateRow(selRows[i].getIndex(), data)
            num = parseInt(num) + parseInt(fillStep)
        }
    }
    table.redraw(true);
    autoSave();
});

// delete dummy row
let dummyRow = table.searchRows("path", "=", "");
table.deleteRow(dummyRow);

// load csv button
document.getElementById("load_btn").addEventListener("click", function(){
    const userConfirm = document.getElementById("user_confirm")
    userConfirm.style.zIndex = 999;
    var toast = new bootstrap.Toast(userConfirm)
    toast.show()
});
document.getElementById("toast_clear").addEventListener("click", function(){
    document.getElementById("clear").click()
    document.getElementById("formFile").click()
});
document.getElementById("toast_append").addEventListener("click", function(){
    document.getElementById("formFile").click()
});
document.getElementById("formFile").addEventListener('change', function (e) {
    e.preventDefault();
    let csv = this.files[0];
    let reader = new FileReader();
    reader.onload = function (e) {
        let text = e.target.result;
        addData(csvToJSON(text), true)
    };
    reader.readAsText(csv);
});

// convert title to field name
function titleToField(title) {
    const obj = {
        "FILE PATH":"path",
        "FRAME IN":"framein",
        "FRAME OUT":"frameout",
        "FRAME RANGE":"framerange",
        "SEQUENCE PAD":"pad",
        "TIMECODE IN":"timecodein",
        "TIMECODE OUT":"timecodeout",
        "WIDTH":"width",
        "HEIGHT":"height",
        "EXTENTION":"ext",
        "FPS":"fps",
        "CODEC":"codec",
        "SHOT NAME":"shotname",
        "TRIM IN(frame)":"trimin",
        "TRIM OUT(frame)":"trimout",
        "TRIM IN(timecode)":"trimintc",
        "TRIM OUT(timecode)":"trimouttc",
        "COLORSPACE IN":"colorin",
        "COLORSPACE OUT":"colorout",
        "RESIZE WIDTH":"rewidth",
        "RESIZE HEIGHT":"reheight",
        "PUBLISH":"pub",
        "LOG":"log"
    }
    if (title in obj) {
        return obj[title]
    }
    return title
}

// convert csv to json
function csvToJSON(text){
    const rows = text.split("\n");
    const jsonArray = [];
    const header = rows[0].split(",");
    for(let i = 1; i < rows.length; i++){
        let obj = {};
        let row = rows[i].split(",");
        for(let j=0; j < header.length; j++){
            obj[titleToField(JSON.parse(header[j]))] = JSON.parse(row[j]);
        }
        jsonArray.push(obj);
    }
    return jsonArray;
}

// append data form
document.querySelector(".form-floating").addEventListener('submit', handleSubmit);

function handleSubmit(event) {
    event.preventDefault();
    document.querySelector(".loading").style.display = 'block';
    const pathInput = document.querySelector("#floatingInputValue");
    let path = pathInput.value;
    if (!path) {
        document.querySelector(".loading").style.display = 'none';
        return
    }
    fetch("/search", {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        body: new URLSearchParams({
            path:path
        })
    })
    .then(response => response.json())
    .catch(error => console.error('Error:', error))
    .then(function(data) {
        document.querySelector(".loading").style.display = 'none';
        if ("error" in data) {
            message(data["error"], "alert-danger")
            return
        }
        addData(data["data"], true)
    })
}

// add data from object
function addData(data, option) {
    show_all()
    addOption = ""
    if (!option) {
        addOption = "addAll"
    }
    for (var i = 0; i < data.length; i++) {
        item = {
            id:JSON.parse(lastIndex),
            path:data[i]["path"],
            framein:JSON.parse(data[i]["framein"]),
            frameout:JSON.parse(data[i]["frameout"]),
            framerange:JSON.parse(data[i]["framerange"]),
            pad:JSON.parse(data[i]["pad"]),
            timecodein:data[i]["timecodein"],
            timecodeout:data[i]["timecodeout"],
            width:JSON.parse(data[i]["width"]),
            height:JSON.parse(data[i]["height"]),
            ext:data[i]["ext"],
            fps:JSON.parse(data[i]["fps"]),
            codec:data[i]["codec"],
            shotname:data[i]["shotname"],
            trimin:JSON.parse(data[i]["trimin"]),
            trimout:JSON.parse(data[i]["trimout"]),
            trimintc:data[i]["trimintc"],
            trimouttc:data[i]["trimouttc"],
            colorin:data[i]["colorin"],
            colorout:data[i]["colorout"],
            rewidth:JSON.parse(data[i]["rewidth"]),
            reheight:JSON.parse(data[i]["reheight"]),
            pub:JSON.parse(data[i]["pub"]),
            log:data[i]["log"]
        }
        let value = checkItem(item)
        if (!value) {
            continue
        }
        table.addRow(item, true);
        lastIndex = lastIndex + 1;
    }
    table.redraw(true);
    autoSave()
};

function checkItem(item) {
    if (addOption == "addAll") {
        return true;
    } else if (addOption == "noAll") {
        return false;
    }
    let sameNames = table.searchData("path", "=", item.path);
    if (sameNames.length == 0) {
        return true;
    }
    let value = confirm(item.path+"\nSame file exists.\nDo you want add it?");
    let option = confirm("Apply to all items?");
    if (option) {
        if (value) {
            addOption = "addAll"
        } else {
            addOption = "noAll"
        }
    }
    return value
}

// add padding to 0
function fillZero(p, n){
    if (!n) {
        return n
    }
    let width = String(p);
    let num = String(n);
    return num.length >= width ? num:new Array(width-num.length+1).join('0')+num;
}

// Autosave
function autoSave(){
    if (!autosaveCheck.checked) {
        return
    }
    let array = table.getData();
    fetch("/autosave", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(array)
    })
    .then(response => response.json())
    .then(function(data) {
        if ("error" in data) {
            message(data["error"], "alert-danger")
            return
        }
        console.log("autosaved :", data["data"])
    })
}

// trigger download of data.csv file
document.getElementById("save_btn").addEventListener("click", function(){
    let current = new Date().toJSON().slice(0,9).replace(/-/g,'');
    table.download("csv", "patch_"+current+".csv");
});

// undo button
document.getElementById("undo_btn").addEventListener("click", function(){
    table.undo();
  });
  
// redo button
document.getElementById("redo_btn").addEventListener("click", function(){
    table.redo();
});

// show and hide column
const columns = {
    frame : ["framein", "frameout", "framerange", "pad"],
    timecode : ["timecodein", "timecodeout"],
    format : ["width", "height"],
    ext : ["ext"],
    video : ["fps", "codec"]
}

// show input column
function showColumn(col) {
    for (let i = 0; i < columns[col].length; i++) {
        table.showColumn(columns[col][i])
    }
    document.getElementById(col).classList.remove("active")
}

// hide input column
function hideColumn(col) {
    for (let i = 0; i < columns[col].length; i++) {
        table.hideColumn(columns[col][i])
    }
    document.getElementById(col).classList.add("active")
}

// hide all columns
function hide_all() {
    for (let [k, v] of Object.entries(columns)) {
        hideColumn(k)
    }
}

// show all columns
function show_all() {
    for (let [k, v] of Object.entries(columns)) {
        showColumn(k)
    }
}

// click hide all button
document.getElementById("hide_all").addEventListener("click", function(){
    hide_all();
    table.redraw(true);
});

// click show all button
document.getElementById("show_all").addEventListener("click", function(){
    show_all();
    table.redraw(true);
});

// click frame button
document.getElementById("frame").addEventListener("click", function(){
    if (table.getColumn(columns["frame"][0]).isVisible()) {
        hideColumn("frame")
    } else {
        showColumn("frame")
    }
    table.redraw(true);
});

// click timecode button
document.getElementById("timecode").addEventListener("click", function(){
    if (table.getColumn(columns["timecode"][0]).isVisible()) {
        hideColumn("timecode")
    } else {
        showColumn("timecode")
    }
    table.redraw(true);
});

// click format button
document.getElementById("format").addEventListener("click", function(){
    if (table.getColumn(columns["format"][0]).isVisible()) {
        hideColumn("format")
    } else {
        showColumn("format")
    }
    table.redraw(true);
});

// click ext button
document.getElementById("ext").addEventListener("click", function(){
    if (table.getColumn(columns["ext"][0]).isVisible()) {
        hideColumn("ext")
    } else {
        showColumn("ext")
    }
    table.redraw(true);
});

// click video button
document.getElementById("video").addEventListener("click", function(){
    if (table.getColumn(columns["video"][0]).isVisible()) {
        hideColumn("video")
    } else {
        showColumn("video")
    }
    table.redraw(true);
});

// select row on "select all" button click
document.getElementById("select_all").addEventListener("click", function(){
    table.selectRow("visible");
});

// deselect row on "deselect all" button click
document.getElementById("deselect_all").addEventListener("click", function(){
    table.deselectRow();
});

// Update filters on value change
document.getElementById("filter-value").addEventListener("keyup", function(){
    table.setFilter("path", "keywords", document.getElementById("filter-value").value, {matchAll:true});
});

// click redraw button
document.getElementById("redraw").addEventListener("click", function(){
    table.redraw(true);
});

// Reset table contents on "Reset the table" button click
document.getElementById("clear").addEventListener("click", function(){
    table.clearData()
    document.getElementById("total-stats").innerHTML = table.getDataCount();
    autoSave()
});

// shortcut
window.addEventListener("keydown", function(e) {
    if (e.keyCode == 27) { // esc Event
        table.deselectRow();
    }
});

// Publish
document.getElementById("pub_btn").addEventListener("click", function(){
    message("Item Published", "alert-success")
    let array = table.getData();
    let now = new Date();
    let dateString = now.getUTCFullYear() + "-" +
    ("0" + (now.getUTCMonth()+1)).slice(-2) + "-" +
    ("0" + now.getUTCDate()).slice(-2) + " " +
    ("0" + now.getUTCHours()).slice(-2) + ":" +
    ("0" + now.getUTCMinutes()).slice(-2) + ":" +
    ("0" + now.getUTCSeconds()).slice(-2);
    for (let i = 0; i < array.length; i++) {
        if (array[i].pub == false) {
            continue
        }
        array[i].log = `Sent at : ${dateString}`
    }
    table.updateData(array)
    autoSave()
    fetch("/publish", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(array)
    })
    .then(response => response.json())
    .then(function(data) {
        if ("error" in data) {
            message(data["error"], "alert-danger")
            return
        }
        let newData = data["data"];
        for (let i = 0; i < newData.length; i++) {
            if (newData[i].pub == false) {
                continue
            }
            let path = newData[i].path;
            let log = newData[i].log;
            let searchDatas = table.searchData("shotname", "=", newData[i].shotname)
            for (let i = 0; i < searchDatas.length; i++) {
                if (searchDatas[i].path == path) {
                    searchDatas[i].log = log
                    table.updateRow(searchDatas[i].id, searchDatas[i])
                }
            }
        }
        autoSave()
        message("Publish done", "alert-success")
    })
})

// log
document.getElementById("log_btn").addEventListener("click", function(){
    location.href = "/log";
})

// Message
function message(text, cls) {
    const alertMsg = document.getElementById("alert_msg")
    const msg = document.getElementById("msg")
    alertMsg.classList.add(cls)
    alertMsg.classList.remove("hide")
    msg.innerHTML = text
    setTimeout(function(){
        alertMsg.classList.add("hide");
        alertMsg.classList.remove(cls);
    }, 3000);
}

// mouse right click position
let mouseX = 0;
let mouseY = 0;
function printMousePos(event) {
    mouseX = event.pageX;
    mouseY = event.pageY;
}
document.addEventListener("contextmenu", printMousePos);
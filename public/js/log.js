// log
document.getElementById("home_btn").addEventListener("click", function(){
    console.log("A")
    location.href = "/";
})

// table data
let tabledata = [{id:0, shotname:"", task:"", subtask:"", cmdstr:"", start:"", end:"", duration:"", errorlog:""}];

// last index number of table
let lastIndex = 0

// make Table
let table = new Tabulator("#log_table", {
    data:tabledata,
    autoColumns:true,
    layout:"fitColumns",
    addRowPos:"top",
    height:"auto",
    pagination:"local",
    paginationSize:50,
    paginationSizeSelector:[50, 100, 200],
    autoColumnsDefinitions:[
        {title:"ID", field:"id", headerVertical:true, hozAlign:"center", visible:false},
        {title:"SHOTNAME", field:"shotname", headerVertical:true, hozAlign:"center"},
        {title:"TASK", field:"task", headerVertical:true, hozAlign:"center"},
        {title:"SUBTASK", field:"subtask", headerVertical:true, hozAlign:"center"},
        {title:"CMD_STRING", field:"cmdstr", headerVertical:true, hozAlign:"center"},
        {title:"START", field:"start", headerVertical:true, hozAlign:"center"},
        {title:"END", field:"end", headerVertical:true, hozAlign:"center"},
        {title:"DURATION", field:"duration", headerVertical:true, hozAlign:"center"},
        {title:"ERR_LOG", field:"errorlog", headerVertical:true, hozAlign:"center"},
    ],
});

// delete dummy row
let dummyRow = table.searchRows("shotname", "=", "");
table.deleteRow(dummyRow);

// set init data
init()

function init() {
    fetch("/logupdate")
    .then(response => response.json())
    .then(data => addData(data["data"]))
};

// add data from object
function addData(data) {
    for (var i = 0; i < data.length; i++) {
        item = {
            id:JSON.parse(lastIndex),
            shotname:data[i]["shotname"],
            task:data[i]["task"],
            subtask:data[i]["subtask"],
            cmdstr:data[i]["cmdstr"],
            start:data[i]["start"],
            end:data[i]["end"],
            duration:data[i]["duration"],
            errorlog:data[i]["errorlog"],
        }
        table.addRow(item, true);
        lastIndex = lastIndex + 1;
    }
    table.redraw(true);
};

// Update filters on value change
document.getElementById("filter-value").addEventListener("keyup", function(){
    table.setFilter("shotname", "keywords", document.getElementById("filter-value").value, {matchAll:true});
});

async function POST(url, body) {
    try {
        const response = await fetch(url, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: body
        });
        return await response.json();
    } catch(error) {
        return {error: error.message || "Unknown error"};
    }
}

async function DELETE(url) {
    try {
        const response = await fetch(url, {
            method: "DELETE",
        });
        return await response.json();
    } catch(error) {
        return {error: error.message || "Unknown error"};
    }
}

async function plusOne(id) {
    const countTag = document.querySelector("#itemid-"+id);
    const curCount = parseInt(countTag.textContent, 10);
    const reqbody = {"id": parseInt(id, 10)}

    response = await POST("//192.168.0.142:12000/api/v1/consumable-items/actions/plus-one/invoke", JSON.stringify(reqbody));
    countTag.textContent = (curCount + 1).toString()
    //console.log(response)
}

async function minusOne(id) {
    const countTag = document.querySelector("#itemid-"+id);
    const curCount = parseInt(countTag.textContent, 10);
    const reqbody = {"id": parseInt(id, 10)}

    if(curCount == 0) {
        return
    }
    response = await POST("//192.168.0.142:12000/api/v1/consumable-items/actions/minus-one/invoke", JSON.stringify(reqbody));
    countTag.textContent = (curCount - 1).toString()
    //console.log(response)
}

async function deleteOne(id) {
    response = await DELETE("/api/v1/consumable-items/"+id);
    if(response.status !== "ok") {
        return;
    }
    const row = document.querySelector("#row-itemid"+id);
    const name = document.querySelector("#itemid"+id+"-name").textContent;
    const resultArea = document.querySelector("#actionResult");
    row.remove();
    resultArea.textContent = "Remove successfully: " + name;
    //console.log(response)

}

async function addItem() {
    const name = document.querySelector("#name").value;
    const count = document.querySelector("#count").value;

    if(name === "" || count === "") {
        console.log("invalid input");
        console.log("name:", name);
        console.log("count:", count);
        return
    }

    const reqbody = {"name":name, "count":parseInt(count, 10)}
    const result = document.querySelector("#addResult")

    console.log(reqbody)
    response = await POST("/api/v1/consumable-items", JSON.stringify(reqbody));

    if(response.status === "ok") {
        result.textContent = "successfully added => " + name + " @ " + count
    }

}
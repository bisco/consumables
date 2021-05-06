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

async function plusOne(id) {
    const countTag = document.querySelector("#itemid-"+id);
    const curCount = parseInt(countTag.textContent, 10);
    const reqbody = {"id": parseInt(id, 10)}

    response = await POST("//192.168.0.142:12000/api/v1/consumable-items/actions/plus-one/invoke", JSON.stringify(reqbody));
    countTag.textContent = (curCount + 1).toString()
    console.log(response)
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
    console.log(response)
}

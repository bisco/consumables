//------------------------
// HTTP Method
//------------------------
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

async function PATCH(url, body) {
    try {
        const response = await fetch(url, {
            method: "PATCH",
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

//------------------------
// Functions for buttons at "/"
//------------------------
async function plusOne(id) {
    const countTag = document.querySelector("#itemid"+id+"-count");
    const curCount = parseInt(countTag.textContent, 10);
    const reqbody = {"id": parseInt(id, 10)}

    response = await POST("/api/v1/consumable-items/actions/plus-one/invoke", JSON.stringify(reqbody));
    countTag.textContent = (curCount + 1).toString()
    //console.log(response)
}

async function minusOne(id) {
    const countTag = document.querySelector("#itemid"+id+"-count");
    const curCount = parseInt(countTag.textContent, 10);
    const reqbody = {"id": parseInt(id, 10)}

    if(curCount == 0) {
        return
    }
    response = await POST("/api/v1/consumable-items/actions/minus-one/invoke", JSON.stringify(reqbody));
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

//------------------------
// add entry at "/add"
//------------------------
async function addItem() {
    const name = document.querySelector("#name").value;
    const count = document.querySelector("#count").value;
    const category = document.querySelector("#category").value;
    const subcategory = document.querySelector("#subcategory").value;
    const result = document.querySelector("#addResult")

    if(name === "" || count === "") {
        result.textContent = "add fail: name or count is blank";
        return;
    }

    const reqbody = {"name":name, "count":parseInt(count, 10), "category":category, "subcategory": subcategory};

    console.log(reqbody)
    response = await POST("/api/v1/consumable-items", JSON.stringify(reqbody));

    if(response.status === "ok") {
        result.textContent = "successfully added => " + name + " @ " + count
    }

}

//------------------------
// modify entry
//------------------------
async function modifyItem() {
    const name = document.querySelector("#name").value;
    const count = document.querySelector("#count").value;
    const category = document.querySelector("#category").value;
    const subcategory = document.querySelector("#subcategory").value;
    const id = document.querySelector("#id").value;

    const result = document.querySelector("#modifyResult")

    if(name === "" || count === "" || id === "") {
        result.textContent = "add fail: id or name or count is blank";
        return;
    }

    const reqbody = {"name": name, 
                     "count": parseInt(count, 10),
                     "category": category,
                     "subcategory": subcategory};

    console.log(reqbody)
    response = await PATCH("/api/v1/consumable-items/"+id, JSON.stringify(reqbody));

    if(response.status === "ok") {
        result.textContent = "successfully modified"
    }

}

//------------------------
// format strings
//------------------------
function convertTZ() {
    let currentPath = location.pathname;
    if(currentPath !== "/") {
        return;
    }
    document.querySelectorAll(".updated-at").forEach((tag) => {
        ts = parseInt(tag.textContent, 10);
        d = new Date(ts);
       tag.textContent = d.toLocaleDateString() + " " + d.toLocaleTimeString();
    });
}

function changeNavigation() {
    let currentPath = location.pathname;
    let naviSelector;
    if(currentPath === "/") {
        naviSelector = "#navigation-home";
    } else if(currentPath === "/add"){ // add
        naviSelector = "#navigation-add";
    } else {
        return;
    }
    let nav = document.querySelector(naviSelector);
    nav.classList.add("is-selected");
}


//------------------------
// 絞り込み検索
//------------------------
function filter(className) {
    document.querySelectorAll(className).forEach((t) => {
        t.classList.remove("u-off-screen");
    });
}

function unfilter(className) {
    document.querySelectorAll(className).forEach((t) => {
        t.classList.add("u-off-screen");
    });
}

function unfilterAll() {
    document.querySelectorAll(".p-chip__dismiss").forEach((t) => {
        if(t.childNodes[0].className === "p-icon--success") {
            t.click();
        }
    });    
}

function toggleTag(id) {
    const stateTag = document.querySelector("#"+id+"-state");
    let state = stateTag.classList[0];
    if(state === "p-icon--success") {
        stateTag.classList.remove("p-icon--success");
        stateTag.classList.add("p-icon--error");
        unfilter("."+id+"-items");
    } else {
        stateTag.classList.remove("p-icon--error");
        stateTag.classList.add("p-icon--success");
        filter("."+id+"-items");
    }
}

/**
  Toggles the necessary aria- attributes' values on the accordion panels
  and handles to show or hide them.
  @param {HTMLElement} element The tab that acts as the handles.
  @param {Boolean} show Whether to show or hide the accordion panel.
*/
function toggleExpanded(element, show) {
    var target = document.getElementById(element.getAttribute('aria-controls'));
  
    if (target) {
      element.setAttribute('aria-expanded', show);
      target.setAttribute('aria-hidden', !show);
      if(show) {
        unfilterAll();
      }
    }
  }
  
  /**
    Attaches event listeners for the accordion open and close click events.
    @param {HTMLElement} accordionContainer The accordion container element.
  */
  function setupAccordion(accordionContainer) {
    // Finds any open panels within the container and closes them.
    function closeAllPanels() {
      var openPanels = accordionContainer.querySelectorAll('[aria-expanded=true]');
  
      for (var i = 0, l = openPanels.length; i < l; i++) {
        toggleExpanded(openPanels[i], false);
      }
    }
  
    // Set up an event listener on the container so that panels can be added
    // and removed and events do not need to be managed separately.
    accordionContainer.addEventListener('click', function (event) {
      var target = event.target;
  
      if (target.closest) {
        target = target.closest('[class*="p-accordion__tab"]');
      } else if (target.msMatchesSelector) {
        // IE friendly `Element.closest` equivalent
        // as in https://developer.mozilla.org/en-US/docs/Web/API/Element/closest
        do {
          if (target.msMatchesSelector('[class*="p-accordion__tab"]')) {
            break;
          }
          target = target.parentElement || target.parentNode;
        } while (target !== null && target.nodeType === 1);
      }
  
      if (target) {
        var isTargetOpen = target.getAttribute('aria-expanded') === 'true';
        closeAllPanels();
  
        // Toggle visibility of the target panel.
        toggleExpanded(target, !isTargetOpen);
      }
    });
  }
  
  // Setup all accordions on the page.
  var accordions = document.querySelectorAll('.p-accordion');
  
  for (var i = 0, l = accordions.length; i < l; i++) {
    setupAccordion(accordions[i]);
  }

  
//------------------------
// page initialization
//------------------------
convertTZ();
changeNavigation();
